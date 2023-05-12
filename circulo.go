package geometric

import (
	"math"
)

type Circulo struct{ Radio float64 }

func (cir *Circulo) calculo_area() float64 {
	return math.Pi * (math.Pow(cir.Radio, 2))
}
func (cir *Circulo) calculo_perimetro() float64 {
	return 2 * math.Pi * cir.Radio
}
