package main

import "fmt"

// 使用值接收者实现接口和使用指针接受者实现接口的区别

// 接口嵌套
type animal interface {
	mover //move()
}
type mover interface {
	move()
}

type person struct {
	name string
	age int8
}

//func (p person) move()  {
//	fmt.Printf("%s在跑...\n", p.name)
//}

func (p *person) move()  { // 指针类型接受者
	fmt.Printf("%s在跑...\n", p.name)
}

func main()  {
	var m mover
	p1 := person{
		name: "a",
		age:  10,
	}

	p2 := &person{ // person类型指针
		name: "b",
		age:  16,
	}
	// m = p1 // 值类型接受者
	m = &p1 // 指针类型接受者
	m.move()
	fmt.Println(m)
	m = p2
	m.move()
	fmt.Println(m)
}
