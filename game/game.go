package game

import (
	"fmt"
	"pingpong/player"
	
	"time"
)

// Game struct with advanced tournament logic
type Game struct {
	Player1     *player.Player
	Player2     *player.Player
	IsBonusRound bool
	RoundNumber int
	Winner      *player.Player
	History     []string
}

// ToggleBonusRound function for special rounds
func (g *Game) ToggleBonusRound() {
	g.IsBonusRound = !g.IsBonusRound
}

// PlayRound function to simulate a round of the game
func (g *Game) PlayRound() {
	attacker := g.Player1
	defender := g.Player2

	attackNum := attacker.Attack()
	defenseSet := defender.Defend()

	if !contains(defenseSet, attackNum) {
		attacker.Score++
		g.History = append(g.History, fmt.Sprintf("%s attacks %d, scores! Current score: %d", attacker.Name, attackNum, attacker.Score))

		if g.IsBonusRound {
			attacker.Score++ // Bonus point
			g.History = append(g.History, fmt.Sprintf("Bonus round: %s gets extra point!", attacker.Name))
		}
	} else {
		// Swap roles properly
		g.Player1, g.Player2 = g.Player2, g.Player1
		g.History = append(g.History, fmt.Sprintf("%s defends successfully. Switching roles.", defender.Name))
	}

	g.RoundNumber++
}


// PlayGame starts the game and simulates the rounds
func (g *Game) PlayGame() {
	for g.Player1.Score < 5 && g.Player2.Score < 5 {
		g.PlayRound()
		time.Sleep(1 * time.Second) // Pause to simulate time between rounds
	}

	// Declare winner
	if g.Player1.Score >= 5 {
		g.Winner = g.Player1
	} else {
		g.Winner = g.Player2
	}

	// Final logs
	g.History = append(g.History, fmt.Sprintf("Game Over! Winner: %s\n", g.Winner.Name))
}

// Helper function to check if a number exists in the defense array
func contains(arr []int, num int) bool {
	for _, val := range arr {
		if val == num {
			return true
		}
	}
	return false
}
