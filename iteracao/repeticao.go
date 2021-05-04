package iteracao


import (
	"strings"
)

const quantidadeRepeticoes = 5

func Repetidor(text string) string {

	var repeticoes string

	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticoes += text
	}

	return repeticoes
}

func RepetidorWithLimit(text string, limit int) (repeticoes string) {
	for i := 0; i < limit; i++ {
		repeticoes += text
	}

	return
}

func StringRepeater(text string, limit int) string {
	return strings.Repeat(text, limit)
}
