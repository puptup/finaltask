package dbrepo

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
	TaskID int    `json:"task_id"`
	From   string `json:"from"`
	To     string `json:"to"`
}

type GroupsResponse struct {
	Groups []Group `json:"groups"`
}

type TasksResponse struct {
	Tasks []Task `json:"tasks"`
}
