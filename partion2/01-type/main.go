package main

import "fmt"

func main() {

	//基于内置的基本类型，定义新的类型；
	type MyInt int
	var m MyInt = 12
	fmt.Printf("type:%T\n", m) //type:main.MyInt
	fmt.Println(m)             //12

	//类型别名
	type MyIntAlias = MyInt
	type intAlias = int
	var a1 MyIntAlias
	var a2 intAlias
	fmt.Printf("type:%T\n", a1) //type:main.MyInt
	fmt.Printf("type:%T\n", a2) //type:int
}
