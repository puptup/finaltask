package dbrepo

func getTimeframeByTaskID(task_id int) ([]Timeframe, error) {
	var Timeframes []Timeframe
	resultTimeframes, err := repo.Query("SELECT * from time_frames where task_id = $1", task_id)
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

func getTaskByGroupID(group_id int) ([]Task, error) {
	var Tasks []Task
	resultTasks, err := repo.Query("SELECT * from tasks where group_id = $1", group_id)
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

		task.Timeframes, err = getTimeframeByTaskID(task.TaskID)
		if err != nil {
			return Tasks, err
		}
		Tasks = append(Tasks, task)
	}
	return Tasks, nil
}

//GetGroups позволяет получить список всех групп из БД
func GetGroups() (GroupsResponse, error) {
	var GroupsResp GroupsResponse

	resultGroups, err := repo.Query("SELECT * from groups;")
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

		group.Tasks, err = getTaskByGroupID(group.GroupID)
		if err != nil {
			return GroupsResp, err
		}
		Groups = append(Groups, group)
	}

	GroupsResp.Groups = Groups
	return GroupsResp, nil
}

func GetTasks() (TasksResponse, error) {
	var TasksResp TasksResponse

	resultTasks, err := repo.Query("SELECT * from tasks;")
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

		task.Timeframes, err = getTimeframeByTaskID(task.TaskID)
		if err != nil {
			return TasksResp, err
		}
		Tasks = append(Tasks, task)
	}

	TasksResp.Tasks = Tasks
	return TasksResp, nil
}

//PostGroup позволяет запостить новую группу в БД
func PostGroup(title string) (Group, error) {
	var group Group
	lastInsertId := 0
	err := repo.QueryRow("INSERT INTO groups(title) values($1) RETURNING group_id", title).Scan(&lastInsertId)
	if err != nil {
		return group, err
	}

	group.GroupID = lastInsertId
	group.Title = title
	return group, nil
}

func PutGroup(id int, title string) (Group, error) {
	var group Group
	_, err := repo.Exec("update groups set title = $1 where group_id = $2;", title, id)
	if err != nil {
		return group, err
	}

	group.GroupID = id
	group.Title = title
	return group, nil
}

func DeleteGroup(id int) error {
	_, err := repo.Exec("DELETE FROM groups WHERE group_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func PostTask(title string, group_id int) (Task, error) {
	var task Task
	lastInsertId := 0
	err := repo.QueryRow("INSERT INTO tasks(title,group_id) values($1,$2) RETURNING task_id", title, group_id).Scan(&lastInsertId)
	if err != nil {
		return task, err
	}

	task.Title = title
	task.GroupID = group_id
	task.TaskID = lastInsertId
	return task, nil
}

func PutTask(task_id, group_id int, title string) (Task, error) {
	var task Task

	_, err := repo.Exec("update tasks set title = $1, group_id = $2 where task_id = $3;", title, group_id, task_id)
	if err != nil {
		return task, err
	}

	task.Title = title
	task.GroupID = group_id
	task.TaskID = task_id
	return task, nil
}

func DeleteTask(id int) error {
	_, err := repo.Exec("DELETE FROM tasks WHERE task_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func PostTimeFrame(task_id int, from, to string) (Timeframe, error) {
	var timefr Timeframe

	_, err := repo.Exec("INSERT INTO time_frames(task_id,start_at,end_at) values($1,$2,$3)", task_id, from, to)
	if err != nil {
		return timefr, err
	}

	timefr.TaskID = task_id
	timefr.From = from
	timefr.To = to
	return timefr, nil
}

func DeleteTimeFrame(id int) error {
	_, err := repo.Exec("DELETE FROM time_frames WHERE task_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
