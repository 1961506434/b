package main
import (
	"fmt"
	"math"
	"unicode"
	"math/rand"
	"time"
	"sort"
)

// 遍历字符串
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for  _,r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func demo(){
	var a, b = 3, 4
	var c int
	c=int(math.Sqrt(float64(a*a+b*b)))
	fmt.Print( c)
}

func demo1(){
	var a =2
	var c=true
	var b =3.123
	var d="hello你好"
	fmt.Printf("%v(%c)%v(%c)%v(%c)%v(%c)",a,a,b,b,c,c,d,d)
}
func countHanzi(s string) (num int) {
	var cnt int = 0
	for _, c := range s {
		if unicode.Is(unicode.Han, c) {
			cnt++
		}
	}
	return cnt
}
func demo2() {
	a := 1
	var b float64 = 1
	c := true
	d := "来来来"
	fmt.Printf("%v:%T\n", a, a)
	fmt.Printf("%v:%T\n", b, b)
	fmt.Printf("%v:%T\n", c, c)
	fmt.Printf("%v:%T\n", d, d)
}
func demo3(){
	s:="一起舞蹈"
	count:=0
	for _,d :=range s   {
		if len(string(d))>=3 {
			count++
		}
	}
	fmt.Printf("%v:%T",count,count)
}
func demo6() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}
func main() {
//fmt.Println(countHanzi("啦啦啦啦啦啦"))
//
//	var name  = "奥术大师多"
//	fmt.Println(len(name))
//	demo2()
//	demo3()
//	gotoDemo2()
//breakDemo1()
//	continueDemo()
//	qiuhe()
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
func demo12(){
		type Map map[string][]int
		m := make(Map)
		s := []int{1, 2}
		s = append(s, 3)
		fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
func demo11(){

		var sliceMap = make(map[string][]string, 3)
		fmt.Println(sliceMap)
		fmt.Println("after init")
		key := "中国"
		value, ok := sliceMap[key]
		if !ok {
			value = make([]string, 0, 2)
		}
		value = append(value, "北京", "上海")
		sliceMap[key] = value
		fmt.Println(sliceMap)

}
func demo10(){

		var mapSlice = make([]map[string]string, 3)
		for index, value := range mapSlice {
			fmt.Printf("index:%d value:%v\n", index, value)
		}
		fmt.Println("after init")
		// 对切片中的map元素进行初始化
		mapSlice[0] = make(map[string]string, 10)
		mapSlice[0]["name"] = "小王子"
		mapSlice[0]["password"] = "123456"
		mapSlice[0]["address"] = "沙河"
		for index, value := range mapSlice {
			fmt.Printf("index:%d value:%v\n", index, value)
		}

}
func demo8(){
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	newSlice = append(newSlice, 60)
	fmt.Println(slice, newSlice)
}
func demo9(){

		rand.Seed(time.Now().UnixNano()) //初始化随机数种子

		var scoreMap = make(map[string]int, 200)

		for i := 0; i < 100; i++ {
			key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
			value := rand.Intn(100)          //生成0~99的随机整数
			scoreMap[key] = value
		}
		//取出map中的所有key存入切片keys
		var keys = make([]string, 0, 200)
		for key := range scoreMap {
			keys = append(keys, key)
		}
		//对切片进行排序
		sort.Strings(keys)
		//按照排序后的key遍历map
		for _, key := range keys {
			fmt.Println(key, scoreMap[key])
		}

}
func demo7(){

		scoreMap := map[string]int{}
		scoreMap["张三"] = 90
		scoreMap["小明"] = 100
		// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
		v, ok := scoreMap["张"]
		if ok {
			fmt.Println(v)
		} else {
			fmt.Println("查无此人")
		}


}
func breakDemo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}
func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}
func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}
func demo4() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a) //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆
}
func jiujiu(){

	for i:=1;i<=9;i++  {
		for j:=1;j<=i ;j++  {
			fmt.Printf("%v*%v=%v\t",i,j,i*j)
		}
		fmt.Println()
	}
}
func qiuhe(){
	a:=[4]int{1,2,3,4}
	for i:=0;i< len(a);i++  {
		for j:=0;j<i ;j++  {
			if a[i]+a[j]==5 {
				fmt.Printf("(%v,%v)",i,j)
			}
		}
	}
}
func demo5() {
	a := [5]int{1, 23, 345, 4, 22}
	s := a[1:3]  // s := a[low:high]
	fmt.Println(len(a))
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4]  // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
}