package player

import (
	
	"math/rand"
	"time"
)

// Enum-like structure for strategy types
type Strategy int

const (
	Defensive Strategy = iota
	Aggressive
	Balanced
)

// Player struct with added strategy and attributes
type Player struct {
	Name          string
	DefenseLength int
	Score         int
	Strategy      Strategy
	SkillLevel    int // 1-10
	Stamina       int // 1-100
	Ready         bool // Track if player is ready
}

// NewPlayer function to create a new player with default attributes
func NewPlayer(name string, defenseLength, skillLevel, stamina int, strategy Strategy) *Player {
	return &Player{
		Name:          name,
		DefenseLength: defenseLength,
		SkillLevel:    skillLevel,
		Stamina:       stamina,
		Strategy:      strategy,
		Ready:         false, // Players are initially not ready
	}
}

// Attack method based on player's strategy
func (p *Player) Attack() int {
	rand.Seed(time.Now().UnixNano())

	// Adjust attack based on strategy and stamina
	attackStrength := 0
	switch p.Strategy {
	case Defensive:
		attackStrength = rand.Intn(5) + 1
	case Aggressive:
		attackStrength = rand.Intn(5) + 6
	case Balanced:
		attackStrength = rand.Intn(10) + 1
	}

	// If stamina is low, decrease attack strength
	if p.Stamina < 50 {
		attackStrength -= 1
	}

	// Ensure attack strength is positive
	if attackStrength < 1 {
		attackStrength = 1
	}

	return attackStrength
}

// Defend method generates defense set based on player's defense length
func (p *Player) Defend() []int {
	rand.Seed(time.Now().UnixNano())
	defenseSet := make([]int, p.DefenseLength)
	for i := 0; i < p.DefenseLength; i++ {
		defenseSet[i] = rand.Intn(10) + 1
	}
	return defenseSet
}
