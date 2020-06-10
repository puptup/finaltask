package dbrepo

import (
	"time"
)

//getTimeframeByTaskID allows you to get timeframe by ID
func (repo *DBRepo) getTimeframeByTaskID(task_id int) ([]Timeframe, error) {
	var Timeframes []Timeframe
	resultTimeframes, err := repo.DB.Query("SELECT * from time_frames where task_id = $1", task_id)
	if err != nil {
		return Timeframes, err
	}
	defer resultTimeframes.Close()

	for resultTimeframes.Next() {
		var timeframe Timeframe
		err := resultTimeframes.Scan(&timeframe.TaskID, &timeframe.From, &timeframe.To)
		if err != nil {
			return Timeframes, err
		}

		Timeframes = append(Timeframes, timeframe)
	}
	return Timeframes, nil
}

//getTaskByGroupID allows you to get task by ID
func (repo *DBRepo) getTaskByGroupID(group_id int) ([]Task, error) {
	var Tasks []Task
	resultTasks, err := repo.DB.Query("SELECT * from tasks where group_id = $1", group_id)
	if err != nil {
		return Tasks, err
	}
	defer resultTasks.Close()

	for resultTasks.Next() {
		var task Task
		err := resultTasks.Scan(&task.TaskID, &task.Title, &task.GroupID)
		if err != nil {
			return Tasks, err
		}

		task.Timeframes, err = repo.getTimeframeByTaskID(task.TaskID)
		if err != nil {
			return Tasks, err
		}
		Tasks = append(Tasks, task)
	}
	return Tasks, nil
}

//GetGroups allows you to get all groups from the database
func (repo *DBRepo) GetGroups() (GroupsResponse, error) {
	var GroupsResp GroupsResponse

	resultGroups, err := repo.DB.Query("SELECT * from groups;")
	if err != nil {
		return GroupsResp, err
	}
	defer resultGroups.Close()

	var Groups []Group
	for resultGroups.Next() {
		var group Group
		err := resultGroups.Scan(&group.GroupID, &group.Title)
		if err != nil {
			return GroupsResp, err
		}

		group.Tasks, err = repo.getTaskByGroupID(group.GroupID)
		if err != nil {
			return GroupsResp, err
		}
		Groups = append(Groups, group)
	}

	GroupsResp.Groups = Groups
	return GroupsResp, nil
}

//GetTasks allows you to get all tasks from the database
func (repo *DBRepo) GetTasks() (TasksResponse, error) {
	var TasksResp TasksResponse

	resultTasks, err := repo.DB.Query("SELECT * from tasks;")
	if err != nil {
		return TasksResp, err
	}
	defer resultTasks.Close()

	var Tasks []Task
	for resultTasks.Next() {
		var task Task
		err := resultTasks.Scan(&task.TaskID, &task.Title, &task.GroupID)
		if err != nil {
			return TasksResp, err
		}

		task.Timeframes, err = repo.getTimeframeByTaskID(task.TaskID)
		if err != nil {
			return TasksResp, err
		}
		Tasks = append(Tasks, task)
	}

	TasksResp.Tasks = Tasks
	return TasksResp, nil
}

//PostGroup allows you to post a new group in the database
func (repo *DBRepo) PostGroup(title string) (Group, error) {
	var group Group
	lastInsertId := 0
	err := repo.DB.QueryRow("INSERT INTO groups(title) values($1) RETURNING group_id", title).Scan(&lastInsertId)
	if err != nil {
		return group, err
	}

	group.GroupID = lastInsertId
	group.Title = title
	return group, nil
}

//PutGroup allows you to update an existing group in the database
func (repo *DBRepo) PutGroup(id int, title string) (Group, error) {
	var group Group
	_, err := repo.DB.Exec("update groups set title = $1 where group_id = $2;", title, id)
	if err != nil {
		return group, err
	}

	group.GroupID = id
	group.Title = title
	return group, nil
}

//DeleteGroup allows you to delete group in the database
func (repo *DBRepo) DeleteGroup(id int) error {
	_, err := repo.DB.Exec("DELETE FROM groups WHERE group_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

//PostTask allows you to post a new task in the database
func (repo *DBRepo) PostTask(title string, group_id int) (Task, error) {
	var task Task
	lastInsertId := 0
	err := repo.DB.QueryRow("INSERT INTO tasks(title,group_id) values($1,$2) RETURNING task_id", title, group_id).Scan(&lastInsertId)
	if err != nil {
		return task, err
	}

	task.Title = title
	task.GroupID = group_id
	task.TaskID = lastInsertId
	return task, nil
}

//PutTask  allows you to update an existing task in the database
func (repo *DBRepo) PutTask(task_id, group_id int, title string) (Task, error) {
	var task Task

	_, err := repo.DB.Exec("update tasks set title = $1, group_id = $2 where task_id = $3;", title, group_id, task_id)
	if err != nil {
		return task, err
	}

	task.Title = title
	task.GroupID = group_id
	task.TaskID = task_id
	return task, nil
}

//DeleteTask allows you to delete task in the database
func (repo *DBRepo) DeleteTask(id int) error {
	_, err := repo.DB.Exec("DELETE FROM tasks WHERE task_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

//PostTimeFrame allows you to post a new time frame in the database
func (repo *DBRepo) PostTimeFrame(task_id int, from, to time.Time) (Timeframe, error) {
	var timefr Timeframe

	_, err := repo.DB.Exec("INSERT INTO time_frames(task_id,start_at,end_at) values($1,$2,$3)", task_id, from, to)
	if err != nil {
		return timefr, err
	}

	timefr.TaskID = task_id
	timefr.From = from
	timefr.To = to
	return timefr, nil
}

//DeleteTimeFrame allows you to delete time frame in the database
func (repo *DBRepo) DeleteTimeFrame(id int) error {
	_, err := repo.DB.Exec("DELETE FROM time_frames WHERE task_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
