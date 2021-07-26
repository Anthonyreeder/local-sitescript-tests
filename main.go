package main

import (
	"fmt"
)

const (
	LoggingIn string = "Logging in"
)

func main() {
	//	Demo()
	Testval := "hi"
	fmt.Println(fmt.Sprint(LoggingIn, Testval))

	//task := Task{}
	//task.Test()
}

type Task struct {
	Val string
}

func (task *Task) Test() {
	task.Val = "123"
	task.Test3(task.Val, task.Test2(task.Val))
}

func (task *Task) Test2(value string) bool {
	task.Val = task.Val + "4"
	return true
}

func (task *Task) Test3(v string, vv bool) {
	println("Value is " + task.Val)
}
