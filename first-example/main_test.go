package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
    stdOut := os.Stdout
    pr, pw, _ := os.Pipe()
    os.Stdout = pw

    var wg sync.WaitGroup
    wg.Add(1)
    go printSomething("epsilon", &wg)
    wg.Wait()

    pw.Close()
    os.Stdout = stdOut
    content, _ := io.ReadAll(pr)
    result := string(content)
    if !strings.Contains(result, "epsilon") {
        t.Fatal("printSomething() did not print 'epsilon")
    }
}