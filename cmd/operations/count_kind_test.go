package operations

import (
	"fmt"
	"testing"
	"time"
)

func TestCountKind(t *testing.T) {
	// Used for profiling
	start := time.Now()
	countKind("../../out/wp", []string{"", "", "variable_name"}, true, true)
	duration := time.Since(start)
	fmt.Printf("Execution time: %v\n", duration)
}
