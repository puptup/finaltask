package dbrepo

import (
	"database/sql"
	"time"
)

type DBWorker interface {
	Remover
	Getter
	Poster
	Putter
	Initializer
}

type Initializer interface {
	DBInit() *sql.DB
}

type Remover interface {
	DeleteGroup(id int) error
	DeleteTask(id int) error
	DeleteTimeFrame(id int) error
}

type Getter interface {
	GetGroups() (GroupsResponse, error)
	GetTasks() (TasksResponse, error)
}

type Poster interface {
	PostGroup(title string) (Group, error)
	PostTask(title string, group_id int) (Task, error)
	PostTimeFrame(task_id int, from, to time.Time) (Timeframe, error)
}

type Putter interface {
	PutGroup(id int, title string) (Group, error)
	PutTask(task_id, group_id int, title string) (Task, error)
}
