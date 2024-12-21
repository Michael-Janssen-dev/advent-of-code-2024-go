package core

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{
		x, y,
	}
}

func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

func (p Point) Subtract(other Point) Point {
	return Point{p.X - other.X, p.Y - other.Y}
}

func (p Point) North() Point {
	return Point{p.X, p.Y - 1}
}

func (p Point) East() Point {
	return Point{p.X + 1, p.Y}
}

func (p Point) South() Point {
	return Point{p.X, p.Y + 1}
}

func (p Point) West() Point {
	return Point{p.X - 1, p.Y}
}

var NORTH = Point{0, -1}
var EAST = Point{1, 0}
var SOUTH = Point{0, 1}
var WEST = Point{-1, 0}

func (p Point) Cardinal() []Point {
	result := make([]Point, 4)
	result[0] = p.North()
	result[1] = p.East()
	result[2] = p.South()
	result[3] = p.West()
	return result
}

func (p Point) Clockwise() Point {
	return Point{-p.Y, p.X}
}

func (p Point) CounterClockwise() Point {
	return Point{p.Y, -p.X}
}
