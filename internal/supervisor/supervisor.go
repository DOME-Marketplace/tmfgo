package supervisor

import (
	"log/slog"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

// Run acts as an init process (PID 1), launching a child process and forwarding
// signals to it. This is typically used in container environments where
// the application might be PID 1.
//
// It sets up the child process to share stdout/stderr, places it in the same
// process group, and captures system signals (SIGINT, SIGTERM, SIGHUP) to
// gracefully relay them to the child, while continuously reaping zombie processes.
func Run(args []string) {
	ourPid := os.Getpid()

	// Get the name of our executable, to be able to restart it automatically
	ourExecPath, err := os.Executable()
	if err != nil {
		slog.Error("Failed to get executable path", slog.Any("error", err))
		panic(err)
	}

	slog.Info("We are the INIT process!", "PID", ourPid, "executable", ourExecPath, "args", args)

	// Pass to child all arguments except the "-init" flag, so the child runs as a normal process.
	childArgs := make([]string, 0, len(args))
	for _, a := range args {
		if a != "-init" && a != "--init" {
			childArgs = append(childArgs, a)
		}
	}
	cmd := exec.Command(ourExecPath, childArgs...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// Set ProcessGroupID for child process as init process. Both will be under same process group
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// Start the child (a fork of ourselves) without waiting for termination
	slog.Info("INIT: starting child process")
	if err := cmd.Start(); err != nil {
		slog.Error("INIT: failed to start child process", "error", err)
		return
	}

	// We need notification of all relevant signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	done := make(chan bool, 1)

	// Forward all signals to the child in a goroutine
	go func() {
		for sig := range sigs {
			// Forward the signal to the child process
			slog.Info("INIT: forwarding signal to child process", "signal", sig, "PID", cmd.Process.Pid)
			err := cmd.Process.Signal(sig)
			if err != nil {
				slog.Error("INIT: failed to forward signal to child process", "signal", sig, "PID", cmd.Process.Pid, "error", err)
			}

			// If the signal was SIGTERM or SIGINT, wait 10 seconds for the child to terminate and send a KILL signal
			if sig == syscall.SIGTERM || sig == syscall.SIGINT {
				go func() {
					// Wait 10 seconds for the child process to finish
					time.Sleep(10 * time.Second)
					// Kill the child immediately
					_ = cmd.Process.Kill()
				}()

				slog.Info("INIT: using DONE channel to terminate init process")
				done <- true
			}
		}
	}()

	// Enter in a goroutine an infinite loop reaping periodically the zombie children
	go func() {
		for {
			var ws syscall.WaitStatus
			pid, err := syscall.Wait4(-1, &ws, syscall.WNOHANG, nil)
			if pid <= 0 || err != nil {
				time.Sleep(5 * time.Second)
			} else {
				slog.Info("INIT: reaped zombie child with PID", "PID", pid)
			}
		}
	}()

	slog.Info("INIT: awaiting signal to terminate init process")
	<-done

	// Wait for the child process to finish and release its resources
	slog.Info("INIT: waiting for child process to finish")
	_, _ = cmd.Process.Wait()

	slog.Info("INIT: exiting init process")
}
