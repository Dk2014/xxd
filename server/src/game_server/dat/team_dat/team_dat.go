package team_dat

import (
	"core/mysql"
)

func Load(db *mysql.Connection) {
	loadTeamship(db)
}

func loadTeamship(db *mysql.Connection) {
	sql := "select `level`, `needs_relationship`, `health`, `attack`, `defence` from `teamship` order by `level` asc;"
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}

	arrTeamshipStuff = []*TeamshipStuff{}

	iLevel := res.Map("level")
	iNeedsRelationship := res.Map("needs_relationship")
	iHealth := res.Map("health")
	iAttack := res.Map("attack")
	iDefence := res.Map("defence")

	for _, row := range res.Rows {
		arrTeamshipStuff = append(arrTeamshipStuff, &TeamshipStuff{
			Level:             row.Int16(iLevel),
			NeedsRelationship: row.Int32(iNeedsRelationship),
			Health:            row.Int32(iHealth),
			Attack:            row.Int32(iAttack),
			Defence:           row.Int32(iDefence),
		})
	}
}

var (
	arrTeamshipStuff []*TeamshipStuff
)

type TeamshipStuff struct {
	Level             int16 // 团队配合等级
	NeedsRelationship int32 // 所需友情数
	Health            int32 // 单项生命值增量
	Attack            int32 // 单项攻击增量
	Defence           int32 // 单项防御增量
}

func GetTeamshipStuff(level int16) *TeamshipStuff {
	return arrTeamshipStuff[level /*this level is started from 0*/]
}
