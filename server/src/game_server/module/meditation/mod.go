package meditation

import (
	"core/log"
	"core/net"
	"core/time"
	"game_server/api/protocol/notify_api"
	"game_server/dat/mail_dat"
	"game_server/dat/meditation_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
)

func init() {
	module.Meditation = MeditationMod{}
}

type MeditationMod struct{}

//登录时更新打坐状态（离线未满10分钟）或者发通知（满10分钟）
func (mod MeditationMod) LoginUpdateMeditationState(session *net.Session) {
	state := module.State(session)
	if module.Player.IsOpenFunc(state.Database, player_dat.FUNC_MEDITATION) {
		meditationState := state.Database.Lookup.PlayerMeditationState(state.PlayerId)

		if meditationState.StartTimestamp != meditation_dat.MEDITATION_STOP { //处于打坐状态
			if time.GetNowTime() > meditationState.StartTimestamp { //打坐状态
				session.Send(&notify_api.MeditationState_Out{})
			} else { //离线自动打坐但时间未足够
				meditationState.StartTimestamp = meditation_dat.MEDITATION_STOP
				state.Database.Update.PlayerMeditationState(meditationState)
			}
		}
	}
}

//在城镇中离线设置离线打坐状态
func (mod MeditationMod) LogoutStartMeditation(state *module.SessionState) {
	if state.TownId > 0 && module.Player.IsOpenFunc(state.Database, player_dat.FUNC_MEDITATION) {
		startMeditation(state.Database, time.GetNowTime()+meditation_dat.AUTO_MEDITATE_DELAY, false)
	}
}

func (mod MeditationMod) InMeditationState(db *mdb.Database) bool {
	if !module.Player.IsOpenFunc(db, player_dat.FUNC_MEDITATION) {
		return false
	}
	meditationState := db.Lookup.PlayerMeditationState(db.PlayerId())
	if meditationState == nil {
		//FIXME 存在玩家等级足够但功能相关数据结构没有初始化
		log.Errorf("打坐状态异常 %d", db.PlayerId())
		return false
	}
	return meditationState.StartTimestamp != meditation_dat.MEDITATION_STOP
}

func (mod MeditationMod) OpenFunc(db *mdb.Database) {
	db.Insert.PlayerMeditationState(&mdb.PlayerMeditationState{
		Pid: db.PlayerId(),
	})
	rpc.RemoteMailSend(db.PlayerId(), &mail_dat.MailDaZuoTiXing{})
}
