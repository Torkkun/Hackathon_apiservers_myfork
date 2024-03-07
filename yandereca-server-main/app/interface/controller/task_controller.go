package controller

import (
	"fmt"
	"log"

	"yandereca.tech/yandereca/domain"
	"yandereca.tech/yandereca/interface/database"
	"yandereca.tech/yandereca/usecase"
)

type TaskController struct {
	Interactor usecase.TaskInteractor
}

func NewTaskController(sqlHandler database.SqlHandler) *TaskController {
	return &TaskController{
		Interactor: usecase.TaskInteractor{
			TaskRepository: &database.TaskRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TaskController) Create(c Context) {
	// テストオブジェクト

	task := domain.Todo{}
	err := c.Bind(&task)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(task.TaskId)
	u := domain.TodoSuccessMessage{}

	task_id, err := controller.Interactor.Create(task)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	u.TaskId = task_id.TaskId
	u.Message = "success"
	u.Result = true
	c.JSON(201, u)
}

func (controller *TaskController) TestCreate(c Context) {
	// テストオブジェクト
	task := domain.Todo{}
	task.Desc = "hogehoge"
	task.Task = "example task"

	u := domain.TodoSuccessMessage{}

	task_id, err := controller.Interactor.Create(task)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	u.TaskId = task_id.TaskId
	u.Message = "success"
	u.Result = true
	c.JSON(201, u)

}
func (controller *TaskController) Read(c Context) {

	tasks, err := controller.Interactor.TaskRepository.FindById("example-user-id")
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, tasks)
}

type ResultProgress struct {
	Progress float64 `json:"progress"`
}

func (TaskController *TaskController) Calc(c Context) {
	identifier := c.DefaultQuery("task_id", "example-user-id")
	progress, err := TaskController.Interactor.CalcProgress(identifier)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(201, ResultProgress{Progress: progress})
}
