package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}

	c := make(chan int)
	go sum(s[:len(s)/2], c) // sum the first half
	go sum(s[len(s)/2:], c) // sum the second half
	x, y := <-c, <-c        // receive from c

	fmt.Println(x, y, x+y)
}
