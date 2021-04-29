package iteracao

import (
	"strings"
)

const DEFAULT_TIMES_REPEAT = 5

func Repetidor(text string) string {
	var newText string

	for i := 0; i < DEFAULT_TIMES_REPEAT; i++ {
		newText += text
	}

	return newText
}

func RepetidorWithLimit(text string, limit int) string {
	return strings.Repeat(text,limit)
}

