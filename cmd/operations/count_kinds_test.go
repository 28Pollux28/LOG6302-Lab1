package operations

import (
	"fmt"
	"testing"
	"time"
)

func TestCountKinds(t *testing.T) {
	// Used for profiling
	start := time.Now()
	countKinds("../../out/wp", []string{"", "", "variable_name", "name"}, true, true)
	duration := time.Since(start)
	fmt.Printf("Execution time: %v\n", duration)
}
