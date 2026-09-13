package pdp

import (
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func TestPDPConcurrentAccess(t *testing.T) {
	// Create a temporary Starlark policy file
	tmpDir := t.TempDir()
	policyFile := filepath.Join(tmpDir, "test_policy.star")
	content := `
def authorize():
    return True
`
	if err := os.WriteFile(policyFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write policy file: %v", err)
	}

	pdp, err := NewPDPService(&Config{
		PolicyFileName: policyFile,
	})
	if err != nil {
		t.Fatalf("failed to create PDP service: %v", err)
	}

	const goroutines = 20
	const iterations = 500

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for range goroutines {
		go func() {
			defer wg.Done()
			for range iterations {
				input := StarTMFMap{
					"key": "value",
				}
				authorized, err := pdp.Authorize(input)
				if err != nil {
					t.Errorf("Authorize error: %v", err)
				}
				if !authorized {
					t.Errorf("expected authorized to be true")
				}
			}
		}()
	}

	wg.Wait()

	if count := pdp.threadPoolCounter.Load(); count == 0 {
		t.Errorf("expected threadPoolCounter > 0, got %d", count)
	}
}
