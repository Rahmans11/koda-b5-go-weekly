package geometristemplate

type IGeometric interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	Area := 3.14 * c.Radius * c.Radius
	return Area
}

type Rectangle struct {
	Tinggi float64
	Lebar  float64
}

func (r Rectangle) Area() float64 {
	Area := r.Tinggi * r.Lebar
	return Area
}

func CalculateArea(geometric []IGeometric) float64 {
	var totalArea float64
	for _, v := range geometric {
		totalArea += v.Area()
	}
	return totalArea
}
