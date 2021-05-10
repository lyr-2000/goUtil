package main

import "fmt"

func main() {
	//var str string = "hello_world_App"
	//s :=util.Upfirst( (str))
	//1tln(util.Md5([]byte(str)))
	// fmt.Println(util.GetCurrentProjectPath())
	var h map[string]interface{}
	h = map[string]interface{} {"hello":"world"}
	h["hello"] = "1"
	fmt.Println(h["hello"])
	for k,v := range h {
		fmt.Println(k)
		fmt.Println(v)
	}
	//fmt.Println(str)
	//fmt.Println(s)
	//fmt.Println(util.LowFirst(s))
	//
	//fmt.Println(util.GetRandomString(6))
}
