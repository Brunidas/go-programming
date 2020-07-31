package palabra

import (
	"fmt"
	"testing"

	"gitlab.com/go-programming/010-Ninja-nivel13/02/cita"
)

func TestConteoUso(t *testing.T) {
	m := ConteoUso("Uno Dos Tres Tres Tres")
	for k, v := range m {
		switch k {
		case "Uno":
			if v != 1 {
				t.Error("Esperaba", 1, "Obtuvo", v)
			}
		case "Dos":
			if v != 1 {
				t.Error("Esperaba", 1, "Obtuvo", v)
			}
		case "Tres":
			if v != 3 {
				t.Error("Esperaba", 3, "Obtuvo", v)
			}
		}
	}
}

func TestConteo(t *testing.T) {
	n := Conteo(cita.Cuando)
	if n != 355 {
		t.Error("Esperaba", 355, "Obtuvo", n)
	}
}

func BenchmarkConteo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Conteo(cita.Cuando)
	}
}
func ExampleConteo() {
	fmt.Println(Conteo("Uno Dos Tres"))
	// Output:
	// 3
}

func BenchmarkConteoUso(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConteoUso(cita.Cuando)
	}
}
