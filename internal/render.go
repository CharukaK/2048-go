package internal

import (
	"github.com/CharukaK/2048-go/internal/components"
	"github.com/CharukaK/2048-go/internal/state"
	"github.com/gdamore/tcell"
)

// func Makebox(s tcell.Screen, num int) {
// 	startX := 5 + num
// 	startY := 5 + num
// 	w, h := 11, 7
// 	gl := ' '
// 	st := tcell.StyleDefault
//
// 	rgb := tcell.NewHexColor(int32(0xbcac9f))
// 	st = st.Background(rgb)
//
// 	if !(s.Colors() > 256 && s.Colors() > 1) {
// 		gl = '@'
// 	}
//
// 	for row := 0; row < w; row++ {
// 		for col := 0; col < h; col++ {
// 			if row == 0 || row == w-1 {
// 				s.SetCell(startX+row, startY+col, st, gl)
// 			} else if col == 0 || col == h-1 {
// 				s.SetCell(startX+row, startY+col, st, gl)
// 			}
// 		}
// 	}
//
// 	s.SetContent(startX+w/2, startY+h/2, '2', nil, st)
// }

const (
	// NOTE: keep these values odd, as most of the renderings are based on that assumption
	blockH = 7
	blockW = 11
)

func Render(st state.GameState, scr tcell.Screen) {
	topOffset, leftOffset := 10, 10
	board := components.NewBoard(topOffset, leftOffset, blockH, blockW, st)
	board.Render(scr)

	// render individual cells
	for row := 0; row < len(st.Board); row++ {
		for col := 0; col < len(st.Board); col++ {

			cell := components.NewCell(
				leftOffset+1+col*blockW,
				topOffset+1+row*blockH,
				blockH-2,
				blockW-2,
				st.Board[row][col],
			)

			cell.Render(scr)
		}
	}

	scr.Show()
}
