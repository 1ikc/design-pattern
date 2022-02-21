package decorator

type IDraw interface {
	Draw() string
}

type Square struct {}
func (s Square) Draw() string {
	return "Square"
}

type ColorSquare struct {
	s IDraw
	color string
}
func (c ColorSquare) Draw() string {
	return c.color + c.s.Draw()
}

func NewColorSquare(s IDraw, color string) ColorSquare {
	return ColorSquare{
		color: color,
		s: s,
	}
}