package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func DoScrapedTitlesMatch(t *testing.T, buf []byte, expected map[string]bool) {
	matches, err := GetMatches(string(buf))
	if err != nil {
		fmt.Println("error on GetMatches")
		fmt.Println(err)
		return
	}

	for _, m := range matches {
		if expected[m] {
			t.Logf("OK: %s", m)
			delete(expected, m) // delete so we see titles not found
		} else {
			t.Errorf("False positive:\n\t %s", m)
		}
	}

	for exp := range expected {
		// titles not found that should have been
		t.Errorf("Not found:\n\t %s", exp)
	}
}

func TestUnicode(t *testing.T) {
	expected := map[string]bool{
		"The Bigs":                           true,
		"Killzone":                           true,
		"MLB 10 The Show":                    true,
		"Castlevania Lament of Innocence":    true,
		"Grand Theft Auto Vice City Stories": true,
	}

	buf, err := ioutil.ReadFile("phrases-from-text/test-html-unicode.html")
	if err != nil {
		fmt.Println("Error on reading file")
		fmt.Println(err)
		return
	}

	DoScrapedTitlesMatch(t, buf, expected)
}

func TestInfiniteLoop(t *testing.T) {
	expected := map[string]bool{
	// expect nothing
	}

	buf, err := ioutil.ReadFile("phrases-from-text/test-html-infinite-loop.html")
	if err != nil {
		fmt.Println("Error on reading file")
		fmt.Println(err)
		return
	}

	DoScrapedTitlesMatch(t, buf, expected)
}

func TestRepeatTitles(t *testing.T) {
	expected := map[string]bool{
		"Bases Loaded":   true,
		"Bases Loaded 3": true,
		"Bases Loaded 4": true,
	}

	buf, err := ioutil.ReadFile("phrases-from-text/test-html-repeat-titles.html")
	if err != nil {
		fmt.Println("Error on reading file")
		fmt.Println(err)
		return
	}

	DoScrapedTitlesMatch(t, buf, expected)
}
