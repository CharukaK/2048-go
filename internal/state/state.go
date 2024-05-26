package state

type GameState struct {
	Board [][]int
}

func NewGameState(size int) *GameState {
	board := make([][]int, size)

	for i := range board {
		board[i] = make([]int, size)
	}

	return &GameState{
		Board: board,
	}
}
