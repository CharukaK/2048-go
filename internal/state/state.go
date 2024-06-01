package state

import (
	"container/list"
	"log"
	"math/rand"
)

type MoveEvent int

const (
	MoveUp MoveEvent = iota
	MoveDown
	MoveLeft
	MoveRight
)

type GameState struct {
	Board  [][]int
	Filled bool
	won    bool
}

type emptyEntry struct {
	row int
	col int
}

func (gs *GameState) MakeMove(ev MoveEvent) {
	emptyList := make([]emptyEntry, 0)
	switch ev {
	case MoveUp:
		for col := 0; col < len(gs.Board); col++ {
			var colList list.List
			for row := 0; row < len(gs.Board)-1; row++ {
				if gs.Board[row][col] != 0 {
					colList.PushBack(gs.Board[row][col])
				}

				gs.Board[row][col] = 0
			}

			ri := 0
			for x := colList.Front(); x != nil; x = x.Next() {
				next := x.Next()
				val, ok := x.Value.(int)
				if !ok {
					log.Fatal("error parsing board value")
				}
				if next != nil && x.Value == next.Value {
					gs.Board[ri][col] = val * 2
					x = next
				} else {
					gs.Board[ri][col] = val
				}
				ri++
			}

			for row := colList.Len() - 1; row < len(gs.Board); row++ {
				emptyList = append(emptyList, emptyEntry{row, col})
			}
		}
	case MoveDown:
	case MoveRight:
	case MoveLeft:
	}

	if len(emptyList) == 0 {
		gs.Filled = true
	} else {
        entryI := rand.Intn(len(emptyList))

        gs.Board[emptyList[entryI].row][emptyList[entryI].col] = 2
	}
}

func NewGameState(size int) *GameState {
	board := make([][]int, size)

	for i := range board {
		board[i] = make([]int, size)
	}

	startBlockX := rand.Intn(size)
	startBlockY := rand.Intn(size)

	board[startBlockY][startBlockX] = 2

	return &GameState{
		Board: board,
	}
}
