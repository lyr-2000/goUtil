package main

import (
	"fmt"
	"goUtil/util"
)
func main() {
	var str string = "hello_world_App"
	//s :=util.Upfirst( (str))
	fmt.Println(util.SnakeToWord(str))

	//fmt.Println(str)
	//fmt.Println(s)
	//fmt.Println(util.LowFirst(s))
	//
	//fmt.Println(util.GetRandomString(6))
}
