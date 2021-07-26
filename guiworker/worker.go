package guiworker

type Task struct {
	OnSuccess *func()
	Run       *func()
}

var tasks = make(chan Task, 50)

type Worker struct {
}

func Submit(task Task) {
	tasks <- task
}

func DoWork() {
	for task := range tasks {
		task.Run()
	}
}
