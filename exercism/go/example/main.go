package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var ball = make(chan string)

	kickBall := func(ctx context.Context, playerName string) {
	LOOP:
		for {
			var who string
			select {
			case who = <-ball:
			case <-ctx.Done():
				break LOOP
			}

			fmt.Println(who, "kicked the ball.")
			time.Sleep(time.Second)
			select {
			case ball <- playerName:
			case <-ctx.Done():
				break LOOP
			}
		}
		fmt.Println(playerName, "exited the game")
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)

	go kickBall(ctx, "John")
	go kickBall(ctx, "Alice")
	go kickBall(ctx, "Bob")
	go kickBall(ctx, "Emily")
	ball <- "referee" // kick off
	<-ctx.Done()
	time.Sleep(time.Second * 2)
}
