package figuras

type Circulo struct {
	Radio float64
	Pi    float64
}

func (ci *Circulo) perimetro() float64 {
	return 2 * ci.Pi * ci.Radio
}

func (ci *Circulo) area() float64 {
	return ci.Pi * (ci.Radio * ci.Radio)
}
