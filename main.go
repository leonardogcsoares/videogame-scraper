package main

import (
	"fmt"
	"io/ioutil"
	"videogame-scraper/scrape"
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

	matches, err := scrape.GetMatches(string(buf))
	fmt.Print("\n\n\n")
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
