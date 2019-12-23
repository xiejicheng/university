package main

import "fmt"

// 1.下列代码 x已声明，y没有声明，判断每条语句的对错
/*
1. x, _ := f()     false   x已经声明，不能再用 :=
2. x, _ = f()      true
3. x, y := f()     true    多值赋值， := 左边变量无论声明与否都正确
4. x, y = f()      false   y未声明
 */

//2.下面代码输出什么
/*
func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func main() {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
}
// 输出：0 1
increaseA()返回参数是匿名
increaseB() 是具名
*/

//3.下面代码输出什么?

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	var a A = Work{3}
	s := a.(Work)
	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())
}

//输出：13 23