package main

import "fmt"

//map的声明
var map1 map[string]int

//map的声明并初始化
var map2 map[string]int = map[string]int{"huang": 17, "zhou": 16}
var map3 = map[string]int{"huang": 17, "zhou": 16}

func main() {

	//使用make获取map
	map4 := make(map[string]int, 8)
	fmt.Println(map4)

	//设置map的值
	map4["xiaohuang"] = 13

	//从map中获取值
	age := map4["xiaohuang"]
	fmt.Println(age)

	//判断指定key是否存在
	//即使不存在对应key值，也能获取到对应类型的零值
	value, ok := map4["huang"]
	fmt.Println(value, ok) //0 false

	//使用range遍历map
	//适应两个参数接收时，获取key-value，使用一个参数接收时，返回key
	for key, value := range map4 {
		fmt.Println(key, value)
	}
	for key := range map4 {
		fmt.Println(key, map4[key])
	}

	//删除指定key
	map4["huang"] = 14
	fmt.Println(map4)
	delete(map4, "huang")
	fmt.Println(map4)

	//元素为map类型的切片
	var mapSlice = make([]map[string]string, 8)
	mapSlice[0] = make(map[string]string, 3)
	mapSlice[0]["name"] = "huang"
	mapSlice[0]["age"] = "17"
	fmt.Println(mapSlice)

	//值为切片类型的map
	var sliceMap = make(map[string][]int, 3)
	sliceMap["huang"] = make([]int, 3)
	sliceMap["huang"][1] = 12
	fmt.Println(sliceMap)
}
