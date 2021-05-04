package iteracao

import (
	"fmt"
	"testing"
)

func assertTest(t *testing.T, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("Result %s, Expected %s", result, expected)
	}
}


func ExampleRepetidor(){
	repeticao := Repetidor("a")
	fmt.Println(repeticao)
	// Output 'aaaaa'
}

func TestRepetidor(t *testing.T) {
	repeticoes := Repetidor("a")
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("Resultado %s, Esperado %s", repeticoes, esperado)
	}
}

func TestRepetidorWithLimit(t *testing.T) {
	repeticaoComLimite := RepetidorWithLimit("a", 10)
	esperado := "aaaaaaaaaa"

	assertTest(t, repeticaoComLimite, esperado)
}

func TestStringRepeater(t *testing.T) {
	t.Run("Repetir 'lambda' três vezes", func(t *testing.T){
		repeticao := StringRepeater("Lambda",3)
		esperado := "LambdaLambdaLambda"

		assertTest(t, repeticao, esperado)
	})

	t.Run("Não Repetir se a quantidade for zero", func(t *testing.T){
		resultado := StringRepeater("Lambda",0)
		esperado := ""

		assertTest(t, resultado, esperado)
	})
}

func BenchmarkRepetidor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetidor("a")
	}
}
