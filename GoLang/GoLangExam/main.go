package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
)

/*
============1==============

(1) 2 , 3
    1 , 3
	3 , 3

(2) 10 20 30
    dead lock

(3) recoverd: trying to close a nil channel

==============2=================

(1)

*/

func main() {
	question := flag.Int("question", 3, "question to run: 2 or 3")
	workers := flag.Int("workers", 3, "number of Question 3 workers")
	flag.Parse()
	if *question == 2 {
		runQuestion2Examples()
		return
	}
	if *question != 3 {
		fmt.Fprintln(os.Stderr, "question must be 2 or 3")
		os.Exit(1)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	if err := runQuestion3Demo(ctx, *workers); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
