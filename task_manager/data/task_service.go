package data

import(
	"task_manager/models"
	"errors"
)

var tasks = make(map[int]models.Task)
var currID = 1

func GetAllTasks() []models.Task {
	var all []models.Task
	for _, t := range tasks {
		all = append(all, t)
	}

	return all
}

func GetTaskByID(id int) (models.Task, error) {
	if task, exists := tasks[id]; exists{
		return task,nil
	}

	return models.Task{},errors.New("task not found")
}

func CreateTask(task models.Task) models.Task {
	task.ID = currID
	tasks[task.ID] = task
	currID++

	return task
}

func UpdateTask(id int, updatedTask models.Task) (models.Task, error){
	if task, exists := tasks[id]; exists{
		updatedTask.ID = task.ID
		tasks[id] = updatedTask

		return updatedTask,nil
	}

	return models.Task{},errors.New("task not found")
}

func DeleteTask(id int) error{
	if _,exists := tasks[id]; exists{
		delete(tasks,id)
		
		return nil
	}

	return errors.New("task not found")
}