package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleSoma() {
	numeros := []int{10, 15, 5, 18, 2}
	soma := Soma(numeros)

	fmt.Println(soma)
	// Output: 50
}

func assertTest(t *testing.T, result int, expect int, dados []int) {

	t.Helper()

	if result != expect {
		t.Errorf("Result %d, Expected %d, Given %v", result, expect, dados)
	}
}

func TestGenerateSlice(t *testing.T) {
	assertResults := func(t *testing.T, result, expected []int, given int) {
		t.Helper()

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Result %v, Expected: %v, Given: %d", result, expected, given)
		}
	}

	t.Run("Deve gerar um slice vazio ao passar o numero 0", func(t *testing.T) {
		limite := 0

		resultado := GenerateSlice(limite)
		esperado := []int{}

		assertResults(t, resultado, esperado, limite)
	})

	t.Run("Deve Gerar um Slice de 10 numeros", func(t *testing.T) {
		limite := 10

		resultado := GenerateSlice(limite)
		esperado :=  []int{1,2,3,4,5,6,7,8,9,10}

		assertResults(t, resultado, esperado, limite)
	})
}

func TestSoma(t *testing.T) {
	t.Run("Soma de uma coleção de 5 numeros", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}
		resultado := Soma(numeros)
		esperado := 15

		assertTest(t, resultado, esperado, numeros)
	})

	t.Run("Soma de uma coleção de três numeros", func(t *testing.T) {
		numeros := []int{10, 20, 30}
		resultado := Soma(numeros)
		esperado := 60

		assertTest(t, resultado, esperado, numeros)
	})

	t.Run("A Soma de uma coleção vazia deve ser 0", func(t *testing.T) {
		numeros := []int{}
		resultado := Soma(numeros)
		esperado := 0

		assertTest(t, resultado, esperado, numeros)
	})

}

func BenchmarkSoma(b *testing.B) {
	b.Run("Soma de uma coleção de 100 numeros", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			numeros := GenerateSlice(100)
			Soma(numeros)
		}
	})

	b.Run("Soma de uma coleção de 1000 numeros", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			numeros := GenerateSlice(1000)
			Soma(numeros)
		}
	})
}
