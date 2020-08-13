package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student2 struct{}

func (stu *Student2) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo People = &Student2{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}