package dbrepo

import (
	"database/sql"
	"time"
)

type DBWorker interface {
	Deleter
	Getter
	Poster
	Putter
	Initializer
}

type Initializer interface {
	DBInit() *sql.DB
}

type Deleter interface {
	DeleteGroup(id int) error
	DeleteTask(id int) error
	DeleteTimeFrame(id int) error
}

type Getter interface {
	GetGroups() (GroupsResponse, error)
	getTaskByGroupID(group_id int) ([]Task, error)
	getTimeframeByTaskID(task_id int) ([]Timeframe, error)
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

type Group struct {
	GroupID int    `json:"id"`
	Title   string `json:"title"`
	Tasks   []Task `json:"tasks,omitempty"`
}

type Task struct {
	TaskID     int         `json:"id"`
	Title      string      `json:"title"`
	GroupID    int         `json:"group_id"`
	Timeframes []Timeframe `json:"time_frames,omitempty"`
}

type Timeframe struct {
	TaskID int       `json:"task_id"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
}

type GroupsResponse struct {
	Groups []Group `json:"groups"`
}

type TasksResponse struct {
	Tasks []Task `json:"tasks"`
}
