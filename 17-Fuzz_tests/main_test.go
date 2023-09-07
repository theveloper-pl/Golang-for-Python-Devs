package main

import (
	"fmt"
	"testing"
	"os"
)

var file *os.File

func init() {
	var err error
	file, err = os.Create("/home/ubuntu/tutorial/17-Fuzz_tests/fuzz_log.txt") // Open file for writing, will truncate existing content
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file: %v", err))
	}
}

func FuzzDoSomething(f *testing.F) {
	defer file.Close() // Ensure the file is closed when the fuzzing is done

	f.Fuzz(func(t *testing.T, a int, b int) {
		
		_, err := fmt.Fprintf(file, "Testing with values a=%d, b=%d\n", a, b) // Write to the file
		if err != nil {
			t.Fatalf("Failed to write to log file: %v", err)
		}

		result := doSomething(a, b)
		expected := a + b
		if result != expected {
			t.Fatalf("For inputs (%d, %d), expected %d but got %d", a, b, expected, result)
		}
	})
}
