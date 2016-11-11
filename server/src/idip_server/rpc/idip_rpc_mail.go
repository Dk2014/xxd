package rpc

import (
	"game_server/mdb"
)

/*
	发放邮件
*/
type Args_IdipSendMail struct {
	RPCArgTag
	OpenId      string // openid
	MailTitle   string // 邮件标题
	MailContent string // 邮件内容
	SendTime    uint64 // 邮件发送时间
	EndTime     uint64 // 邮件自动删除时间
	ItemDetail  MailItem
}

type MailItem struct {
	ItemId1  uint64 // 邮件赠送道具1ID
	ItemNum1 uint32 // 邮件赠送道具1数量
	ItemId2  uint64 // 邮件赠送道具2ID
	ItemNum2 uint32 // 邮件赠送道具2数量
	ItemId3  uint64 // 邮件赠送道具3ID
	ItemNum3 uint32 // 邮件赠送道具3数量
	ItemId4  uint64 // 邮件赠送道具4ID
	ItemNum4 uint32 // 邮件赠送道具4数量
}

type Reply_IdipSendMail struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func RemoteIdipSendMail(openId string, sid int, mailTitle, mailContent string, sendTime, endTime uint64, mailItem MailItem, callback func(*Reply_IdipSendMail, error)) {
	reply := &Reply_IdipSendMail{}
	args := &Args_IdipSendMail{
		OpenId:      openId,
		MailTitle:   mailTitle,
		MailContent: mailContent,
		SendTime:    sendTime,
		EndTime:     endTime,
		ItemDetail:  mailItem,
	}
	Remote.Call(sid, mdb.RPC_Remote_IdipSendMail, args, reply, func(err error) {
		callback(reply, err)
	})

}

/*
	发放邮件(AQ)
*/
type Args_IdipAqSendMail struct {
	RPCArgTag
	OpenId      string // openid
	MailContent string // 邮件内容
}

type Reply_IdipAqSendMail struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func RemoteIdipAqSendMail(openId string, sid int, mailContent string, callback func(*Reply_IdipAqSendMail, error)) {
	reply := &Reply_IdipAqSendMail{}
	args := &Args_IdipSendMail{
		OpenId:      openId,
		MailContent: mailContent,
	}
	Remote.Call(sid, mdb.RPC_Remote_IdipAqSendMail, args, reply, func(err error) {
		callback(reply, err)
	})

}

/*
	发放全区邮件
*/

func RemoteIdipSendAllMail(sid int, mailTitle, mailContent string, sendTime, endTime uint64, mailItem MailItem, callback func(*Reply_IdipSendMail, error)) {
	reply := &Reply_IdipSendMail{}
	args := &Args_IdipSendMail{
		MailTitle:   mailTitle,
		MailContent: mailContent,
		SendTime:    sendTime,
		EndTime:     endTime,
		ItemDetail:  mailItem,
	}
	Remote.Call(sid, mdb.RPC_Remote_IdipSendAllMail, args, reply, func(err error) {
		callback(reply, err)
	})
}
