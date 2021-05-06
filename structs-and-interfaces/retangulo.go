package	formas

type Retangulo struct {
	Largura, Altura float64
}

func (r Retangulo) Area() float64{
	return r.Altura * r.Largura
}


func (r Retangulo) Perimetro() float64 {
	return 2 * (r.Altura + r.Largura)
}

