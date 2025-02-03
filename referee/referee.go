package referee

import (
	"fmt"
	"math/rand"  // ✅ For generating random numbers
	"time"       // ✅ For seeding randomness
	"pingpong/player"
)

type Referee struct {
	Players []*player.Player
	History []string
}

// StartTournament orchestrates the full championship
func (r *Referee) StartTournament() {
	fmt.Println("🏓 Championship Started!")

	if len(r.Players) != 8 {
		fmt.Printf("❌ Error: Expected 8 players, but got %d\n", len(r.Players))
		return
	}

	fmt.Println("\n🎯 Quarter-Finals")
	roundWinners := r.playRound(r.Players)

	fmt.Println("\n🔥 Semi-Finals")
	roundWinners = r.playRound(roundWinners)

	fmt.Println("\n🏆 Final Match")
	champion := r.playRound(roundWinners)[0]

	// Store champion in history
	result := fmt.Sprintf("🏆 Champion: %s 🏅", champion.Name)
	r.History = append(r.History, result)

	fmt.Println("\n🥇", result)
}

// playRound handles each round, returning winners
func (r *Referee) playRound(players []*player.Player) []*player.Player {
	numMatches := len(players) / 2
	winners := make([]*player.Player, 0, numMatches)

	// Seed random number generator for each match
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numMatches; i++ {
		player1 := players[i*2]
		player2 := players[i*2+1]

		// Use random selection for the winner (change this to more sophisticated logic if needed)
		winner := player1
		if rand.Intn(2) == 1 { // 50% chance for player2 to win
			winner = player2
		}

		// Store match result in history
		matchResult := fmt.Sprintf("⚔️ %s vs %s -> Winner: %s", player1.Name, player2.Name, winner.Name)
		r.History = append(r.History, matchResult)

		fmt.Println(matchResult)
		winners = append(winners, winner)
	}

	return winners
}

// ShowTournamentHistory prints the full history of the tournament
func (r *Referee) ShowTournamentHistory() {
	fmt.Println("\n📜 Tournament History:")
	for _, record := range r.History {
		fmt.Println(record)
	}
}
