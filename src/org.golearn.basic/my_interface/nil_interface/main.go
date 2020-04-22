package main

import "fmt"

// 定义的空借口可以是任意一种类型
func main()  {
	// 空借口作为map的value（value可以使任意类型）
	var m = make(map[string]interface{}, 16)
	m["name"] = "123"
	m["age"] = 10
	m["hobby"] = []string{"a", "b", "c"}
	// 类型断言
	var x interface{}
	x = "hello"
	x = 100
	x = false

	// 类型断言
	switch v := x.(type) {
	case string:
		fmt.Printf("是字符串类型的值，value:%#v\n", v)
	case bool:
		fmt.Printf("是布尔类型的值，value:%#v\n", v)



	}


}