package driving_sword

import (
	"core/fail"
	"core/net"
	"game_server/api/protocol/driving_sword_api"
	"game_server/dat/driving_sword_dat"
	"game_server/dat/item_dat"
	"game_server/dat/quest_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/xdlog"
)

func init() {
	driving_sword_api.SetInHandler(DrivingSwordAPI{})
}

type DrivingSwordAPI struct {
}

func (this DrivingSwordAPI) CloudMapInfo(session *net.Session, in *driving_sword_api.CloudMapInfo_In) {
	// TODO guard the function openation
	state := module.State(session)

	drivingInfo := state.Database.Lookup.PlayerDrivingSwordInfo(state.PlayerId)

	//刷新行动点相关数据
	drivingInfo = refreshAllowedActionBought(state.Database, drivingInfo)
	drivingInfo = refreshAllowedAction(state.Database, drivingInfo)

	session.Send(&driving_sword_api.CloudMapInfo_Out{
		CurrentCloud:      drivingInfo.CurrentCloud,
		HighestCloud:      drivingInfo.HighestCloud,
		CurrentX:          uint8(drivingInfo.CurrentX),
		CurrentY:          uint8(drivingInfo.CurrentY),
		AllowedAction:     drivingInfo.AllowedAction,
		DailyActionBought: drivingInfo.DailyActionBought,
		Map:               cacheCloudMap(state, drivingInfo.CurrentCloud),
	})
}

func (this DrivingSwordAPI) CloudClimb(session *net.Session, in *driving_sword_api.CloudClimb_In) {
	state := module.State(session)
	db := state.Database

	drivingInfo := db.Lookup.PlayerDrivingSwordInfo(db.PlayerId())

	//检查是否已达上限
	fail.When(drivingInfo.HighestCloud >= driving_sword_dat.HIGHEST_CLOUD_LEVEL_LIMIT, "there is nothing more you can climb")

	//扣除石符
	module.Item.DelItemByItemId(db, item_dat.ITEM_DRIVING_SWORD_CLOUD_STONE, 1, 0, xdlog.ET_CLOUD_CLIMB)

	//爬上去！
	drivingInfo.HighestCloud++
	db.Update.PlayerDrivingSwordInfo(drivingInfo)

	session.Send(&driving_sword_api.CloudClimb_Out{})
}

func (this DrivingSwordAPI) CloudTeleport(session *net.Session, in *driving_sword_api.CloudTeleport_In) {
	state := module.State(session)
	db := state.Database

	dest_cloud := in.Cloud

	drivingInfo := db.Lookup.PlayerDrivingSwordInfo(db.PlayerId())

	fail.When(dest_cloud == drivingInfo.CurrentCloud, "can't teleport to the same cloud level")

	//check if player is standing in the hole
	hole := lookupPlayerDrivingCommonEvent(db, drivingInfo.CurrentCloud, uint8(drivingInfo.CurrentX), uint8(drivingInfo.CurrentY))
	fail.When(hole == nil || hole.EventType != int8(driving_sword_api.COMMON_EVENT_HOLE), "player hasn't been standing in the hole")

	//check if this guy can reach this level
	fail.When(dest_cloud < 1 || dest_cloud > drivingInfo.HighestCloud, "the destination cloud hasn't opened yet!")

	//update driving sword basical information
	drivingInfo.CurrentCloud = dest_cloud
	//update the born_x and born_y
	cloud := driving_sword_dat.GetCloudLevel(dest_cloud)
	drivingInfo.CurrentX = cloud.BornX
	drivingInfo.CurrentY = cloud.BornY
	db.Update.PlayerDrivingSwordInfo(drivingInfo)

	//if this is the first time stepping in, initialize the cloud
	if lookupPlayerCloudShadow(db, dest_cloud) == nil {
		punchTheShadow(state, uint8(cloud.BornX), uint8(cloud.BornY))
	}

	session.Send(&driving_sword_api.CloudTeleport_Out{
		Map: cacheCloudMap(state, dest_cloud),
	})
}

func (this DrivingSwordAPI) AreaTeleport(session *net.Session, in *driving_sword_api.AreaTeleport_In) {
	state := module.State(session)
	db := state.Database

	drivingInfo := db.Lookup.PlayerDrivingSwordInfo(db.PlayerId())

	teleport := lookupPlayerDrivingCommonEvent(db, drivingInfo.CurrentCloud, uint8(drivingInfo.CurrentX), uint8(drivingInfo.CurrentY))
	fail.When(teleport == nil || teleport.EventType != int8(driving_sword_api.COMMON_EVENT_TELEPORT), "player haven't been standing in teleport point!")

	//搜索下一个传送点
	var next_teleport *mdb.PlayerDrivingSwordEvent = nil
	db.Select.PlayerDrivingSwordEvent(func(row *mdb.PlayerDrivingSwordEventRow) {
		if row.CloudId() == drivingInfo.CurrentCloud && row.EventType() == int8(driving_sword_api.COMMON_EVENT_TELEPORT) && row.DataId() != teleport.DataId {
			event := row.GoObject()
			if next_teleport == nil {
				next_teleport = event
			} else if next_teleport.DataId < teleport.DataId {
				//range over的情况
				if next_teleport.DataId > event.DataId {
					next_teleport = event
				}
			} else if event.DataId > teleport.DataId && event.DataId < next_teleport.DataId {
				//不断逼近当前传送点
				next_teleport = event
			}
		}
	})

	//检查是否可传送
	fail.When(next_teleport == nil, "there is no more area teleport for transmitting")

	drivingInfo.CurrentX = next_teleport.X
	drivingInfo.CurrentY = next_teleport.Y
	db.Update.PlayerDrivingSwordInfo(drivingInfo)

	session.Send(&driving_sword_api.AreaTeleport_Out{
		Events: punchTheShadow(state, uint8(drivingInfo.CurrentX), uint8(drivingInfo.CurrentY)),
	})
}

func (this DrivingSwordAPI) AreaMove(session *net.Session, in *driving_sword_api.AreaMove_In) {
	state := module.State(session)
	db := state.Database

	//移动方向
	direction := in.Direction

	drivingInfo := db.Lookup.PlayerDrivingSwordInfo(db.PlayerId())

	//检查贱气是否充足
	fail.When(drivingInfo.AllowedAction <= 0, "there are not enough allowed action point can be substract")
	drivingInfo.AllowedAction--

	//尝试性移动一哈哈
	cloud := driving_sword_dat.GetCloudLevel(drivingInfo.CurrentCloud)
	switch direction {
	case driving_sword_api.MOVING_DIRECTION_NORTH:
		drivingInfo.CurrentY++
		fail.When(drivingInfo.CurrentY >= cloud.Height, "you've already been the northorn limit")
	case driving_sword_api.MOVING_DIRECTION_SOUTH:
		fail.When(drivingInfo.CurrentY <= 0, "you've already been the southern limit")
		drivingInfo.CurrentY--
	case driving_sword_api.MOVING_DIRECTION_WEST:
		fail.When(drivingInfo.CurrentX <= 0, "you've already been the western limit")
		drivingInfo.CurrentX--
	case driving_sword_api.MOVING_DIRECTION_EAST:
		drivingInfo.CurrentX++
		fail.When(drivingInfo.CurrentX >= cloud.Width, "you've already been the eastern limit")
	default:
		fail.When(true, "wrong direction you wanna move toward!")
	}

	//检查目的地是否有障碍
	dest_event := lookupPlayerDrivingCommonEvent(db, drivingInfo.CurrentCloud, uint8(drivingInfo.CurrentX), uint8(drivingInfo.CurrentY))
	fail.When(dest_event != nil && dest_event.EventType == int8(driving_sword_api.COMMON_EVENT_OBSTACLE), "the destination location is an obstacle")

	db.Update.PlayerDrivingSwordInfo(drivingInfo)

	module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_DRIVING_SWORD)

	session.Send(&driving_sword_api.AreaMove_Out{
		Events: punchTheShadow(state, uint8(drivingInfo.CurrentX), uint8(drivingInfo.CurrentY)),
	})
}

//Deprecated
func (this DrivingSwordAPI) ExplorerStartBattle(session *net.Session, in *driving_sword_api.ExplorerStartBattle_In) {
	//TODO
	session.Send(&driving_sword_api.ExplorerStartBattle_Out{})
}

//Deprecated
func (this DrivingSwordAPI) ExplorerAward(session *net.Session, in *driving_sword_api.ExplorerAward_In) {
	//TODO
	session.Send(&driving_sword_api.ExplorerAward_Out{})
}

func (this DrivingSwordAPI) ExplorerGarrison(session *net.Session, in *driving_sword_api.ExplorerGarrison_In) {
	doExplorerGarrison(session, in.RoleId)
}

func (this DrivingSwordAPI) VisitMountain(session *net.Session, in *driving_sword_api.VisitMountain_In) {
	visitingMoutian(session, true, nil)
}

func (this DrivingSwordAPI) VisiterStartBattle(session *net.Session, in *driving_sword_api.VisiterStartBattle_In) {
	startBattleVisting(session)
}

//Deprecated
func (this DrivingSwordAPI) VisiterAward(session *net.Session, in *driving_sword_api.VisiterAward_In) {
	//VisitingAward(session)
}

func (this DrivingSwordAPI) MountainTreasureOpen(session *net.Session, in *driving_sword_api.MountainTreasureOpen_In) {
	openMountainTreasure(session, xdlog.ET_MOUNTAIN_TREASURE_OPEN)
}

func (this DrivingSwordAPI) ListGarrisons(session *net.Session, in *driving_sword_api.ListGarrisons_In) {
	doListGarrisons(session)
}

func (this DrivingSwordAPI) AwardGarrison(session *net.Session, in *driving_sword_api.AwardGarrison_In) {
	doAwardGarrison(session, in.RoleId, xdlog.ET_AWARD_GARRISON)
}

func (this DrivingSwordAPI) EndGarrison(session *net.Session, in *driving_sword_api.EndGarrison_In) {
	state := module.State(session)
	x, y, status, cloud := doEndGarrison(state.Database, in.RoleId)
	session.Send(&driving_sword_api.EndGarrison_Out{
		X:       x,
		Y:       y,
		Status:  driving_sword_api.ExploringMountainStatus(status),
		CloudId: cloud,
	})
}

func (this DrivingSwordAPI) BuyAllowedAction(session *net.Session, in *driving_sword_api.BuyAllowedAction_In) {
	buyActionTimes(session)
}
