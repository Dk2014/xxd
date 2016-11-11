package clique_escort_rpc

import (
	"core/net"
	"core/time"
	"game_server/api/protocol/clique_escort_api"
	"game_server/dat/channel_dat"
	"game_server/mdb"
	"game_server/module"
)

func getCliqueBoatMessages(session *net.Session) {
	state := module.State(session)
	out := &clique_escort_api.GetCliqueBoatMessages_Out{}
	num := 0
	state.Database.Select.PlayerGlobalCliqueEscortMessage(func(row *mdb.PlayerGlobalCliqueEscortMessageRow) {
		num++
		if num > 10 {
			state.Database.Delete.PlayerGlobalCliqueEscortMessage(row.GoObject())
			return
		}
		out.Messages = append(out.Messages, clique_escort_api.GetCliqueBoatMessages_Out_Messages{
			Message: mdbMsgToRespMsg(row.GoObject()),
		})
	})
	session.Send(out)
}

//func newMessage(session *net.Session, db *mdb.Database, msg MessageTpl) {
func newMessage(db *mdb.Database, msg channel_dat.MessageTpl) {
	dbMsg := &mdb.PlayerGlobalCliqueEscortMessage{
		Pid:        db.PlayerId(),
		Timestamp:  time.GetNowTime(),
		TplId:      msg.GetTplId(),
		Parameters: string(msg.GetParameters()),
	}
	db.Insert.PlayerGlobalCliqueEscortMessage(dbMsg)
	if session, online := module.Player.GetPlayerOnline(db.PlayerId()); online {
		out := &clique_escort_api.SendCliqueBoatMessage_Out{
			Message: mdbMsgToRespMsg(dbMsg),
		}
		session.Send(out)
	}
}
