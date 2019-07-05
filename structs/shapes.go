package structs

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(x, y float64) float64 {
	return 2 * (x + y)
}

func Area(x, y float64) float64 {
	return x * y
}
