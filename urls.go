package main

import (
	"fmt"
	"strings"
)

var (
	T20_MATCH_RESULTS_URLS = map[string]string{

		// T20 2008 Main Page
		//"T20_2008_MAIN": "https://www.espncricinfo.com/ci/engine/series/313494.html?view=records",

		// Team Records -> Match Results
		"T20_2008_MATCHRESULTS": "https://stats.espncricinfo.com/ci/engine/records/team/match_results.html?id=3519;type=tournament",
		"T20_2009_MATCHRESULTS": "https://www.espncricinfo.com/ci/engine/records/team/match_results.html?id=4801;type=tournament",
		"T20_2010_MATCHRESULTS": "https://www.espncricinfo.com/ci/engine/records/team/match_results.html?id=5319;type=tournament",
		"T20_2011_MATCHRESULTS": "https://www.espncricinfo.com/ci/engine/records/team/match_results.html?id=5969;type=tournament",

		"T20_2012_MATCHRESULTS": "https://stats.espncricinfo.com/ci/engine/records/team/match_results.html?id=6680;type=tournament",
		"T20_2013_MATCHRESULTS": "https://stats.espncricinfo.com/ci/engine/records/team/match_results.html?id=7720;type=tournament",
		"T20_2014_MATCHRESULTS": "https://stats.espncricinfo.com/ci/engine/records/team/match_results.html?id=8827;type=tournament",
		"T20_2015_MATCHRESULTS": "https://stats.espncricinfo.com/ci/engine/records/team/match_results.html?id=9657;type=tournament",
		"T20_2016_MATCHRESULTS": "https://stats.espncricinfo.com/ci/engine/records/team/match_results.html?id=11001;type=tournament",
		"T20_2017_MATCHRESULTS": "https://stats.espncricinfo.com/ci/engine/records/team/match_results.html?id=11701;type=tournament",
		"T20_2018_MATCHRESULTS": "https://stats.espncricinfo.com/ci/engine/records/team/match_results.html?id=12210;type=tournament",
		"T20_2019_MATCHRESULTS": "https://stats.espncricinfo.com/ci/engine/records/team/match_results.html?id=12741;type=tournament",
		"T20_2020_MATCHRESULTS": "https://stats.espncricinfo.com/ci/engine/records/team/match_results.html?id=13533;type=tournament",
		"T20_2021_MATCHRESULTS": "https://stats.espncricinfo.com/ci/engine/records/team/match_results.html?id=13840;type=tournament",
		"T20_2022_MATCHRESULTS": "https://stats.espncricinfo.com/ci/engine/records/team/match_results.html?id=14452;type=tournament",
	}

	T20_HIGHEST_TOTALS_URLS = map[string]string{
		//Highest TotalS
		"T20_2008_HIGHESTTOTALS": "https://stats.espncricinfo.com/ci/engine/records/team/highest_innings_totals.html?id=3519;type=tournament",
	}
)

func GetMapKey(m map[string]string, value string) (key string) {
	for k, v := range m {
		if v == value {
			key = k
			return
		}
	}
	return
}

func GetPath(url string) string {
	key := GetMapKey(T20_MATCH_RESULTS_URLS, url)
	if len(strings.TrimSpace(key)) > 0 {
		arr := strings.Split(key, "_")
		s := fmt.Sprintf("%s\\%s\\%s.json", arr[0], arr[1], strings.Replace(strings.ToLower(arr[2]), "_", "", 1))
		return s
	}

	return ""
}
