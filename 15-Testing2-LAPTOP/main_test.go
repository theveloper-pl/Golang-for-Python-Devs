package main

import (
	"testing"
)


// Table driven tests

// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out

func TestDoSomething(t *testing.T) {
	cases := []struct {
		a     	 int
		b    	 int
		result 	 int
	}{
		{
			a:      1,
			b:      2,
			result: 3,
		},
		{
			a:      2,
			b:      3,
			result: 5,
		},
		{
			a:      5,
			b:      5,
			result: 10,
		},
	}

	for _, c := range cases {
		result := doSomething(c.a, c.b)
		if result != c.result {
			t.Errorf("Expected %d, got %d", c.result, result)
		} else {
			t.Logf("Testing: %d + %d = %d", c.a, c.b, c.result) 
		}
	}

}


// func TestCleanup(t *testing.T) {
// 	t.Cleanup(func() {t.Log("Cleanup")},)
// }




// go test -bench=.
// output means that the loop ran 10,000,000 times at a speed of 1.07 nanosecond per loop.
func BenchmarkAdd(b *testing.B){
    for i :=0; i < b.N ; i++{
        doSomething(4, 6)
    }
}

