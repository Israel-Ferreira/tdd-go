package formas

import "testing"

func TestArea(t *testing.T) {

	verificaArea := func(t *testing.T, forma Forma, esperado float64) {
		t.Helper()

		resultado := forma.Area()

		if resultado != esperado {
			t.Errorf("Resultado %.2f, Esperado %.2f", resultado, esperado)
		}
	}

	t.Run("Retangulo", func(t *testing.T) {
		retangulo := Retangulo{12, 10}
		esperado := 120.0

		verificaArea(t, retangulo, esperado)
	})

	t.Run("Area Circulo", func(t *testing.T) {
		circulo := Circulo{10}

		esperado := 314.1592653589793
		verificaArea(t, circulo, esperado)
	})

}

func TestArea2(t *testing.T) {
	testesArea := []struct {
		nome     string
		forma    Forma
		esperado float64
	}{
		{forma: Retangulo{12, 6}, esperado: 72.0, nome: "Retângulo"},
		{forma: Circulo{10}, esperado: 314.1592653589793, nome: "Circulo"},
		{forma: Triangulo{20, 10}, esperado: 100, nome: "Triângulo"},
	}

	for _, tt := range testesArea {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := tt.forma.Area()
			if resultado != tt.esperado {
				t.Errorf("resultado %.2f, esperado %.2f", resultado, tt.esperado)
			}
		})
	}
}
