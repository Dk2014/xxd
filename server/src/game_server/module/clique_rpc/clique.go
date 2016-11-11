package clique_rpc

import (
	"core/fail"
	"core/i18l"
	"core/net"
	"core/time"
	"core/valid_str"
	"game_server/api/protocol/clique_api"
	"game_server/dat/channel_dat"
	"game_server/dat/clique_dat"
	"game_server/dat/mail_dat"
	"game_server/dat/player_dat"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/xdlog"
)

func createClique(session *net.Session, name, announce string) {
	err := valid_str.StripEmoji(name)
	fail.When(err != nil, err)
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo != nil && playerCliqueInfo.CliqueId > 0 {
		//已加入帮派
		session.Send(&clique_api.CreateClique_Out{
			Result: clique_api.CREATE_CLIQUE_RESULT_HAVE_CLIQUE,
		})

		return
	}
	if nameOccupied := CacheCheckCliqueByName(name); nameOccupied {
		session.Send(&clique_api.CreateClique_Out{
			Result: clique_api.CREATE_CLIQUE_RESULT_DUP_NAME,
		})
		return
	}
	clique := &mdb.GlobalClique{
		Name:                name,
		Announce:            announce,
		OwnerPid:            state.PlayerId,
		CenterBuildingLevel: 1,
		OwnerLoginTime:      time.GetNowTime(),
		ContribClearTime:    time.GetNowTime(),
	}
	state.Database.Insert.GlobalClique(clique)
	if playerCliqueInfo == nil {
		playerCliqueInfo = &mdb.PlayerGlobalCliqueInfo{
			Pid:      state.PlayerId,
			CliqueId: clique.Id,
			JoinTime: time.GetNowTime(),
		}
		state.Database.Insert.PlayerGlobalCliqueInfo(playerCliqueInfo)
		playerCliqueBuilding := &mdb.PlayerGlobalCliqueBuilding{
			Pid: state.PlayerId,
		}
		state.Database.Insert.PlayerGlobalCliqueBuilding(playerCliqueBuilding)
	} else {
		playerCliqueInfo.CliqueId = clique.Id
		playerCliqueInfo.JoinTime = time.GetNowTime()
		state.Database.Update.PlayerGlobalCliqueInfo(playerCliqueInfo)
	}

	module.Chan.AddWorldTplMessage(state.PlayerId, state.PlayerNick, module.CHANNEL_CLIQUE_MESSAGE, channel_dat.MessageFoundClique{
		Player:    channel_dat.ParamPlayer{state.PlayerNick, state.PlayerId},
		Clique:    channel_dat.ParamClique{clique.Name, clique.Id},
		DummyLink: channel_dat.ParamClique{i18l.T.Tran("立即申请加入"), clique.Id},
	})
	rpc.RemoteDecMoney(state.PlayerId, clique_dat.CREATE_CLIQUE_COST, player_dat.INGOT, 0, xdlog.ET_CLIQUE, func() {})
	session.Send(&clique_api.CreateClique_Out{
		Result: clique_api.CREATE_CLIQUE_RESULT_SUCCESS,
	})
	//把帮主加入频道
	JoinCliqueChannel(clique.Id, session)
	joinCliqueCleanup(state.Database, playerCliqueInfo, clique)
	return
}

func applyJoinClique(session *net.Session, cliqueId int64) {
	state := module.State(session)

	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil {
		playerCliqueInfo = &mdb.PlayerGlobalCliqueInfo{
			Pid: state.PlayerId,
		}
		state.Database.Insert.PlayerGlobalCliqueInfo(playerCliqueInfo)
		playerCliqueBuilding := &mdb.PlayerGlobalCliqueBuilding{
			Pid: state.PlayerId,
		}
		state.Database.Insert.PlayerGlobalCliqueBuilding(playerCliqueBuilding)
	}

	//已加入帮派
	if playerCliqueInfo.CliqueId > 0 {
		session.Send(&clique_api.ApplyJoinClique_Out{
			Result:   clique_api.APPLY_CLIQUE_RESULT_ALREADY_JOIN,
			CliqueId: cliqueId,
		})
		return
	}

	//超过申请的帮派个数
	if ApplyFull(state.PlayerId) {
		session.Send(&clique_api.ApplyJoinClique_Out{
			Result:   clique_api.APPLY_CLIQUE_RESULT_TOO_MUCH_APPLY,
			CliqueId: cliqueId,
		})
		return
	}

	cliqueCache := CacheGetCliqueInfo(cliqueId)
	if cliqueCache == nil {
		//帮派不存在
		session.Send(&clique_api.ApplyJoinClique_Out{
			Result:   clique_api.APPLY_CLIQUE_RESULT_NOT_EXIST,
			CliqueId: cliqueId,
		})
		return
	}

	if len(cliqueCache.JoinApplies) >= 200 {
		//帮派申请人数过多
		session.Send(&clique_api.ApplyJoinClique_Out{
			Result:   clique_api.APPLY_CLIQUE_RESULT_REFUSE,
			CliqueId: cliqueId,
		})
		return
	}
	cliqueInfo := cliqueInfoLookUp(state.Database, cliqueId)
	playerInfo := global.GetPlayerInfo(state.PlayerId)
	cliqueLevelDat := clique_dat.GetCenterBuildingLevelInfo(cliqueInfo.CenterBuildingLevel)
	if cliqueCache.MemberNum < cliqueLevelDat.MaxMember && cliqueInfo.AutoAudit > 0 && playerInfo.RoleLevel >= cliqueInfo.AutoAuditLevel {
		//自动审核加入
		playerCliqueInfo.CliqueId = cliqueInfo.Id
		playerCliqueInfo.JoinTime = time.GetNowTime()
		state.Database.Update.PlayerGlobalCliqueInfo(playerCliqueInfo)
		//批准后，加入帮派channel,以便广播
		JoinCliqueChannel(cliqueId, session)
		joinSuccess := &clique_api.NotifyJoincliqueSuccess_Out{}
		joinSuccess.Pidlist = append(joinSuccess.Pidlist, clique_api.NotifyJoincliqueSuccess_Out_Pidlist{
			Pid:    state.PlayerId,
			RoleId: playerInfo.RoleId,
			Level:  playerInfo.RoleLevel,
			Nick:   state.PlayerNick,
		})
		module.CliqueRPC.AddCliqueNews(playerCliqueInfo.CliqueId, channel_dat.MessageCliqueNewMember{
			NewMember: channel_dat.ParamPlayer{
				Nick: state.PlayerNick,
				Pid:  state.PlayerId,
			},
		})
		module.CliqueRPC.Broadcast(cliqueId, joinSuccess)

		session.Send(&clique_api.ApplyJoinClique_Out{
			Result:   clique_api.APPLY_CLIQUE_RESULT_SUCCESS,
			CliqueId: cliqueId,
		})
		//加入后，发邮件通知
		rpc.RemoteMailSend(state.PlayerId, &mail_dat.Mailcliquejoin{Name: cliqueInfo.Name})

		joinCliqueCleanup(state.Database, playerCliqueInfo, cliqueInfo)
		return
	}

	playerRank := module.ArenaRPC.GetPlayerRank(state.PlayerId)
	g_CliqueTable.AddApply(state.PlayerId, cliqueId, &JoinApply{
		Pid:       state.PlayerId,
		Level:     playerInfo.RoleLevel,
		Nick:      string(state.PlayerNick),
		Timestamp: time.GetNowTime(),
		ArenaRank: playerRank,
	})

	//插入用户申请队列里

	AddApply(state.PlayerId, cliqueId)

	session.Send(&clique_api.ApplyJoinClique_Out{
		Result:   clique_api.APPLY_CLIQUE_RESULT_SUCCESS,
		CliqueId: cliqueId,
	})
	return
}

func cancelApplyClique(session *net.Session, cliqueId int64) {
	state := module.State(session)
	out := &clique_api.CancelApplyClique_Out{CliqueId: cliqueId}

	cliqueCache := CacheGetCliqueInfo(cliqueId)
	if cliqueCache == nil {
		out.Result = clique_api.CANCEL_APPLY_CLIQUE_RESULT_NOT_EXIST
		session.Send(out)
		return
	}
	g_CliqueTable.DeleteApply(state.PlayerId, cliqueId)
	//删除用户申请列表里的ID
	DeleteApply(state.PlayerId, cliqueId)

	out.Result = clique_api.CANCEL_APPLY_CLIQUE_RESULT_SUCCESS
	session.Send(out)
}

//成员管理
func mangeMember(sessoin *net.Session, in *clique_api.MangeMember_In) {
	state := module.State(sessoin)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)

	out := &clique_api.MangeMember_Out{
		Pid:    in.Pid,
		Action: in.Action,
		Result: clique_api.MANGE_MEMBER_RESULT_NO_PERMISSION,
	}
	cliqueInfo := cliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)
	var result bool

	switch in.Action {
	case clique_api.MANGE_MEMBER_ACTION_SET_OWNER: //传位，要求被传人为副帮主

		if !isOwer(cliqueInfo, state.PlayerId) || !isManger(cliqueInfo, in.Pid) {
			break
		}

		cliqueInfo.OwnerPid = in.Pid //帮主位置替换
		if cliqueInfo.MangerPid1 == in.Pid {
			cliqueInfo.MangerPid1 = 0
		} else {
			cliqueInfo.MangerPid2 = 0
		}
		state.Database.Update.GlobalClique(cliqueInfo)

		result = true

		playerInfo := global.GetPlayerInfo(in.Pid)
		rpc.RemoteMailSend(in.Pid, &mail_dat.Mailcliquechangeowner{
			Name: string(state.PlayerNick),
		})
		module.CliqueRPC.AddCliqueNews(playerCliqueInfo.CliqueId, channel_dat.MessageCliqueAssignOwner{
			OldOwner: channel_dat.ParamPlayer{
				Nick: state.PlayerNick,
				Pid:  state.PlayerId,
			},
			NewOwner: channel_dat.ParamPlayer{
				Nick: playerInfo.PlayerNick,
				Pid:  in.Pid,
			},
		})
	case clique_api.MANGE_MEMBER_ACTION_SET_MANGER: //设置副版主

		if !isOwer(cliqueInfo, state.PlayerId) || isManger(cliqueInfo, in.Pid) {
			break
		}
		// 已有2个副版主
		fail.When((cliqueInfo.MangerPid1 > 0 && cliqueInfo.MangerPid2 > 0), "已有2个副版主")
		if cliqueInfo.MangerPid1 > 0 {
			cliqueInfo.MangerPid2 = in.Pid
		} else {
			cliqueInfo.MangerPid1 = in.Pid
		}
		state.Database.Update.GlobalClique(cliqueInfo)
		playerInfo := global.GetPlayerInfo(in.Pid)
		rpc.RemoteMailSend(in.Pid, &mail_dat.Mailcliquebemanger{})
		module.CliqueRPC.AddCliqueNews(playerCliqueInfo.CliqueId, channel_dat.MessageCliqueAssignManger{
			Player: channel_dat.ParamPlayer{
				Nick: playerInfo.PlayerNick,
				Pid:  in.Pid,
			},
		})
		result = true
	case clique_api.MANGE_MEMBER_ACTION_SET_MEMBER: //把副版主设置为成员
		if !isOwer(cliqueInfo, state.PlayerId) {
			out.Result = clique_api.MANGE_MEMBER_RESULT_NO_PERMISSION
			break
		}

		if !isManger(cliqueInfo, in.Pid) {
			out.Result = clique_api.MANGE_MEMBER_RESULT_NOT_EXIST
			break
		}

		if cliqueInfo.MangerPid1 == in.Pid {
			cliqueInfo.MangerPid1 = 0
		} else {
			cliqueInfo.MangerPid2 = 0
		}
		state.Database.Update.GlobalClique(cliqueInfo)

		result = true

		playerInfo := global.GetPlayerInfo(in.Pid)
		rpc.RemoteMailSend(in.Pid, &mail_dat.Mailcliquecancelmanager{})
		module.CliqueRPC.AddCliqueNews(playerCliqueInfo.CliqueId, channel_dat.MessageCliqueFireManger{
			Manger: channel_dat.ParamPlayer{
				Nick: playerInfo.PlayerNick,
				Pid:  in.Pid,
			},
		})

	case clique_api.MANGE_MEMBER_ACTION_KICKOFF:
		if !isOwer(cliqueInfo, state.PlayerId) {
			break
		}
		exist := true
		//deleteCliqueMember(state, in.Pid, playerCliqueInfo.CliqueId)
		state.Database.AgentExecute(in.Pid, func(agentDB *mdb.Database) {
			kickInfo := agentDB.Lookup.PlayerGlobalCliqueInfo(in.Pid)
			if kickInfo == nil || kickInfo.CliqueId != cliqueInfo.Id {
				exist = false
				return
			}
			kickBuildingInfo := agentDB.Lookup.PlayerGlobalCliqueBuilding(in.Pid)
			//加入不满48小时的玩家退出会清空建筑捐献 FIXME 去除魔数
			if time.GetNowTime()-kickInfo.JoinTime <= 86400*2 {
				withdrawPlayerCliqueContrib(kickBuildingInfo, cliqueInfo)
			}
			resetPlayerCliqueState(kickInfo, kickBuildingInfo, cliqueInfo)
			resetPlayerCliqueContrib(agentDB)
			agentDB.Update.PlayerGlobalCliqueInfo(kickInfo)
			agentDB.Update.PlayerGlobalCliqueBuilding(kickBuildingInfo)
			leaveCliqueCleanup(agentDB)
		})
		if !exist {
			out.Result = clique_api.MANGE_MEMBER_RESULT_NOT_EXIST
			break
		}

		if cliqueInfo.MangerPid1 == in.Pid {
			cliqueInfo.MangerPid1 = 0
		} else if cliqueInfo.MangerPid2 == in.Pid {
			cliqueInfo.MangerPid2 = 0
		}
		state.Database.Update.GlobalClique(cliqueInfo)

		rpc.RemoteMailSend(in.Pid, &mail_dat.Mailcliqueleave{
			Name: cliqueInfo.Name,
		})

		module.CliqueRPC.Broadcast(cliqueInfo.Id, &clique_api.NotifyLeaveClique_Out{
			Pid:    in.Pid,
			Reason: clique_api.NOTIFY_LEAVE_CLIQUE_REASON_KICKOFF,
		})
		if memberSession, online := module.Player.GetPlayerOnline(in.Pid); online {
			module.CliqueRPC.LeaveCliqueChannel(cliqueInfo.Id, memberSession)
		}

		playerInfo := global.GetPlayerInfo(in.Pid)
		module.CliqueRPC.AddCliqueNews(playerCliqueInfo.CliqueId, channel_dat.MessageCliqueKickMember{
			Manger: channel_dat.ParamPlayer{
				Nick: state.PlayerNick,
				Pid:  state.PlayerId,
			},
			Member: channel_dat.ParamPlayer{
				Nick: playerInfo.PlayerNick,
				Pid:  in.Pid,
			},
		})

		result = true

	default:
		fail.When(true, "wrong   clique   MangeMember  type ")
	}

	if result {
		out.Result = clique_api.MANGE_MEMBER_RESULT_SUCCESS
		if in.Action != clique_api.MANGE_MEMBER_ACTION_KICKOFF {
			module.CliqueRPC.Broadcast(cliqueInfo.Id, &clique_api.NotifyCliqueMangerAction_Out{
				Actiontype: in.Action,
				Pid:        in.Pid,
			})
		}
	}

	sessoin.Send(out)
	return
}

//解散帮派
func destoryClique(session *net.Session, in *clique_api.DestoryClique_In) {
	state := module.State(session)
	out := &clique_api.DestoryClique_Out{}
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		session.Send(out)
		return
	}
	cliqueInfo := cliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)
	//帮主
	fail.When(!isOwer(cliqueInfo, state.PlayerId), "just only Ower can Destory Clique")

	cliqueCache := CacheGetCliqueInfo(playerCliqueInfo.CliqueId)
	//删除玩家申请
	for pid, _ := range cliqueCache.JoinApplies {
		DeleteApply(pid, playerCliqueInfo.CliqueId)
	}
	for _, member := range cliqueCache.Members {
		if member != nil {
			state.Database.AgentExecute(member.Pid, func(agentDB *mdb.Database) {
				kickInfo := agentDB.Lookup.PlayerGlobalCliqueInfo(member.Pid)
				kickBuildingInfo := agentDB.Lookup.PlayerGlobalCliqueBuilding(member.Pid)
				resetPlayerCliqueState(kickInfo, kickBuildingInfo, cliqueInfo)
				resetPlayerCliqueContrib(agentDB)
				agentDB.Update.PlayerGlobalCliqueInfo(kickInfo)
				agentDB.Update.PlayerGlobalCliqueBuilding(kickBuildingInfo)
				leaveCliqueCleanup(agentDB)
				//清楚帮派任务
				rpc.RemoteMailSend(member.Pid, &mail_dat.Mailcliquedestory{
					Name: cliqueInfo.Name,
				})
			})
			if memberSession, online := module.Player.GetPlayerOnline(member.Pid); online {
				memberSession.Send(&clique_api.NotifyLeaveClique_Out{
					Pid:    member.Pid,
					Reason: clique_api.NOTIFY_LEAVE_CLIQUE_REASON_COLLAPSE,
				})
			}
		}
	}
	//删全局的信息
	state.Database.Delete.GlobalClique(cliqueInfo)
	session.Send(out)

	return
}

//更新帮派公告
func updateAnnounce(session *net.Session, in *clique_api.UpdateAnnounce_In) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	fail.When(playerCliqueInfo == nil || !isCliqueAuthority(state, state.PlayerId, playerCliqueInfo.CliqueId), "no clique info")

	cliqueInfo := cliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)
	cliqueInfo.Announce = string(in.Content)
	state.Database.Update.GlobalClique(cliqueInfo)

	session.Send(&clique_api.UpdateAnnounce_Out{})

	//通知群成员
	module.CliqueRPC.Broadcast(playerCliqueInfo.CliqueId, &clique_api.NotifyCliqueAnnounce_Out{
		Announce: in.Content,
	})
	module.CliqueRPC.AddCliqueNews(playerCliqueInfo.CliqueId, channel_dat.MessageCliqueNewAnnc{})

	return
}

//离开帮派
func leaveClique(session *net.Session) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	fail.When(playerCliqueInfo == nil, "未加入任何帮派")

	if playerCliqueInfo.CliqueId <= 0 {
		//无需退出
		session.Send(&clique_api.LeaveClique_Out{Success: true})
		return
	}

	cliqueInfo := cliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)

	if isOwer(cliqueInfo, state.PlayerId) {
		//帮主不能退出
		session.Send(&clique_api.LeaveClique_Out{Success: false})
		return
	}

	if cliqueInfo.MangerPid1 == state.PlayerId {
		cliqueInfo.MangerPid1 = 0
	} else if cliqueInfo.MangerPid2 == state.PlayerId {
		cliqueInfo.MangerPid2 = 0
	}
	module.CliqueRPC.Broadcast(cliqueInfo.Id, &clique_api.NotifyLeaveClique_Out{
		Pid:    state.PlayerId,
		Reason: clique_api.NOTIFY_LEAVE_CLIQUE_REASON_LEAVE,
	})
	module.CliqueRPC.LeaveCliqueChannel(cliqueInfo.Id, session)

	playerCliqueBuilding := state.Database.Lookup.PlayerGlobalCliqueBuilding(state.PlayerId)
	//加入不满48小时的玩家退出会清空建筑捐献 FIXME 去除魔数
	if time.GetNowTime()-playerCliqueInfo.JoinTime <= 86400*2 {
		withdrawPlayerCliqueContrib(playerCliqueBuilding, cliqueInfo)
	}

	module.CliqueRPC.AddCliqueNews(playerCliqueInfo.CliqueId, channel_dat.MessageCliqueMemberLeave{
		Player: channel_dat.ParamPlayer{
			Nick: state.PlayerNick,
			Pid:  state.PlayerId,
		},
	})

	resetPlayerCliqueState(playerCliqueInfo, playerCliqueBuilding, cliqueInfo)
	resetPlayerCliqueContrib(state.Database)
	state.Database.Update.PlayerGlobalCliqueInfo(playerCliqueInfo)
	state.Database.Update.PlayerGlobalCliqueBuilding(playerCliqueBuilding)
	state.Database.Update.GlobalClique(cliqueInfo)
	session.Send(&clique_api.LeaveClique_Out{Success: true})
	leaveCliqueCleanup(state.Database)

	return
}

//帮派列表
func listClique(session *net.Session, in *clique_api.ListClique_In) {
	state := module.State(session)
	out := &clique_api.ListClique_Out{}
	var (
		count, limit, index, num int16
	)

	//清过期的
	if cliqueid, ok := g_PlayerCliqueApplyTable.applyMap[state.PlayerId]; ok {
		for _, id := range cliqueid {
			if id > 0 {
				DeleteTimeOutApply(id)
			}
		}
	}

	//申请过的帮派
	if existCliqueid, exist := g_PlayerCliqueApplyTable.applyMap[state.PlayerId]; exist {
		for _, values := range existCliqueid {
			if values > 0 {
				out.AppliedCliques = append(out.AppliedCliques, clique_api.ListClique_Out_AppliedCliques{CliqueId: values})
				count++
			}
		}
	}

	//帮派信息
	if count > 0 {
		state.Database.Select.GlobalClique(func(row *mdb.GlobalCliqueRow) {
			for _, cliques := range out.AppliedCliques {
				if cliques.CliqueId == row.Id() {
					cliqueCache := CacheGetCliqueInfo(row.Id())
					playerInfo := global.GetPlayerInfo(row.OwnerPid())
					out.Cliques = append(out.Cliques, clique_api.ListClique_Out_Cliques{
						Id:        row.Id(),
						Name:      []byte(row.Name()),
						Level:     row.CenterBuildingLevel(),
						MemberNum: cliqueCache.MemberNum,
						OwnerNick: []byte(playerInfo.PlayerNick),
						OwnerPid:  row.OwnerPid(),
						Announce:  []byte(row.Announce()),
					})
				}
			}
		})
	}

	recordNum := clique_dat.CLIQUE_CLIENT_NUM - count
	offset := recordNum * (in.Offset / clique_dat.CLIQUE_CLIENT_NUM)

	//帮派信息
	state.Database.Select.GlobalClique(func(row *mdb.GlobalCliqueRow) {
		if num >= offset && limit < recordNum {
			var exist bool
			for _, v := range out.AppliedCliques {
				if v.CliqueId == row.Id() {
					exist = true
					break
				}
			}
			//在申请列表里存在，就跳过
			if !exist {
				cliqueCache := CacheGetCliqueInfo(row.Id())
				playerInfo := global.GetPlayerInfo(row.OwnerPid())
				out.Cliques = append(out.Cliques, clique_api.ListClique_Out_Cliques{
					Id:        row.Id(),
					Name:      []byte(row.Name()),
					Level:     row.CenterBuildingLevel(),
					MemberNum: cliqueCache.MemberNum,
					OwnerNick: []byte(playerInfo.PlayerNick),
					OwnerPid:  row.OwnerPid(),
					Announce:  []byte(row.Announce()),
				})

				limit++
			}
		}

		var justexist bool //此记录是否在置顶中
		for _, v := range out.AppliedCliques {
			if v.CliqueId == row.Id() {
				justexist = true
				break
			}
		}
		if !justexist {
			num++
		}

		index++
	})

	//总的帮派数
	out.Total = index

	session.Send(out)

	return
}

//帮派基本信息(帮外玩家申请加入时用)
func cliquePublicInfo(session *net.Session, in *clique_api.CliquePublicInfo_In) {
	state := module.State(session)

	cliqueInfo := cliqueInfoLookUp(state.Database, in.CliqueId)
	if cliqueInfo != nil {
		playerInfo := global.GetPlayerInfo(cliqueInfo.OwnerPid)
		cliqueCache := CacheGetCliqueInfo(in.CliqueId)
		out := &clique_api.CliquePublicInfo_Out{
			CliqueId:             in.CliqueId,
			Exist:                true,
			Name:                 []byte(cliqueInfo.Name),
			OwnerNick:            playerInfo.PlayerNick,
			OwnerPid:             cliqueInfo.OwnerPid,
			MemberNum:            cliqueCache.MemberNum,
			Level:                cliqueInfo.CenterBuildingLevel,
			Announce:             []byte(cliqueInfo.Announce),
			CenterBuildingLevel:  cliqueInfo.CenterBuildingLevel,
			TempleBuildingLevel:  cliqueInfo.TempleBuildingLevel,
			BankBuildingLevel:    cliqueInfo.BankBuildingLevel,
			HealthBuildingLevel:  cliqueInfo.HealthBuildingLevel,
			AttackBuildingLevel:  cliqueInfo.AttackBuildingLevel,
			DefenseBuildingLevel: cliqueInfo.DefenseBuildingLevel,
		}

		//申请过的帮派
		FetchPlayerCliqueApply(state.PlayerId, func(cliqueid int64) {
			out.AppliedCliques = append(out.AppliedCliques, clique_api.CliquePublicInfo_Out_AppliedCliques{CliqueId: cliqueid})
		})

		session.Send(out)
		return
	}

	out := &clique_api.CliquePublicInfo_Out{
		CliqueId: in.CliqueId,
		Exist:    false,
	}
	session.Send(out)
}

//帮派信息
func cliqueInfo(session *net.Session) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	//没有加入帮派，或者此刻已经被帮派踢掉
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		session.Send(&clique_api.CliqueInfo_Out{})
		return
	}

	cliqueInfo := cliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)
	out := &clique_api.CliqueInfo_Out{
		CliqueId:             cliqueInfo.Id,
		Name:                 []byte(cliqueInfo.Name),
		Announce:             []byte(cliqueInfo.Announce),
		TotalDonateCoins:     cliqueInfo.TotalDonateCoins,
		OwnerLoginTime:       cliqueInfo.OwnerLoginTime,
		OwnerPid:             cliqueInfo.OwnerPid,
		MangerPid1:           cliqueInfo.MangerPid1,
		MangerPid2:           cliqueInfo.MangerPid2,
		CenterBuildingCoins:  cliqueInfo.CenterBuildingCoins,
		TempleBuildingCoins:  cliqueInfo.TempleBuildingCoins,
		BankBuildingCoins:    cliqueInfo.BankBuildingCoins,
		HealthBuildingCoins:  cliqueInfo.HealthBuildingCoins,
		AttackBuildingCoins:  cliqueInfo.AttackBuildingCoins,
		DefenseBuildingCoins: cliqueInfo.DefenseBuildingCoins,
		CenterBuildingLevel:  cliqueInfo.CenterBuildingLevel,
		TempleBuildingLevel:  cliqueInfo.TempleBuildingLevel,
		BankBuildingLevel:    cliqueInfo.BankBuildingLevel,
		HealthBuildingLevel:  cliqueInfo.HealthBuildingLevel,
		AttackBuildingLevel:  cliqueInfo.AttackBuildingLevel,
		DefenseBuildingLevel: cliqueInfo.DefenseBuildingLevel,
		Contrib:              playerCliqueInfo.Contrib,
	}

	//缓存里面找成员信息
	cliqueCache := CacheGetCliqueInfo(playerCliqueInfo.CliqueId)
	for _, member := range cliqueCache.Members {
		if member != nil {
			playerInfo := global.GetPlayerInfo(member.Pid)
			out.Members = append(out.Members, clique_api.CliqueInfo_Out_Members{
				Pid:     member.Pid,
				RoleId:  playerInfo.RoleId,
				Level:   playerInfo.RoleLevel,
				Nick:    []byte(playerInfo.PlayerNick),
				Contrib: member.TotalContrib,
			})
		}
	}

	session.Send(out)
	return
}

//帮派申请列表
func listApply(session *net.Session, in *clique_api.ListApply_In) {
	state := module.State(session)
	out := &clique_api.ListApply_Out{}
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	// 没有加入帮派，或者此刻已经被帮派踢掉 ||检查权限
	if playerCliqueInfo == nil || !isCliqueAuthority(state, state.PlayerId, playerCliqueInfo.CliqueId) {
		session.Send(out)
		return
	}
	//清过期的
	DeleteTimeOutApply(playerCliqueInfo.CliqueId)
	getCliqueApplys(playerCliqueInfo.CliqueId, in.Limit, in.Offset, out)

	cliqueInfo := cliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)
	out.AutoAudit = cliqueInfo.AutoAudit > 0
	out.Level = cliqueInfo.AutoAuditLevel

	session.Send(out)
	return
}

//弹劾帮主
func electOwner(session *net.Session) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	fail.When(playerCliqueInfo == nil, "PlayerGlobalCliqueInfo is nil")
	if playerCliqueInfo.CliqueId <= 0 {
		//当前没有加入任何帮派
		session.Send(&clique_api.ElectOwner_Out{
			Success: false,
		})
		return
	}
	cliqueInfo := cliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)
	if cliqueInfo.OwnerPid == state.PlayerId {
		//已经是帮主
		session.Send(&clique_api.ElectOwner_Out{
			Success: false,
		})
		return
	}
	//FIXME 用常量 代替 字面值
	if time.GetNowTime()-cliqueInfo.OwnerLoginTime < clique_dat.CLIQUE_OWNER_MACX_OFFLINE_TIME {
		//帮主离线时间不够
		session.Send(&clique_api.ElectOwner_Out{
			Success: false,
		})
		return
	}
	oldOwnerPid := cliqueInfo.OwnerPid
	cliqueInfo.OwnerPid = state.PlayerId
	switch state.PlayerId {
	case cliqueInfo.MangerPid1:
		cliqueInfo.MangerPid1 = 0
	case cliqueInfo.MangerPid2:
		cliqueInfo.MangerPid2 = 0
	}
	cliqueInfo.OwnerLoginTime = time.GetNowTime()
	state.Database.Update.GlobalClique(cliqueInfo)
	newOwnerInfo := global.GetPlayerInfo(state.PlayerId)
	rpc.RemoteMailSend(oldOwnerPid, &mail_dat.Mailcliquecancelowner{
		Name: string(newOwnerInfo.PlayerNick),
	})
	//TODO 扣钱
	session.Send(&clique_api.ElectOwner_Out{
		Success: true,
	})

	//通知群成员老帮主被新帮主弹劾
	module.CliqueRPC.Broadcast(playerCliqueInfo.Pid, &clique_api.NotifyCliqueElectowner_Out{
		Ownerid: state.PlayerId,
	})

	playerInfo := global.GetPlayerInfo(oldOwnerPid)
	module.CliqueRPC.AddCliqueNews(playerCliqueInfo.CliqueId, channel_dat.MessageCliqueElectOwner{
		NewOwner: channel_dat.ParamPlayer{
			Nick: state.PlayerNick,
			Pid:  state.PlayerId,
		},
		Player2: channel_dat.ParamPlayer{
			Nick: playerInfo.PlayerNick,
			Pid:  oldOwnerPid,
		},
	})

	return
}

//帮派基本信息(帮外玩家申请加入时用)
func cliquePublicInfoSummary(session *net.Session, cliqueId int64) {

	state := module.State(session)
	out := &clique_api.CliquePublicInfoSummary_Out{}
	cliqueInfo := cliqueInfoLookUp(state.Database, cliqueId)

	if cliqueInfo != nil {
		cliqueCache := CacheGetCliqueInfo(cliqueId)
		playerInfo := global.GetPlayerInfo(cliqueInfo.OwnerPid)
		out = &clique_api.CliquePublicInfoSummary_Out{
			CliqueId:  cliqueInfo.Id,
			Name:      []byte(cliqueInfo.Name),
			Level:     cliqueInfo.CenterBuildingLevel,
			MemberNum: cliqueCache.MemberNum,
			OwnerNick: []byte(playerInfo.PlayerNick),
			OwnerPid:  cliqueInfo.OwnerPid,
			Announce:  []byte(cliqueInfo.Announce),
		}

		//申请过的帮派
		if cliqueid, ok := g_PlayerCliqueApplyTable.applyMap[state.PlayerId]; ok {
			for _, values := range cliqueid {
				if values > 0 {
					out.Cliques = append(out.Cliques, clique_api.CliquePublicInfoSummary_Out_Cliques{CliqueId: values})
				}
			}
		}
	}

	session.Send(out)
	return
}

func cliqueAutoAudit(session *net.Session, level int16, enable bool) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		//没有加入帮派
		return
	}
	cliqueInfo := cliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)
	if !isOwer(cliqueInfo, state.PlayerId) && !isManger(cliqueInfo, state.PlayerId) {
		//没有权限
		return
	}
	if enable {
		cliqueInfo.AutoAudit = 1
	} else {
		cliqueInfo.AutoAudit = 0
	}
	cliqueInfo.AutoAuditLevel = level
	state.Database.Update.GlobalClique(cliqueInfo)
}

//帮派招募公告
func cliqueRecruitment(session *net.Session) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	out := &clique_api.CliqueRecruitment_Out{
		Result: clique_api.CLIQUE_RECUITMENT_RESULT_SUCCESS,
	}
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		//没有加入帮派
		out.Result = clique_api.CLIQUE_RECUITMENT_RESULT_NO_PERMISSION
		session.Send(out)
		return
	}
	cliqueInfo := cliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)
	if !isOwer(cliqueInfo, state.PlayerId) && !isManger(cliqueInfo, state.PlayerId) {
		//没有权限
		out.Result = clique_api.CLIQUE_RECUITMENT_RESULT_NO_PERMISSION
		session.Send(out)
		return
	}
	now := time.GetNowTime()
	if cliqueInfo.RecruitTime+clique_dat.CLIQUE_RECRUITMENT_CD > now {
		out.Result = clique_api.CLIQUE_RECUITMENT_RESULT_CD
		out.Timestamp = cliqueInfo.RecruitTime
		session.Send(out)
		return
	}
	cliqueInfo.RecruitTime = now
	state.Database.Update.GlobalClique(cliqueInfo)
	out.Timestamp = now
	session.Send(out)

	rpc.RemoteDecMoney(state.PlayerId, clique_dat.CLIQUE_RECRUITMENT_COST, player_dat.INGOT, 0 /*FIXME tlog */, xdlog.ET_CLIQUE, func() {
		module.Chan.AddWorldTplMessage(state.PlayerId, state.PlayerNick, module.CHANNEL_CLIQUE_MESSAGE, channel_dat.MessageCliqueRecruitment{
			Clique:    channel_dat.ParamClique{cliqueInfo.Name, cliqueInfo.Id},
			DummyLink: channel_dat.ParamClique{i18l.T.Tran("立即申请加入"), cliqueInfo.Id},
		})

	})
}

func quickApply(session *net.Session) {
	state := module.State(session)
	playerCliqueInfo := getPlayerGlobalCliqueInfo(state.Database)
	if playerCliqueInfo != nil && playerCliqueInfo.CliqueId > 0 {
		//已经加入帮派
		return
	}
	var (
		requiLevel = int16(-1)
		cliqueId   = int64(0)
		memberNum  = int16(0)
	)
	playerInfo := global.GetPlayerInfo(state.PlayerId)
	FetchClique(func(cid int64, cliqueCache *CliqueInfo) bool {
		if !cliqueCache.AutoAudit || cliqueCache.AutoAuditMinLevel > playerInfo.RoleLevel {
			//没有开启自动审核或者等级不足
			return true
		}

		baseBuildingDat := clique_dat.GetCenterBuildingLevelInfo(cliqueCache.Level)
		if baseBuildingDat.MaxMember <= cliqueCache.MemberNum {
			//人数已满
			return true
		}

		if cliqueCache.AutoAuditMinLevel > requiLevel ||
			(cliqueCache.AutoAuditMinLevel == requiLevel && memberNum > cliqueCache.MemberNum) {
			requiLevel = cliqueCache.AutoAuditMinLevel
			cliqueId = cid
			memberNum = cliqueCache.MemberNum
		}

		return true
	})
	//加入帮派
	if cliqueId > 0 {
		JoinCliqueChannel(cliqueId, session)
		playerCliqueInfo.CliqueId = cliqueId
		playerCliqueInfo.JoinTime = time.GetNowTime()
		state.Database.Update.PlayerGlobalCliqueInfo(playerCliqueInfo)
		joinSuccess := &clique_api.NotifyJoincliqueSuccess_Out{}
		joinSuccess.Pidlist = append(joinSuccess.Pidlist, clique_api.NotifyJoincliqueSuccess_Out_Pidlist{
			Pid:    state.PlayerId,
			RoleId: playerInfo.RoleId,
			Level:  playerInfo.RoleLevel,
			Nick:   state.PlayerNick,
		})
		module.CliqueRPC.Broadcast(cliqueId, joinSuccess)
		session.Send(&clique_api.QuickApply_Out{
			Success: true,
		})
		cliqueInfo := cliqueInfoLookUp(state.Database, cliqueId)
		joinCliqueCleanup(state.Database, playerCliqueInfo, cliqueInfo)
	} else {
		session.Send(&clique_api.QuickApply_Out{
			Success: false,
		})
	}
}

//帮派贡献增加
func addCliqueContrib(db *mdb.Database, contrib int64) bool {
	record := db.Lookup.PlayerGlobalCliqueInfo(db.PlayerId())
	if record == nil || record.CliqueId <= 0 {
		return false
	}
	record.Contrib += contrib
	record.TotalContrib += contrib
	db.Update.PlayerGlobalCliqueInfo(record)
	if session, online := module.Player.GetPlayerOnline(db.PlayerId()); online {
		session.Send(&clique_api.NotifyContribChange_Out{
			Value: record.Contrib,
		})
	}
	return true
}

//帮派贡献扣除
func delCliqueContrib(db *mdb.Database, contrib int64) bool {
	record := db.Lookup.PlayerGlobalCliqueInfo(db.PlayerId())
	if record != nil && record.Contrib >= contrib {
		record.Contrib -= contrib
		//record.Updatetime = time.GetNowTime()
		db.Update.PlayerGlobalCliqueInfo(record)
		if session, online := module.Player.GetPlayerOnline(db.PlayerId()); online {
			session.Send(&clique_api.NotifyContribChange_Out{
				Value: record.Contrib,
			})
		}
		return true
	}
	return false
}

//帮派贡献清零
func CleanCliqueContrib(session *net.Session, cliqueId int64) {
	state := module.State(session)
	cliqueInfo := cliqueInfoLookUp(state.Database, cliqueId)
	fail.When(cliqueInfo == nil, "clique not found")
	cliqueInfo.Contrib = 0
	state.Database.Update.GlobalClique(cliqueInfo)
}

func otherClique(session *net.Session, page int) {
	state := module.State(session)
	out := &clique_api.OtherClique_Out{}
	// allClique := sortCliqueList{}
	// state.Database.Select.GlobalClique(func(row *mdb.GlobalCliqueRow) {
	// 	allClique = append(allClique, row.Id())
	// })
	// sort.Sort(allClique)
	// out.TotalNum = int16(len(allClique))
	// page -= 1
	// start := page * clique_dat.CLIQUE_LIST_PAGE_SIZE

	// //TODO cache 下面的结果
	// for index, clique_id := range allClique[start:] {
	// 	cliqueInfo := CacheGetCliqueInfo(clique_id)
	// 	cliqueDbInfo := state.Database.Lookup.GlobalClique(clique_id)
	// 	out.CliqueList = append(out.CliqueList, clique_api.OtherClique_Out_CliqueList{
	// 		Rank:            int16(start + index + 1),
	// 		CliqueId:        clique_id,
	// 		CliqueLevel:     cliqueInfo.Level,
	// 		Name:            []byte(cliqueInfo.Name),
	// 		CliqueMemberNum: cliqueInfo.MemberNum,
	// 		OwnerPid:        cliqueDbInfo.OwnerPid,
	// 		OwnerName:       global.GetPlayerInfo(cliqueDbInfo.OwnerPid).PlayerNick,
	// 	})
	// 	if len(out.CliqueList) >= clique_dat.CLIQUE_LIST_PAGE_SIZE {
	// 		break
	// 	}
	// }
	// session.Send(out)

	start, pageNum, cliquePageList := CacheGetPageCliqueInfo(page)
	out.TotalNum = int16(pageNum)
	for index, clique_id := range cliquePageList {
		cliqueInfo := CacheGetCliqueInfo(clique_id)
		cliqueDbInfo := state.Database.Lookup.GlobalClique(clique_id)
		out.CliqueList = append(out.CliqueList, clique_api.OtherClique_Out_CliqueList{
			Rank:            int16(start + index + 1),
			CliqueId:        clique_id,
			CliqueLevel:     cliqueInfo.Level,
			Name:            []byte(cliqueInfo.Name),
			CliqueMemberNum: cliqueInfo.MemberNum,
			OwnerPid:        cliqueDbInfo.OwnerPid,
			OwnerName:       global.GetPlayerInfo(cliqueDbInfo.OwnerPid).PlayerNick,
		})
	}
	session.Send(out)

}
