package main

import "fmt"

func main2() {
	a := 10
	b := &a
	fmt.Printf("a:%d prt:%d \n", a, &a)
	fmt.Printf("b:%p type:%T \n", b, b)
	fmt.Println(&b)
	fmt.Print(&a)

}
type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		num:=stu
		m[stu.name] = &num
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
func main1() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(a)       // 0
	fmt.Println(b)       // false
}
