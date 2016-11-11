package rpc

import (
	"game_server/mdb"
)

/*
	发送单人邮件
*/
type Args_XdgmSendMail struct {
	RPCArgTag
	Pid     int64
	Title   string
	Content string
	Attach  string
	Endtime int64
}

type Reply_XdgmSendMail struct {
}

func RemoteXdgmSendMail(pid int64, sid int, mailtitle, mailcontent, mailattach string, endtime int64, callback func(*Reply_XdgmSendMail, error)) {
	reply := &Reply_XdgmSendMail{}
	args := &Args_XdgmSendMail{
		Pid:     pid,
		Title:   mailtitle,
		Content: mailcontent,
		Attach:  mailattach,
		Endtime: endtime,
	}
	Remote.Call(sid, mdb.RPC_Remote_XdgmSendMail, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	服务器群发邮件
*/
type Args_XdgmSendMailAll struct {
	RPCArgTag
	Title       string
	Content     string
	Attach      string
	Begintime   int64
	Endtime     int64
	MinLevel    int16
	MaxLevel    int16
	MinVipLevel int16
	MaxVipLevel int16
}

type Reply_XdgmSendMailAll struct {
}

func RemoteXdgmSendMailAll(sid int, title, content, attach string, begintime, endtime int64, minlevel, maxlevel, minviplevel, maxviplevel int16, callback func(*Reply_XdgmSendMailAll, error)) {
	reply := &Reply_XdgmSendMailAll{}
	args := &Args_XdgmSendMailAll{
		Title:       title,
		Content:     content,
		Attach:      attach,
		Begintime:   begintime,
		Endtime:     endtime,
		MinLevel:    minlevel,
		MaxLevel:    maxlevel,
		MinVipLevel: minviplevel,
		MaxVipLevel: maxlevel,
	}
	Remote.Call(sid, mdb.RPC_Remote_XdgmSendMailAll, args, reply, func(err error) {
		callback(reply, err)
	})
}
