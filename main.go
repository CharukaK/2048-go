package main

import (
	"fmt"
	"log"
	"time"

	"github.com/CharukaK/2048-go/internal"
	"github.com/CharukaK/2048-go/internal/state"
	"github.com/gdamore/tcell"
)

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
	gs := state.NewGameState(size)

    internal.Render(*gs, scr)
	// scr.SetContent(0, 0, 'H', nil, tcell.StyleDefault)
	// scr.SetContent(1, 0, 'i', nil, tcell.StyleDefault)
	// scr.SetContent(2, 0, '!', nil, tcell.StyleDefault)
	//
	// scr.Show()

	// Makebox(scr, 2)

	scr.Show()
	time.Sleep(5 * time.Second)
	scr.Fini()
}
