//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

//editing function producer(): now it doesn't return slices from tweets,
//it accepts channel with Tweets function producer() reads data from stream and enters to the channel

func producer(stream Stream, tweets chan *Tweet) /*(tweets []*Tweet)*/ {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF{
			close(tweets)
			return
		}
		tweets<-tweet/*= append(tweets, tweet)*/
	}
}

func consumer(tweets /*[]*Tweet*/ chan *Tweet) {
	for  t := range tweets {
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

	// Producer

	tweets := /*producer(stream)*/ make(chan *Tweet)
	//after exchanging the slice to the channel, we can launch goroutine
	//so producer and consumer can run concurrently
	go producer(stream, tweets)
	// Consumer
	consumer(tweets)

	fmt.Printf("Process took %s\n", time.Since(start))
}
