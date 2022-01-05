package main

import (
	"fmt"
	"time"
)

func chanrecv(c chan int) {
	for {
		n := <-c
		fmt.Println(n)
	}
}

func chandemo() {
	c := make(chan int)
	go chanrecv(c)
	c <- 1
	c <- 1
	time.Sleep(time.Microsecond * 101)

}

func main() {
	chandemo()
}
