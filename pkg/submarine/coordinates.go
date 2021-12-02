package submarine

type Coordinates struct {
	Position int
	Depth    int
	Aim      int
}

func BuildCoordinates(position int, depth int, aim int) Coordinates {
	return Coordinates{Position: position, Depth: depth, Aim: aim}
}

func ZeroCoordinates() Coordinates {
	return Coordinates{}
}
