package maps

import (
	"testing"
)



func assertResult(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("Resultado: %s, Esperado: %s", result, expected)
	}
}


func compararErro(t *testing.T, result, expected error){
	t.Helper()

	if result != expected {
		
	}
}

func TestBusca(t *testing.T) {
	dicionario := Dicionario{"teste": "Isso é apenas um teste"}

	t.Run("Palavra Conhecida", func(t *testing.T) {
		resultado, _:= dicionario.Busca("teste")
		esperado := "Isso é apenas um teste"

		assertResult(t, resultado, esperado)
	})

	t.Run("Palavra desconhecida", func(t *testing.T) {
		_, err := dicionario.Busca("desconhecida")


		if err == nil {
			t.Fatal("É esperado que um erro seja obtido")
		}
	})

}
