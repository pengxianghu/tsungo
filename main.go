package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello, this is tsungo.")
}

func makeRequest() {

}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	os.Exit(0)
}
