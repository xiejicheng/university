package main

import "fmt"

/*
1.关于 channel，下面语法正确的是（）
A. var ch chan int
B. ch := make(chan int)
C. <- ch
D. ch <-


答案：ABC
A、B都是声明channel
C 读取channel
写 channel必须带上值，所以D错误
 */

/*
2.下面这段代码输出什么？
A. 0
B. 1
C. Compilation error

type person struct {
	name string
}

func main() {
	var m map[person] int
	p := person{"mike"}
	fmt.Println(m[p])
}

答案：输出 0
打印一个 map 中不存在的值时，返回元素类型的零值，m中不存在p
*/

/*
3.下面这段代码输出什么？
A. 18
B. 5
C. Compilation error
 */
func hello(num ...int) {
	num[0] = 18
}

func main() {
	i := []int{5,6,7}
	hello(i...)
	fmt.Println(i[0])
}
/*
答案：18
可变函数
 */





