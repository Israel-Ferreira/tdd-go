package main

import (
	"bytes"
	"testing"
)

func TestCumprimento(t *testing.T) {
	buffer := bytes.Buffer{}

	Cumprimenta(&buffer, "Israel")

	resultado := buffer.String()
	esperado :=  "Olá, Israel \n"

	if resultado != esperado {
		t.Errorf("Resultado: %s, Esperado: %s", resultado, esperado)
	}
}
