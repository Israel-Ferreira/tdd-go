package formas

type Triangulo struct {
	Largura, Altura float64
}

func (t Triangulo) Area() float64 {
	return (t.Altura * t.Largura) / 2
}
