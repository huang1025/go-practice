package main

import (
	_ "debug/dwarf"
	"fmt"
	pkg "github.com/huang1025/go-practice/partion2/04-package"
)

func main() {
	pkg.Xixi()
	pkg.Haha()
	fmt.Println("hehe")
}

func init() {
	fmt.Println("main-init")
}
