package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	limit := 5
	listURL := make(chan int, 200)
	results := make(chan int, 200)
	defer close(listURL)

	for i := 0; i < 200; i++ {
		listURL <- i
	}

	go crawlData(listURL, results, limit)

	for c := range results {
		fmt.Println("read value", c, "from results")
	}
	fmt.Println("Crawl finished!")
}

func crawlData(listURL chan int, results chan int, limit int) {
	defer close(results)

	if limit > len(listURL) {
		limit = len(listURL)
	}

	for i := 1; i <= limit; i++ {
		results <- <-listURL
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
	}
}
