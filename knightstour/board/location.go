package board

import "fmt"

// Location is the X and Y coordinates of a location on the board
type Location struct {
	X int
	Y int
}

var (
	xaxisLabels = map[int]string{
		0: "H",
		1: "G",
		2: "F",
		3: "E",
		4: "D",
		5: "C",
		6: "B",
		7: "A",
	}
	yaxisLabels = map[int]string{
		0: "8",
		1: "7",
		2: "6",
		3: "5",
		4: "4",
		5: "3",
		6: "2",
		7: "1",
	}
)

func (l Location) String() string {
	x := xaxisLabels[l.X]
	y := yaxisLabels[l.Y]
	if x == "" {
		x = "?"
	}
	if y == "" {
		y = "?"
	}
	return fmt.Sprintf("%v%v", x, y)
}

// IsValid returns a boolean indicating the move is off the board
func (l Location) IsValid() bool {
	if l.X < 0 || l.X > Size-1 || l.Y < 0 || l.Y > Size-1 {
		return false
	}
	return true
}
