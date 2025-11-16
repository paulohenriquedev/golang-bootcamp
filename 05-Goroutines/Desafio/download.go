package main

import (
	"fmt"
	"time"
)

func video(c chan int) {

	for i := 0; i <= 100; i++ {
		fmt.Printf("Downloading video %d%%\n", i)
		time.Sleep(time.Millisecond * 50)
	}
	// envia o valor final pelo canal
	c <- 100

}

func music(c chan int) {

	for i := 0; i <= 100; i++ {
		fmt.Printf("Downloading music %d%%\n", i)
		time.Sleep(time.Millisecond * 50)
	}
	c <- 100
}

func main() {

	c1 := make(chan int)
	c2 := make(chan int)
	go video(c1)
	go music(c2)
	x, y := <-c1, <-c2 // receive from c

	fmt.Println(x, y)
}
