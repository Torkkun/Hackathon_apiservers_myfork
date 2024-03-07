package database

import (
	"log"
	"time"

	"github.com/google/uuid"
	"yandereca.tech/yandereca/domain"
)

type TaskRepository struct {
	SqlHandler
}

func (repo *TaskRepository) Store(task domain.Todo) (id string, err error) {
	//TaskIDの設定が無い時は新規発行
	if len(task.TaskId) == 0 {
		uid, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}
		uid_str := uid.String()
		task.TaskId = uid_str
	}
	//PoCの段階では同一ユーザーとして扱う
	task.UserId = "example-user-id"
	_, err = repo.Execute(
		"INSERT INTO tasks(task_id, name, description,is_done,user_id,task_status) values($1,$2,$3,$4,$5,$6) ON CONFLICT ON CONSTRAINT task_pkey DO UPDATE SET name=$2, description=$3, is_done=$4, user_id=$5, task_status=$6",
		task.TaskId, task.Task, task.Desc, task.IsDone, "example-user-id", task.Status,
	)
	if err != nil {
		return "", err
	}
	id = task.TaskId
	return
}
func (repo *TaskRepository) FindById(identifier string) (data domain.Todos, err error) {
	rows, err := repo.Query("select * from tasks where user_id = $1", identifier)

	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var task_id string
		var name string
		var desc string
		var is_done bool
		var task_status int
		var deadline time.Time
		var created_at time.Time
		var updated_at time.Time
		var user_id string
		if err := rows.Scan(&task_id, &name, &desc, &is_done, &user_id, &task_status, &deadline, &created_at, &updated_at); err != nil {
			log.Println(err)
			continue
		}
		task := domain.Todo{
			Task:     name,
			TaskId:   task_id,
			Desc:     desc,
			IsDone:   is_done,
			UserId:   user_id,
			Deadline: deadline.String(),
			Status:   task_status,
		}

		data = append(data, task)
	}
	return
}

func (repo *TaskRepository) CalcProgress(identifier string) (progress float64, err error) {
	datas, err := repo.FindById(identifier)
	var count = 0

	for _, data := range datas {
		if data.IsDone {
			count = count + 1
		}
	}

	progress = float64(count) / float64(len(datas))
	return
}
