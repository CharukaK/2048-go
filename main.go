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
			if k == tcell.KeyEsc || k == tcell.KeyCtrlC {
                return
			}
		}
	}

	// scr.SetContent(0, 0, 'H', nil, tcell.StyleDefault)
	// scr.SetContent(1, 0, 'i', nil, tcell.StyleDefault)
	// scr.SetContent(2, 0, '!', nil, tcell.StyleDefault)
	//
	// scr.Show()

	// Makebox(scr, 2)

}
