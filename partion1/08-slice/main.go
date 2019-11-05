package main

import "fmt"

func main() {

	//切片的声明
	var s1 []string //nil 数组声明但未初始化，则为一个nil；

	//切片的声明并初始化
	var s2 []int = []int{} //[] 数组声明并初始化空值，那么将是一个空数组；
	var s3 = []int{1, 2, 3}

	//从数组获取切片
	//[startIndex,endIndex]是一个左闭右开区间；
	// startIndex不写表示为0，endIndex不写表示为数组长度；
	//最终切片的size=endIndex-startIndex，切片容量=数组长度-startIndex；
	a := [5]int{55, 56, 57, 58, 59}
	s4 := a[1:4] //[56 57 58]
	s5 := a[1:]  //[56 57 58 59]
	s6 := a[:4]  //[55 56 57 58]
	s7 := a[:]   //[55 56 57 58 59]

	//切片在切片
	//因为基于同一个底层数组，所以看似丢掉的数据还能别找回来；
	a2 := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	fmt.Printf("a2:%v type:%T len:%d  cap:%d\n", a2, a2, len(a2), cap(a2)) //a2:[北京 上海 广州 深圳 成都 重庆] type:[6]string len:6  cap:6
	s8 := a2[1:3]
	fmt.Printf("s8:%v type:%T len:%d  cap:%d\n", s8, s8, len(s8), cap(s8)) //s8:[上海 广州] type:[]string len:2  cap:5
	s9 := s8[1:5]
	fmt.Printf("s9:%v type:%T len:%d  cap:%d\n", s9, s9, len(s9), cap(s9)) //s9:[广州 深圳 成都 重庆] type:[]string len:4  cap:4

	//使用make创建切片
	s10 := make([]int, 2, 10)
	fmt.Println(s10)      //[0 0]
	fmt.Println(len(s10)) //2
	fmt.Println(cap(s10)) //10

	//切片遍历
	for index, value := range s10 {
		fmt.Println(index, value)
	}

	//append:为切片增加元素，操作的是底层数组，如果切片从数组而来，那么之前的数组也将受影响；
	//增加元素时，长度增加容量不变，如果长度达到容量，那么将触发扩容；
	//可以一次添加多个元素；
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
	numSlice = append(numSlice, 1, 2, 3)
	fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)

	//从数组而来的切片，增加元素，将影响原数组；
	aa := [...]int{1, 2, 3}
	bb := aa[:2]
	bb = append(bb, 4)
	fmt.Println(aa, bb)

	//使用copy复制切片
	//切片直接赋值给另一个切片只是引用的赋值，而该函数可以复制底层数组；
	//并且拷贝只拷贝目标切片长度个元素；
	co := make([]int, 3, 5)
	copy(co, numSlice)
	fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", co, len(co), cap(co), co)

	//从切片中删除元素
	co = append(co[:2], co[3:]...)
	fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", co, len(co), cap(co), co)

	fmt.Println(s1, s2, s3, s4, s5, s6, s7)
}
