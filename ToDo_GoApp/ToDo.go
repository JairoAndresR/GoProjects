package main

import "fmt"

type task struct {  // Define task structure
	name string
	description string
	completed bool
}

// Note: In go arrays are not dynamics but we have Slices
type taskList struct {
	tasks []*task
}

func (t *taskList) addToList (tl *task){
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) removeFromList (index int){
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *task) changeToCompleted(){ // Set completed attribute to true
	t.completed=true
}

func (t *task) updateDescription(description string){ // Receives a description and Change the task description
	t.description=description
}

func (t *task) updateName(name string){ // Receives a name and Change the task name
	t.name=name
}

func main()  {
	t1 := &task{	// Create a new task
		name: "Complete go course",
		description: "Complete weekly class",
	}
	t2 := &task{	// Create a new task
		name: "Complete Python course",
		description: "Complete weekly class",
	}
	t3 := &task{	// Create a new task
		name: "Complete NodeJS course",
		description: "Complete weekly class",
	}

	//fmt.Printf("%+v\n", t) // Print with the format "name: taskName description: taskDescription..."
	//t.changeToCompleted() // change t to completed
	//t.updateName("Finish go course") //"Update t name"
	//t.updateDescription("Complete course now") //"Update t description"
	//fmt.Printf("%+v\n", t) // Print task
	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	//fmt.Println(lista.tasks[0])
	lista.addToList(t3)
	//fmt.Println(len(lista.tasks))
	//lista.removeFromList(1)
	//fmt.Println(len(lista.tasks))
	for i:=0; i<len(lista.tasks);i++ {
		fmt.Println("Index: ", i, "Task name: ", lista.tasks[i].name)
	}

	for index, tarea := range lista.tasks {
		fmt.Println("Index: ", index, "Task name: ", tarea.name)
	}

}