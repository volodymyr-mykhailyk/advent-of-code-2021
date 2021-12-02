package submarine

type Coordinates struct {
	Position int
	Depth    int
}

func BuildCoordinates(position int, depth int) Coordinates {
	return Coordinates{Position: position, Depth: depth}
}
