package main

import "fmt"

// 1.下面的代码输出什么
/*
func main() {
	str := "hello"
	str[0] = 'x'
	fmt.Println(str)
}


编译出错，因为Go 语言中的字符串是只读的。
*/

// 2.下面这个代码输出什么

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	p := 1
	incr(&p)
	fmt.Println(p)
}
// 输出：2

// 3.对 add() 函数调用正确的是（）
func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

/*
A. add(1,2)
B. add(1,3,7)
C. add([]int{1,2})
D. add([]int{1,3,7}...)



答案：ABD，可变函数
 */

