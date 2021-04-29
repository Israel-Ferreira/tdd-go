package iteracao

import (
	"testing"
)

func assertTest(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("Result: %s, Expected %s", result, expected)
	}
}

func TestRepetidor(t *testing.T) {
	resultado := Repetidor("a")
	esperado := "aaaaa"

	assertTest(t, resultado, esperado)
}

func TestRepetidorWithLimit(t *testing.T) {
	resultado := RepetidorWithLimit("CIAO", 10)
	esperado := "CIAOCIAOCIAOCIAOCIAOCIAOCIAOCIAOCIAOCIAO"

	assertTest(t, resultado, esperado)
}



func BenchmarkRepetidor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetidor("a")
	}
}

func BenchmarkRepetidorWithLimit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepetidorWithLimit("a", 1200)
	}
}

