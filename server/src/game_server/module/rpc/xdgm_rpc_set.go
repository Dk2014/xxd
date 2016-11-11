package rpc

import (
	"core/fail"
	"core/mysql"
	"fmt"
	. "game_server/config"
	"game_server/dat/payments_rule_dat"
	"game_server/mdb"
	. "game_server/rpc"
)

/*
	修改充值返利规则
*/
type Args_XdgmSetPaymentsPresent struct {
	RPCArgTag
	Rule      string
	BeginTime int64
	EndTime   int64
}

type Reply_XdgmSetPaymentsPresent struct {
}

var PresentDB *mysql.Connection
var PresentStmt *mysql.Stmt

func (this *RemoteServe) XdgmSetPaymentsPresent(args *Args_XdgmSetPaymentsPresent, reply *Reply_XdgmSetPaymentsPresent) error {
	return Remote.Serve(mdb.RPC_Remote_XdgmSetPaymentsPresent, args, mdb.TRANS_TAG_RPC_Serve_XdgmSetPaymentsPresent, func() error {
		defer PresentStmt.Close()
		defer PresentDB.Close()
		var err error
		if PresentDB, err = mysql.Connect(GetDBConfig()); err != nil {
			fail.When(true, fmt.Sprintf("onlineCount connect mysql error: ", err))
		}
		PresentStmt = mdb.PreparePresentRuleInfoToPresentRuleTable(PresentDB, "payments_rule")
		err = mdb.DoUpdateRule(PresentStmt, args.Rule, args.BeginTime, args.EndTime)
		if err != nil {
			PresentDB = nil
			fail.When(true, err)
		}
		payments_rule_dat.ModifyRule(args.Rule, args.BeginTime, args.EndTime)
		return nil
	})
}
