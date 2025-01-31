package referee

import (
	"fmt"
	"pingpong/game"
	"pingpong/player"
)

type Referee struct {
	Players []*player.Player
}

// StartTournament runs the championship
func (r *Referee) StartTournament() {
	fmt.Println("🏓 Championship Started!")

	// Round 1
	game1 := game.Game{Player1: r.Players[0], Player2: r.Players[1]}
	winner1 := game1.PlayGame()

	game2 := game.Game{Player1: r.Players[2], Player2: r.Players[3]}
	winner2 := game2.PlayGame()

	// Finals
	final := game.Game{Player1: &player.Player{Name: winner1}, Player2: &player.Player{Name: winner2}}
	champion := final.PlayGame()

	fmt.Printf("🎉 The champion is %s!\n", champion)
}
