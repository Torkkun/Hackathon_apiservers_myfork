package domain

type Todo struct {
	Task     string `json:"name"`
	Desc     string `json:"desc"`
	Deadline string `json:"deadline"`
	TaskId   string `json:"task_id"`
	IsDone   bool   `json:"is_done"`
	Status   int    `json:"status"`
	UserId   string `json:"user_id"`
}

type TodoSuccessMessage struct {
	Message string `json:"message"`
	Result  bool   `json:"result"`
	TaskId  string `json:"task_id"`
}

type Todos []Todo
