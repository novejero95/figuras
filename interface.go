package figuras

import "fmt"

type Forma interface {
	perimetro() float64
	area() float64
}

func CalcularArea(forma Forma) {
	fmt.Println(forma.area())
}

func CalcularPerimetro(forma Forma) {
	fmt.Println(forma.perimetro())
}
