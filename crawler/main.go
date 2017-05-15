package main

import (
	"fmt"
	"os"
	"time"
)

const usage = `
usage:
	crawler <starting-url>
`

func worker(linkq chan string, resultsq chan []string) {
	for link := range linkq {
		plinks, err := getPageLinks(link)
		if err != nil {
			fmt.Printf("ERROR fetching %s: %v\n", link, err)
			continue
		}
		fmt.Printf("%s (%d links)\n", link, len(plinks.Links))
		time.Sleep(time.Millisecond * 500)
		if len(plinks.Links) > 0 {
			// create an inline go routine function. separate from the worker, only job is to write results to resultsq
			// unblocks the worker
			go func(links []string) {
				resultsq <- plinks.Links
			}(plinks.Links)
		}

	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}
	// in this wrong version, all of the workers are trying to write to seen map
	nWorkers := 1000
	linkq := make(chan string, 10000)
	resultsq := make(chan []string)
	for i := 0; i < nWorkers; i++ {
		go worker(linkq, resultsq)
	}

	linkq <- os.Args[1]

	seen := map[string]bool{}
	for links := range resultsq {
		for _, link := range links {
			if !seen[link] {
				seen[link] = true
				linkq <- link
			}
		}
	}

}
