package global

import (
	"game_server/dat/event_dat"
	"game_server/dat/vip_dat"
)

var (
	playerInfoDataTable *playerInfoTable
)

func Start(mysqlInfo map[string]interface{}) {
	playerInfoDataTable = &playerInfoTable{
		tables:            make(map[int64]PlayerInfo),
		nickMap:           make(map[string]int64),
		openIdMap:         make(map[string]int64),
		VIP:               make([]int32, vip_dat.MAX_VIP_LEVEL+1),
		GroupBuyInfoTable: make(map[int16]int32),
	}
	playerInfoDataTable.load(mysqlInfo)
}

func GetPlayerInfo(pid int64) *PlayerInfo {
	return playerInfoDataTable.GetPlayerInfo(pid)
}

func UpdatePlayerInfoSets(sets map[int64]map[int]int64) {
	playerInfoDataTable.UpdatePlayerInfoSets(sets)
}

func UpdatePlayerInfo(pid int64, info *PlayerInfo) {
	playerInfoDataTable.UpdatePlayerInfo(pid, info)
}

func UpdatePlayerLoginTime(pid int64, time int64) {
	playerInfoDataTable.UpdateLastLoginTime(pid, time)
}

func GetPlayerIdWithOpenId(openid string) (pid int64, ok bool) {
	return playerInfoDataTable.GetPlayerIdWithOpenId(openid)
}

func GetPlayerIdWithNick(nick string) (pid int64, ok bool) {
	return playerInfoDataTable.GetPlayerIdWithNick(nick)
}

func GetPlayerVIPTotal() int {
	return int(playerInfoDataTable.VIP[0])
}

func GetPlayerVIP() []int32 {
	return playerInfoDataTable.VIP
}

func ExistPlayerInfo(pid int64) bool {
	return playerInfoDataTable.ExistPlayerInfo(pid)
}

func GetGroupBuyCount() int32 {
	return playerInfoDataTable.GroupBuyInfoTable[event_dat.EVENT_GROUP_BUY]
}

func OperateGroupBuyCount(isAdd bool) int32 {
	if isAdd {
		playerInfoDataTable.Lock()
		playerInfoDataTable.GroupBuyInfoTable[int16(event_dat.EVENT_GROUP_BUY)]++
		playerInfoDataTable.Unlock()
	}
	return GetGroupBuyCount()
}
