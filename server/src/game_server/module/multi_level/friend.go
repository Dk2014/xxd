package multi_level

import (
	"core/fail"
	"core/net"
	"fmt"
	"game_server/api/protocol/multi_level_api"
	"game_server/module"
	"game_server/module/rpc"
)

func getOnlineFriends(session *net.Session) {
	rpc.RemoteGetFriendsWithMultiLevel(module.State(session).PlayerId, func(rsp *multi_level_api.GetOnlineFriend_Out) {
		session.Send(rsp)
	})
}

func inviteFriend(session *net.Session, friendPid int64) {
	state := module.State(session)
	fail.When(state.MultiLevelRoomId == 0, "room not found. can't invite friend")

	room, ok := multiDataTable.getRoom(state.MultiLevelRoomId)
	fail.When(!ok, "room not found. can't invite friend")
	fail.When(room.OnFighting, "room state: fighting. can't invite friend")

	if len(room.Partners) == MAX_ROOM_NUM {
		return
	}

	rpc.RemoteInviteFriendWithMultiLevel(state.PlayerNick, state.PlayerId, friendPid, room.Id, room.LevelId, func(reply *rpc.Reply_InviteFriendWithMultiLevel) {
		// 邀请者邀请反馈
		session.Send(&multi_level_api.InviteIntoRoom_Out{
			IsOffline: reply.IsOnline,
		})
	})
}

func refuseRoomInvite(session *net.Session, roomId int64, inviterId int64) {
	state := module.State(session)
	fmt.Println(state.PlayerNick)
	rpc.RemoteRefuseMultiLevelRoomInvite(roomId, inviterId, state.PlayerNick)
}
