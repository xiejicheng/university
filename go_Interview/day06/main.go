package main

import "fmt"

/*
1.通过指针变量 p 访问其成员变量 name, 有哪几种方式？
A. p.name
B. (&p).name
C. (*p).name
D. p->name
 */
/*
答案是 AC , & 取址运算符，* 指针解引用
 */

// 2.以下代码能否通过编译？如果通过，输出什么？
type MyInt1 int
type MyInt2 = int

func main() {
	var i int = 0
	var i1 MyInt1 = i
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}
/*
编译不通过，cannot use i (type int) as type MyInt1 in assignment
考察类型别名与类型定义的区别
17行代码是基于类型 int 创建了新类型 MyInt1
18行代码是创建了 int 类型变量的别名 MyInt2

22行代码相当于将 int 类型的变量赋值给 MyInt1 类型的变量，Go 是强类型语言，编译会不通过
23行的MyInt2 是 int 的别名，本质还是 int, 因此可以赋值

22行代码可以使用强制类型转化 var i1 MyInt1 = MyInt1(i)
 */