package main

type AttackType int

const (
	AttackStraight      AttackType = 0
	AttackCombination   AttackType = 1
	AttackBruteForce    AttackType = 3
	AttackHybridDictMap AttackType = 6
	AttackHybridMapDict AttackType = 7
)

type Attack struct {
	Type     AttackType
	Straight *StraightAttack `json:",omitempty"`
}

type StraightAttack struct {
	Rules []string
	Dict  string
}
