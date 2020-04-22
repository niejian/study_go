package main

import "fmt"

// 接口：抽象的结构体

// 为什么要使用接口
type dog struct {

}
// 方法
func (d dog) say()  {
	fmt.Println("旺旺旺...")
}

type cat struct {

}
func (c cat) say()  {
	fmt.Println("喵喵喵...")
}

// 接口的定义
// 一个抽象类型，只要实现了say()方法的类型就是sayer类型
type sayer interface {
	say()
}
// 接口不管是什么类型，只关注应该实现什么方法
// 打的函数
func da(arg sayer)  {
	arg.say()
}

func main()  {

	c1 := cat{}
	da(c1)
	d1 := dog{}
	da(d1)

}