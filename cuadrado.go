package geometric

type Cuadrado struct {
	Altura float64
	Ancho  float64
}

func (cua *Cuadrado) calculo_area() float64 {
	return cua.Altura * cua.Ancho
}
func (cua *Cuadrado) calculo_perimetro() float64 {
	return 2*cua.Altura + 2*cua.Ancho
}
