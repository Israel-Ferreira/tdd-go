package main

import (
	"fmt"
)


func Ola(name, idioma string) string {

	prefix := greetingsPrefix(idioma)

	if name != "" {
		return prefix + name
	}

	return prefix + "Mundo"
}


func greetingsPrefix(language string) (prefix string) {
	switch language {
	case "Inglês":
		prefix = "Hello, "
	case "Espanhol":
		prefix = "Hola, "
	case "Francês":
		prefix = "Bonjour, "
	case "Italiano":
		prefix = "Ciao, "
	default:
		prefix = "Olá, "
	}

	return	
}


func main() {
	fmt.Println(Ola("", ""))
}
