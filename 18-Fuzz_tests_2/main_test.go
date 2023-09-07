package main

import (
	"fmt"
	"testing"
	"os"
)

var file *os.File

func toHumanReadable(b []byte) string {
    readable := make([]rune, 0, len(b))
    for _, byteVal := range b {
        if byteVal >= 32 && byteVal <= 126 { // Printable ASCII range
            readable = append(readable, rune(byteVal))
        } else {
            readable = append(readable, []rune(fmt.Sprintf("\\x%02x", byteVal))...)
        }
    }
    return string(readable)
}



func init() {
	var err error
	file, err = os.Create("/home/ubuntu/tutorial/18-Fuzz_tests_2/fuzz_log.txt") // Open file for writing, will truncate existing content
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file: %v", err))
	}
}

func FuzzDoSomething(f *testing.F) {

    f.Add("exampleA1", "exampleB1")
    f.Add("exampleA2", "exampleB2")
	defer file.Close() // Ensure the file is closed when the fuzzing is done

	f.Fuzz(func(t *testing.T, a string, b string) {
		
        _, err := fmt.Fprintf(file, "Testing with values a=%s, b=%s\n", toHumanReadable([]byte(a)), toHumanReadable([]byte(b)))
		if err != nil {
			t.Fatalf("Failed to write to log file: %v", err)
		}

		result := doSomething(a, b)
		expected := fmt.Sprint(a, b)
		if result != expected {
			t.Fatalf("For inputs (%s, %s), expected %s but got %s", a, b, expected, result)
		}
	})
}
