package main

import (
	"fmt"
	"reflect"
)


// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}
type cat struct {
name string
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (c cat) move() {
	fmt.Println("猫会动")
}

func main4() {
	var x animal
	x = cat{name: "花花"}
	x.move()
	x.say()
}
func main5() {
	// 定义一个空接口x
	var x interface{}
	s := "Hello 沙河"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
}
func main8() {
	var x interface{}
	x = "Hello 沙河"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
	justifyType("asd")
}
func justifyType(x interface{}) {
	switch v:=x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}
//func main() {
//	var a float32 = 3.14
//	var b int64 = 100
//	reflectValue(a) // type is float32, value is 3.140000
//	reflectValue(b) // type is int64, value is 100
//	// 将int类型的原始值转换为reflect.Value类型
//	c := reflect.ValueOf(10)
//	fmt.Print( c,c,c.Kind()) // type c :reflect.Value
//}
//func main() {
//	// *int类型空指针
//	var a *int
//	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
//	// nil值
//	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
//	// 实例化一个匿名结构体
//	b := struct{}{}
//	// 尝试从结构体中查找"abc"字段
//	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
//	// 尝试从结构体中查找"abc"方法
//	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
//	// map
//	c := map[string]int{}
//	// 尝试从map中查找一个不存在的键
//	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
//}
type student9 struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

//func main() {
//	stu1 := student9{
//		Name:  "小王子",
//		Score: 90,
//	}
//
//	t := reflect.TypeOf(stu1)
//	fmt.Println(t.Name(), t.Kind()) // student struct
//	// 通过for循环遍历结构体的所有字段信息
//	for i := 0; i < t.NumField(); i++ {
//		field := t.Field(i)
//		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
//	}
//
//	// 通过字段名获取指定结构体字段信息
//	if scoreField, ok := t.FieldByName("Score"); ok {
//		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
//	}
//}
func (s student9) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student9) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}
func main(){
	printMethod(student9{})
}

