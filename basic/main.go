package main

import (
	"fmt"
	"strconv"
	"time"
)


func main(){
	fmt.Println("3+5="+strconv.Itoa(add(3,5)))
	fmt.Println("you sleep "+strconv.Itoa(sleep(3,5))+"hour")
}


func add(a, b int) int {
		return a + b
}

func addFalse(a, b int) int {
	return a + b
}

func sleep(a, b int) int {
	time.Sleep(time.Duration(a+b) * time.Second)
return a + b
}

func sleepSkip(a, b int) int {
	time.Sleep(time.Duration(a+b) * time.Second)
	return a + b
}