package main

import "fmt"

type Student struct {
	name string
	age int
	id int
}
type Class struct {
	class []*Student
}

func stu(name string,age,id int)*Student{
	return &Student{
		name:name,
		age:age,
		id:id,
	}
}
func cls()*Class{
	return &Class{
		class:make([]*Student,0),
	}
}

//添加学生
func (s *Class) add(newStu *Student){
	s.class=append(s.class,newStu)
}
//修改学生
func (s *Class) update(newStu *Student){
	for key, value := range s.class {
		if value.id==newStu.id {
			s.class[key]=newStu
			return
		}
	}
	fmt.Print("查无此人")
}
//查找学生
func (s *Class) search(){
	for _, value := range s.class {
		fmt.Println(value)
	}
}
//删除学生
func (s *Class) delete(id int){
	for key, value := range s.class {
		if value.id==id {
			s.class=append(s.class[:key],s.class[key+1:]...)
		}
	}
}
func appen(a,b int)int{
	return a+b
}