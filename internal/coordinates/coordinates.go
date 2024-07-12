package coordinates

import "fmt"

type Coordinates struct {
	X, Y uint8
}

func (c *Coordinates) String() string {
	return fmt.Sprintf("x: %v, y: %v", c.X, c.Y)
}
