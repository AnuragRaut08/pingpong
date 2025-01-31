package player

import (
	"testing"
)

func TestAttack(t *testing.T) {
	player := Player{Name: "Rishi", DefenseLength: 8}
	attackNum := player.Attack()

	if attackNum < 1 || attackNum > 10 {
		t.Errorf("Attack number out of range: got %d", attackNum)
	}
}

func TestDefend(t *testing.T) {
	player := Player{Name: "Suraj", DefenseLength: 7}
	defenseSet := player.Defend()

	if len(defenseSet) != player.DefenseLength {
		t.Errorf("Defense set length mismatch: expected %d, got %d", player.DefenseLength, len(defenseSet))
	}
}
