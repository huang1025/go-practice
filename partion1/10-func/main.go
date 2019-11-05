package main

import "fmt"

func main() {

	//函数的调用，通过函数名（）的方式调用
	sayHello()

	//可变参数调用
	intSum(1)
	intSum(1, 2)
	intSum(1, 2, 3)

	//匿名函数
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(1, 2)

	//高阶函数
	returnAdd()(1, 2)
}

//函数定义
func sayHello() {
	fmt.Println("hello world")
}

//函数参数
func sayHello2(firstName string, lastName string) {
	fmt.Println(firstName, lastName)
}

//函数参数类型简写
func sayHello3(firstName, lastName string) {
	fmt.Println(firstName, lastName)
}

//可变参数
func intSum(num ...int) {
	sum := 0
	for _, value := range num {
		sum += value
	}
	fmt.Println(sum)
}

//单个返回值
func intsum2(num ...int) int {
	sum := 0
	for _, value := range num {
		sum += value
	}
	return sum
}

//多个返回值
func intsum3(num ...int) (int, int) {
	sum := intsum2(num...)
	size := len(num)
	return sum, size
}

//返回值命名
func intsum4(num ...int) (sum, size int) {
	sum = intsum2(num...)
	size = len(num)
	return
}

//高阶函数-函数作为参数
func add(x, y int) int {
	return x + y
}
func doAdd(x, y int, add func(int, int) int) int {
	return add(x, y)
}

//高阶函数-函数作为返回值
func returnAdd() func(int, int) int {
	return add
}

//闭包
func adder(base int) func(int) int {
	return func(adder int) int {
		base += adder
		return base
	}
}
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
