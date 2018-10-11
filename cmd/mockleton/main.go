package main

import (
	"fmt"
	"github.com/troykinsella/mockleton"
	"os"
)

func main() {
	m := mockleton.New()

	exitCode, err := m.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "mockleton: %s", err.Error())
		os.Exit(255)
	}

	os.Exit(exitCode)
}
