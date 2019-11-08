package main

import (
	"fmt"
)

func main() {

	//将dog类型的指针 赋值给接口类型；
	//方法的接受者是对象类型，需要源对象的一个拷贝，现在能拿到一个源对象的指针的拷贝，通过这个指针可以拿到源对象，在拷贝一份就可以了；
	d := new(dog)
	d.name = "旺财"
	var dogDayer sayer = d
	dogDayer.say()

	//将dog类型 赋值给接口类型
	d2 := *d
	dogDayer = d2
	d2.say()

	//方法的接受者是类型的指针，将类型的指针 赋值给接口，并调用方法；
	c := new(cat)
	c.name = "小喵"
	var catSayer sayer = c
	catSayer.say()

	//方法的接收者是类型的指针，则不能将类型的对象赋值给接口；
	//原因很简单，类型的对象调用方法时，将类型的对象复制一份，而方法的接收者是指针，复制过去的对象不可能拿到原对象的指针；
	//c2 := *c
	//catSayer = c2 //Type does not implements sayer,method had a pointer receiver

	//接口嵌套
	d3 := new(dog)
	d3.name = "旺财3"
	var s1 sayer = d3
	s1.say()
	var m1 mover = d3
	m1.move()
	var a animal = d3
	a.say()
	a.move()

	//空接口，可以接受任意类型参数
	m := map[string]interface{}{}
	m["name"] = "huang"
	m["age"] = 17
	fmt.Println(m) //map[age:17 name:huang]

	//类型断言，返回的第一个参数是断言成功后转换成断言类型，第二个参数是断言是否成功
	var name interface{}
	name = 12
	value, ok := name.(string)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("断言失败，value：" + value)
	}

	//使用switch-case进行多次断言
	switch name.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("other")
	}
}

type sayer interface {
	say()
}

type dog struct {
	name string
}

func (d dog) say() {
	fmt.Println(d.name + ":汪汪汪")
}

func (d dog) move() {
	fmt.Println(d.name + ":跳来跳去")
}

type cat struct {
	name string
}

func (c *cat) say() {
	fmt.Println(c.name + ":喵喵喵")
}

type mover interface {
	move()
}

type animal interface {
	sayer
	mover
}
