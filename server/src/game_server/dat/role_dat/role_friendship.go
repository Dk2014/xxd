// 角色羁绊数据

package role_dat

import (
	"core/mysql"
)

var (
	g_RoleFriendshipStuffes map[int8][]*RoleFriendshipStuff // key: role_id
)

// 角色羁绊数据
type RoleFriendshipStuff struct {
	RoleId            int8
	FriendshipLevel   int16
	RequiredRoleLevel int16
	FavouriteItem     int16
	FavouriteCount    int32
	Health            int32
	Attack            int32
	Defend            int32
	Cultivation       int32
}

func loadRoleFriendship(db *mysql.Connection) {
	sql := "select `role_id`, `friendship_level`, `required_role_level`, `favourite_item`, `favourite_count`, `health`, `attack`, `defend`, `cultivation` from `role_friendship` order by `role_id`, `friendship_level`"
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}

	g_RoleFriendshipStuffes = make(map[int8][]*RoleFriendshipStuff)

	iRoleId := res.Map("role_id")
	iFriendshipLevel := res.Map("friendship_level")
	iRequiredRoleLevel := res.Map("required_role_level")
	iFavouriteItem := res.Map("favourite_item")
	iFavouriteCount := res.Map("favourite_count")
	iHealth := res.Map("health")
	iAttack := res.Map("attack")
	iDefend := res.Map("defend")
	iCultivation := res.Map("cultivation")

	for _, row := range res.Rows {
		roleID := row.Int8(iRoleId)

		// ATTENTION: 这里的羁绊等级是从零开始算起的，实际策划案是从一算起的，在对外的接口函数处进行了处理！
		g_RoleFriendshipStuffes[roleID] = append(g_RoleFriendshipStuffes[roleID], &RoleFriendshipStuff{
			RoleId:            row.Int8(iRoleId),
			FriendshipLevel:   row.Int16(iFriendshipLevel),
			RequiredRoleLevel: row.Int16(iRequiredRoleLevel),
			FavouriteItem:     row.Int16(iFavouriteItem),
			FavouriteCount:    row.Int32(iFavouriteCount),
			Health:            row.Int32(iHealth),
			Attack:            row.Int32(iAttack),
			Defend:            row.Int32(iDefend),
			Cultivation:       row.Int32(iCultivation),
		})
	}
}
