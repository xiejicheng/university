package main

// 1.下面这段代码输出什么
/*
func main() {
	a := 5
	b := 8.1
	fmt.Println(a + b)
}

编译错误，a的类型是 int，b的类型是 float，两个不同类型的数值不能相加，编译报错
*/

// 2.下面代码输出什么
/*
func main() {
	a := [5] int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t[0])
}

答案：4

 */

// 3.下面这段代码输出什么
/*func main() {
	a := [2] int{5, 6}
	b := [3] int{5, 6}
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

Go 的数组是值类型，可比较，数组的长度也是数组类型的组成部分，所以a和b是不同的类型，是不能比较，答案是编译错误
}*/
