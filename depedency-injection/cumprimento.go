package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Cumprimenta(writer io.Writer, nome string) {
	fmt.Fprintf(writer, "Olá, %s \n", nome)
}

func HandlerMeuCumprimento(w http.ResponseWriter, r *http.Request) {
	Cumprimenta(w, "mundo")
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(HandlerMeuCumprimento))

	if err != nil {
		log.Fatalln("Erro")
	}
}
