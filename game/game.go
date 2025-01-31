package game

import (
	"fmt"
	"pingpong/player" // ✅ Correct import
)

// Game struct
type Game struct {
	Player1 *player.Player
	Player2 *player.Player
}

// PlayGame simulates the ping pong match
func (g *Game) PlayGame() string {
	attacker, defender := g.Player1, g.Player2

	for {
		attackNum := attacker.Attack()
		defenseSet := defender.Defend()

		// Check if attack number is in defense set
		if !contains(defenseSet, attackNum) {
			attacker.Score++
			fmt.Printf("%s scores! Current score: %d\n", attacker.Name, attacker.Score)
			if attacker.Score == 5 {
				fmt.Printf("🏆 %s wins the game!\n", attacker.Name)
				return attacker.Name
			}
		} else {
			attacker, defender = defender, attacker
		}
	}
}

// Helper function to check if an attack exists in defense array
func contains(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}
