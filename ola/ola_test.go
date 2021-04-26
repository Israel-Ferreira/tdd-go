package main

import (
	"testing"
)

func mostrarErros(t *testing.T, resultado, esperado string) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("Resultado %s, Esperado %s", resultado, esperado)
	}
}

func TestOla(t *testing.T) {

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Chris")
		esperado := "Olá, Chris"

		mostrarErros(t, resultado, esperado)
	})

	t.Run("diz 'Olá, Mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Olá, Mundo"

		mostrarErros(t, resultado, esperado)
	})

}
