package main

import (
	"fmt"
	"org.golearn.basic/model_test"
)

func main() {
	fmt.Println(Add(1, 2))
	fmt.Println("调用其模块")
	fmt.Println(model_test.Add_x(1, 2))
}
