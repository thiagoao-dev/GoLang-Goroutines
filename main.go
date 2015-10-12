package main

import "github.com/thiagoaugustus/GoLang-Goroutines/functions"

func main() {
	// Print in sequence
	functions.PrintMe(1)
	functions.PrintMe(2)
	// Treats as threads
	go functions.PrintMe(3)
	functions.PrintMe(4)
	// The main function will end earlier than goroutines
	go functions.PrintMe(5)
	go functions.PrintMe(6)
}
