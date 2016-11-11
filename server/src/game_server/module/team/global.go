package team

import (
	"core/fail"
	"game_server/mdb"
	"game_server/module"
)

func init() {
	module.Team = TeamMod{}
}

type TeamMod struct {
}

func (mod TeamMod) IncRelationship(db *mdb.Database, value int32) {
	fail.When(value < 0, "invald relationship incremental value")
	teamInfo := db.Lookup.PlayerTeamInfo(db.PlayerId())
	teamInfo.Relationship += value
	db.Update.PlayerTeamInfo(teamInfo)
}

func (mod TeamMod) GetFormPosArray(db *mdb.Database) (posArray [9]*int8) {
	formationInfo := db.Lookup.PlayerFormation(db.PlayerId())
	return getFormPosArray(formationInfo)
}

func (mod TeamMod) DownFormRole(db *mdb.Database, roleId int8) {
	formationInfo := db.Lookup.PlayerFormation(db.PlayerId())
	posArray := getFormPosArray(formationInfo)
	for pos, rolePtr := range posArray {
		if *rolePtr == roleId {
			downFormation(db, int8(pos))
		}
	}
}

func (mod TeamMod) IsRoleInForm(db *mdb.Database, roleId int8) bool {
	formationInfo := db.Lookup.PlayerFormation(db.PlayerId())
	posArray := getFormPosArray(formationInfo)
	for _, rolePtr := range posArray {
		if *rolePtr == roleId {
			return true
		}
	}
	return false
}
