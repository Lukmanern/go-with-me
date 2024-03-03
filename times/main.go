package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("program start . . .")

	go printOut(15 * time.Second)
	go printOut(13 * time.Second)
	go printOut(11 * time.Second)

	time.Sleep(20 * time.Second)
	fmt.Println("program finish . . .")
}

func printOut(s time.Duration) {
	time.Sleep(s)
	fmt.Println("print after sleep:", s.Seconds())
}
