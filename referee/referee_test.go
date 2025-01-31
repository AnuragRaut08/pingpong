package referee

import (
	"pingpong/player"
	"testing"
)

func TestStartTournament(t *testing.T) {
	players := []*player.Player{
		{Name: "Rishi", DefenseLength: 8},
		{Name: "Suraj", DefenseLength: 7},
		{Name: "Nikhil", DefenseLength: 6},
		{Name: "Rohan", DefenseLength: 4},
	}

	ref := Referee{Players: players}
	ref.StartTournament()
}
