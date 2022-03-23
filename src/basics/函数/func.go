package main

import (
	"fmt"
)

func main() {
	calculationTest()
}

// 函数定义1：带参数且有返回值
func test1(x int, y int) int {
	return x + y
}

// 函数定义2：有参数但无返回值
func test2(x int, y int) {
	fmt.Println(x + y)
}

// 函数定义3：无参数且无返回值
func test3() {
	fmt.Println("hello world")
}

// 函数定义4：使用参数类型简写
func test4(x, y int) int {
	return x + y
}

// 函数定义5：函数返回值命名,可以在函数体中直接使用,通过return返回
func test5(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	// return 这里可以直接return,不需要指定sum和sub
	return sum, sub
}

// 函数多返回值,如果是多返回值必须使用括号
func test6(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 可变参数指的是参数数量不固定
func test7(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// 可变参数搭配固定参数时,必须放在参数最后一位
func test8(x int, y ...int) int {
	for _, v := range y {
		x += v
	}
	return x
}

// 当函数返回值类型为slice时,nil可以看做是一个有效的slice，没必要显式返回一个长度为0的切片
func test9(x string) []int {
	if x == "" {
		return nil // 没必要返回[]int{}
	}
	return []int{0}
}

// 定义函数类型calculation：接受两个int类型参数并且返回一个int类型.凡是满足这个条件的函数都是calculation类型的函数,比如Csum()/Csub(),像C#中委托的定义
type calculation func(int, int) int

// calculation类型的函数Sum
func sum(x, y int) int {
	return x + y
}

// calculation类型的函数Sub
func sub(x, y int) int {
	return x - y
}

// 函数类型测试
func calculationTest() {
	var a calculation = sum         // 将Sum赋值给函数类型变量a
	fmt.Printf("type of a:%T\n", a) // type of a:main.calculation
	fmt.Println(a(1, 2))            // 像调用sum一样调用a

	var b calculation = sub         // 将Sub赋值给函数类型变量b
	fmt.Printf("type of b:%T\n", b) // type of b:main.calculation
	fmt.Println(b(1, 2))            // 像调用sub一样调用b
}

// 将函数作为参数传递，该函数接收两个int类型变量x/y,一个函数参数sum
func functionAsArgument(x, y int, sum func(int, int) int) int {
	return sum(x, y)
}

// 按照正常调用函数的方式调用即可
func functionAsArgumentTest() {
	ret2 := functionAsArgument(10, 20, test1)
	fmt.Println(ret2) //30
}

// 接收一个切片参数patients,返回一个函数
func functionAsTheReturnValue(patients []string) func(string) bool {
	// 定义匿名函数并返回
	return func(name string) bool {
		for _, soul := range patients {
			if soul == name {
				return true
			}
		}
		return false
	}
}

func functionAsTheReturnValueTest() {
	testValue := []string{"a", "b", "c", "d", "e", "f"}
	result := functionAsTheReturnValue(testValue)
	// 调用筛选器函数获取字母是否已存在
	fmt.Println(result("a"))  // true
	fmt.Println(result("ff")) // false
}

// 匿名函数
func AnonymousFunc() {
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	// 通过变量调用匿名函数
	add(1, 2)

	// 匿名函数作为立即执行函数,一般用于匿名函数只用于一次的情况下就不需要再指定变量存储
	func(x, y int) {
		fmt.Println(x - y)
	}(1, 2)
}

// 闭包示例1
func Add() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func AddTest() {
	var f = Add()
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60
}

// 闭包示例2
func Add2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func Add2Test() {
	var f = Add2(10)
	fmt.Println(f(10)) //20
	fmt.Println(f(20)) //40
	fmt.Println(f(30)) //70
}
