package main

import "fmt"

//正确写法
func main() {
	slice := [] int{0,1,2,3}
	m := make(map[int]*int)

	for key, val := range slice {
		value := val
		m[key] = &value
	}

	for k, v := range m {
		fmt.Println(k, "-->", *v)
	}
}
