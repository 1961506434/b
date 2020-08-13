package main

import (
	"awesomeProject/p2"
	"fmt"
)

func main(){
	s1:=stu("张三",17,1)
	s2:=stu("李四",18,1)
	Class:=cls()
	Class.add(s1)
	//Class.search()
	Class.update(s2)
	//Class.search()
	Class.delete(1)
	Class.search()
	fmt.Print(p2.Add(2,3))
}
