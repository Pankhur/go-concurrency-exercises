//////////////////////////////////////////////////////////////////////
//
// Your task is to change the code to limit the crawler to at most one
// page per second, while maintaining concurrency (in other words,
// Crawl() must be called concurrently)
//
// @hint: you can achieve this by adding 3 lines
//

package main

import (
	"fmt"
	"sync"
	"time"
)

//Create a buffer channel
var sem = make(chan struct{}, 1)

// Crawl uses `fetcher` from the `mockfetcher.go` file to imitate a
// real crawler. It crawls until the maximum depth has reached.
func Crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	//Implement semaphore concept
	sem <- struct{}{}
	body, urls, err := fetcher.Fetch(url)
	//Since it's mentioned one page per second so we have to introduce second as well
	time.Sleep(time.Second)
	<-sem
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	wg.Add(len(urls))
	
	for _, u := range urls {
		
		// Do not remove the `go` keyword, as Crawl() must be
		// called concurrently
		go Crawl(u, depth-1, wg)
		
	}
	return
}

func main() {
	var wg sync.WaitGroup
	//sem := make(chan struct{})
	
	wg.Add(1)
	//sem <- struct{}{}
	Crawl("http://golang.org/", 4, &wg)
	//<-sem
	wg.Wait()
}
