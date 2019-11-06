package main

import "fmt"

func main() {

	//结构体本身就是一种数据类型，所以声明和变量声明没有什么不同；
	//只声明，并不会分配内存，实例化后才会分配内存；
	var huang person
	huang.name = "huang"
	huang.city = "上海"
	huang.age = 17
	fmt.Println(huang) //{huang 17 上海}

	//匿名结构体
	var user struct {
		name string
		age  int
	}
	user.name = "hehe"
	user.age = 23
	fmt.Printf("type:%T\n", user) //type:struct { name string; age int }

	//使用new实例化结构体，获取到的是结构体的实例的指针
	p := new(person)
	p.name = "pp"
	fmt.Printf("type:%T\n", p) //type:*main.person

	//go语言中，可以直接使用结构的指针，访问结构体的成员
	fmt.Println(p.name)    //pp
	fmt.Println((*p).name) //pp

	//也可以直接使用结构的指针设置成员
	var stu struct {
		name string
		age  int
	}
	stuPointer := &stu
	stuPointer.name = "huang"
	stuPointer.age = 17
	fmt.Println(stuPointer)  //&{huang 17}
	fmt.Println(*stuPointer) //{huang 17}

	//没有初始化的结构体，成员都是对应的零值
	var p2 person
	fmt.Println(p2) //{ 0 } 前后都是空字符串

	//使用键值对初始化
	p3 := person{name: "zhou", age: 16}
	fmt.Println(p3)

	//使用值列表初始化，要求给出所有字段值
	p4 := person{"huang", 17, "shanghai"}
	fmt.Println(p4)
}

//定义结构体
type person struct {
	name string
	age  int
	city string
}

//相同类型的字段也可以简写
type person2 struct {
	name, city string
	age        int
}
