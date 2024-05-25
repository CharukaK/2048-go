package main

import (
	"fmt"
	"log"
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

	scr.SetContent(0, 0, 'H', nil, tcell.StyleDefault)
	scr.SetContent(1, 0, 'i', nil, tcell.StyleDefault)
	scr.SetContent(2, 0, '!', nil, tcell.StyleDefault)

    scr.Show()

    time.Sleep(5 * time.Second)
    scr.Fini()
}

