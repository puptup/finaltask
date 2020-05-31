package dbrepo

import (
	"log"
)

func GetGroups() GroupsResponse {
	result, err := repo.Query("SELECT * from groups;")
	if err != nil {
		log.Panic(err)
	}
	defer result.Close()

	var Groups []Group

	for result.Next() {
		var group Group
		err := result.Scan(&group.GroupID, &group.Title)
		if err != nil {
			log.Panic(err)
		}
		Groups = append(Groups, group)
	}

	var GroupsResp GroupsResponse
	GroupsResp.Groups = Groups
	return GroupsResp
}

func PostGroup() {

}
