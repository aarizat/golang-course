package main

import "fmt"


type taskList struct {
	tasks []*task
}

func (t *taskList) appendList (tk *	task) {
	t.tasks = append(t.tasks, tk)
}

func (t *taskList) deleteTask (ix int) {
	t.tasks = append(t.tasks[:ix], t.tasks[ix+1:]...)
}

func (t *taskList) printTasks () {
	for _, task := range t.tasks {
		fmt.Println("Name", task.name)
		fmt.Println("Description", task.name)
	}
}

func (t *taskList) printCompleted () {
	for _, task := range t.tasks {
		if task.completed {
			fmt.Println("Name", task.name)
			fmt.Println("Description", task.name)
		}
	}
}

type task struct {
	name string
	description string
	completed bool
}

func (t *task) markCompleted() {
	t.completed = true
}

func (t *task) updateDescription (description string) {
	t.description = description
}

func (t *task) updateName (name string) {
	t.name = name
}

func main() {
	t1 := &task{
		name: "task 1",
		description: "Learn Go programming language",
	}

	t2 := &task{
		name: "task 2",
		description: "Learn Python programming language",
		completed: true,
	}

	t3 := &task{
		name: "task 3",
		description: "Learn NodeJs programming language",
	}

	list := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	fmt.Println(list.tasks[0])
	list.appendList(t3)

	fmt.Println(len(list.tasks))
	//list.deleteTask(1)
	//fmt.Println(len(list.tasks))

	// print tasks
	list.printTasks()
	list.printCompleted()

	mapTasks := make(map[string]*taskList)
	mapTasks["Andres"] = list

	fmt.Println("Andres' Tasks")
	mapTasks["Andres"].printTasks()
}