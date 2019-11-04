package main

import "fmt"

func main() {

	//if条件判断的基本写法；
	score := 65
	if score > 90 {
		fmt.Println("good")
	} else if score > 80 {
		fmt.Println("not bad")
	} else {
		fmt.Println("bad")
	}

	//if条件判断特殊写法；
	if score := 89; score > 90 {
		fmt.Println("good")
	} else if score > 80 {
		fmt.Println("not bad")
	} else {
		fmt.Println("bad")
	}
}
