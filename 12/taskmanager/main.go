package main

import (
	in "fmt"
	t "taskmanager/task"
)

func main() {
	task1 := t.New("Apply new settings")
	task2 := t.New("Clean up some disk space")
	task3 := t.New("Install some drivers idk")
	task4 := t.New("Check the new keyboard")
	task5 := t.New("clean up old deives")

	task1.Done = true
	task2.Done = true
	task5.Done = true

	tasks := []t.Task{task1, task2, task3, task4, task5}

	in.Println("=== Taskmanager ===")

	for _, t := range tasks {
		mark := "[ ]"

		if t.Done {
			mark = "[x]"
		}

		in.Printf("%s %s\n", mark, t.Title)
	}
}
