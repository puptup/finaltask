package handlers

import "time"

type Group struct {
	GroupID int    `json:"id"`
	Title   string `json:"titile"`
	Tasks   []Task `json:"tasks"`
}

type Task struct {
	TaskID     int         `json:"id"`
	Title      string      `json:"titile"`
	GroupID    string      `json:"group"`
	Timeframes []Timeframe `json:"time_frames"`
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
