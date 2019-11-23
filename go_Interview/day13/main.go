package main

/*
1.定义一个包内全局字符串变量，下面语法正确的是？
A. var str string
B. str := ""
C. str = ""
D. var str = ""


答案：AD
B 只支持局部变量声明
C是赋值，str必须在这之前已经声明
 */

// 2.下面这段代码输出什么？
/*
func hello(i int) {
	fmt.Println(i)
}

func main() {
	i := 5
	defer hello(i)
	i = i + 10
}
//答案 ：5
hello() 函数在执行 defer 语句的时候会保存一份副本，在实际调用 hello() 函数时用
 */

// 3.下面代码会输出什么?
/*
type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}
输出：
showA
showB

结构体嵌套
 */