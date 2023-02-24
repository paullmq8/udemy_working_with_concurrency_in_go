package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
    defer wg.Done()
    msg = s
}

func main() {
    msg = "Hello, World!"
    wg.Add(2)
    go updateMessage("Hello, universe!")
    go updateMessage("Hello, cosmos!")
    wg.Wait()
    fmt.Println(msg)
}

/*
func updateMessage(s string, mu *sync.Mutex) {
    defer wg.Done()
    mu.Lock()
    msg = s
    mu.Unlock()
}

func main() {
    msg = "Hello, World!"
    wg.Add(2)
    var mutex sync.Mutex
    go updateMessage("Hello, universe!", &mutex)
    go updateMessage("Hello, cosmos!", &mutex)
    wg.Wait()
    fmt.Println(msg)
}
*/