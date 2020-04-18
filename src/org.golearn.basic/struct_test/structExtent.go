package main

import "fmt"

// 结构体的继承

type Animal struct {
	name string
}

func (a *Animal) move()  {
	fmt.Printf("%s move \n", a.name)
}

type Dog struct {
	Feet int8
	*Animal // 匿名嵌套（嵌套指针）
}

func (d *Dog) wang() {
	// DOG结构体并没有定义name这个熟悉，因此就去找嵌套结构体中的name熟悉
	fmt.Printf("%s wang \n", d.name)

}

func main()  {
	d := Dog{
		Feet:   4,
		Animal: &Animal{
			name:"旺财",
		},
	}

	d.wang()
	d.move()
}
