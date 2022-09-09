package main

import (
	"fmt"
	"strconv"
)

func main() {
    fmt.Println("hello world!")

    var i int = 1;
    fmt.Println(i)
    fmt.Printf("%v, %T \n", i, i)

    var j = strconv.Itoa(i)

    fmt.Println(j)
    fmt.Printf("%v, %T \n", j, j)

    const (
        a = iota
        b
        c
    )

    fmt.Println(a, b, c)

    //array
    arr := [3]int{1, 2, 3}

    fmt.Println(arr)

    //pointer
    p1 := [3]int{1, 2, 3}
    p2 := &p1

    p2[0] = 100

    fmt.Println(p1, p2)

    //Map
    m := map[string]int{"a": 1, "b": 2}
    fmt.Println(m)

    //Struct
    type Person struct {
        name string
        age int
    }
    Minh := Person{name: "Minh", age: 20}

    

    //Handle Panic
    Panicker()

    fmt.Println(Minh)
}

func Panicker(){
    defer func()  {
        if err := recover(); err != nil {
            fmt.Println("Error", err)
        }
    }()
    panic("Panic")
}