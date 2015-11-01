package main

import (
	"fmt"
)

type Task struct {
	ID int
	Detail string
	done bool
}

func NewTask(id int, detail string) *Task {
	task := &Task{
		ID: id,
		Detail: detail,
		done: false,
	}
	return task
}

func main() {
	var tasks []*Task
	var tasks2 []*Task
	var tasks3 []*Task

	task1 := NewTask(1, "buy something")
	fmt.Printf("%+v\n", task1)

	tasks = append(tasks, task1)

	task2 := NewTask(2, "hoge")
	tasks = append(tasks, task2)

	task3 := NewTask(3, "3333")
	tasks = append(tasks, task3)

	tasks2 = tasks

	fmt.Printf("%+v\n", tasks2)

	var lenn int
	lenn = len(tasks2)
	fmt.Printf("num : %d\n", lenn)
	fmt.Printf("num : %d\n", len(tasks3))
}
