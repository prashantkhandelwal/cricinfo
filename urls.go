package main

import (
	"fmt"
	"strings"
)

var (
	T20_URLS = map[string]string{

		// T20 2008 Main Page
		"T20_2008_MAIN": "https://www.espncricinfo.com/ci/engine/series/313494.html?view=records",

		// Team Records -> Match Results
		"T20_2008_MATCHRESULTS": "https://stats.espncricinfo.com/ci/engine/records/team/match_results.html?id=3519;type=tournament",
		"T20_2009_MATCHRESULTS": "https://www.espncricinfo.com/ci/engine/records/team/match_results.html?id=4801;type=tournament",
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
	key := GetMapKey(T20_URLS, url)
	if len(strings.TrimSpace(key)) > 0 {
		arr := strings.Split(key, "_")
		s := fmt.Sprintf("%s\\%s-%s.json", arr[0], strings.Replace(strings.ToLower(arr[2]), "_", "", 1), arr[1])
		return s
	}

	return ""
}
