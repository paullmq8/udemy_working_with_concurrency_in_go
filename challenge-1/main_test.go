package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
    localMsg := "Hello, universe!"
    updateMessage(localMsg)
    if msg != localMsg {
        println("Test_updateMessage failed")
        t.Errorf("Expected %s, got %s", localMsg, msg)
    }
}

func Test_printMessage(t *testing.T) {
    stdOut := os.Stdout
    pr, pw, _ := os.Pipe()
    os.Stdout = pw

    localMsg := "Hello, universe!"
    updateMessage(localMsg)
    printMessage()
    pw.Close()
    os.Stdout = stdOut
    content, _ := io.ReadAll(pr)
    if string(content) != localMsg + "\n" {
        t.Errorf("Expected %s, got %s", localMsg, string(content))
    }
}

func Test_main(t *testing.T) {
	// arr := []string{"Hello, universe!", "Hello, cosmos!", "Hello, world!"}
    stdOut := os.Stdout
    pr, pw, _ := os.Pipe()
    os.Stdout = pw
    main()
    pw.Close()
    os.Stdout = stdOut
    content, _ := io.ReadAll(pr)
    res := string(content)
    res_slice := strings.Split(res[:len(res) - 1], "\n")
    for _, v := range res_slice {
        if v != "Hello, universe!" && v != "Hello, cosmos!" && v != "Hello, world!" {
            t.Errorf("Expected %s, got %s", "Hello, universe! or Hello, cosmos! or Hello, world!", v)
        }
    }
}