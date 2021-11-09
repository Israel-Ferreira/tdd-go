package concorrencia

import "time"


type resultado struct { 
	string
	bool
}

type VerificadorWebsite func(string) bool

func VerificaWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
	resultados := make(map[string]bool)

	canalResultado := make(chan resultado)


	for _, url := range urls {
		go func(urlStr string){
			canalResultado <- resultado{urlStr, vw(urlStr)}
		}(url)
		
	}


	for i := 0; i < len(urls); i++ {
		resultado := <-canalResultado
		resultados[resultado.string] = resultado.bool
	}



	time.Sleep(2 * time.Second)

	return resultados
}
