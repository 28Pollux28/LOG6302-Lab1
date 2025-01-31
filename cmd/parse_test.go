package cmd

import (
	"os"
	"path/filepath"
	"sync"
	"testing"
)

// Used for profiling
func TestParseDir(t *testing.T) {

	tempDir := t.TempDir()

	dataDir := os.Getenv("DATA")
	if dataDir == "" {
		dataDir = "../data"
	}

	outputDirPath := filepath.Join(tempDir, "output")

	var wg sync.WaitGroup
	parseDir(dataDir, outputDirPath, true, false, &wg)
	wg.Wait()
}
