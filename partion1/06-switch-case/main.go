package main

import "fmt"

func main() {

	switchDemo1(1)
	switchDemo1(5)

	switchDemo2(2)
	switchDemo2(4)
	switchDemo2(7)

	switchDemo3(1)
	switchDemo3(-1)

	switchDemo4(1)
}

//switch-case
func switchDemo1(num int) {
	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("hehe")
	}
}

//一个分支多个值；
func switchDemo2(num int) {
	switch num {
	case 1, 2, 3:
		fmt.Println("hehe")
	case 4, 5, 6:
		fmt.Println("haha")
	default:
		fmt.Println("cici")
	}
}

//分支使用表达式；
//此时switch省略变量值；
func switchDemo3(num int) {
	switch {
	case num < 0:
		fmt.Println("less than 0")
	default:
		fmt.Println("great than 0")
	}
}

//fallthrough
//fallthrough所在分支满足条件，则后面的分支不论是否满足，必定会执行；
func switchDemo4(num int) {
	switch num {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("hehe")
	}
}
