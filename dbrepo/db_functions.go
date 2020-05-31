package dbrepo

import (
	"log"
)

func onErrPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}
func selecter(query string, args ...interface{}) {

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

		resultTasks, err := repo.Query("SELECT * from tasks where group_id = $1", group.GroupID)
		onErrPanic(err)

		defer resultTasks.Close()
		var Tasks []Task
		for resultTasks.Next() {
			var task Task
			err := resultTasks.Scan(&task.TaskID, &task.Title, &task.GroupID)
			onErrPanic(err)

			resultTimeframes, err := repo.Query("SELECT * from time_frames where task_id = $1", task.TaskID)
			onErrPanic(err)

			var Timeframes []Timeframe
			for resultTimeframes.Next() {
				var timeframe Timeframe
				err := resultTimeframes.Scan(&timeframe.TaskID, &timeframe.From, &timeframe.To)
				onErrPanic(err)

				Timeframes = append(Timeframes, timeframe)
			}
			task.Timeframes = Timeframes
			Tasks = append(Tasks, task)
		}

		group.Tasks = Tasks
		Groups = append(Groups, group)
	}

	var GroupsResp GroupsResponse
	GroupsResp.Groups = Groups
	return GroupsResp
}

//PostGroup позволяет запостить новую группу в БД
func PostGroup() {

}

func PutGroup() {

}

func DeleteGroup() {

}

func GetTasks() {

}

func PostTask() {

}

func PutTask() {

}

func DeleteTask() {

}

func PostTimeFrame() {

}

func DeleteTimeFrame() {

}
