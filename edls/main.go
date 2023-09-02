package main

import (
	"flag"
	"fmt"
)

func main() {
	// Filter patterns flags
	flagPattern := flag.String("p", "", "Filter by pattern")
	flagAll := flag.Bool("a", false, "Sort all files, including hidden files")
	flagNumberRecord := flag.Int("n", 0, "Sort a spesific number of files")

	// Order Flags
	hasOrderByTime := flag.Bool("t", false, "Order files by time, older first")
	hasOrderBySize := flag.Bool("s", false, "Order files by size, smallest first")
	hasOrderReverse := flag.Bool("r", false, "Reverse order while sorting")

	// Function for index the flags
	flag.Parse()
	fmt.Println("pattern:", *flagPattern)
	fmt.Println("All:", *flagAll)
	fmt.Println("Records:", *flagNumberRecord)
	fmt.Println("OByTime:", *hasOrderByTime)
	fmt.Println("OBySize:", *hasOrderBySize)
	fmt.Println("OReverse:", *hasOrderReverse)
}
