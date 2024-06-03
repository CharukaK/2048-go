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
	foreground tcell.Color
}

type ColorEntry struct {
	background, foreground tcell.Color
}

var colorMap = map[int]ColorEntry{
	2:    {foreground: tcell.NewHexColor(int32(0x776e65)), background: tcell.NewHexColor(int32(0xeee4da))},
	4:    {foreground: tcell.NewHexColor(int32(0x776e65)), background: tcell.NewHexColor(int32(0xede0c8))},
	8:    {foreground: tcell.NewHexColor(int32(0xf9f6f2)), background: tcell.NewHexColor(int32(0xf2b179))},
	16:   {foreground: tcell.NewHexColor(int32(0xf9f6f2)), background: tcell.NewHexColor(int32(0xf59563))},
	32:   {foreground: tcell.NewHexColor(int32(0xf9f6f2)), background: tcell.NewHexColor(int32(0xf67c5f))},
	64:   {foreground: tcell.NewHexColor(int32(0xf9f6f2)), background: tcell.NewHexColor(int32(0xf65e3b))},
	128:  {foreground: tcell.NewHexColor(int32(0xf9f6f2)), background: tcell.NewHexColor(int32(0xedcf72))},
	256:  {foreground: tcell.NewHexColor(int32(0xf9f6f2)), background: tcell.NewHexColor(int32(0xedcc61))},
	512:  {foreground: tcell.NewHexColor(int32(0xf9f6f2)), background: tcell.NewHexColor(int32(0xedc850))},
	1024: {foreground: tcell.NewHexColor(int32(0xf9f6f2)), background: tcell.NewHexColor(int32(0xedc53f))},
	2048: {foreground: tcell.NewHexColor(int32(0xf9f6f2)), background: tcell.NewHexColor(int32(0xedc22e))},
	// 4096:  {foreground: int32(0xeee4da), background: int32(0x776e65)},
	// 8192:  {foreground: int32(0xeee4da), background: int32(0x776e65)},
	// 16384: {foreground: int32(0xeee4da), background: int32(0x776e65)},
}

func (c *Cell) Render(s tcell.Screen) {
	gl := ' '
	st := tcell.StyleDefault
	st = st.Background(c.background)
	st = st.Foreground(c.foreground)

	for row := 0; row <= c.blockH; row++ {
		ri := len(c.val) - 1
		for col := c.blockW; col >= 0; col-- {
			if row == c.blockH && ri >= 0 {
				gl = c.val[ri]
				ri--
			}
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

	foreground := tcell.ColorBlack
	background := tcell.ColorBlack
	color, ok := colorMap[val]

	if ok {
		foreground = color.foreground
		background = color.background
	}

	return &Cell{
		x:          x,
		y:          y,
		val:        chars,
		blockH:     cellH,
		blockW:     cellW,
		background: background,
		foreground: foreground,
	}
}
