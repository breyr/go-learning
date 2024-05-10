package structs

type Reactangle struct {
	Width  float64
	Height float64
}

func Perimeter(rect Reactangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

func Area(rect Reactangle) float64 {
	return rect.Width * rect.Height
}
