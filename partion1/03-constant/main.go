package main

import "fmt"

// 常量声明和变量类似，但是常量在定义时，必须赋值；
//const a1
const a2 int = 12
const a3 = "hehe"

// 批量声明变量；
const (
	b1 string = "haha"
	b2        = 12
	b3        = 1.23
)

// 批量声明变量时，如果省略了值，则表示和上面一行的值相同；
const (
	c1 = 100
	c2 //100
	c3 //100
	c4 = 10
	c5 //10
	c6 //10
	//c7,c8 当两个变量声明在同一行，并且都省略了值，编译将不通过；
	c7, c8 = 12, 12
)

// iota;
// 在const关键字出现时被重置为0；
// const中每新增一行常量声明，将使iota计数一次，iota可以理解为const语句块中的行索引；
const (
	d1 = iota //0
	d2        //1
	d3        //2
)

// 使用匿名变量跳过某些值，iota也会进行计数；
// 不管有没有使用iota，const中每出现一行iota都会计数；
const (
	e1 = 1    //1
	e2        //1
	e3 = iota //2
	_
	e4 //4
)

// 多个变量
const (
	f1, f2 = iota, iota //0,0
	f3, f4              //1,1
	f5, f6              //2,2
)
const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

func main() {
	fmt.Println(c1, c2, c3, c4, c5, c6)
	fmt.Println(d1, d2, d3)
	fmt.Println(e1, e2, e3, e4)
	fmt.Println(f1, f2, f3, f4, f5, f6)
	fmt.Println(a, b, c, d, e, f)
}
