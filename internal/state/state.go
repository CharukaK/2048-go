package state

import (
	"fmt"
	"math/rand"

	"github.com/CharukaK/2048-go/internal/logger"
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

type entry struct {
	merged bool
	val    int
}

type inputPlaceHolder struct {
	row int
	col int
}

func (gs *GameState) MakeMove(ev MoveEvent) {
	logger.PrintInfo(fmt.Sprintf("Before state: %v", gs.Board))
	moved := false
	freeEntries := make([]inputPlaceHolder, 0)
	switch ev {
	case MoveUp:
		logger.PrintInfo("Move up event triggered")
		for col := 0; col < len(gs.Board); col++ {
			entries := make([]entry, 0)
			for row := 0; row < len(gs.Board); row++ {
				if gs.Board[row][col] == 0 {
					continue
				}

				if len(entries) > 0 &&
					!entries[len(entries)-1].merged &&
					gs.Board[row][col] == entries[len(entries)-1].val {

					entries[len(entries)-1].val = gs.Board[row][col] + entries[len(entries)-1].val
					entries[len(entries)-1].merged = true
				} else {
					entries = append(entries, entry{merged: false, val: gs.Board[row][col]})
				}
			}

			target := 0
			for _, ent := range entries {
				if gs.Board[target][col] != ent.val {
					moved = true
				}
				gs.Board[target][col] = ent.val
				target++
			}

			for i := target; i < len(gs.Board); i++ {
				if gs.Board[i][col] != 0 {
					moved = true
				}
				gs.Board[i][col] = 0
				freeEntries = append(freeEntries, inputPlaceHolder{i, col})
			}
		}
	case MoveDown:
		logger.PrintInfo("Move down event triggered")
		for col := 0; col < len(gs.Board); col++ {
			entries := make([]entry, 0)
			for row := len(gs.Board) - 1; row >= 0; row-- {
				if gs.Board[row][col] == 0 {
					continue
				}

				if len(entries) > 0 &&
					!entries[len(entries)-1].merged &&
					gs.Board[row][col] == entries[len(entries)-1].val {

					entries[len(entries)-1].val = gs.Board[row][col] + entries[len(entries)-1].val
					entries[len(entries)-1].merged = true
				} else {
					entries = append(entries, entry{merged: false, val: gs.Board[row][col]})
				}
			}

			target := len(gs.Board) - 1
			for _, ent := range entries {
				if gs.Board[target][col] != ent.val {
					moved = true
				}
				gs.Board[target][col] = ent.val
				target--
			}

			for i := target; i >= 0; i-- {
				if gs.Board[i][col] != 0 {
					moved = true
				}
				gs.Board[i][col] = 0
				freeEntries = append(freeEntries, inputPlaceHolder{target, col})
			}
		}
	case MoveRight:
		logger.PrintInfo("Move right event triggered")
		for row := 0; row < len(gs.Board); row++ {
			entries := make([]entry, 0)
			for col := len(gs.Board) - 1; col >= 0; col-- {
				if gs.Board[row][col] == 0 {
					continue
				}

				if len(entries) > 0 &&
					!entries[len(entries)-1].merged &&
					gs.Board[row][col] == entries[len(entries)-1].val {

					entries[len(entries)-1].val = gs.Board[row][col] + entries[len(entries)-1].val
					entries[len(entries)-1].merged = true
				} else {
					entries = append(entries, entry{merged: false, val: gs.Board[row][col]})
				}
			}

			target := len(gs.Board) - 1
			for _, ent := range entries {
				if gs.Board[row][target] != ent.val {
					moved = true
				}
				gs.Board[row][target] = ent.val
				target--
			}

			for i := target; i >= 0; i-- {
				if gs.Board[row][i] != 0 {
					moved = true
				}
				gs.Board[row][i] = 0
				freeEntries = append(freeEntries, inputPlaceHolder{row, i})
			}
		}
	case MoveLeft:
		logger.PrintInfo("Move left event triggered")
		for row := 0; row < len(gs.Board); row++ {
			entries := make([]entry, 0)
			for col := 0; col < len(gs.Board); col++ {
				if gs.Board[row][col] == 0 {
					continue
				}

				if len(entries) > 0 &&
					!entries[len(entries)-1].merged &&
					gs.Board[row][col] == entries[len(entries)-1].val {

					entries[len(entries)-1].val = gs.Board[row][col] + entries[len(entries)-1].val
					entries[len(entries)-1].merged = true
				} else {
					entries = append(entries, entry{merged: false, val: gs.Board[row][col]})
				}
			}

			target := 0
			for _, ent := range entries {
				if gs.Board[row][target] != ent.val {
					moved = true
				}
				gs.Board[row][target] = ent.val
				target++
			}

			for i := target; i < len(gs.Board); i++ {
				if gs.Board[row][i] != 0 {
					moved = true
				}
				gs.Board[row][i] = 0
				freeEntries = append(freeEntries, inputPlaceHolder{row, i})
			}
		}
	}
	logger.PrintInfo(fmt.Sprintf("after state: %v", gs.Board))

	if moved {
		if len(freeEntries) == 0 {
			gs.Filled = true
			return
		}
		pos := rand.Intn(len(freeEntries))
		gs.Board[freeEntries[pos].row][freeEntries[pos].col] = 2
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
