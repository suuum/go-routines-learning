package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, tweets chan *Tweet, quit chan bool) {

	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			quit <- true
			break
		}

		tweets <- tweet
	}
}

func consumer(tweet chan *Tweet, quit chan bool) {
	for {
		select {
		case t := <-tweet:
			{
				if t.IsTalkingAboutGo() {
					fmt.Println(t.Username, "\ttweets about golang")
					continue
				}

				fmt.Println(t.Username, "\tdoes not tweet about golang")
			}
		case <-quit:
			return
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	// Modification starts from here
	// Hint: this can be resolved via channels
	// Producer
	tweets := make(chan *Tweet)
	quit := make(chan bool)
	go producer(stream, tweets, quit)
	// Consumer
	consumer(tweets, quit)

	fmt.Printf("Process took %s\n", time.Since(start))
}
