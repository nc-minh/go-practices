package main

import (
	"fmt"
	"test/singleton"
	"time"
)

func main() {
	for i := 0 ; i < 10 ; i++ {
		go func ()  {
			fmt.Printf("%p \n", singleton.GetInstance())
		}()
	}

	time.Sleep(time.Second * 10)
}