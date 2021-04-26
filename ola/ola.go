package main

import (
	"fmt"
)

func OlaWithName(name string) string {
	return "Olá, " + name
}

func Ola() string {
	return "Olá, Mundo"
}

func main() {
	fmt.Println(Ola())
}
