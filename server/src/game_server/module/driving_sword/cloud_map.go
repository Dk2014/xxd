package driving_sword

import (
	"game_server/api/protocol/driving_sword_api"
	"game_server/dat/driving_sword_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"

	"core/fail"
	"core/time"

	"core/log"

	"encoding/json"
	"math"
	"math/rand"
)

//刷新行动点购买时间
func refreshAllowedActionBought(db *mdb.Database, drivingInfo *mdb.PlayerDrivingSwordInfo) *mdb.PlayerDrivingSwordInfo {
	if !time.IsToday(drivingInfo.ActionBuyTime) {
		drivingInfo.DailyActionBought = 0
		drivingInfo.ActionBuyTime = time.GetNowTime()
		db.Update.PlayerDrivingSwordInfo(drivingInfo)
	}
	return drivingInfo
}

//刷新行动点
func refreshAllowedAction(db *mdb.Database, drivingInfo *mdb.PlayerDrivingSwordInfo) *mdb.PlayerDrivingSwordInfo {
	if !time.IsInPointHour(player_dat.RESET_DRIVING_SWORD_ACTION_TIMES_IN_HOUR, drivingInfo.ActionRefreshTime) {
		if drivingInfo.AllowedAction < driving_sword_dat.DAILY_ADDITIONAL_ACTION_COUNT {
			drivingInfo.AllowedAction = driving_sword_dat.DAILY_ADDITIONAL_ACTION_COUNT
		}
		drivingInfo.ActionRefreshTime = time.GetNowTime()
		db.Update.PlayerDrivingSwordInfo(drivingInfo)
	}
	return drivingInfo
}

//将阴影区和事件列表装载进内存
func cacheCloudMap(state *module.SessionState, cloud_id int16) (cloud driving_sword_api.CloudMap) {
	db := state.Database
	shadows := lookupPlayerCloudShadow(db, cloud_id)
	if shadows == nil {
		module.DrivingSword.ResetCloudData(state, cloud_id)
		shadows = lookupPlayerCloudShadow(db, cloud_id)
	}

	//first initialize the field in SessionState
	state.DrivingSwordMapCache = &module.DrivingSwordMapCache{
		UnknowTeleports: map[uint16]*mdb.PlayerDrivingSwordEvent{},
	}

	//give out shadows
	cloud.Shadows = []byte(shadows.Shadows)

	//give out all events
	//common event
	db.Select.PlayerDrivingSwordEvent(func(row *mdb.PlayerDrivingSwordEventRow) {
		if row.CloudId() == cloud_id {
			event := row.GoObject()
			if driving_sword_api.CommonEvent(event.EventType) == driving_sword_api.COMMON_EVENT_UNKNOW_TELEPORT {
				state.DrivingSwordMapCache.UnknowTeleports[(uint16(event.X)<<8)|uint16(event.Y)] = event
			} else {
				cloud.Events.Common = append(cloud.Events.Common, driving_sword_api.EventAreas_Common{
					X:     uint8(event.X),
					Y:     uint8(event.Y),
					Id:    event.DataId,
					Event: driving_sword_api.CommonEvent(event.EventType),
				})
			}
		}
	})
	//exploring event
	db.Select.PlayerDrivingSwordEventExploring(func(row *mdb.PlayerDrivingSwordEventExploringRow) {
		if row.CloudId() == cloud_id {
			event := row.GoObject()
			garrison_time := event.GarrisonTime
			if event.Status == int8(driving_sword_api.EXPLORING_MOUNTAIN_STATUS_IN_GARRISON) {
				garrison_time, _ = calcGarrisonTime(event)
			}
			cloud.Events.ExploringStatus = append(cloud.Events.ExploringStatus, driving_sword_api.EventAreas_ExploringStatus{
				X:            uint8(event.X),
				Y:            uint8(event.Y),
				Id:           event.DataId,
				Status:       driving_sword_api.ExploringMountainStatus(event.Status),
				GarrisonTime: garrison_time,
			})
		}
	})
	//visiting event
	db.Select.PlayerDrivingSwordEventVisiting(func(row *mdb.PlayerDrivingSwordEventVisitingRow) {
		if row.CloudId() == cloud_id {
			event := row.GoObject()
			target := &targetStatusInfo{}

			if event.TargetPid != 0 {
				err := json.Unmarshal([]byte(event.TargetStatus), target)
				fail.When(err != nil, err)
			}

			cloud.Events.VisitingStatus = append(cloud.Events.VisitingStatus, driving_sword_api.EventAreas_VisitingStatus{
				X:      uint8(event.X),
				Y:      uint8(event.Y),
				Id:     event.DataId,
				Status: driving_sword_api.VisitingMountainStatus(event.Status),
				Pid:    event.TargetPid,

				Nick:            []byte(target.Nick),
				RoleId:          target.RoleId,
				Level:           target.Level,
				FightNum:        target.FightNum,
				FashionId:       target.FashionId,
				FriendshipLevel: target.FriendshipLevel,
			})
		}
	})
	//treasure event
	db.Select.PlayerDrivingSwordEventTreasure(func(row *mdb.PlayerDrivingSwordEventTreasureRow) {
		if row.CloudId() == cloud_id {
			event := row.GoObject()
			cloud.Events.TreasureProgress = append(cloud.Events.TreasureProgress, driving_sword_api.EventAreas_TreasureProgress{
				X:        uint8(event.X),
				Y:        uint8(event.Y),
				Id:       event.DataId,
				Progress: event.Progress,
			})
		}
	})

	return
}

func roundArea(x uint8, y uint8, radius uint8, width, height uint8) (ltx, lty, rbx, rby uint8) {
	if 0+radius > x {
		ltx = 0
	} else {
		ltx = x - radius
	}

	if 0+radius > y {
		rby = 0
	} else {
		rby = y - radius
	}

	if width-radius-1 < x {
		rbx = width - 1
	} else {
		rbx = x + radius
	}

	if height-radius-1 < y {
		lty = height - 1
	} else {
		lty = y + radius
	}

	return
}

func roundPoint(x, y int, width, height uint8) (nx, ny uint8) {
	x = int(math.Max(float64(x), 0))
	x = int(math.Min(float64(x), float64(width-1)))
	y = int(math.Max(float64(y), 0))
	y = int(math.Min(float64(y), float64(height-1)))
	nx = uint8(x)
	ny = uint8(y)
	return
}

func markCross(data []byte, eventx, eventy uint8, width, height uint8) (count uint16) {
	var x, y uint8
	var ind uint
	//left
	x, y = roundPoint(int(eventx)-1, int(eventy), width, height)
	if !isMarkedOrShaded(data, x, y, width) {
		ind = uint(y)*uint(width) + uint(x)
		data[ind/8] |= (1 << (ind % 8))
		count++
	}
	//right
	x, y = roundPoint(int(eventx)+1, int(eventy), width, height)
	if !isMarkedOrShaded(data, x, y, width) {
		ind = uint(y)*uint(width) + uint(x)
		data[ind/8] |= (1 << (ind % 8))
		count++
	}
	//top
	x, y = roundPoint(int(eventx), int(eventy)+1, width, height)
	if !isMarkedOrShaded(data, x, y, width) {
		ind = uint(y)*uint(width) + uint(x)
		data[ind/8] |= (1 << (ind % 8))
		count++
	}
	//bottom
	x, y = roundPoint(int(eventx), int(eventy)-1, width, height)
	if !isMarkedOrShaded(data, x, y, width) {
		ind = uint(y)*uint(width) + uint(x)
		data[ind/8] |= (1 << (ind % 8))
		count++
	}
	//middle
	x = eventx
	y = eventy
	if !isMarkedOrShaded(data, x, y, width) {
		ind = uint(y)*uint(width) + uint(x)
		data[ind/8] |= (1 << (ind % 8))
		count++
	}
	return
}

func markCrossEmpty(data []byte, eventx, eventy uint8, width, height uint8, shadows []byte) (count uint16) {
	var x, y uint8
	var ind uint
	//left
	x, y = roundPoint(int(eventx)-1, int(eventy), width, height)
	if !isMarkedOrShaded(data, x, y, width) && isMarkedOrShaded(shadows, x, y, width) {
		ind = uint(y)*uint(width) + uint(x)
		data[ind/8] |= (1 << (ind % 8))
		count++
	}
	//right
	x, y = roundPoint(int(eventx)+1, int(eventy), width, height)
	if !isMarkedOrShaded(data, x, y, width) && isMarkedOrShaded(shadows, x, y, width) {
		ind = uint(y)*uint(width) + uint(x)
		data[ind/8] |= (1 << (ind % 8))
		count++
	}
	//top
	x, y = roundPoint(int(eventx), int(eventy)+1, width, height)
	if !isMarkedOrShaded(data, x, y, width) && isMarkedOrShaded(shadows, x, y, width) {
		ind = uint(y)*uint(width) + uint(x)
		data[ind/8] |= (1 << (ind % 8))
		count++
	}
	//bottom
	x, y = roundPoint(int(eventx), int(eventy)-1, width, height)
	if !isMarkedOrShaded(data, x, y, width) && isMarkedOrShaded(shadows, x, y, width) {
		ind = uint(y)*uint(width) + uint(x)
		data[ind/8] |= (1 << (ind % 8))
		count++
	}
	return
}

func isMarkedOrShaded(data []byte, x, y uint8, width uint8) bool {
	ind := uint(y)*uint(width) + uint(x)
	if ind/8 >= uint(len(data)) {
		log.Debugf("driving_sword.isMarkOrShaded -- x: %v, y: %v, width: %v, len: %v, ind: %v, ind/8: %v", x, y, width, len(data), ind, ind/8)
	}
	return data[ind/8]&(1<<(ind%8)) != 0
}

func isCrossMarked(data []byte, eventx, eventy uint8, width, height uint8) bool {
	lx, ly := roundPoint(int(eventx)-1, int(eventy), width, height)
	rx, ry := roundPoint(int(eventx)+1, int(eventy), width, height)
	tx, ty := roundPoint(int(eventx), int(eventy)+1, width, height)
	bx, by := roundPoint(int(eventx), int(eventy)-1, width, height)
	return isMarkedOrShaded(data, lx, ly, width) || isMarkedOrShaded(data, rx, ry, width) || isMarkedOrShaded(data, tx, ty, width) || isMarkedOrShaded(data, bx, by, width) || isMarkedOrShaded(data, eventx, eventy, width)
}

//打开相关阴影
func punchTheShadow(state *module.SessionState, heart_x, heart_y uint8) (events driving_sword_api.EventAreas) {
	db := state.Database
	drivingState := state.DrivingSwordMapCache

	const radius = 1

	drivingInfo := db.Lookup.PlayerDrivingSwordInfo(db.PlayerId())
	drivingDat := driving_sword_dat.GetCloudLevel(drivingInfo.CurrentCloud)

	ltx, lty, rbx, rby := roundArea(heart_x, heart_y, radius, uint8(drivingDat.Width), uint8(drivingDat.Height))

	//select and initialize shadows
	shadows := lookupPlayerCloudShadow(db, drivingInfo.CurrentCloud)
	eventmask := lookupPlayerCloudEventMask(db, drivingInfo.CurrentCloud)

	if shadows == nil {
		shadows = &mdb.PlayerDrivingSwordMap{
			Pid:     db.PlayerId(),
			CloudId: drivingInfo.CurrentCloud,
			Shadows: make([]byte, int(math.Ceil(float64(uint16(drivingDat.Width)*uint16(drivingDat.Height))/8.0))),
		}
		eventmask = &mdb.PlayerDrivingSwordEventmask{
			Pid:     db.PlayerId(),
			CloudId: drivingInfo.CurrentCloud,
			Events:  make([]byte, int(math.Ceil(float64(uint16(drivingDat.Width)*uint16(drivingDat.Height))/8.0))),
		}

		mapseed := int(math.Abs(float64(rand.Int())))
		teleport_i := 1
		for y := uint8(0); y < uint8(drivingDat.Height); y++ {
			for x := uint8(0); x < uint8(drivingDat.Width); x++ {
				event_num := driving_sword_dat.GetCloudMapEvent(drivingInfo.CurrentCloud, mapseed, x, y)
				if event_num != 0 {
					ind := uint(y)*uint(drivingDat.Width) + uint(x)
					eventmask.Events[ind/8] |= (1 << (ind % 8))
				}
				//teleport position
				if event_num == 3 {
					teleport := &mdb.PlayerDrivingSwordEvent{
						Pid:       db.PlayerId(),
						CloudId:   drivingInfo.CurrentCloud,
						X:         int8(x),
						Y:         int8(y),
						EventType: int8(driving_sword_api.COMMON_EVENT_UNKNOW_TELEPORT),
						DataId:    int8(teleport_i),
					}
					teleport_i++
					db.Insert.PlayerDrivingSwordEvent(teleport)
				}
			}
		}

		db.Insert.PlayerDrivingSwordMap(shadows)
		db.Insert.PlayerDrivingSwordEventmask(eventmask)

		//recache the map info after initialize this map
		cacheCloudMap(state, drivingInfo.CurrentCloud)
		drivingState = state.DrivingSwordMapCache
	}

	for y := uint16(rby); y <= uint16(lty); y++ {
		for x := uint16(ltx); x <= uint16(rbx); x++ {
			ind := y*uint16(drivingDat.Width) + x
			if shadows.Shadows[ind/8]&(1<<(ind%8)) == 0 {
				shadows.Shadows[ind/8] |= (1 << (ind % 8))

				if eventmask.Events[ind/8]&(1<<(ind%8)) == 0 {
					//there is no event
					continue
				}

				if x == uint16(drivingDat.BornX) && y == uint16(drivingDat.BornY) {
					//the birth place
				} else if x == uint16(drivingDat.HoleX) && y == uint16(drivingDat.HoleY) {
					//first check if this is a hole

					//insert the hole info
					playerHole := &mdb.PlayerDrivingSwordEvent{
						Pid:       db.PlayerId(),
						CloudId:   drivingInfo.CurrentCloud,
						X:         int8(x),
						Y:         int8(y),
						EventType: int8(driving_sword_api.COMMON_EVENT_HOLE),
						DataId:    0,
					}
					db.Insert.PlayerDrivingSwordEvent(playerHole)
					events.Common = append(events.Common, driving_sword_api.EventAreas_Common{
						X:     uint8(playerHole.X),
						Y:     uint8(playerHole.Y),
						Id:    playerHole.DataId,
						Event: driving_sword_api.CommonEvent(playerHole.EventType),
					})

				} else if teleport, exist := drivingState.UnknowTeleports[(x<<8)|y]; exist {
					teleport.EventType = int8(driving_sword_api.COMMON_EVENT_TELEPORT)
					db.Update.PlayerDrivingSwordEvent(teleport)

					delete(drivingState.UnknowTeleports, (x<<8)|y)

					events.Common = append(events.Common, driving_sword_api.EventAreas_Common{
						X:     uint8(teleport.X),
						Y:     uint8(teleport.Y),
						Id:    teleport.DataId,
						Event: driving_sword_api.CommonEvent(teleport.EventType),
					})
				} else {

					obstacle_reminds := driving_sword_dat.CountObstacleByCloud(drivingInfo.CurrentCloud) - int(shadows.ObstacleCount)

					exploring_reminds := driving_sword_dat.CountExploringByCloud(drivingInfo.CurrentCloud) - int(shadows.ExploringCount)

					visiting_reminds := driving_sword_dat.CountVisitingByCloud(drivingInfo.CurrentCloud) - int(shadows.VisitingCount)

					treasure_reminds := driving_sword_dat.CountTreasureByCloud(drivingInfo.CurrentCloud) - int(shadows.TreasureCount)

					if (obstacle_reminds + exploring_reminds + visiting_reminds + treasure_reminds) > 0 {
						probability := rand.Intn(obstacle_reminds + exploring_reminds + visiting_reminds + treasure_reminds)

						if probability < int(obstacle_reminds) {
							//generate obstacle
							shadows.ObstacleCount++
							playerObstacle := &mdb.PlayerDrivingSwordEvent{
								Pid:       db.PlayerId(),
								CloudId:   drivingInfo.CurrentCloud,
								X:         int8(x),
								Y:         int8(y),
								EventType: int8(driving_sword_api.COMMON_EVENT_OBSTACLE),
								DataId:    shadows.ObstacleCount,
							}
							db.Insert.PlayerDrivingSwordEvent(playerObstacle)
							events.Common = append(events.Common, driving_sword_api.EventAreas_Common{
								X:     uint8(playerObstacle.X),
								Y:     uint8(playerObstacle.Y),
								Id:    playerObstacle.DataId,
								Event: driving_sword_api.CommonEvent(playerObstacle.EventType),
							})
						} else if probability -= int(obstacle_reminds); probability < int(exploring_reminds) {
							//generate exploring
							shadows.ExploringCount++
							playerExploring := &mdb.PlayerDrivingSwordEventExploring{
								Pid:     db.PlayerId(),
								CloudId: drivingInfo.CurrentCloud,
								X:       int8(x),
								Y:       int8(y),
								DataId:  shadows.ExploringCount,
								Status:  int8(driving_sword_api.EXPLORING_MOUNTAIN_STATUS_UNEXPLORED),
							}
							db.Insert.PlayerDrivingSwordEventExploring(playerExploring)
							events.ExploringStatus = append(events.ExploringStatus, driving_sword_api.EventAreas_ExploringStatus{
								X:            uint8(playerExploring.X),
								Y:            uint8(playerExploring.Y),
								Id:           playerExploring.DataId,
								Status:       driving_sword_api.ExploringMountainStatus(playerExploring.Status),
								GarrisonTime: playerExploring.GarrisonTime,
							})
						} else if probability -= int(exploring_reminds); probability < int(visiting_reminds) {
							//generate visiting
							shadows.VisitingCount++
							playerVisiting := &mdb.PlayerDrivingSwordEventVisiting{
								Pid:     db.PlayerId(),
								CloudId: drivingInfo.CurrentCloud,
								X:       int8(x),
								Y:       int8(y),
								DataId:  shadows.VisitingCount,
								Status:  int8(driving_sword_api.VISITING_MOUNTAIN_STATUS_UNWIN),
							}
							db.Insert.PlayerDrivingSwordEventVisiting(playerVisiting)
							events.VisitingStatus = append(events.VisitingStatus, driving_sword_api.EventAreas_VisitingStatus{
								X:      uint8(playerVisiting.X),
								Y:      uint8(playerVisiting.Y),
								Id:     playerVisiting.DataId,
								Status: driving_sword_api.VisitingMountainStatus(playerVisiting.Status),
								Pid:    playerVisiting.TargetPid,
							})
						} else if probability -= int(visiting_reminds); probability < int(treasure_reminds) {
							//generate treasure
							shadows.TreasureCount++
							playerTreasure := &mdb.PlayerDrivingSwordEventTreasure{
								Pid:     db.PlayerId(),
								CloudId: drivingInfo.CurrentCloud,
								X:       int8(x),
								Y:       int8(y),
								DataId:  shadows.TreasureCount,
							}
							db.Insert.PlayerDrivingSwordEventTreasure(playerTreasure)
							events.TreasureProgress = append(events.TreasureProgress, driving_sword_api.EventAreas_TreasureProgress{
								X:        uint8(playerTreasure.X),
								Y:        uint8(playerTreasure.Y),
								Id:       playerTreasure.DataId,
								Progress: playerTreasure.Progress,
							})
						}
					}
				}

			}
		}
	}
	//update database
	db.Update.PlayerDrivingSwordMap(shadows)

	return
}

func lookupPlayerCloudShadow(db *mdb.Database, cloud_id int16) (cloud_shadow *mdb.PlayerDrivingSwordMap) {
	db.Select.PlayerDrivingSwordMap(func(row *mdb.PlayerDrivingSwordMapRow) {
		if row.CloudId() == cloud_id {
			cloud_shadow = row.GoObject()
			row.Break()
		}
	})
	return
}

func lookupPlayerCloudEventMask(db *mdb.Database, cloud_id int16) (eventmask *mdb.PlayerDrivingSwordEventmask) {
	db.Select.PlayerDrivingSwordEventmask(func(row *mdb.PlayerDrivingSwordEventmaskRow) {
		if row.CloudId() == cloud_id {
			eventmask = row.GoObject()
			row.Break()
		}
	})
	return
}

func lookupPlayerDrivingCommonEvent(db *mdb.Database, cloud_id int16, x uint8, y uint8) (event *mdb.PlayerDrivingSwordEvent) {
	db.Select.PlayerDrivingSwordEvent(func(row *mdb.PlayerDrivingSwordEventRow) {
		if row.CloudId() == cloud_id && uint8(row.X()) == x && uint8(row.Y()) == y {
			event = row.GoObject()
			row.Break()
		}
	})
	return
}

func lookupPlayerDrivingExploringEvent(db *mdb.Database, cloud_id int16, x uint8, y uint8) (event *mdb.PlayerDrivingSwordEventExploring) {
	db.Select.PlayerDrivingSwordEventExploring(func(row *mdb.PlayerDrivingSwordEventExploringRow) {
		if row.CloudId() == cloud_id && uint8(row.X()) == x && uint8(row.Y()) == y {
			event = row.GoObject()
			row.Break()
		}
	})
	return
}

func lookupPlayerDrivingVisitingEvent(db *mdb.Database, cloud_id int16, x uint8, y uint8) (event *mdb.PlayerDrivingSwordEventVisiting) {
	db.Select.PlayerDrivingSwordEventVisiting(func(row *mdb.PlayerDrivingSwordEventVisitingRow) {
		if row.CloudId() == cloud_id && uint8(row.X()) == x && uint8(row.Y()) == y {
			event = row.GoObject()
			row.Break()
		}
	})
	return
}

func lookupPlayerDrivingTreasureEvent(db *mdb.Database, cloud_id int16, x uint8, y uint8) (event *mdb.PlayerDrivingSwordEventTreasure) {
	db.Select.PlayerDrivingSwordEventTreasure(func(row *mdb.PlayerDrivingSwordEventTreasureRow) {
		if row.CloudId() == cloud_id && uint8(row.X()) == x && uint8(row.Y()) == y {
			event = row.GoObject()
			row.Break()
		}
	})
	return
}
