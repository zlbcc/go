package main

import (
	"fmt"
	"strconv"
)

func main() {
	var obs string = "321321"
	int, _ := strconv.Atoi(obs)
	//strconv.Atoi(obs)
	//var obs1 = strconv.Atoi()
	fmt.Println(int)
}
