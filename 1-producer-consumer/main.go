//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
	"sync"
)

func producer(stream Stream, prod_tweet chan *Tweet, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(prod_tweet)
			return
		}


		prod_tweet <- tweet
	}
	
}

func consumer(prod_tweet chan *Tweet, wg *sync.WaitGroup) {
	defer wg.Done()
	for  t := range prod_tweet {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	prod_tweet := make(chan *Tweet)
	// Producer
	var wg sync.WaitGroup
	wg.Add(2)
	go producer(stream, prod_tweet, &wg)

	// Consumer
	go consumer(prod_tweet, &wg)
	wg.Wait()

	fmt.Printf("Process took %s\n", time.Since(start))
}
