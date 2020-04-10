package main

import "fmt"

// slice 切片
// 将拥有相同类型元素的可变长度的序列，在数组层面上做的一层封装，可以动态扩容
// 切片是一个引用类型，内部结构包含地址、长度和容量。切片一般用玉快速操作一块数据集合
func main() {
	// 声明切片类型语法
	// var name []T
	// len() 切片长度 cap() 切片容量

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("a:%v, len(a): %v, cap(a) %v \n", a, len(a), cap(a))
	// 从数组a中选出 1 <= index < 3的元素
	// len = high -low
	// cap: 底层数组容量
	// 切片信息实际就是 2,3,4,5
	s := a[1:3]
	fmt.Printf("s:%v, len(s): %v, cap(s) %v \n", s, len(s), cap(s))
	s2 := s[3:4] // [2,3,4,5] 5
	fmt.Printf("s2:%v, len(s2): %v, cap(s2) %v \n", s2, len(s2), cap(s2))

}
