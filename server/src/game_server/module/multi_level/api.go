package multi_level

import (
	"core/net"
	"game_server/api/protocol/multi_level_api"
	"game_server/module"
)

func init() {
	multi_level_api.SetInHandler(MultiLevelAPI{})
}

type MultiLevelAPI struct {
}

func (api MultiLevelAPI) CreateRoom(session *net.Session, in *multi_level_api.CreateRoom_In) {
	session.Send(&multi_level_api.CreateRoom_Out{
		Status: createRoom(session, in.LevelId),
	})
}

func (api MultiLevelAPI) EnterRoom(session *net.Session, in *multi_level_api.EnterRoom_In) {
	enterRoom(session, in.RoomId)
}

func (api MultiLevelAPI) LeaveRoom(session *net.Session, in *multi_level_api.LeaveRoom_In) {
	leaveRoom(session)
}

func (api MultiLevelAPI) AutoEnterRoom(session *net.Session, in *multi_level_api.AutoEnterRoom_In) {
	autoEnterRoom(session, in.LevelId)
}

func (api MultiLevelAPI) ChangeBuddy(session *net.Session, in *multi_level_api.ChangeBuddy_In) {
	session.Send(&multi_level_api.ChangeBuddy_Out{
		Success: changeBuddy(module.State(session), in.BuddyRoleId),
	})
}

func (api MultiLevelAPI) GetFormInfo(session *net.Session, in *multi_level_api.GetFormInfo_In) {
	rsp := &multi_level_api.GetFormInfo_Out{}
	getFormInfo(module.State(session), rsp)
	session.Send(rsp)
}

func (api MultiLevelAPI) ChangeForm(session *net.Session, in *multi_level_api.ChangeForm_In) {
	session.Send(&multi_level_api.ChangeForm_Out{
		Success: changeForm(module.State(session), in.BuddyRoleRow),
	})
}

func (api MultiLevelAPI) GetInfo(session *net.Session, in *multi_level_api.GetInfo_In) {
	dailyNum, lock := getInfo(module.State(session))
	session.Send(&multi_level_api.GetInfo_Out{
		DailyNum: dailyNum,
		Lock:     lock,
	})
}

func (api MultiLevelAPI) GetOnlineFriend(session *net.Session, in *multi_level_api.GetOnlineFriend_In) {
	getOnlineFriends(session)
}

func (api MultiLevelAPI) InviteIntoRoom(session *net.Session, in *multi_level_api.InviteIntoRoom_In) {
	inviteFriend(session, in.Pid)
}

func (api MultiLevelAPI) RefuseRoomInvite(session *net.Session, in *multi_level_api.RefuseRoomInvite_In) {
	refuseRoomInvite(session, in.RoomId, in.InviterId)
}

func (api MultiLevelAPI) CancelAutoMatch(session *net.Session, in *multi_level_api.CancelAutoMatch_In) {
	cancelAutoMatch(module.State(session))
}
