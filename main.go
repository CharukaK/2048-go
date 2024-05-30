package main

import (
	"fmt"
	"log"

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
	quit := func() {
		panickVal := recover()
		scr.Fini()
		if panickVal != nil {
			panic(panickVal)
		}
	}
	defer quit()

	gs := state.NewGameState(size)
	for {
		internal.Render(*gs, scr)
		scr.Show()

		event := scr.PollEvent()

		switch ev := event.(type) {
		case *tcell.EventResize:
			scr.Sync()
		case *tcell.EventKey:
			k := ev.Key()

			if k == tcell.KeyCtrlC || k == tcell.KeyEsc {
				return
			}

			if k == tcell.KeyUp || ev.Rune() == 'k' {
				gs.MakeMove(state.MoveUp)
			} else if k == tcell.KeyDown || ev.Rune() == 'j' {
				gs.MakeMove(state.MoveDown)
			} else if k == tcell.KeyLeft || ev.Rune() == 'h' {
				gs.MakeMove(state.MoveLeft)
			} else if k == tcell.KeyRight || ev.Rune() == 'l' {
				gs.MakeMove(state.MoveRight)
			}
		}
	}
}
