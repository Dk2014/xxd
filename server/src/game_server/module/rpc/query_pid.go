package rpc

import (
	"fmt"
	"game_server/global"
	"game_server/mdb"
	"rpc_common"
)

func (this *RemoteServe) QueryPid(args *rpc_common.Args_QueryPid, reply *rpc_common.Reply_QueryPid) error {
	return Remote.Serve(mdb.RPC_Remote_QueryPid, args, mdb.TRANS_TAG_RPC_Serve_QueryPid, func() error {
		pid, ok := global.GetPlayerIdWithOpenId(args.OpenId)
		if !ok {
			return fmt.Errorf("can not find open id [%s", args.OpenId)
		}
		reply.Pid = pid
		return nil
	})
}
