package main

import (
	"fmt"
	"reflect"
)

// 变量声明；
var a1 string
var a2 int
var a3 bool

// 变量批量声明；
var (
	b1     float32
	b2     string
	b3, b4 int
)

// 变量的声明与初始化；
var c1 string = "hehe"
var c2 int = 17
var c3 bool = true

// 类型推导；
var (
	d1 = 12     //int
	d2 = 12.3   //float64
	d3 = "hehe" //string
	d4 = true   //bool
)

// 变量声明简写(只能在函数内部使用)
//f1:=12

// 匿名变量（匿名变量不占命名空间，不会分配内存，所以也不存在重复声明）
//在使用多重赋值时，如果不想要某个值，使用匿名变量，用下划线表示；
var _, name = foo()

func main() {
	fmt.Println(a1, a2, a3)
	fmt.Println(b1, b2, b3, b4)
	fmt.Println(c1, c2, c3)
	fmt.Println(typeof(d1), typeof(d2), typeof(d3), typeof(d4))

	// 变量声明简写(只能在函数内部使用)
	f1 := 12
	f2 := "hehe"
	fmt.Println(f1, f2)

	// 匿名变量（多重赋值下使用）
	g1, g2 := foo()
	_, g3 := g1, g2
	fmt.Println(g3) //huang

}

func typeof(p interface{}) string {
	return reflect.TypeOf(p).String()
}

func foo() (int, string) {
	return 17, "huang"
}
