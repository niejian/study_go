package main

import "fmt"

// go中的结构体
/*
 type 类名 struct {
	字段名 字段类型
	字段名 字段类型
	字段名 字段类型
	...
	}
*/

type person struct {
	name string
	city string // 同一类型的字段可以写一行
	age int8
}

// 定义结构体方法（java中类的成员方法）
// func (结构体变量名 结构体) 方法名
func (p person) Dream()  {
	fmt.Printf("%s的梦想是xxx \n", p.name)
}

// 修改成员变量(指针类型)
func (p *person) SetAge(age int8)  {
	p.age = age
}

// go方法是值拷贝的，并不会改变对象p中的age的值
func (p person) SetAge2(age int8)  {
	p.age = age
}

// 嵌套结构体 对象里面包含对象
type Student struct {
	Name string
	age int8
	Gender string
	Address Address
	Email Email
}
type Address struct {
	Province string
	City string
	UpdateTime string
}
type Email struct {
	Addr string
	UpdateTime string
}



func main()  {
	// 结构体的实例化
	var p person
	p.name = "zhangsan"
	p.city = "gz"
	p.age = 20
	fmt.Printf("p = %v \n", p)
	fmt.Printf("p = %T \n", p)
	// 键值对初始化
	p3 := person{
		name: "a",
		city: "b",
		age:  0,
	}

	// 匿名结构体, 定义在方法中中结构体
	var user struct{name string; age int}
	user.name = "lisi"
	user.age = 10
	fmt.Printf("匿名结构体 user = %v \n", user) //匿名结构体 user = {lisi 10}

	// 指针类型结构体
	var p2 = new(person)
	fmt.Printf("%T\n", p2)

	fmt.Printf("p2 = %#v\n", p2) // &main.person{name:"", city:"", age:0}
	// 去结构体的实例化地址
	p13 := &person{}
	fmt.Printf("%T\n", p13)     //*main.person
	fmt.Printf("p3=%#v\n", p13) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "七米"
	p3.age = 30 // 等价 (*p3).age = 10
	p3.city = "成都"

	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"七米", city:"成都", age:30}
	test1()
	fmt.Println("====结构体的构造函数（指针类型）===")
	p4 := newPerson("lisi", "gz", 30)
	fmt.Printf("p4 = %T \n", p4)
	fmt.Printf("p4 = %v \n", p4)
	fmt.Println("===结构体的方法和接受者==")
	p5 := newPerson("lisi", "zg", 20)
	p5.Dream()
	p5.SetAge(15)
	fmt.Println(p5.age)
	// 函数SetAge2传递只是一个值，并不是指针
	p5.SetAge2(125)
	fmt.Println(p5.age)
	fmt.Println("===嵌套结构体==")

	stu := Student{
		Name:    "张三",
		age:     30,
		Gender:  "男",
		Address: Address{
			Province:   "广东",
			City:       "广州",
			UpdateTime: "2020-04-18",
		},
		Email: Email{
			Addr:       "niejian900@1",
			UpdateTime: "2020-04-18 12",
		},
	}

	fmt.Printf("stu = %#v\n", stu)
	fmt.Println(stu.Address.Province) // 间接访问结构体中字段
	fmt.Println(stu.Province) // 直接访问结构体中字段


}

type students struct {
	name string
	age int
}

func test1()  {
	m := make(map[string]*students)
	stus := []students{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

// 构造函数（当结构体字段较多，值拷贝消耗过多，使用指针类型）
// 使用指针类型，返回结构体的地址
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

