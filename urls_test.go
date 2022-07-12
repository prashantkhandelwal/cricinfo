package main

import (
	"strings"
	"testing"
)

func TestGetMapKey(t *testing.T) {
	key := GetMapKey(T20_URLS, "https://stats.espncricinfo.com/ci/engine/records/team/match_results.html?id=3519;type=tournament")
	if len(strings.TrimSpace(key)) == 0 {
		t.Fatalf("ERRORGetMapKey does not retun key.")
	}
}

func TestGetPath(t *testing.T) {
	// Key: T20_2022_MATCHRESULTS
	var data = "https://stats.espncricinfo.com/ci/engine/records/team/match_results.html?id=14452;type=tournament"
	var want = "T20\\matchresults-2022.json"
	result := GetPath(data)
	if want != result {
		t.Fatalf("GetPath: Want %s but got %s", want, result)
	}
}
