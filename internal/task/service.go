package task

func CreateTask(task Task) Task {
	task.ID = idCounter
	idCounter++
	tasks = append(tasks, task)
	return task
}

func GetTasks() []Task {
	return tasks
}

func UpdateTask(id int, updatedTask Task) (Task, bool) {
	for i, t := range tasks {
		if t.ID == id {
			tasks[i] = updatedTask
			tasks[i].ID = id
			return tasks[i], true
		}
	}
	return Task{}, false
}

func DeleteTask(id int) bool {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}
