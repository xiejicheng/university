package main

// 下面三段代码能否通过编译，不能的话原因是什么；如果能，输出什么。

/*
不能通过编译，
 */
/*func main() {
	list := new([]int)
	list := append(list, 1)
	fmt.Println(list)
}*/

/*func main() {
	s1 := []int{1,2,3}
	s2 := []int{4,5}
	s1 := append(s1, s2)
	fmt.Println(s1)
}*/

/*
var (
	size := 1024
	max_size = size * 2
)

func main() {
	fmt.Println(size, max_size)
}
 */
