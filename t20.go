package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

type T20MatchResult struct {
	TeamA     string
	TeamB     string
	Winner    string
	Margin    string
	Ground    string
	Match     string
	MatchDate string
	Scorecard string
}

var data = []T20MatchResult{}

func (t *T20MatchResult) Download(c *colly.Collector) {

	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		d := T20MatchResult{}
		e.ForEach("td", func(i int, h *colly.HTMLElement) {
			switch i {
			case 0:
				d.TeamA = strings.ReplaceAll(strings.TrimSpace(h.Text), " ", "")
				break
			case 1:
				d.TeamB = strings.ReplaceAll(strings.TrimSpace(h.Text), " ", "")
				break
			case 2:
				d.Winner = strings.ReplaceAll(strings.TrimSpace(h.Text), " ", "")
				break
			case 3:
				d.Margin = strings.ReplaceAll(strings.TrimSpace(h.Text), " ", "")
				break
			case 4:
				d.Ground = strings.ReplaceAll(strings.TrimSpace(h.Text), " ", "")
				break
			case 5:
				d.MatchDate = strings.ReplaceAll(strings.TrimSpace(h.Text), " ", "")
				break
			case 6:
				d.Scorecard = strings.ReplaceAll(strings.TrimSpace(h.Text), " ", "")
				break
			}
		})
		data = append(data, d)
	})

	c.Visit(T20_URLS["T20_2008_MATCH_RESULTS"])
}

func (t *T20MatchResult) ToString() {
	for _, d := range data {
		fmt.Println("----------------------------------")
		fmt.Println("Team A : " + d.TeamA)
		fmt.Println("Team B : " + d.TeamB)
		fmt.Println("Margin : " + d.Margin)
		fmt.Println("Ground : " + d.Ground)
		fmt.Println("Match Date : " + d.MatchDate)
		fmt.Println("Winner : " + d.Winner)
		fmt.Println("----------------------------------")
	}
}
