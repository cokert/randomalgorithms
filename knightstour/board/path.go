package board

import "fmt"

type Path []Location

func (p Path) Last() Location {
	return p[len(p)-1]
}

func (p *Path) Push(l Location) {
	*p = append(*p, l)
}

func (p Path) String() string {
	s := ""
	for _, l := range p {
		s = fmt.Sprintf("%v %v", s, l)
	}
	return s
}
