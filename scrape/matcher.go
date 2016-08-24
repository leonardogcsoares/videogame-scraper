package scrape

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

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
