package components

import (
	"strconv"

	"github.com/gdamore/tcell"
)

type Cell struct {
	x          int
	y          int
	val        []rune
	background tcell.Color
}

func (c *Cell) Render(s tcell.Screen) error {
	// gl := ' '
	// st := tcell.StyleDefault
	// st = st.Background(c.background)
	//
	// if !(s.Colors() > 256 && s.Colors() > 1) {
	// 	gl = '@'
	// }

	return nil
}

func NewCell(x, y, val int) *Cell {
	strV := strconv.Itoa(val)
	chars := make([]rune, len(strV))

	for _, r := range strV {
		chars = append(chars, r)
	}

	return &Cell{
		x:          x,
		y:          y,
		val:        chars,
		background: tcell.NewHexColor(int32(val & 0xffffff)),
	}
}
