package main

import "fmt"

func main() {
	fmt.Println("gitflow-demo")
	// 这里打印的是hello，而非hello world
	// 这个时候并没有合并分支，之后突然线上出现一个Bug，需要立即修复bug
	fmt.Println("hello World")
}
