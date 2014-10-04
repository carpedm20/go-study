package main

import "fmt"

var student = map[int]Person {
    0: Person{"carpedm20", 23},
    1: Person{"tunz", 22},
}

func main() {
    fmt.Println(student)
}
