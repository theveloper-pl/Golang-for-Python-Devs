package main

import (
	"testing"
)


func TestSomethingParallel(t *testing.T) {
	t.Run("Test A in Parallel", func(t *testing.T) {
		t.Parallel()
		result := doSomethingVeryLong(2, 5)
		if result != 7 {
			t.Errorf("Failed test A. Expected 7, got %d", result)
		}
	})

	t.Run("Test B in Parallel", func(t *testing.T) {
		t.Parallel()
		result := doSomethingVeryLong(54, 4)
		if result != 58 {
			t.Errorf("Failed test B. Expected 58, got %d", result)
		}
	})
}


