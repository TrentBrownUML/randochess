package rules

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"prushton.com/randochess/v2/board"
)

type Ruleset struct {
	Name       string `json:"name"`
	PieceRules map[board.PieceType]func(board.Board, int, int) bool
	Width      int
	Height     int
	Move       func(*board.Board, int, int)
	GetWinner  func(board.Board) board.Team
	InitBoard  func(*board.Board) error
}

func (self Ruleset) MarshalJSON() ([]byte, error) {
	type MarshalableRuleset struct {
		Name string `json:"name"`
	}

	return json.Marshal(MarshalableRuleset{Name: self.Name})
}

func SelectRuleset(name string) (Ruleset, error) {
	if name == "Random" {
		keys := make([]string, 0, len(AllRulesets))
		for k := range AllRulesets {
			keys = append(keys, k)
		}
		randomKey := keys[rand.Intn(len(keys))]
		return AllRulesets[randomKey], nil
	}

	ruleset, exists := AllRulesets[name]
	if !exists {
		return Ruleset{}, fmt.Errorf("Invalid name")
	}

	return ruleset, nil
}
