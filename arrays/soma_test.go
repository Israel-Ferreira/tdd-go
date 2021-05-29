package arrays

import (
	"testing"
)

func assertResult(t *testing.T, result, expected int) {
	t.Helper()

	if result != expected {
		t.Errorf("Resultado: %d, Esperado: %d", result, expected)
	}
}

func TestSoma(t *testing.T) {

	t.Run("coleção de 5 números", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}

		resultado := Soma(numeros)
		esperado := 15

		assertResult(t, resultado, esperado)
	})

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5, 6, 7}
		resultado := Soma(numeros)
		esperado := 28

		assertResult(t, resultado, esperado)
	})

}
