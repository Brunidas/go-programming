package decir

import (
	"fmt"
	"testing"
)

func TestSaludar(t *testing.T) {
	s := Saludar("Bruno")
	if s != "Bienvenido querido Bruno" {
		t.Error("Expected", "Bienvenido querido Bruno", "Got", s)
	}
}

func ExampleSaludar() {
	fmt.Println(Saludar("Bruno"))
	//Output:
	// Bienvenido querido Bruno
}

func BenchmarkSaludar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Saludar("Bruno")
	}
}
