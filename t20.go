package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type T20MatchResult struct {
	TeamA     string
	TeamB     string
	Winner    string
	Margin    string
	Ground    string
	MatchDate string
	Scorecard string
}

var data = []T20MatchResult{}
var d = T20MatchResult{}

func (t *T20MatchResult) DownloadMatchResults(c *colly.Collector, url string) {
	c.OnHTML("#ciHomeContentlhs", func(e *colly.HTMLElement) {
		gq := e.DOM
		cols := gq.Find(".engineTable > tbody tr").First().Children().Length()
		rows := gq.Find(".pnl650M > div > table:nth-child(5) > tbody > tr").Length()

		for i := 1; i < rows; i++ {
			selector := fmt.Sprintf(".pnl650M > div > table:nth-child(5) > tbody > tr:nth-child(%v)", i)
			d = T20MatchResult{}
			for j := 0; j < cols; j++ {
				switch j {
				case 0:
					d.TeamA = gq.Find(selector).Children().Eq(j).Text()
				case 1:
					d.TeamB = gq.Find(selector).Children().Eq(j).Text()
				case 2:
					d.Winner = gq.Find(selector).Children().Eq(j).Text()
				case 3:
					d.Margin = gq.Find(selector).Children().Eq(j).Text()
				case 4:
					d.Ground = gq.Find(selector).Children().Eq(j).Text()
				case 5:
					d.MatchDate = gq.Find(selector).Children().Eq(j).Text()
				case 6:
					d.Scorecard = gq.Find(selector).Children().Eq(j).Text()
				}
			}
			data = append(data, d)
		}
	})

	err := c.Visit(url)
	if err != nil {
		fmt.Printf("ERROR: Cannot visit %s", url)
		return
	}
	c.Wait()

	var w Wrap[T20MatchResult]
	w.Data = data

	_, err = w.Save(GetPath(url), JSON)
	if err != nil {
		fmt.Println(err)
	}

	data = []T20MatchResult{}
}

// TODO: make this function generic
func (t *T20MatchResult) ToString() {
	for _, d := range data {
		fmt.Println("----------------------------------")
		fmt.Println("Team A : " + d.TeamA)
		fmt.Println("Team B : " + d.TeamB)
		fmt.Println("Winner : " + d.Winner)
		fmt.Println("Margin : " + d.Margin)
		fmt.Println("Ground : " + d.Ground)
		fmt.Println("Match Date : " + d.MatchDate)
		fmt.Println("ScoreCard: " + d.Scorecard)
		fmt.Println("----------------------------------")

	}
}
