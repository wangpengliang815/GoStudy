// @title 《Go语言编程》 -使用sync包
// @description
// @author wangpengliang
// @date 2022-03-31 17:25:12

package main

import (
	"fmt"
	"sync"
)

// 声明全局等待组变量
var wait sync.WaitGroup

func printHello() {
	fmt.Println("hello")
	wait.Done()
}

// 使用sync包
func run() {
	wait.Add(1)
	go printHello()
	fmt.Println("end...")
	wait.Wait()
}

func main() {
	run() // 先输出:end... 在输出:hello  存在随机性
}
