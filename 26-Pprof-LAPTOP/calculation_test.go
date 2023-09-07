package calculation

import (
	"testing"
)

/*




pprof jest narzędziem do profilowania aplikacji napisanych w języku Go.
 Dzięki profilowaniu możemy zbierać dane na temat wydajności naszej aplikacji, takie jak zużycie procesora, alokację pamięci czy blokady wątków. 
 Po zebraniu tych danych możemy je wizualizować i analizować za pomocą różnych narzędzi, takich jak narzędzie pprof wbudowane w Go.

Graphviz to zestaw narzędzi do wizualizacji grafów. Jest używany przez pprof do generowania graficznych reprezentacji danych profilowania, takich jak grafy wywołań funkcji. 
Aby korzystać z tej funkcji w pprof, musimy mieć zainstalowany Graphviz.

Aby przeprowadzić profilowanie CPU i wygenerować wyniki:
1. Wykonaj test i zbierz dane o profilowaniu CPU: go test -run TestCalculation -cpuprofile cpu.prof -bench .
2. Analizuj zebrane dane za pomocą pprof: go tool pprof cpu.prof
3. Użyj polecenia 'web' w pprof, aby wygenerować wizualizację (potrzebujesz Graphviz): web


Więcej komend można dostać pod ,,go help,,


*/

// TestCalculation testuje funkcję Calulaction.
func TestCalculation(t *testing.T) {
	result := Calulaction(2, 3)
	// Oczekujemy wyniku 26 dla danych wejściowych 2 i 3. Jeśli wynik się różni, zgłaszamy błąd.
	if result != 26 {
		t.Error("Expected 26, got ", result)
	}
}

// BenchmarkCalculation mierzy wydajność funkcji Calulaction.
func BenchmarkCalculation(b *testing.B) {
	// b.N to liczba iteracji ustalana przez framework benchmarkingowy.
	// Funkcja Calulaction zostanie wykonana b.N razy, aby zmierzyć jej wydajność.
	for i := 0; i < b.N; i++ {
		Calulaction(2, 3)
	}
}
