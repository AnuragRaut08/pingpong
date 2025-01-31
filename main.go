package main

import (
	"pingpong/player"   // ✅ Correct import
	"pingpong/referee"  // ✅ Correct import
)

func main() {
	players := []*player.Player{
		{Name: "Rishi", DefenseLength: 8},
		{Name: "Suraj", DefenseLength: 7},
		{Name: "Nikhil", DefenseLength: 6},
		{Name: "Rohan", DefenseLength: 4},
	}

	ref := referee.Referee{Players: players}
	ref.StartTournament()
}
