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
	AtBats     int
	Hits       int
}

type Pitcher struct {
	Name          string
	Pitching      int
	Strikeouts    int
	OppBattingAvg float64
	OppAtBats     int
	OppHits       int
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

func simulateAtBat(hitter *Hitter, pitcher *Pitcher) string {
	chance := rand.Intn(200-1+1) + 1
	//fmt.Printf("debug: chance %d", chance)
	outcome := chance - hitter.Hitting
	hitter.AtBats++
	pitcher.OppAtBats++

	if outcome < pitcher.Pitching {
		hitter.Hits++
		pitcher.OppHits++
		fmt.Printf("Result: %s for %s\n", "Hit", hitter.Name)
		//fmt.Printf("debug: hitter %d hits. pitcher %d opphits\n", hitter.Hits, pitcher.OppHits)
		return "Hit"
	} else {
		pitcher.Strikeouts++
		fmt.Printf("Result: %s for %s\n", "Strikeout", hitter.Name)
		//fmt.Printf("debug: pitcher %d strikeouts\n", pitcher.Strikeouts)
		return "Strikeout"
	}
}

func simulateTopInning(home Team, away Team, inning int) int {
	i := inning
	outs := 0
	runs := 0
	for outs < 3 {
		for j := range away.Hitters {
			result := simulateAtBat(&away.Hitters[j], &home.Pitchers[0])
			if result == "Strikeout" {
				outs++
				if outs >= 3 {
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
		for j := range home.Hitters {
			result := simulateAtBat(&home.Hitters[j], &away.Pitchers[0])
			if result == "Strikeout" {
				outs++
				if outs >= 3 {
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

func simulateInning(home Team, away Team, inning int) (int, int) {
	i := inning
	awayRuns := 0
	homeRuns := 0
	awayRuns += simulateTopInning(home, away, i)
	homeRuns += simulateBottomInning(home, away, i)
	return homeRuns, awayRuns
}

func Statistics(team Team, pitchers []Pitcher, hitters []Hitter) {
	fmt.Printf("Team: %s\n", team.Name)
	for _, pitcher := range pitchers {
		pitcher.OppBattingAvg = float64(pitcher.OppHits) / float64(pitcher.OppAtBats)
		fmt.Printf("P | %s | %f | %d | %d |\n", pitcher.Name, pitcher.OppBattingAvg, pitcher.OppHits, pitcher.Strikeouts)
	}
	for _, hitter := range hitters {
		hitter.BattingAvg = float64(hitter.Hits) / float64(hitter.AtBats)
		fmt.Printf("H | %s | %f | %d |\n", hitter.Name, hitter.BattingAvg, hitter.Hits)
	}
}

func main() {
	homeHitters := []Hitter{
		{Name: "Josh", Hitting: 30, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Jeff", Hitting: 20, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Bobby", Hitting: 10, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Hill", Hitting: 45, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Hank", Hitting: 25, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Todd", Hitting: 15, BattingAvg: 0.000, AtBats: 0, Hits: 0},
	}

	awayHitters := []Hitter{
		{Name: "Cory", Hitting: 31, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Duye", Hitting: 27, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Elvis", Hitting: 29, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Noah", Hitting: 35, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Nate", Hitting: 23, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Kyle", Hitting: 17, BattingAvg: 0.000, AtBats: 0, Hits: 0},
	}

	homePitchers := []Pitcher{
		{Name: "Nolan", Pitching: 69, OppBattingAvg: 0.000, Strikeouts: 0, OppAtBats: 0, OppHits: 0},
		{Name: "Brandon", Pitching: 20, OppBattingAvg: 0.000, Strikeouts: 0, OppAtBats: 0, OppHits: 0},
		{Name: "Eddy", Pitching: 15, OppBattingAvg: 0.000, Strikeouts: 0, OppAtBats: 0, OppHits: 0},
	}

	awayPitchers := []Pitcher{
		{Name: "Chase", Pitching: 70, OppBattingAvg: 0.000, Strikeouts: 0, OppAtBats: 0, OppHits: 0},
		{Name: "Wade", Pitching: 25, OppBattingAvg: 0.000, Strikeouts: 0, OppAtBats: 0, OppHits: 0},
		{Name: "Gabe", Pitching: 20, OppBattingAvg: 0.000, Strikeouts: 0, OppAtBats: 0, OppHits: 0},
	}

	homeTeam := Team{Name: "Home", Hitters: homeHitters, Pitchers: homePitchers}
	awayTeam := Team{Name: "Away", Hitters: awayHitters, Pitchers: awayPitchers}

	game := Game{Innings: 9, Score: [2]int{0, 0}}

	homeScore := 0
	awayScore := 0

	for i := 1; i < game.Innings; i++ {
		homeScore, awayScore = simulateInning(homeTeam, awayTeam, i)
		game.Score[0] += homeScore
		game.Score[1] += awayScore
		fmt.Print("Inning: ", i, " Home: ", game.Score[0], " Away: ", game.Score[1], "\n")
	}

	inning := game.Innings
	for game.Score[0] == game.Score[1] {
		inning++
		homeScore, awayScore = simulateInning(homeTeam, awayTeam, inning)
		game.Score[0] += homeScore
		game.Score[1] += awayScore
		fmt.Print("Inning: ", inning, " Home: ", game.Score[0], " Away: ", game.Score[1], "\n")
	}

	Statistics(homeTeam, homePitchers, homeHitters)
	Statistics(awayTeam, awayPitchers, awayHitters)
	fmt.Printf("Final Score: Home %d - Away %d\n", game.Score[0], game.Score[1])
}
