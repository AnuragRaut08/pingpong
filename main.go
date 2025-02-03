package main

import (
	"fmt"
	"pingpong/player"
	"pingpong/referee"
)

func main() {
	players := []*player.Player{
		{Name: "Rishi", DefenseLength: 8},
		{Name: "Suraj", DefenseLength: 7},
		{Name: "Nikhil", DefenseLength: 6},
		{Name: "Rohan", DefenseLength: 4},
		{Name: "Amit", DefenseLength: 9},
		{Name: "Raj", DefenseLength: 5},
		{Name: "Sam", DefenseLength: 7},
		{Name: "Dev", DefenseLength: 6},
	}

	ref := referee.Referee{Players: players}

	fmt.Println("\nStarting Ping Pong Tournament!")
	ref.StartTournament()

	// ✅ This should now work without errors
	ref.ShowTournamentHistory()
}
