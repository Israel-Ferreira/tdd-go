package inteiros

import (
	"testing"
)

func TestAdicionador(t *testing.T){
	soma :=   Adicionador(2,2)
	resultado := 4

	if soma != resultado {
		t.Errorf("Esperado %d, resultado %d", soma, resultado)
	}

}