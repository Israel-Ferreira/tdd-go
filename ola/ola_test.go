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
		resultado := Ola("Chris", "")
		esperado := "Olá, Chris"

		mostrarErros(t, resultado, esperado)
	})

	t.Run("diz 'Olá, Mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, Mundo"

		mostrarErros(t, resultado, esperado)
	})

	t.Run("Em espanhol", func(t *testing.T) {
		resultado := Ola("", "Espanhol")
		esperado := "Hola, Mundo"

		mostrarErros(t, resultado, esperado)
	})

	t.Run("Em espanhol com nome", func(t *testing.T) {
		resultado := Ola("Eloide", "Espanhol")
		esperado := "Hola, Eloide"

		mostrarErros(t, resultado, esperado)
	})

	t.Run("Em Francês", func(t *testing.T) {
		resultado := Ola("", "Francês")
		esperado := "Bonjour, Mundo"

		mostrarErros(t, resultado, esperado)
	})

	t.Run("Em Francês com nome", func(t *testing.T) {
		resultado := Ola("Jean", "Francês")
		esperado := "Bonjour, Jean"

		mostrarErros(t, resultado, esperado)
	})


	t.Run("Em Inglês", func(t *testing.T) {
		resultado := Ola("", "Inglês")
		esperado  := "Hello, Mundo"

		mostrarErros(t, resultado, esperado)
	})

	t.Run("Em Inglês com nome", func(t *testing.T){
		resultado := Ola("Claire", "Inglês")
		esperado := "Hello, Claire"

		mostrarErros(t, resultado, esperado)
	})


	t.Run("Em Italiano", func(t *testing.T){
		resultado := Ola("", "Italiano")
		esperado := "Ciao, Mundo"

		mostrarErros(t, resultado, esperado)
	})


	t.Run("Em Italiano com nome", func(t *testing.T){
		resultado := Ola("Bianca", "Italiano")
		esperado := "Ciao, Bianca"

		mostrarErros(t, resultado, esperado)
	})

}
