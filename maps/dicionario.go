package maps

import "errors"

type Dicionario map[string]string


var ErrNaoEncontrado error = errors.New("Não foi possível achar a palavra que você procura")


func (d Dicionario) Busca(term string) (string, error){
	definition, exists := d[term]

	if !exists {
		return "", ErrNaoEncontrado
	}

	return definition, nil
}

