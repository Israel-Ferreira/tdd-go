package concorrencia

import (
	"reflect"
	"testing"
	"time"
)

func mockVerficadorSite(url string) bool {
	return url != "tthp://wubbaluba.cry"
}

func slowStubVerificadorSite(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestVerificaWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"tthp://wubbaluba.cry",
	}

	expected := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"tthp://wubbaluba.cry":       false,
	}

	resultado := VerificaWebsites(mockVerficadorSite, websites)

	if !reflect.DeepEqual(expected, resultado) {
		t.Errorf("Não deu certo  : - (")
	}
}

func BenchmarkVerificaWebsite(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "xnd:mdb:jr.wannaluba.cry"
	}

	for i := 0; i < b.N; i++ {
		VerificaWebsites(slowStubVerificadorSite,urls)
	}
}
