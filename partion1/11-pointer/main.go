package main

import "fmt"

func main() {
	a := 10
	modify1(a)
	fmt.Println(a) //10
	modify2(&a)
	fmt.Println(a) //20

	s := []int{2: 12}
	modifySlice1(s)
	fmt.Println(s)
	modifySlice2(&s)
	fmt.Println(s)

	//new:开辟一块指定类型的内存空间，并且返回该空间的地址；
	n := new(int)
	fmt.Println(n)  //fmt.Println(n)
	fmt.Println(*n) //0
}

func modify1(a int) {
	a += a
}
func modify2(a *int) {
	*a = 2 * *a
}

func modifySlice1(s []int) {
	s = []int{2: 22}
}
func modifySlice2(s *[]int) {
	*s = []int{2: 22}
}
