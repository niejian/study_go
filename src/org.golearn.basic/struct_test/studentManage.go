package main

import (
	"fmt"
	"os"
)

/**
学院信息管理
1. 添加学生
2. 编辑学生
3. 展示所有
 */

func showMenu()  {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1. 添加学员")
	fmt.Println("2. 编辑学员信息")
	fmt.Println("3. 展示所有学员信息")
	fmt.Println("4. 退出系统")

}

func getInput() *student{
	var(
		id int
		name string
		class string
	)
	fmt.Println("请按要求输入以下信息")
	fmt.Printf("请输入学号：")
	fmt.Scanf("%d", &id)
	fmt.Printf("请输入姓名：")
	fmt.Scanf("%s", &name)
	fmt.Printf("请输入班级：")
	fmt.Scanf("%s", &class)

	return newStudent1(id, name, class)

}

func main()  {
	// 1. 打印系统菜单
	showMenu()
	mgr := newStuMgr()
	for  {

		// 等待用户的选项
		var input int
		fmt.Printf("请输入操作序号：")

		fmt.Scanf("%d", &input)
		fmt.Println("用户输入的是：", input)
		//
		switch input {
		case 1:
			// 添加学员
			stu := getInput()
			mgr.addStudent(stu)
		case 2:
			// 编辑学员信息
			stu := getInput()
			mgr.modifyStu(stu)
		case 3:
			// 展示所有学员信息
			mgr.showAllStudents()
		case 4:
			// 退出系统
			os.Exit(0)
		default:
			// 展示所有学员信息
			mgr.showAllStudents()

		}

	}

}

