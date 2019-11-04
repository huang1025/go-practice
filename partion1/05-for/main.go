package main

import "fmt"

func main() {

	//for基本格式；
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//for-省略初始语句
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	//for-省略初始语句和结束语句；
	j := 10
	for j > 0 {
		fmt.Println(j)
		j--
	}

	//for-无限循环
	num := 12
	for {
		fmt.Println(num)
		num--
		if num == 1 {
			break
		}
	}

	//for-range
	name := "huang_jiangling"
	for index, value := range name {
		fmt.Println(index, value)
	}
}
