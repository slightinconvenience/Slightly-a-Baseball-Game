// contains the main loop of the game

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name       string
	BattingAvg float64
	OnBasePct  float64
}

type Pitcher struct {
	Name          string
	StrikeoutRate float64
	WalkRate      float64
}

func simulateAtBat(player Player, pitcher Pitcher) string {
	rand.Seed(time.Now().UnixNano())
	outcome := rand.Float64()

	if outcome < pitcher.StrikeoutRate {
		return "Strikeout"
	} else if outcome < pitcher.StrikeoutRate+pitcher.WalkRate {
		return "Walk"
	} else if outcome < pitcher.StrikeoutRate+pitcher.WalkRate+player.BattingAvg {
		return "Hit"
	} else {
		return "Out"
	}
}

func main() {
	player := Player{Name: "Joe", BattingAvg: 0.300, OnBasePct: 0.370}
	pitcher := Pitcher{Name: "Smith", StrikeoutRate: 0.250, WalkRate: 0.100}

	result := simulateAtBat(player, pitcher)
	fmt.Printf("Result of at-bat: %s\n", result)
}
