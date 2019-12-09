package main

// 1.切片a、b、c 的长度和容量分别是多少？
/*
func main() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c : s[1:2:cap(s)]
}
a、b、c 的长度和容量分别是 0 3、2 3、1 2。
知识点：数组或切片的截取操作。截取操作有带 2 个或者 3 个参数，形如：[i:j] 和 [i:j:k]，假设截取对象的底层数组长度为 l。
在操作符 [i:j] 中，如果 i 省略，默认 0，如果 j 省略，默认底层数组的长度，截取得到的切片长度和容量计算方法是 j-i、l-i。
操作符 [i:j:k]，k 主要是用来限制切片的容量，但是不能大于数组的长度 l，截取得到的切片长度和容量计算方法是 j-i、k-i。
 */

//2.下面代码 A B 两处应该怎么修改才能顺利编译？
/*
func main() {
	var m map[string]int  // A
	m["a"] = 1
	if v := m["b"]; v != nil {  //B
		fmt.Println(v)

	}
}

A处 使用mack()初始化 map
m := make(map[string]int)
B处  if v,ok := m["b"]; ok
当 key 为 b 的元素不存在的时候，v 会返回值类型对应的零值，k 返回 false。
*/


