package model

type Position struct {
	X int
	Y int
}

func NewPosition(x int, y int) *Position {
	return &Position{X: x, Y: y}
}

func (p *Position) GetX() int {
	return p.X
}

func (p *Position) SetX(x int) {
	p.X = x
}

func (p *Position) GetY() int {
	return p.Y
}

func (p *Position) SetY(y int) {
	p.Y = y
}
