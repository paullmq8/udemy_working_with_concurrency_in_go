package main

import "testing"

func Test_updateMessage(t *testing.T) {
    msg = "Hello, world!"
    updatedMsg := "Goodbye, cruel world!"
    wg.Add(2)
    go updateMessage("x")
    go updateMessage(updatedMsg)
    wg.Wait()
    
    if msg != updatedMsg {
        t.Errorf("Expected msg to be %s, but got %s", updatedMsg, msg)
    }
}