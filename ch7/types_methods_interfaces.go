package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

// ----- 7.1 -----

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

// ----- 7.2 -----

func (l League) MatchResult(firstTeam string, firstScore int, secondTeam string, secondScore int) {
	if _, ok := l.Teams[firstTeam]; !ok {
		return
	}

	if _, ok := l.Teams[secondTeam]; !ok {
		return
	}

	if firstScore == secondScore {
		return
	}

	if firstScore > secondScore {
		l.Wins[firstTeam]++
	} else {
		l.Wins[secondTeam]++
	}
}

// ----- 7.3 -----

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	for _, v := range r.Ranking() {
		io.WriteString(w, v)
		w.Write([]byte("\n"))
	}
}

func (l League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))
	for k := range l.Teams {
		names = append(names, k)
	}

	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] < l.Wins[names[j]]
	})

	return names
}

func main() {
	// ----- 7.2 -----

	l := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Italy": {
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"France": {
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"India": {
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Nigeria": {
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}

	l.MatchResult("Italy", 50, "France", 70)
	l.MatchResult("India", 85, "Nigeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Nigeria", 110)
	l.MatchResult("Italy", 65, "Nigeria", 70)
	l.MatchResult("France", 95, "India", 80)
	results := l.Ranking()

	fmt.Println(results)

	// ----- 7.3 -----

	RankPrinter(l, os.Stdout)
}
