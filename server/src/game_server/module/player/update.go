// 定时更新游戏服发生改变的玩家排行榜数据
// 并推送到互动服
package player

import (
	"core/log"
	"game_server/api/protocol/player_api"
	"game_server/module/player_common"
	"game_server/module/rpc"
	"sync"
	"time"
)

var (
	updateFightNum         *UpdateData
	updateLevel            *UpdateData
	updateMainRoleFightNum *UpdateData
	updateIngot            *UpdateData
)

func init() {
	updateFightNum = NewUpdateData(func(datas []*player_common.PlayerRankData) {
		rpc.RemotePlayerUpdateRankDatas(player_api.RANK_TYPE_FIGHTNUM, datas)
	})
	updateLevel = NewUpdateData(func(datas []*player_common.PlayerRankData) {
		rpc.RemotePlayerUpdateRankDatas(player_api.RANK_TYPE_LEVEL, datas)
	})
	updateMainRoleFightNum = NewUpdateData(func(datas []*player_common.PlayerRankData) {
		rpc.RemotePlayerUpdateRankDatas(player_api.RANK_TYPE_MAINROLE_FIGHTNUM, datas)
	})
	updateIngot = NewUpdateData(func(datas []*player_common.PlayerRankData) {
		rpc.RemotePlayerUpdateRankDatas(player_api.RANK_TYPE_INGOT, datas)
	})
}

type UpdateData struct {
	sync.Mutex
	callTimer      *time.Timer
	datas          []*player_common.PlayerRankData
	updateCallback func([]*player_common.PlayerRankData)
}

func NewUpdateData(updateCallback func([]*player_common.PlayerRankData)) *UpdateData {
	obj := new(UpdateData)
	obj.updateCallback = updateCallback
	obj.reset()
	return obj
}

func (this *UpdateData) reset() {
	this.datas = []*player_common.PlayerRankData{}
	this.callTimer = time.AfterFunc(5*time.Second, this.update)
}

func (this *UpdateData) update() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("despair_rpc.timer.update error", err)
		}
	}()

	this.Lock()
	defer this.Unlock()

	if len(this.datas) > 0 && this.updateCallback != nil {
		this.updateCallback(this.datas)
	}
	this.reset()
}

func (this *UpdateData) AddData(data *player_common.PlayerRankData) {
	this.Lock()
	defer this.Unlock()
	this.datas = append(this.datas, data)
}
