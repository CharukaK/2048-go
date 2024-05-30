package components

import (
	"strconv"

	"github.com/gdamore/tcell"
)


type Cell struct {
	x          int
	y          int
	blockH     int
	blockW     int
	val        []rune
	background tcell.Color
}

func (c *Cell) Render(s tcell.Screen) {
	gl := ' '
	st := tcell.StyleDefault
	st = st.Background(tcell.Color100)

	for row := 0; row <= c.blockH; row++ {
		for col := 0; col <= c.blockW; col++ {
			s.SetCell(c.x+col, c.y+row, st, gl)
		}
	}
}

func NewCell(x, y, cellH, cellW, val int) *Cell {
	strV := strconv.Itoa(val)
	chars := make([]rune, len(strV))

	for _, r := range strV {
		chars = append(chars, r)
	}

	return &Cell{
		x:          x,
		y:          y,
		val:        chars,
		blockH:     cellH,
		blockW:     cellW,
		background: tcell.NewHexColor(int32(val & 0xffffff)),
	}
}
