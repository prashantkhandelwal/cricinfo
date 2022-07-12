package main

import (
	"reflect"
	"testing"

	"github.com/gocolly/colly/v2"
)

func TestNewColly(t *testing.T) {
	c := NewScrapper()

	want := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.CacheDir("cache"),
		colly.Async(true))

	if reflect.TypeOf(c) != reflect.TypeOf(want) {
		t.Errorf("Error in initializing NewColly")
	}
}
