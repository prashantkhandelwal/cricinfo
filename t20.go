package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

func (t *T20MatchResult) Download(c *colly.Collector) {

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
					break
				case 1:
					d.TeamB = gq.Find(selector).Children().Eq(j).Text()
					break
				case 2:
					d.Winner = gq.Find(selector).Children().Eq(j).Text()
					break
				case 3:
					d.Margin = gq.Find(selector).Children().Eq(j).Text()
					break
				case 4:
					d.Ground = gq.Find(selector).Children().Eq(j).Text()
					break
				case 5:
					d.MatchDate = gq.Find(selector).Children().Eq(j).Text()
					break
				case 6:
					d.Scorecard = gq.Find(selector).Children().Eq(j).Text()
					break
				}
			}
			data = append(data, d)
		}
	})

	c.Visit(T20_URLS["T20_2008_MATCH_RESULTS"])
}

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

func (t *T20MatchResult) Save(path string, format Format) (bool, error) {

	if len(strings.TrimSpace(path)) == 0 {
		return false, errors.New("Parameter: path cannot be empty")
	}
	dir, fname := filepath.Split(path)
	_, err := os.Stat(dir)
	if err != nil {
		fmt.Println("Path does not exists. Creating...")
		err := os.MkdirAll(dir, os.ModeSticky|os.ModeDir|os.ModePerm)
		if err != nil {
			return false, errors.New("ERROR: Cannot create path due to permission - " + err.Error())
		}
	}
	if !strings.Contains(fname, ".json") {
		return false, errors.New("ERROR: Please provide a valid json file name.")
	}

	if format == JSON {
		json, err := json.MarshalIndent(&data, "", "  ")
		if err != nil {
			return false, err
		}
		f, err := os.Create(path)
		if err != nil {
			return false, errors.New("ERROR: Cannot create file - " + err.Error())
		}

		_, err = f.Write(json)
		if err != nil {
			return false, errors.New("ERROR: Cannot write to the file - " + err.Error())
		}

		err = f.Close()
		if err != nil {
			return false, errors.New("ERROR: Cannot close the file - " + err.Error())
		}
	}

	if format == CSV {

	}

	return true, nil
}
