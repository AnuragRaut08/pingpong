package game

import (
	"testing"
	"pingpong/player"
)

// TestPlayGame ensures that the game plays until a winner is determined
func TestPlayGame(t *testing.T) {
	player1 := &player.Player{Name: "Rishi", DefenseLength: 7}
	player2 := &player.Player{Name: "Suraj", DefenseLength: 6}

	game := Game{Player1: player1, Player2: player2}

	// Just call the function, don't assign it to a variable
	game.PlayGame()

	if game.Winner == nil {
		t.Errorf("Expected a winner, but got nil")
	}

	if game.Winner.Score < 5 {
		t.Errorf("Expected winner to have at least 5 points, but got %d", game.Winner.Score)
	}
}
