package main

import "fmt"

func main() {

	//方法的接受者不管是类型还是类型的指针，类型的实例和类型实例的指针都能调用
	stu := student{name: "huang"}
	stu.study()
	(&stu).study()
	tea := teacher{name: "zhou"}
	tea.teach()
	(&tea).teach()

	//方法的接受者是类型，那么调用者会被复制一份到方法的接受者参数中，方法不论怎样都改变不了原始值
	stu2 := new(student)
	stu2.name = "hehe"
	stu2.changeName("huang")
	fmt.Println(stu2) //&{hehe }
	(*stu2).changeName("huang_jiangling")
	fmt.Println(stu2) //&{hehe }

	//当方法的接受者是类型的指针，那么调用者的指针会被拷贝到方法接受者参数中，方法中对调用者的更改将被保留
	tea2 := new(teacher)
	tea2.name = "haha"
	tea2.changeName("zhou")
	fmt.Println(tea2) //&{zhou 0}
	(*tea2).changeName("zhou_haiyan")
	fmt.Println(tea2) //&{zhou_haiyan 0}
}

type student struct {
	name, class string
}

func (s student) study() {
	fmt.Println(s.name + "学习...")
}

func (s student) changeName(name string) {
	s.name = name
}

type teacher struct {
	name string
	age  int
}

func (t *teacher) teach() {
	fmt.Println(t.name + "teach...")
}

func (t *teacher) changeName(name string) {
	t.name = name
}
