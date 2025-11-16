package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 200)
	}
}

func letters() {
	for l := 'a'; l < 'j'; l++ {
		fmt.Printf("%c ", l)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	go numbers()
	go letters()
	time.Sleep(5 * time.Second)
	fmt.Printf("end of execution!\n")
}
