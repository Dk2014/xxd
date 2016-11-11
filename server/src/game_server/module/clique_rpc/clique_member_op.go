package clique_rpc

import (
	//"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/clique_api"
	"game_server/dat/channel_dat"
	"game_server/dat/clique_dat"
	"game_server/dat/mail_dat"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
)

//帮派管理处理申请
func processJoinApply(session *net.Session, in *clique_api.ProcessJoinApply_In) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	out := &clique_api.ProcessJoinApply_Out{}

	//没有加入帮派
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		session.Send(out)
		return
	}

	cliqueId := playerCliqueInfo.CliqueId
	cliqueInfo := cliqueInfoLookUp(state.Database, cliqueId)
	cliqueCache := CacheGetCliqueInfo(cliqueId)

	//帮主或者副帮主才有权限更新
	if !isOwer(cliqueInfo, state.PlayerId) && !isManger(cliqueInfo, state.PlayerId) {
		session.Send(out)
		return
	}

	cliqueLevelDat := clique_dat.GetCenterBuildingLevelInfo(cliqueInfo.CenterBuildingLevel)

	//获取申请列表
	var applylist = map[int64]int64{}
	cliqueInfoMap := CacheGetCliqueInfo(cliqueId)
	for _, v := range cliqueInfoMap.JoinApplies {
		applylist[v.Pid] = v.Pid
	}

	var newMemberCliqueNews []channel_dat.MessageCliqueNewMember

	var addCnt int16 //防止帮派成员超员
	if in.Agree {
		joinSuccess := &clique_api.NotifyJoincliqueSuccess_Out{}
		for _, pidStruct := range in.Pidlist {
			var exist bool
			state.Database.AgentExecute(pidStruct.Pid, func(agentDB *mdb.Database) {
				playerApplyCliqueInfo := agentDB.Lookup.PlayerGlobalCliqueInfo(pidStruct.Pid)
				//已加入帮派
				if playerApplyCliqueInfo != nil && playerApplyCliqueInfo.CliqueId > 0 {
					exist = true
				}
			})

			if exist { //已加帮派
				out.Applylist = append(out.Applylist, clique_api.ProcessJoinApply_Out_Applylist{
					Pid:    pidStruct.Pid,
					Result: clique_api.PROCESS_JOIN_APPLY_RESULT_EXPIRED,
				})

				continue
			}

			//检测是否还在申请列表中
			if _, ok := applylist[pidStruct.Pid]; !ok {
				out.Applylist = append(out.Applylist, clique_api.ProcessJoinApply_Out_Applylist{
					Pid:    pidStruct.Pid,
					Result: clique_api.PROCESS_JOIN_APPLY_RESULT_CANCEL_APPLY,
				})
				continue
			}

			//不管能否加入，都要删除缓存里的记录
			g_CliqueTable.DeleteApply(pidStruct.Pid, cliqueId)
			DeleteApply(pidStruct.Pid, cliqueId)

			if cliqueCache.MemberNum+addCnt >= cliqueLevelDat.MaxMember {
				out.Applylist = append(out.Applylist, clique_api.ProcessJoinApply_Out_Applylist{
					Pid:    pidStruct.Pid,
					Result: clique_api.PROCESS_JOIN_APPLY_RESULT_NO_ROOM,
				})
				break
			} else {
				//2,更新个人帮派信息
				state.Database.AgentExecute(pidStruct.Pid, func(agentDB *mdb.Database) {
					playerCliqueInfo := agentDB.Lookup.PlayerGlobalCliqueInfo(pidStruct.Pid)
					if playerCliqueInfo == nil {
						playerCliqueInfo = &mdb.PlayerGlobalCliqueInfo{
							Pid:      state.PlayerId,
							CliqueId: cliqueId,
							JoinTime: time.GetNowTime(),
						}
						//resetPlayerCliqueState(playerCliqueInfo)
						agentDB.Insert.PlayerGlobalCliqueInfo(playerCliqueInfo)
					} else {
						playerCliqueInfo.JoinTime = time.GetNowTime()
						playerCliqueInfo.CliqueId = cliqueId
						agentDB.Update.PlayerGlobalCliqueInfo(playerCliqueInfo)
					}
					joinCliqueCleanup(agentDB, playerCliqueInfo, cliqueInfo)
				})
				newMemberInfo := global.GetPlayerInfo(pidStruct.Pid)
				newMemberCliqueNews = append(newMemberCliqueNews, channel_dat.MessageCliqueNewMember{
					NewMember: channel_dat.ParamPlayer{
						Nick: newMemberInfo.PlayerNick,
						Pid:  pidStruct.Pid,
					},
				})
				out.Applylist = append(out.Applylist, clique_api.ProcessJoinApply_Out_Applylist{
					Pid:    pidStruct.Pid,
					Result: clique_api.PROCESS_JOIN_APPLY_RESULT_SUCCESS,
				})
				//批准后，加入帮派channel,以便广播
				if memberSession, online := module.Player.GetPlayerOnline(pidStruct.Pid); online {
					JoinCliqueChannel(cliqueId, memberSession)
				}
				playerInfo := global.GetPlayerInfo(pidStruct.Pid)
				joinSuccess.Pidlist = append(joinSuccess.Pidlist, clique_api.NotifyJoincliqueSuccess_Out_Pidlist{
					Pid:    pidStruct.Pid,
					RoleId: playerInfo.RoleId,
					Level:  playerInfo.RoleLevel,
					Nick:   playerInfo.PlayerNick,
				})
				//加入后，发邮件通知
				rpc.RemoteMailSend(pidStruct.Pid, &mail_dat.Mailcliquejoin{Name: cliqueInfo.Name})
				addCnt++
			}
		}
		for i, _ := range newMemberCliqueNews {
			module.CliqueRPC.AddCliqueNews(cliqueId, newMemberCliqueNews[i])
		}

		module.CliqueRPC.Broadcast(cliqueId, joinSuccess)

	} else {

		// 删除申请的消息
		for _, pidStruct := range in.Pidlist {
			g_CliqueTable.DeleteApply(pidStruct.Pid, cliqueId)
			DeleteApply(pidStruct.Pid, cliqueId)
			out.Applylist = append(out.Applylist, clique_api.ProcessJoinApply_Out_Applylist{
				Pid:    pidStruct.Pid,
				Result: clique_api.PROCESS_JOIN_APPLY_RESULT_SUCCESS,
			})
		}
	}

	session.Send(out)

}
