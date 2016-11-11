package clique_rpc

import (
	"game_server/mdb"
	"game_server/module"
)

func init() {
	mdb.Hook.GlobalClique(globalCliqueHook{})
	mdb.Hook.PlayerGlobalCliqueInfo(playerGlobalCliqueInfoHook{})
}

type globalCliqueHook struct{}

func (hook globalCliqueHook) Load(row *mdb.GlobalCliqueRow) {
	g_CliqueTable.addCliqueName(row)
	g_CliqueTable.addClique(row)
}

func (hook globalCliqueHook) Insert(row *mdb.GlobalCliqueRow) {
	//帮派名字缓存
	g_CliqueTable.addCliqueName(row)

	//帮派缓存（内容在玩家帮派信息的钩子出增加） 创建帮派需要保证先插入帮派后更新玩家数据
	g_CliqueTable.addClique(row)
}

func (hook globalCliqueHook) Update(row, old *mdb.GlobalCliqueRow) {
	g_CliqueTable.addClique(row)
}

func (hook globalCliqueHook) Delete(row *mdb.GlobalCliqueRow) {
	//TODO
	//帮派删除时需要做
	// 1. 更新玩家表
	// 2. 更新帮派缓存
	g_CliqueTable.DeleteClique(row.Id())
	g_CliqueTable.deleteCliqueName(row)
}

type playerGlobalCliqueInfoHook struct{}

func (hook playerGlobalCliqueInfoHook) Load(row *mdb.PlayerGlobalCliqueInfoRow) {
	g_CliqueTable.addMember(row)
}

func (hook playerGlobalCliqueInfoHook) Insert(row *mdb.PlayerGlobalCliqueInfoRow) {
	if row.CliqueId() > 0 {
		//1. 帮派成员 cache
		g_CliqueTable.addMember(row)
		//2. 帮派申请 cache
		g_CliqueTable.DeleteApply(row.Pid(), row.CliqueId())
		//3. 玩家申请 cache
		DeleteAllApply(row.Pid())
	}
}

//玩家帮派表变化时更新cache
func (hook playerGlobalCliqueInfoHook) Update(row, old *mdb.PlayerGlobalCliqueInfoRow) {
	pid := row.Pid()
	cliqueId := row.CliqueId()
	oldCliqueId := old.CliqueId()
	if cliqueId > 0 {
		g_CliqueTable.updateContrib(row)
	}

	if cliqueId == oldCliqueId {
		return
	}

	if cliqueId > 0 {
		//1. 帮派成员 cache
		g_CliqueTable.addMember(row)
		//2. 帮派申请 cache
		g_CliqueTable.DeleteApply(pid, cliqueId)
		//3. 玩家申请 cache
		DeleteAllApply(pid)

		//if memberSession, online := module.Player.GetPlayerOnline(pid); online {
		//	JoinCliqueChannel(cliqueId, memberSession)
		//}
	}
	if oldCliqueId > 0 {
		//1. 删除已退出的帮派的缓存
		g_CliqueTable.DeleteMember(oldCliqueId, pid)
		//2. 离开原帮派频道
		if memberSession, online := module.Player.GetPlayerOnline(pid); online {
			module.CliqueRPC.LeaveCliqueChannel(oldCliqueId, memberSession)
		}
	}
}

func (hook playerGlobalCliqueInfoHook) Delete(row *mdb.PlayerGlobalCliqueInfoRow) {
}
