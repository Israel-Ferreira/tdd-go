package main

import (
	"testing"
)

func TestOlaWithName(t *testing.T){
	resultado :=  OlaWithName("Chris")
	esperado := "Olá, Chris"

	if(resultado != esperado){
		t.Errorf("Resultado %s, Esperado %s", resultado, esperado)
	}
}


func TestOla(t *testing.T) {
	resultado := Ola()
	esperado := "Olá, Mundo"

	if resultado != esperado {
		t.Errorf("Resultado %s, Esperado %s", resultado, esperado)
	}
}
