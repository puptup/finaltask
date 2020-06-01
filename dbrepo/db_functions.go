package dbrepo

import (
	"log"
)

func onErrPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func getTimeframeByTaskID(task_id int) []Timeframe {
	resultTimeframes, err := repo.Query("SELECT * from time_frames where task_id = $1", task_id)
	onErrPanic(err)
	defer resultTimeframes.Close()

	var Timeframes []Timeframe
	for resultTimeframes.Next() {
		var timeframe Timeframe
		err := resultTimeframes.Scan(&timeframe.TaskID, &timeframe.From, &timeframe.To)
		onErrPanic(err)

		Timeframes = append(Timeframes, timeframe)
	}
	return Timeframes
}

func getTaskByGroupID(group_id int) []Task {
	resultTasks, err := repo.Query("SELECT * from tasks where group_id = $1", group_id)
	onErrPanic(err)
	defer resultTasks.Close()

	var Tasks []Task
	for resultTasks.Next() {
		var task Task
		err := resultTasks.Scan(&task.TaskID, &task.Title, &task.GroupID)
		onErrPanic(err)

		task.Timeframes = getTimeframeByTaskID(task.TaskID)
		Tasks = append(Tasks, task)
	}
	return Tasks
}

//GetGroups позволяет получить список всех групп из БД
func GetGroups() GroupsResponse {
	resultGroups, err := repo.Query("SELECT * from groups;")
	onErrPanic(err)
	defer resultGroups.Close()

	var Groups []Group
	for resultGroups.Next() {
		var group Group
		err := resultGroups.Scan(&group.GroupID, &group.Title)
		onErrPanic(err)

		group.Tasks = getTaskByGroupID(group.GroupID)
		Groups = append(Groups, group)
	}

	var GroupsResp GroupsResponse
	GroupsResp.Groups = Groups
	return GroupsResp
}

func GetTasks() TasksResponse {
	resultTasks, err := repo.Query("SELECT * from tasks;")
	onErrPanic(err)
	defer resultTasks.Close()

	var Tasks []Task
	for resultTasks.Next() {
		var task Task
		err := resultTasks.Scan(&task.TaskID, &task.Title, &task.GroupID)
		onErrPanic(err)

		task.Timeframes = getTimeframeByTaskID(task.TaskID)
		Tasks = append(Tasks, task)
	}
	var TasksResp TasksResponse
	TasksResp.Tasks = Tasks
	return TasksResp
}

//PostGroup позволяет запостить новую группу в БД
func PostGroup(title string) Group {
	lastInsertId := 0
	err := repo.QueryRow("INSERT INTO groups(title) values($1) RETURNING group_id", title).Scan(&lastInsertId)
	onErrPanic(err)

	var group Group
	group.GroupID = lastInsertId
	group.Title = title
	return group
}

func PutGroup(id int) {

}

func DeleteGroup() {

}

func PostTask(title string, group_id int) Task {
	lastInsertId := 0
	err := repo.QueryRow("INSERT INTO tasks(title,group_id) values($1,$2) RETURNING task_id", title, group_id).Scan(&lastInsertId)
	onErrPanic(err)

	var task Task
	task.Title = title
	task.GroupID = group_id
	task.TaskID = lastInsertId
	return task
}

func PutTask(task_id, group_id int, title string) Task {
	_, err := repo.Exec("update tasks set title = $1, group_id = $2 where task_id = $3;", title, group_id, task_id)
	onErrPanic(err)

	var task Task
	task.Title = title
	task.GroupID = group_id
	task.TaskID = task_id
	return task
}

func DeleteTask() {

}

func PostTimeFrame(task_id int, from, to string) Timeframe {
	_, err := repo.Exec("INSERT INTO time_frames(task_id,from_time,to_time) values($1,$2,$3)", task_id, from, to)
	onErrPanic(err)

	var timefr Timeframe
	timefr.TaskID = task_id
	timefr.From = from
	timefr.To = to
	return timefr
}

func DeleteTimeFrame() {

}
