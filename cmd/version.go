package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	Version   = "1.0.0"
	BuildDate = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of tmfgo",
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("tmfgo version %s (%s/%s) built on %s\n",
			Version, runtime.GOOS, runtime.GOARCH, BuildDate)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
