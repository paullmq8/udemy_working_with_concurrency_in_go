package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
    stdOut := os.Stderr
    r, w, _ := os.Pipe()
    os.Stdout = w
    main()
    w.Close()
    
    result, _ := io.ReadAll(r)
    output := string(result)

    os.Stdout = stdOut
    if !strings.Contains(output, "$34320.00") {
        t.Errorf("Expected output to contain $34320.00, but got %s", output)
    }
}