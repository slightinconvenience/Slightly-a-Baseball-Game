// comments and shit

package main

import (
	"fmt"
	"math"
	"math/rand"
)

type AtBatOutcome int

const (
	Hit AtBatOutcome = iota
	Strikeout
	Walk
	Out
)

type Bases struct {
	First  *Hitter
	Second *Hitter
	Third  *Hitter
}

type Hitter struct {
	Name         string
	Hitting      int
	BattingAvg   float64
	AtBats       int
	Hits         int
	Runs         int
	RunsBattedIn int
}

type Pitcher struct {
	Name          string
	Pitching      int
	Strikeouts    int
	OppBattingAvg float64
	OppAtBats     int
	OppHits       int
	OppRuns       int
	Innings       int
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

func (o AtBatOutcome) String() string {
	return [...]string{"Hit", "Strikeout", "Walk", "Out"}[o]
}

func simulateAtBat(hitter *Hitter, pitcher *Pitcher) AtBatOutcome {
	chance := rand.Intn(200) + 1
	outcome := chance - hitter.Hitting
	hitter.AtBats++
	pitcher.OppAtBats++

	if outcome < pitcher.Pitching {
		hitter.Hits++
		pitcher.OppHits++
		//fmt.Printf("Result: %s for %s\n", "Hit", hitter.Name)
		return Hit
	} else {
		pitcher.Strikeouts++
		//fmt.Printf("Result: %s for %s\n", "Strikeout", hitter.Name)
		return Strikeout
	}
}

func advanceRunners(hitter *Hitter, bases *Bases) int {
	runs := 0
	if bases.Third != nil {
		bases.Third.Runs++
		hitter.RunsBattedIn++
		bases.Third = nil
		runs++
	}
	if bases.Second != nil {
		bases.Third = bases.Second
		bases.Second = nil
	}
	if bases.First != nil {
		bases.Second = bases.First
		bases.First = nil
	}
	bases.First = hitter
	return runs
}

// returns the runs scores in the inning
func simulateTopInning(home Team, away Team, inning int) int {
	i := inning
	outs := 0
	runs := 0
	bases := Bases{}

InningLoop:
	for outs < 3 {
		for j := range away.Hitters {
			result := simulateAtBat(&away.Hitters[j], &home.Pitchers[0])
			switch result {
			case Strikeout:
				outs++
				if outs >= 3 {
					fmt.Printf("End of the top of the %d inning\n", i)
					break InningLoop
				}
			case Hit:
				runs += advanceRunners(&away.Hitters[j], &bases)
			case Walk:
				runs += advanceRunners(&away.Hitters[j], &bases)
			case Out:
				outs++
				if outs >= 3 {
					fmt.Printf("End of the top of the %d inning\n", i)
					break InningLoop
				}
			}
		}
	}
	home.Pitchers[0].OppRuns += runs
	home.Pitchers[0].Innings++
	return runs
}

func simulateBottomInning(home Team, away Team, inning int) int {
	i := inning
	outs := 0
	runs := 0
	bases := &Bases{}

InningLoop:
	for outs < 3 {
		for j := range home.Hitters {
			result := simulateAtBat(&home.Hitters[j], &away.Pitchers[0])
			switch result {
			case Strikeout:
				outs++
				if outs >= 3 {
					fmt.Printf("End of the bottom of the %d inning\n", i)
					break InningLoop
				}
			case Hit:
				runs += advanceRunners(&home.Hitters[j], bases)
			case Walk:
				runs += advanceRunners(&home.Hitters[j], bases)
			case Out:
				outs++
				if outs >= 3 {
					fmt.Printf("End of the bottom of the %d inning\n", i)
					break InningLoop
				}
			}
		}
	}
	away.Pitchers[0].OppRuns += runs
	away.Pitchers[0].Innings++
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
		pitcher.OppBattingAvg = math.Round(pitcher.OppBattingAvg*1000) / 1000
		fmt.Printf("P | Name %s | OppAvg %.3f | Hits %d | SO %d | Runs %d | \n", pitcher.Name, pitcher.OppBattingAvg, pitcher.OppHits, pitcher.Strikeouts, pitcher.OppRuns)
	}
	for _, hitter := range hitters {
		hitter.BattingAvg = float64(hitter.Hits) / float64(hitter.AtBats)
		hitter.BattingAvg = math.Round(hitter.BattingAvg*1000) / 1000
		fmt.Printf("H | Name %s | Avg %.3f | Hits %d | Runs %d | RBIs %d |\n", hitter.Name, hitter.BattingAvg, hitter.Hits, hitter.Runs, hitter.RunsBattedIn)
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

	for i := 0; i < game.Innings; i++ {
		homeScore, awayScore = simulateInning(homeTeam, awayTeam, i+1)
		game.Score[0] += homeScore
		game.Score[1] += awayScore
		fmt.Print("Inning: ", i+1, " Home: ", game.Score[0], " Away: ", game.Score[1], "\n")
	}

	inning := game.Innings
	for game.Score[0] == game.Score[1] {
		inning++
		homeScore, awayScore = simulateInning(homeTeam, awayTeam, inning+1)
		game.Score[0] += homeScore
		game.Score[1] += awayScore
		fmt.Print("Inning: ", inning+1, " Home: ", game.Score[0], " Away: ", game.Score[1], "\n")
	}

	Statistics(homeTeam, homePitchers, homeHitters)
	Statistics(awayTeam, awayPitchers, awayHitters)
	fmt.Printf("Final Score: Home %d - Away %d\n", game.Score[0], game.Score[1])
}
