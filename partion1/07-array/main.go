package main

import "fmt"

//定义数组
var arr1 [2]int
var arr2 [3]string

func main() {

	//数组初始化
	var arr3 [3]int = [3]int{1, 2, 3}
	var arr4 = [3]int{1, 2, 3}

	//自行推断数组长度
	var arr5 = [...]int{1, 2, 3}

	//使用数组索引初始化数组
	//数组的长度为最大索引值+1
	var arr6 = [...]int{1: 1, 10: 10}

	//数组遍历
	for index, value := range arr6 {
		fmt.Println(index, value)
	}

	//多维数组
	//最后一个元素后面的逗号也是不可省略的；
	arr7 := [2][3]int{
		{2, 3, 4},
		{4, 5, 6},
	}

	//多维数组，只有第一层的长度是可以自动推导的；
	arr8 := [...][3]int{
		{2, 3, 4},
		{4, 5, 6},
	}

	//多维数组遍历
	for i, v1 := range arr8 {
		for j, v2 := range v1 {
			fmt.Println(i, j, v2)
		}
	}

	fmt.Println(arr3, arr4, arr5, arr6, arr7)
}
