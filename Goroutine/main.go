package main

import (
	"fmt"
	"time"
)

func main() {
	go count("Minh")
	count("Fish")
}

func count(name string){
	for i := 0; i < 5; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Millisecond * 1000)
	}
}