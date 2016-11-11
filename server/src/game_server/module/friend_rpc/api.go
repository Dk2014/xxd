package friend_rpc

import (
	"core/fail"
	"core/net"
	"game_server/api/protocol/friend_api"
	"game_server/global"
	"game_server/module"
)

func init() {
	friend_api.SetInHandler(FriendAPI{})
}

type FriendAPI struct {
}

func (this FriendAPI) GetFriendList(session *net.Session, in *friend_api.GetFriendList_In) {
	out := GetFriendList(module.State(session))
	session.Send(out)
}

func (this FriendAPI) ListenByNick(session *net.Session, in *friend_api.ListenByNick_In) {
	pid, ok := global.GetPlayerIdWithNick(string(in.Nick))
	fail.When(!ok, "can't found player by nick with ListenByNick")
	out := ListenByPid(module.State(session), pid)
	session.Send(out)
}

func (this FriendAPI) CancelListen(session *net.Session, in *friend_api.CancelListen_In) {
	result := CancelListen(module.State(session), in.Pid)
	session.Send(&friend_api.CancelListen_Out{result})
}

func (this FriendAPI) SendHeart(session *net.Session, in *friend_api.SendHeart_In) {
	SendHeart(module.State(session), in.FriendType, string(in.Nickname), in.Pid)
	session.Send(&friend_api.SendHeart_Out{})
}
func (this FriendAPI) SendHeartToAllFriends(session *net.Session, in *friend_api.SendHeartToAllFriends_In) {
	SendHeartToAllFriends(module.State(session), in)
	session.Send(&friend_api.SendHeartToAllFriends_Out{})
}
func (this FriendAPI) Chat(session *net.Session, in *friend_api.Chat_In) {
	session.Send(&friend_api.Chat_Out{
		Banned: Chat(module.State(session), in),
	})
}

func (this FriendAPI) GetChatHistory(session *net.Session, in *friend_api.GetChatHistory_In) {
	out := GetChatHistory(module.State(session), in)
	session.Send(out)
}

func (this FriendAPI) Block(session *net.Session, in *friend_api.Block_In) {
	Block(module.State(session), in.Pid)
	session.Send(&friend_api.Block_Out{})
}

func (this FriendAPI) CancelBlock(session *net.Session, in *friend_api.CancelBlock_In) {
	CancelBlock(module.State(session), in.Pid)
	session.Send(&friend_api.CancelBlock_Out{})
}

func (this FriendAPI) CleanChatHistory(session *net.Session, in *friend_api.CleanChatHistory_In) {
	CleanChatHistory(module.State(session), in.Pid)
	session.Send(&friend_api.CleanChatHistory_Out{})
}

func (this FriendAPI) CurrentPlatformFriendNum(session *net.Session, in *friend_api.CurrentPlatformFriendNum_In) {
	CurrentPlatformFriendNum(module.State(session), in.Num)
}

//获取赠送爱心情况
func (this FriendAPI) GetSendHeartList(session *net.Session, in *friend_api.GetSendHeartList_In) {
	out := new(friend_api.GetSendHeartList_Out)
	GetSendHeartList(module.State(session), out)
	session.Send(out)
}

//通过openid获取好友信息
func (this FriendAPI) GetPlayerByFacebook(session *net.Session, in *friend_api.GetPlayerByFacebook_In) {
	out := new(friend_api.GetPlayerByFacebook_Out)
	GetPlayerByFacebook(module.State(session), in, out)
	session.Send(out)
}
