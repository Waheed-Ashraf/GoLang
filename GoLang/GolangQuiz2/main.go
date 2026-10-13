package main

import (
	"fmt"
	"sync"
	"time"
)

// Real-life analogy: comparing prices for a phone across several online stores.
//   - Instead of opening each store one by one (slow), we ask ALL stores at once.
//   - Each store lookup is a goroutine.
//   - Answers come back through a channel.
//   - WaitGroup tells us when every store has replied.

type Quote struct {
	Store string
	Price float64
}

type Store struct {
	Name    string
	Price   float64
	Latency time.Duration // pretend network delay
}

func checkStore(s Store, results chan<- Quote, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(s.Latency) // simulate calling the store's website
	fmt.Printf("%-10s replied: $%.2f\n", s.Name, s.Price)

	results <- Quote{Store: s.Name, Price: s.Price}
}

func Check() {
	stores := []Store{
		{"Amazon", 699.99, 700 * time.Millisecond},
		{"eBay", 649.50, 400 * time.Millisecond},
		{"BestBuy", 679.00, 900 * time.Millisecond},
		{"Walmart", 659.99, 500 * time.Millisecond},
		{"Target", 689.00, 300 * time.Millisecond},
	}

	results := make(chan Quote, len(stores))
	var wg sync.WaitGroup

	start := time.Now()

	//Ask every store at the same time
	for _, s := range stores {
		wg.Add(1)
		go checkStore(s, results, &wg)
	}

	// When all stores have answered, close the channel
	go func() {
		wg.Wait()
		close(results)
	}()

	//Collect answers and track the cheapest one
	var best Quote
	for q := range results {
		if best.Store == "" || q.Price < best.Price {
			best = q
		}
	}

	fmt.Println("-----------------------------------")
	fmt.Printf("Best deal: %s at $%.2f\n", best.Store, best.Price)
	fmt.Printf("Checked %d stores in %v\n", len(stores), time.Since(start).Round(100*time.Millisecond))
	fmt.Println("(One by one it would have taken about 2.8s)")
}
