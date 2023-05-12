package geometric

import "fmt"

type Figura interface {
	calculo_area() float64
	calculo_perimetro() float64
}

//Metodo con M mayuscula significa publico
func Medidas(f Figura) {
	fmt.Println("Medidas:", f)
	fmt.Println("Area:", f.calculo_area())
	fmt.Println("Perimetro:", f.calculo_perimetro())
}
