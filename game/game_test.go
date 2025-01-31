package game

import (
	"pingpong/player"
	"testing"
)

func TestPlayGame(t *testing.T) {
	player1 := &player.Player{Name: "Rishi", DefenseLength: 8}
	player2 := &player.Player{Name: "Suraj", DefenseLength: 7}
	game := Game{Player1: player1, Player2: player2}

	winner := game.PlayGame()

	if winner != "Rishi" && winner != "Suraj" {
		t.Errorf("Invalid winner: got %s", winner)
	}
}
