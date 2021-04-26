package main

import (
	"fmt"
)

const PREFIXO_OLA_PORTUGUES = "Olá, "

func Ola(name string) string {

	if name != "" {
		return PREFIXO_OLA_PORTUGUES + name
	}



	return PREFIXO_OLA_PORTUGUES + "Mundo"
}

func main() {
	fmt.Println(Ola(""))
}
