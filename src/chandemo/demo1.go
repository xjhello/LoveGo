package main

import (
	"fmt"
	"time"
)

func chanrecv1(id int, c chan int) {
	for {
		n := <-c
		fmt.Println(id, n)
	}
}

func chandemo1() {
	//c := make(chan int)
	//go chanrecv1(1, c)
	//c <- 1
	//c <- 12121
	//time.Sleep(time.Microsecond * 101)

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		c1 <- 1
		fmt.Println(1)
		c1 <- 1
		fmt.Println(1)
		c1 <- 1
		fmt.Println(1)
	}()

	go func() {
		c2 <- 2
		c2 <- 2
		c2 <- 2
	}()

	select {
	case n := <-c1:
		fmt.Println("c1", n)
	case n := <-c2:
		fmt.Println("c2", n)
	}

	time.Sleep(time.Microsecond * 100)

}

func main() {
	chandemo1()
}
