package internal

import (
	"fmt"
	"strings"
)

const columnWidth = 2
const errorSymbol = "❌"

var emptySpace = strings.Repeat(" ", columnWidth)

func Render(m Map) {
	fmt.Println("┌" + strings.Repeat("─", int(m.Width)*columnWidth) + "┐")

	for i := 0; i < int(m.Height); i++ {
		sb := strings.Builder{}

		for j := 0; j < int(m.Width); j++ {
			coordinates := Coordinates{X: uint8(j), Y: uint8(i)}
			entity, found := m.Find(coordinates)

			if !found {
				sb.WriteString(emptySpace)
				continue
			}

			if entity, ok := entity.(ConsolePrintable); ok {
				sb.WriteRune(entity.GetConsoleSprite())
			} else {
				sb.WriteString(errorSymbol)
			}
		}
		fmt.Println("│" + sb.String() + "│")
	}

	fmt.Println("└" + strings.Repeat("─", int(m.Width)*columnWidth) + "┘")
}
