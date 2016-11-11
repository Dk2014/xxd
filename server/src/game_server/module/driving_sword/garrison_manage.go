package driving_sword

import (
	"game_server/api/protocol/driving_sword_api"
	"game_server/dat/driving_sword_dat"
	"game_server/dat/player_dat"
	"game_server/dat/role_dat"
	"game_server/mdb"
	"game_server/module"

	"core/fail"
	"core/net"
	"core/time"
	"math"
)

func calcGarrisonTime(event *mdb.PlayerDrivingSwordEventExploring) (garrison_time int64, now int64) {
	now = time.GetNowTime()
	garrison_time = event.GarrisonTime + (now - event.GarrisonStart)
	garrison_time = int64(math.Min(float64(garrison_time), float64(driving_sword_dat.GARRISON_TIME_LIMIT)))
	return
}

func refreshGarrisonTime(event *mdb.PlayerDrivingSwordEventExploring) {
	event.GarrisonTime, event.GarrisonStart = calcGarrisonTime(event)
}

func doListGarrisons(session *net.Session) {
	state := module.State(session)
	db := state.Database
	out := &driving_sword_api.ListGarrisons_Out{}

	db.Select.PlayerDrivingSwordEventExploring(func(row *mdb.PlayerDrivingSwordEventExploringRow) {
		event := row.GoObject()
		garrison_time := event.GarrisonTime
		if event.Status == int8(driving_sword_api.EXPLORING_MOUNTAIN_STATUS_IN_GARRISON) {
			garrison_time, _ = calcGarrisonTime(event)
		}
		out.Garrisons = append(out.Garrisons, driving_sword_api.ListGarrisons_Out_Garrisons{
			RoleId:       event.RoleId,
			GarrisonTime: garrison_time,
			AwardedTime:  event.AwardTime,
			Cloud:        event.CloudId,
			EventId:      event.DataId,
			Status:       driving_sword_api.ExploringMountainStatus(row.Status()),
		})
	})

	session.Send(out)
}

func doAwardGarrison(session *net.Session, role_id int8, xdEventType int32) {
	state := module.State(session)
	db := state.Database

	var event *mdb.PlayerDrivingSwordEventExploring
	db.Select.PlayerDrivingSwordEventExploring(func(row *mdb.PlayerDrivingSwordEventExploringRow) {
		if row.RoleId() == role_id {
			event = row.GoObject()
			row.Break()
		}
	})

	fail.When(event == nil, "can't find garrison by role_id")

	refreshGarrisonTime(event)

	fail.When(event.AwardTime+driving_sword_dat.GARRISON_AWARD_TIME_UNIT > event.GarrisonTime, "garrison time is not ready")
	for event.AwardTime+driving_sword_dat.GARRISON_AWARD_TIME_UNIT <= event.GarrisonTime {
		event.AwardTime += driving_sword_dat.GARRISON_AWARD_TIME_UNIT

		exploringDat := driving_sword_dat.GetDrivingSwordExploring(event.CloudId, event.DataId)

		if exploringDat.GarrisonItem != 0 && exploringDat.GarrisonNum > 0 {
			module.Item.AddItem(db, exploringDat.GarrisonItem, int16(exploringDat.GarrisonNum), 0 /*TODO*/, xdEventType, "")
		}
		if exploringDat.GarrisonCoinNum > 0 {
			module.Player.IncMoney(db, state.MoneyState, int64(exploringDat.GarrisonCoinNum), player_dat.COINS, 0 /*TODO*/, xdEventType, "")
		}
	}

	db.Update.PlayerDrivingSwordEventExploring(event)

	session.Send(&driving_sword_api.AwardGarrison_Out{})
}

func doEndGarrison(db *mdb.Database, role_id int8) (x, y uint8, status int8, cloud int16) {

	var event *mdb.PlayerDrivingSwordEventExploring
	db.Select.PlayerDrivingSwordEventExploring(func(row *mdb.PlayerDrivingSwordEventExploringRow) {
		if row.RoleId() == role_id {
			event = row.GoObject()
			x = uint8(event.X)
			y = uint8(event.Y)
			cloud = event.CloudId
			row.Break()
		}
	})

	fail.When(event == nil, "can't find garrison by role_id")

	refreshGarrisonTime(event)

	event.Status = int8(driving_sword_api.EXPLORING_MOUNTAIN_STATUS_TREASURE_EMPTY)
	event.RoleId = 0

	//检查是否不再能够驻守
	if event.AwardTime >= driving_sword_dat.GARRISON_TIME_LIMIT {
		event.Status = int8(driving_sword_api.EXPLORING_MOUNTAIN_STATUS_BROKEN)
	}
	status = event.Status

	db.Update.PlayerDrivingSwordEventExploring(event)
	return
}

func doExplorerGarrison(session *net.Session, role_id int8) {
	state := module.State(session)
	db := state.Database

	fail.When(role_dat.IsMainRole(role_id), "主角不能驻守")
	//检查一下角色是否在客栈
	//role := module.Role.GetBuddyRoleInTeam(db, role_id)
	//role.Status = role_dat.ROLE_STATUS_GARRISION
	//db.Update.PlayerRole(role)

	drivingInfo := db.Lookup.PlayerDrivingSwordInfo(db.PlayerId())

	var event *mdb.PlayerDrivingSwordEventExploring
	is_this_role_busy := false
	db.Select.PlayerDrivingSwordEventExploring(func(row *mdb.PlayerDrivingSwordEventExploringRow) {
		if event == nil && row.CloudId() == drivingInfo.CurrentCloud && row.X() == drivingInfo.CurrentX && row.Y() == drivingInfo.CurrentY {
			event = row.GoObject()
		}
		if row.Status() == int8(driving_sword_api.EXPLORING_MOUNTAIN_STATUS_IN_GARRISON) && row.RoleId() == role_id {
			is_this_role_busy = true
		}
	})

	fail.When(event == nil, "the player current location is not a explorable mountain")
	fail.When(is_this_role_busy, "this role is already garrisoning")

	fail.When(event.Status != int8(driving_sword_api.EXPLORING_MOUNTAIN_STATUS_TREASURE_EMPTY), "this exploring event mountain can't be used to garrison")

	event.RoleId = role_id
	event.GarrisonStart = time.GetNowTime()
	event.Status = int8(driving_sword_api.EXPLORING_MOUNTAIN_STATUS_IN_GARRISON)

	db.Update.PlayerDrivingSwordEventExploring(event)

	session.Send(&driving_sword_api.ExplorerGarrison_Out{})
}
