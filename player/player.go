package player

import (
	"math/rand"
	"time"
)

type Player struct {
	Name          string
	DefenseLength int
	Score         int
}

// Attack: Picks a random number between 1-10
func (p *Player) Attack() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10) + 1
}

// Defend: Generates a random defense array
func (p *Player) Defend() []int {
	rand.Seed(time.Now().UnixNano())
	defense := make([]int, p.DefenseLength)
	for i := range defense {
		defense[i] = rand.Intn(10) + 1
	}
	return defense
}
