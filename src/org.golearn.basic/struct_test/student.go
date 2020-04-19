package main

import "fmt"

type student struct {
	id int
	name string
	class string
}

func newStudent1(id int, name,class string) *student  {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}

}

type stuMgr struct {
	allStudent []*student
}

func newStuMgr() *stuMgr {
	// 构造cap = 100的数组
	return &stuMgr{allStudent: make([] *student, 0, 100)}
}

// stuMgr的成员方法
func (s *stuMgr) addStudent(newStu *student) {
	s.allStudent = append(s.allStudent, newStu)
}

func (s *stuMgr) modifyStu(newStu *student) {
	for i, i2 := range s.allStudent {
		if newStu.id == i2.id {
			s.allStudent[i] = newStu
			return
		}
	}
}

func (s *stuMgr) showAllStudents() {
	for _, v := range s.allStudent {

		fmt.Printf("id: %d, name: %s, class: %s \n", v.id, v.name, v.class)
	}
}


