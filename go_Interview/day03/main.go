package main

import "fmt"

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

/*func main() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
}*/

/*
两段代码分别输出：
[0 0 0 0 0 1 2 3]
[1 2 3 4]

append 向 slice 添加元素
 */
