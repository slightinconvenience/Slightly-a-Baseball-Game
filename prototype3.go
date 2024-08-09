// comments and shit

package main

import (
	"container/list"
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
	Lineup   *list.List
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
	lineup := away.Lineup

TopInningLoop:
	for outs < 3 {
		hitterElement := lineup.Front()
		hitter := hitterElement.Value.(*Hitter)
		result := simulateAtBat(hitter, &home.Pitchers[0])
		switch result {
		case Strikeout:
			fmt.Printf("%s struck out\n", hitter.Name)
			outs++
			lineup.MoveToBack(hitterElement)
			if outs >= 3 {
				fmt.Printf("End of the top of the %d inning\n", i)
				break TopInningLoop
			}
		case Hit:
			fmt.Printf("%s hit the ball\n", hitter.Name)
			lineup.MoveToBack(hitterElement)
			runs += advanceRunners(hitter, &bases)
		case Walk:
			runs += advanceRunners(hitter, &bases)
		case Out:
			outs++
			if outs >= 3 {
				fmt.Printf("End of the top of the %d inning\n", i)
				break TopInningLoop
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
	lineup := home.Lineup

BottomInningLoop:
	for outs < 3 {
		hitterElement := lineup.Front()
		hitter := hitterElement.Value.(*Hitter)
		result := simulateAtBat(hitter, &away.Pitchers[0])
		switch result {
		case Strikeout:
			fmt.Printf("%s struck out\n", hitter.Name)
			outs++
			lineup.MoveToBack(hitterElement)
			if outs >= 3 {
				fmt.Printf("End of the bottom of the %d inning\n", i)
				break BottomInningLoop
			}
		case Hit:
			fmt.Printf("%s hit the ball\n", hitter.Name)
			lineup.MoveToBack(hitterElement)
			runs += advanceRunners(hitter, bases)
		case Walk:
			runs += advanceRunners(hitter, bases)
		case Out:
			outs++
			if outs >= 3 {
				fmt.Printf("End of the bottom of the %d inning\n", i)
				break BottomInningLoop
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

func Statistics(team Team, pitchers []Pitcher, hitters []*Hitter) {
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

func lineUp(hitters []*Hitter) *list.List {
	l := list.New()
	for i := 0; i < len(hitters); i++ {
		l.PushBack(hitters[i])
	}
	return l
}

func main() {
	homeHitters := []*Hitter{
		{Name: "Josh", Hitting: 45, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Jeff", Hitting: 45, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Bobby", Hitting: 35, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Hill", Hitting: 25, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Hank", Hitting: 50, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Todd", Hitting: 40, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Derek", Hitting: 30, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Mike", Hitting: 20, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "John", Hitting: 10, BattingAvg: 0.000, AtBats: 0, Hits: 0},
	}

	homeLineup := lineUp(homeHitters)

	awayHitters := []*Hitter{
		{Name: "Cory", Hitting: 15, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Duye", Hitting: 50, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Elvis", Hitting: 40, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Noah", Hitting: 20, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Nate", Hitting: 55, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Kyle", Hitting: 45, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Evan", Hitting: 35, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Alex", Hitting: 25, BattingAvg: 0.000, AtBats: 0, Hits: 0},
		{Name: "Matt", Hitting: 15, BattingAvg: 0.000, AtBats: 0, Hits: 0},
	}

	awayLineup := lineUp(awayHitters)

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

	homeTeam := Team{Name: "Home", Pitchers: homePitchers, Lineup: homeLineup}
	awayTeam := Team{Name: "Away", Pitchers: awayPitchers, Lineup: awayLineup}

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
