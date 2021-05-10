package carteira

import (
	"errors"
	"fmt"
)

var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Carteira struct {
	saldo Bitcoin
}

func (c *Carteira) Depositar(quantidade Bitcoin) {
	fmt.Printf("O Endereço do saldo no Depositar é %v \n", &c.saldo)
	c.saldo += quantidade
}

func (c *Carteira) Retirar(quantidade Bitcoin) error {
	
	if c.saldo < quantidade {
		return ErroSaldoInsuficiente
	}

	c.saldo -= quantidade
	return nil
}

func (c *Carteira) Saldo() Bitcoin {
	fmt.Printf("O Endereço do saldo no Saldo é %v \n", &c.saldo)
	return c.saldo
}
