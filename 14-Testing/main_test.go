package main

import (
	"testing"
)


/*

Testowane funkcje muszą zaczynać się od słowa Test, a następnie nazwy funkcji, którą testują.
Pliki testowe musza znajdowac się w tym samym folderze co plik testowany lub w podfolderze o nazwie x_test.
Plik testowy musi mieć nazwę x_test.go, gdzie x to nazwa pliku testowanego.


	// t.Log("Something")		// Just print
	// t.Fail() 				// Will show failes test in result
	// t.FailNow() 				// t.Fail + safe stop
	// t.Error("Something")		// t.log() + t.Fail()
	// t.Fatal("Something") 	// t.log() + t.FailNow()
	
}

*/

func TestDoSomething(t *testing.T) {
	t.Log("Running TestDoSomething")       // go test -v
	result := doSomething(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}

}

