package main

import "fmt"

func main() {
	sn1 := struct {
		age int
		name string
	} {age: 11, name: "qq"}

	sn2 := struct {
		age int
		name string
	} {age: 11, name: "qq"}

	if sn1 ==sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string] string
	} {age: 11, m: map[string] string{"a": "1"}}

	sm2 := struct {
		age int
		m   map[string] string
	} {age: 11, m: map[string]string{"a": "1"}}

	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
}

// 以上代码能否编译通过？
/*
不能，编译不通过 invalid operation: sm1 == sm2
考察点：结构体的比较
注意的地方：
1.结构体只能比较是否相等，不能比较大小
2.相同类型的结构体才能够进行比较，结构体的属性类型要相同，属性顺序也要相同。
3.可以通过 == 或 != 进行比较是否相等
可以比较的类型：bool、数值型、字符、指针、数组等
不可以比较：切片、map、函数等
 */
