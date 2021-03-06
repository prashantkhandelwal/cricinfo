package main

import (
	"fmt"
	"strconv"

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

type T20HighestTotals struct {
	Team       string
	Score      string
	Overs      float64
	RunRate    float64
	Innings    int
	Opposition string
	Ground     string
	MatchDate  string
	ScoreCard  string
}

// Highest Totals

var data_highestTotals = []T20HighestTotals{}
var dht = T20HighestTotals{}

func (t *T20HighestTotals) DownloadHighestTotals(c *colly.Collector, url string) {
	c.OnHTML("#ciHomeContentlhs", func(e *colly.HTMLElement) {
		gq := e.DOM
		cols := gq.Find(".engineTable > tbody tr").First().Children().Length()
		rows := gq.Find(".pnl650M > div > table:nth-child(5) > tbody > tr").Length()

		for i := 1; i < rows; i++ {
			selector := fmt.Sprintf(".pnl650M > div > table:nth-child(5) > tbody > tr:nth-child(%v)", i)
			dht = T20HighestTotals{}
			for j := 0; j < cols; j++ {
				switch j {
				case 0:
					dht.Team = gq.Find(selector).Children().Eq(j).Text()
				case 1:
					dht.Score = gq.Find(selector).Children().Eq(j).Text()
				case 2:
					overs, _ := strconv.ParseFloat(gq.Find(selector).Children().Eq(j).Text(), 64)
					dht.Overs = overs
				case 3:
					runrate, _ := strconv.ParseFloat(gq.Find(selector).Children().Eq(j).Text(), 64)
					dht.RunRate = runrate
				case 4:
					innings, _ := strconv.ParseInt(gq.Find(selector).Children().Eq(j).Text(), 0, 32)
					dht.Innings = int(innings)
				case 6:
					dht.Opposition = gq.Find(selector).Children().Eq(j).Text()
				case 7:
					dht.Ground = gq.Find(selector).Children().Eq(j).Text()
				case 8:
					dht.MatchDate = gq.Find(selector).Children().Eq(j).Text()
				case 9:
					dht.ScoreCard = gq.Find(selector).Children().Eq(j).Text()
				}
			}
			data_highestTotals = append(data_highestTotals, dht)
		}
	})

	err := c.Visit(url)
	if err != nil {
		fmt.Printf("ERROR: Cannot visit %s", url)
		return
	}
	c.Wait()

	var w Wrap[T20HighestTotals]
	w.Data = data_highestTotals

	_, err = w.Save(GetPath(T20_HIGHEST_TOTALS_URLS, url), JSON)
	if err != nil {
		fmt.Println(err)
	}

	data_highestTotals = []T20HighestTotals{}
}

// Match Results

var data_matchResults = []T20MatchResult{}
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
			data_matchResults = append(data_matchResults, d)
		}
	})

	err := c.Visit(url)
	if err != nil {
		fmt.Printf("ERROR: Cannot visit %s", url)
		return
	}
	c.Wait()

	var w Wrap[T20MatchResult]
	w.Data = data_matchResults

	_, err = w.Save(GetPath(T20_MATCH_RESULTS_URLS, url), JSON)
	if err != nil {
		fmt.Println(err)
	}

	data_matchResults = []T20MatchResult{}
}

// TODO: make this function generic
func (t *T20MatchResult) ToString() {
	for _, d := range data_matchResults {
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
