// test
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	for v := 10; v < 100; v++ {
		v += v
		fmt.Println(v)
	}
}
