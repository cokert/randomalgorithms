package board

var (
	Size = 4
)

// Board two dimensional array of bool (visited)
type Board [][]bool

// NewBoard returns a new Board with the provided location marked as visited
func NewBoard(s int, l Location) Board {
	Size = s
	b := make([][]bool, Size)
	for i := range b {
		b[i] = make([]bool, Size)
	}
	b[l.X][l.Y] = true
	return b
}

// PossibleMoves returns a slice of Locations
func (b Board) PossibleMoves(current Location) []Location {
	allMoves := getMoves(current)
	p := []Location{}
	for _, m := range allMoves {
		if m.IsValid() && b[m.X][m.Y] == false {
			p = append(p, m)
		}
	}
	return p
}

// AllVisited returns true if all locations have been visited
func (b Board) AllVisited() bool {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if !b[i][j] {
				return false
			}
		}
	}
	return true
}

//Clone returns a copy of the board.
func (b Board) Clone() Board {
	nb := make(Board, len(b))
	for i := range nb {
		nb[i] = make([]bool, len(b))
		for j := range nb[i] {
			nb[i][j] = b[i][j]
		}
	}
	return nb
}

// Visit marks a board location as visited
func (b Board) Visit(l Location) Board {
	b[l.X][l.Y] = true
	return b
}

func getMoves(current Location) []Location {
	m := []Location{}
	m = append(m, Location{X: current.X - 1, Y: current.Y + 2}) //NNW x-1 y+2
	m = append(m, Location{X: current.X + 1, Y: current.Y + 2}) //NNE x+1 y+2
	m = append(m, Location{X: current.X + 2, Y: current.Y + 1}) //ENE x+2 y+1
	m = append(m, Location{X: current.X + 2, Y: current.Y - 1}) //ESE x+2 y-1
	m = append(m, Location{X: current.X + 1, Y: current.Y - 2}) //SSE x+1 y-2
	m = append(m, Location{X: current.X - 1, Y: current.Y - 2}) //SSW x-1 y-2
	m = append(m, Location{X: current.X - 2, Y: current.Y - 1}) //WSW x-2 y-1
	m = append(m, Location{X: current.X - 2, Y: current.Y + 1}) //WNW x-2 y+1

	return m
}
