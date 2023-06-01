package main

import "fmt"

var name = make(map[int]map[string]string)

func main() {
	name[1] = map[string]string{"ak": "ssdsd"}
	name[2] = map[string]string{"ak": "ssdsd"}
	name[3] = map[string]string{"ak": "ssdsd"}
	fmt.Println(name)
}