// comments and shit

package main

import (
	"fmt"
	"math/rand"
)

type Hitter struct {
	Name       string
	Hitting    int
	BattingAvg float64
}

type Pitcher struct {
	Name          string
	Pitching      int
	OppBattingAvg float64
}

type Team struct {
	Name     string
	Hitters  []Hitter
	Pitchers []Pitcher
}

type Game struct {
	Innings int
	Score   [2]int
}

func simulateAtBat(hitter Hitter, pitcher Pitcher) string {
	chance := rand.Intn(250-1+1) + 1
	outcome := chance + hitter.Hitting

	if outcome < pitcher.Pitching {
		fmt.Printf("Result: %s for %s\n", "Hit", hitter.Name)
		return "Hit"
	} else {
		fmt.Printf("Result: %s for %s\n", "Strikeout", hitter.Name)
		return "Strikeout"
	}
}

func simulateTopInning(home Team, away Team, inning int) int {
	i := inning
	outs := 0
	runs := 0
	for outs < 3 {
		for _, Hitter := range away.Hitters {
			result := simulateAtBat(Hitter, home.Pitchers[0])
			if result == "Strikeout" {
				outs++
				if outs > 3 {
					fmt.Printf("End of the top of the %d inning\n", i)
					break
				}
			} else if result == "Hit" {
				runs++
			}
		}
	}
	return runs
}

func simulateBottomInning(home Team, away Team, inning int) int {
	i := inning
	outs := 0
	runs := 0
	for outs < 3 {
		for _, Hitter := range home.Hitters {
			result := simulateAtBat(Hitter, away.Pitchers[0])
			if result == "Strikeout" {
				outs++
				if outs > 3 {
					fmt.Printf("End of the top of the %d inning\n", i)
					break
				}
			} else if result == "Hit" {
				runs++
			}
		}
	}
	return runs
}

func simulateInning(home Team, away Team, inning int) [2]int {
	i := inning
	awayRuns := simulateTopInning(home, away, i)
	homeRuns := simulateBottomInning(home, away, i)
	fmt.Printf("End of the %d inning\n Home - %d Away - %d\n", i, homeRuns, awayRuns)
	score := [2]int{awayRuns, homeRuns}
	return score
}

func main() {
	homeHitters := []Hitter{
		{Name: "Player1", Hitting: 30, BattingAvg: 0.370},
		{Name: "Player2", Hitting: 20, BattingAvg: 0.350},
		{Name: "Player3", Hitting: 10, BattingAvg: 0.320},
		{Name: "Player4", Hitting: 45, BattingAvg: 0.320},
		{Name: "Player5", Hitting: 25, BattingAvg: 0.320},
		{Name: "Player6", Hitting: 15, BattingAvg: 0.320},
	}

	awayHitters := []Hitter{
		{Name: "PlayerA", Hitting: 31, BattingAvg: 0.380},
		{Name: "PlayerB", Hitting: 27, BattingAvg: 0.340},
		{Name: "PlayerC", Hitting: 29, BattingAvg: 0.360},
		{Name: "PlayerD", Hitting: 35, BattingAvg: 0.360},
		{Name: "PlayerE", Hitting: 23, BattingAvg: 0.360},
		{Name: "PlayerF", Hitting: 17, BattingAvg: 0.360},
	}

	homePitchers := []Pitcher{
		{Name: "Pitcher1", Pitching: 25, OppBattingAvg: 0.300},
		{Name: "Pitcher2", Pitching: 20, OppBattingAvg: 0.320},
		{Name: "Pitcher3", Pitching: 15, OppBattingAvg: 0.310},
	}

	awayPitchers := []Pitcher{
		{Name: "PitcherA", Pitching: 30, OppBattingAvg: 0.290},
		{Name: "PitcherB", Pitching: 25, OppBattingAvg: 0.310},
		{Name: "PitcherC", Pitching: 20, OppBattingAvg: 0.300},
	}

	homeTeam := Team{Name: "Home", Hitters: homeHitters, Pitchers: homePitchers}
	awayTeam := Team{Name: "Away", Hitters: awayHitters, Pitchers: awayPitchers}

	game := Game{Innings: 9}

	for i := 0; i < game.Innings; i++ {
		tempScore := simulateInning(homeTeam, awayTeam, i+1)
		game.Score[0] += tempScore[0]
		game.Score[1] += tempScore[1]
	}

	fmt.Printf("Final Score: Away %d - Home %d\n", game.Score[0], game.Score[1])
}
