package main

import "fmt"

/*
1.关于字符串连接，下面语法正确的是？
A. str := 'abc' + '123'
B. str := "abc" + "123"
C. str := '123' + "abc"
D. fmt.Printf("abc%d", 123)



答案：BD.
字符串连接的方式还有：strings.Join()、buffer.WriteString()
 */

// 2.下面这段代码能否编译通过？如果可以，输出什么？
const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func main() {
	fmt.Println(x, y, z, k, p)
}
// 输出 0 2 zz zz 5      知识点：iota的使用

/*
3.下面赋值正确的是？
A. var x = nil
B. var x interface{} = nil
C. var x string = nil
D. var x error = nil
 */
/*


答案 ：BD.
nil 只能赋值给指针、chan、func、interface、map或slice类型的变量
D选项是一种内置接口类型
 */