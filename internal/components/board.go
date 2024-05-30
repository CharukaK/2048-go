package components

import (
	"github.com/CharukaK/2048-go/internal/state"
	"github.com/gdamore/tcell"
)

type Board struct {
	x       int
	y       int
	blockH  int
	blockW  int
	borderC tcell.Color
	size    int
}

func (b *Board) Render(s tcell.Screen) {
	st := tcell.StyleDefault
	st = st.Background(b.borderC)
	gl := ' '

	if !(s.Colors() > 256 && s.Colors() > 1) {
		gl = '⠿'
	}

	fullW := b.size * b.blockW
	fullH := b.size * b.blockH

	for row := 0; row <= fullH; row++ {
		for col := 0; col <= fullW; col++ {
			if col == 0 || col == fullW {
				s.SetCell(b.x+col, b.y+row, st, gl)
			}

			if row == 0 || row == fullH {
				s.SetCell(b.x+col, b.y+row, st, gl)
			}

			if col%b.blockW == 0 || row%b.blockH == 0 {
				s.SetCell(b.x+col, b.y+row, st, gl)

				if !(col == fullW || row == fullH) {
					cell := NewCell(
						b.x+col+1,
						b.y+row+1,
						b.blockH-2,
						b.blockW-2,
						2,
					)

					cell.Render(s)
				}
			}

		}
	}
}

func NewBoard(startX, startY, cellH, cellW int, st state.GameState) *Board {
	return &Board{
		x:       startX,
		y:       startY,
		blockH:  cellH,
		blockW:  cellW,
		size:    len(st.Board),
		borderC: tcell.NewHexColor(int32(0xa0a0a0)),
	}
}
