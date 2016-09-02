package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {

	// buf, err := ioutil.ReadFile("phrases-from-text/test-html1.html")
	// buf, err := ioutil.ReadFile("phrases-from-text/test-html2.html")
	// buf, err := ioutil.ReadFile("phrases-from-text/test-html3.html")
	// buf, err := ioutil.ReadFile("phrases-from-text/test-html4.html")
	// buf, err := ioutil.ReadFile("phrases-from-text/test-html5.html")
	buf, err := ioutil.ReadFile("phrases-from-text/test-html6.html")
	if err != nil {
		fmt.Println("Error on reading file")
		fmt.Println(err)
		return
	}

	matches, err := GetMatches(string(buf))
	if err != nil {
		fmt.Println("error on GetMatches")
		fmt.Println(err)
		return
	}

	for _, m := range matches {
		fmt.Println(m)
	}

	return
}

// GetMatches takes the HTML document string and returns the collection of words that were found
func GetMatches(document string) ([]string, error) {
	var matches []string

	// This retrieves the phrases from the text files in string array format
	phrases, err := lookoutPhrases()
	if err != nil {
		return matches, err
	}

	// This sorts the phrases list by length firstly, and reverse alphabetical secondly.
	sort.Sort(ByLength(phrases))

	// Then transforms everything in the document passed to lowercase
	document = strings.ToLower(document)
	document = strings.Replace(document, ":", "", -1)
	document = strings.Replace(document, "-", " ", -1)
	document = strings.Replace(document, ".", "", -1)
	document = strings.Replace(document, "&nbsp;", "", -1)

	// For each phrase starting from longest to shortest
	for _, p := range phrases {
		// From the item name removes ":" and replaces dashes with empty spaces
		pp := strings.Replace(p, ":", "", -1)
		pp = strings.Replace(pp, "-", " ", -1)
		pp = strings.Replace(pp, ".", "", -1)
		pp = strings.ToLower(pp)

		// Checks if phrase is in the document
		if matchedPhraseInText(pp, document) {
			// If so appends to original phrase to the array to be returned
			matches = append(matches, p)
			// And deletes the modified phrase from the document (so as to not be matched twice or more)
			document = strings.Replace(document, pp, " ", -1)
		}
	}

	return matches, nil
}

func matchedPhraseInText(p, doc string) bool {

	// Looks for word exactly, so as to avoid subset matches
	// Avoids "super castlevania", "super c" problem discussed
	match, _ := regexp.MatchString(
		fmt.Sprintf("\\b%s\\b", p),
		doc,
	)
	return match
}

func lookoutPhrases() ([]string, error) {
	var phrases []string

	file, err := os.Open("slice-of-phrases.txt")
	if err != nil {
		return phrases, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		phrases = append(phrases, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return phrases, err
	}

	return phrases, err
}

// ByLength is a string interface implementaion to be used for sorting
type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	if len(s[i]) == len(s[j]) {
		return s[i] > s[j]
	}
	return len(s[i]) > len(s[j])
}
