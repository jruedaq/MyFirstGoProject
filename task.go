package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) addToList(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) deleteFromList(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) printList() {
	for index, task := range t.tasks {
		fmt.Println("Pocisión", index+1)
		fmt.Println("Nombre", task.name)
		fmt.Println("Descripción", task.description)
		fmt.Println("Completado?", task.complete)
		fmt.Println()
	}
}

func (t *taskList) printTaskCompeteList() {
	for index, task := range t.tasks {
		if task.complete {
			fmt.Println("Pocisión", index+1)
			fmt.Println("Nombre", task.name)
			fmt.Println("Descripción", task.description)
			fmt.Println("Completado?", task.complete)
			fmt.Println()
		}
	}
}

type task struct {
	name        string
	description string
	complete    bool
}

func (t *task) markAsComplete() {
	t.complete = true
}

func (t *task) updateDescription(desc string) {
	t.description = desc
}

func (t *task) updateName(name string) {
	t.name = name
}

func main() {
	t1 := &task{
		name:        "Completar curso de Go",
		description: "Terminar el curso de Go esta semana",
	}

	t2 := &task{
		name:        "Completar curso de Go2",
		description: "Terminar el curso de Go esta semana2",
	}

	t3 := &task{
		name:        "Completar curso de Go3",
		description: "Terminar el curso de Go esta semana3",
	}

	list := &taskList{tasks: []*task{
		t1, t2,
	}}
	list.addToList(t3)

	list.printList()
	fmt.Println("-------------------------------")
	fmt.Println()
	list.tasks[1].markAsComplete()
	list.printTaskCompeteList()

	taskMap := make(map[string]*taskList)
	taskMap["Juan"] = list
}
