package carteira

import (
	"testing"
)

func confirmaErro(t *testing.T, valor error, esperado string) {
	t.Helper()

	if valor == nil {
		t.Fatal("An Error was expected")
	}

	resultado := valor.Error()

	if resultado != esperado {
		t.Errorf("Result %s, Expected %s", resultado, esperado)
	}
}

func confirmaSaldo(t *testing.T, carteira Carteira, esperado Bitcoin){
	t.Helper()

	if carteira.Saldo() != esperado {
		t.Errorf("Result: %s, Expected: %s", carteira.Saldo(), esperado)
	}
}

func TestCarteira(t *testing.T) {

	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}

		carteira.Depositar(Bitcoin(10))
		resultado := carteira.Saldo()

		esperado := Bitcoin(10)

		if resultado != esperado {
			t.Errorf("Resultado %s, Esperado %s", resultado, esperado)
		}
	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}

		carteira.Retirar(Bitcoin(10))

		resultado := carteira.Saldo()
		esperado := Bitcoin(10)

		if resultado != esperado {
			t.Errorf("Resultado %s, Esperado %s", resultado, esperado)
		}
	})

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(30)
		carteira := Carteira{saldoInicial}

		erro := carteira.Retirar(50)

		confirmaSaldo(t, carteira, saldoInicial)
		confirmaErro(t, erro, ErroSaldoInsuficiente.Error())

	})
}
