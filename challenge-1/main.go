package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 27, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	msg = "Hello, world!"

	arr := []string{"Hello, universe!", "Hello, cosmos!", "Hello, world!"}
	var wg sync.WaitGroup
	for _, v := range arr {
		wg.Add(1)
		go func(msg string, wg *sync.WaitGroup) {
			defer wg.Done()
			updateMessage(msg)
			// time.Sleep(1 * time.Second)
			printMessage()
		} (v, &wg)
	}
	wg.Wait()

/* 	updateMessage("Hello, universe!")
	printMessage()

	updateMessage("Hello, cosmos!")
	printMessage()

	updateMessage("Hello, world!")

	printMessage() */
}
