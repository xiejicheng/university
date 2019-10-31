package main

import "fmt"

func main() {
	slice := [] int{0,1,2,3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

/*
答案：
0 -> 3
1 -> 3
2 -> 3
3 -> 3

for range 循环的时候会创建每个元素的副本，而不是元素的引用，
m[key] = &val 取得都是变量 val 的地址，
因为 val 最后被赋值的值为3，输入所有输出都是3
 */