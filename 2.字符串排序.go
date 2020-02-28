package main

import (
	"fmt"
	"strings"
)

func main() {
	//比较字符串
	fmt.Println(strings.Compare("b", "c"))
	fmt.Println(strings.Compare("c", "b"))
	fmt.Println(strings.Compare("a", "a"))

	fmt.Println("~~~~~~~~")
	// go 1.3 之前不支持这么比
	fmt.Println("b" > "a")
	fmt.Println("b" < "a")

}
