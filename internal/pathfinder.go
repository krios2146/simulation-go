package internal

type PathFinder interface {
	FindPath(m Map, c Creature, target Entity) []Coordinates
}

type BFSPathFinder struct{}

func (bfspf BFSPathFinder) FindPath(m Map, c Creature, target Entity) []Coordinates {
	panic("BFSPathFinder not implemented")
}
