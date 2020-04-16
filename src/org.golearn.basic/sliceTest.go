package main

import (
	"fmt"
	"sort"
)

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

	// 完整的切片表達式
	// a[low: high: max] max不填就是數組a的size

	// make() 动态创建切片
	// make([]T, size, cap)
	fmt.Println("动态构造切片 make()====")
	// 给切片b分配了10哥空间实际只是用2个初始化为0
	b := make([]int, 2, 10)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
	// 切片的赋值
	b[0] = 10
	fmt.Println(b)

	fmt.Println("===切片的遍历===")
	// 切片的遍历
	for index, r := range b {
		fmt.Printf("下标：%v, 数值：%v \n", index, r)
	}

	// 为切片添加元素
	var s1 []int
	// 数组中添加元素
	s1 = append(s1, 1)
	s1 = append(s1, 2, 3, 4, 5)
	fmt.Println(s1)
	s3 := []int{5, 6, 7}
	s1 = append(s1, s3...)
	fmt.Println(s1)
	// 查看切片的扩容情况
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)

	}

	fmt.Println("==copy 复制切片==")
	d := []int{1, 2, 3, 4}
	e := make([]int, 4, 4)
	fmt.Println(e)
	copy(e, d)
	fmt.Println(d)
	fmt.Println(e)
	d[0] = 10
	fmt.Println(d[:2])
	fmt.Println(e)

	var f = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		f = append(f, fmt.Sprintf("%v", i))

	}
	fmt.Println(f)

	var g = [...]int{3, 7, 8, 4, 4, 13, 57, 12}
	var h = make([]int, len(g), len(g))
	copy(h, g)
	ls := sort.IntSlice{
		3, 7, 8, 4, 4, 13, 57, 12,
	}

	fmt.Println(ls)
	sort.Ints(ls)
	fmt.Println(ls)

}
