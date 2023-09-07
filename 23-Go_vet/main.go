package main

import (
	"fmt"
	"sync"
)



/*

go vet to narzędzie w ekosystemie Go, które analizuje kod źródłowy w poszukiwaniu podejrzanych konstrukcji.
Działa to podobnie do lintera, ale zamiast koncentrować się na stylu kodu, go vet skupia się na konstrukcjach


Tak niepoprawny kod uda nam się skompilować i uruchomić, przez co nie zauważymy błędu

*/





type Data struct {
    mu    sync.Mutex
    value int
}

func SetValue(d Data, v int) {
    d.mu.Lock()
    d.value = v
    d.mu.Unlock()
}


type Person struct {
    Name string "json:name"      // Poprawnie by było :       Name string `json:"name"`
}



func main() {
    name := "Alice"
    fmt.Printf("Hello, %s, welcome to the world of Go!\n", name, 42)







	

}
