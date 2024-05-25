package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

type GameState struct {
	board [][]string
}

func NewGameState(size int) *GameState {
	board := make([][]string, size)

	for i := range board {
		board[i] = make([]string, size)
	}

	return &GameState{
		board: board,
	}
}

func makebox(s tcell.Screen, num int) {
	startX := 10 + num
	startY := 10 + num
	w, h := 7, 7
	gl := ' '
	st := tcell.StyleDefault

	// lx := 10
	// ly := 10

	rgb := tcell.NewHexColor(int32(rand.Int() & 0xffffff))
	st = st.Background(rgb)

	if !(s.Colors() > 256 && s.Colors() > 1) {
		gl = '@'
	}

	for row := 0; row < w; row++ {
		for col := 0; col < h; col++ {
			if row == 0 || row == w-1 {
				s.SetCell(startX+row, startY+col, st, gl)
			} else if col == 0 || col == h-1 {
				s.SetCell(startX+row, startY+col, st, gl)
			}
		}
	}

	s.SetContent(startX+w/2, startY+h/2, '2', nil, st)
}

func main() {
	fmt.Print("Enter grid size:")
	var size int
	_, err := fmt.Scanln(&size)

	if err != nil {
		log.Fatal("Error reading grid size")
	}

	// gs := NewGameState(size)

	scr, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := scr.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	scr.SetStyle(tcell.StyleDefault)

	// Clear screen
	scr.Clear()

	// scr.SetContent(0, 0, 'H', nil, tcell.StyleDefault)
	// scr.SetContent(1, 0, 'i', nil, tcell.StyleDefault)
	// scr.SetContent(2, 0, '!', nil, tcell.StyleDefault)
	//
	// scr.Show()

	makebox(scr, 2)

	scr.Show()
	time.Sleep(5 * time.Second)
	scr.Fini()
}
