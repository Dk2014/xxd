package battle

import (
	"game_server/api/protocol/battle_api"
	"game_server/battle"
	"game_server/module"
)

func checkGhostUsage(ghostId int16, ghostUsageForOnePlayer map[int16]int8) bool {
	if ghostUsageForOnePlayer == nil {
		return false
	}
	_, used := ghostUsageForOnePlayer[ghostId]
	return used
}

func GetStartBattleResponse(b *battle.BattleState, levelState *module.MissionLevelState, skillUsage map[int64]map[int]int, ghostUsage map[int64]map[int16]int8) *battle_api.StartBattle_Out {
	if ghostUsage == nil {
		ghostUsage = map[int64]map[int16]int8{}
	}
	rspAttackerGroups := make([]battle_api.StartBattle_Out_AttackerGroups, len(b.Attackers.Groups))

	var buffs []battle.Buff
	SelfBuffs := make([]battle_api.StartBattle_Out_AttackerGroups_SelfBuffs, 0)
	BuddyBuffs := make([]battle_api.StartBattle_Out_AttackerGroups_BuddyBuffs, 0)

	for groupIndex, fighters := range b.Attackers.Groups {
		rspAttackers := make([]battle_api.StartBattle_Out_AttackerGroups_Attackers, 0, len(fighters))

		for _, v := range fighters {
			if v == nil {
				continue
			}

			skills := make([]battle_api.StartBattle_Out_AttackerGroups_Attackers_Skills, len(v.Skills))

			for i, skill := range v.Skills {
				if skill == nil {
					skills[i] = battle_api.StartBattle_Out_AttackerGroups_Attackers_Skills{
						Skill: battle_api.SkillInfo{
							SkillId: -2,
						},
					}
					continue
				}
				var restReleaseNum int16 = -1
				if skillUsage != nil {
					if _, isCurrPlayer := skillUsage[v.PlayerId]; isCurrPlayer {
						if num, check := skillUsage[v.PlayerId][skill.SkillId]; check {
							restReleaseNum = int16(num)
						}
					}
				}
				skills[i] = battle_api.StartBattle_Out_AttackerGroups_Attackers_Skills{
					Skill: battle_api.SkillInfo{
						SkillId:  int16(skill.SkillId),
						IncPower: int8(skill.IncPower),
						DecPower: int8(skill.DecPower),
					},
					RestReleaseNum: restReleaseNum,
				}
			}

			ghosts := []battle_api.BattleRole_Ghosts{}
			for _, roleGhost := range v.Ghosts {
				ghostUsed := checkGhostUsage(roleGhost.GhostId, ghostUsage[v.PlayerId])
				var relatedGhost int16
				if roleGhost.RelationGhost != nil {
					relatedGhost = roleGhost.RelationGhost.GhostId
				}
				ghosts = append(ghosts, battle_api.BattleRole_Ghosts{
					GhostId:      roleGhost.GhostId,
					GhostStar:    roleGhost.GhostStar,
					GhostLevel:   roleGhost.GhostLevel,
					GhostSkillId: int32(roleGhost.GhostSkillId),
					Used:         ghostUsed,
					RelatedGhost: relatedGhost,
				})
				if roleGhost.RelationGhost != nil {
					ghosts = append(ghosts, battle_api.BattleRole_Ghosts{
						GhostId:      roleGhost.RelationGhost.GhostId,
						GhostStar:    roleGhost.RelationGhost.GhostStar,
						GhostLevel:   roleGhost.RelationGhost.GhostLevel,
						GhostSkillId: int32(roleGhost.RelationGhost.GhostSkillId),
						Used:         ghostUsed,
					})
				}
			}

			rspAttackers = append(rspAttackers, battle_api.StartBattle_Out_AttackerGroups_Attackers{
				Role: battle_api.BattleRole{
					Kind:                battle_api.FighterKind(v.Kind),
					PlayerId:            v.PlayerId,
					RoleId:              int32(v.RoleId),
					RoleLevel:           int16(v.Level),
					Position:            int32(v.Position),
					FashionId:           v.FashionId,
					FriendshipLevel:     v.FriendshipLevel,
					Health:              int32(v.Health),
					MaxHealth:           int32(v.MaxHealth),
					Power:               int16(v.Power),
					MaxPower:            int16(v.MaxPower),
					SunderValue:         int16(v.GetSunderValue()),
					SunderMaxValue:      int16(v.SunderMaxValue),
					SunderMinHurtRate:   int16(v.SunderMinHurtRate),
					SunderEndHurtRate:   int16(v.SunderEndHurtRate),
					SunderEndDefendRate: int16(v.SunderEndDefendRate),
					Speed:               int32(v.Speed),

					GhostPower:  int32(v.GetGhostPower()),
					Ghosts:      ghosts,
					CanUseGhost: !v.UsedGhostSkill,
					//GhostSkillId:      int16(v.GhostSkillId),
					//GhostSkillRate:    int32(v.GhostSkillRate),
					GhostShieldValue: int32(v.GhostShieldValue),
					//MainGhostId:       int32(v.MainGhostId),
					CouldUseSwordSoul: v.SwordSoulValue > 0,
				},
				Skills: skills,
			})

			if levelState != nil {
				buffs = module.GetBuffInMissionLevelWithRoleId(levelState, v.RoleId)
				if v.Kind == battle.FK_PLAYER {
					for _, buff := range buffs {
						SelfBuffs = append(SelfBuffs, battle_api.StartBattle_Out_AttackerGroups_SelfBuffs{
							Buffer: battle_api.BufferInfo{
								Mode:        battle_api.BuffMode(buff.Mode),
								Keep:        int8(buff.ShowKeep),
								Value:       int32(buff.Value),
								SkillId:     int16(buff.Skill),
								MaxOverride: int8(buff.MaxOverride),
								OverrideNum: int8(buff.OverrideNum),
								Uncleanable: buff.Uncleanable,
							},
						})
					}
				} else {
					for _, buff := range buffs {
						BuddyBuffs = append(BuddyBuffs, battle_api.StartBattle_Out_AttackerGroups_BuddyBuffs{
							Pos: int8(v.Position),
							Buffer: battle_api.BufferInfo{
								Mode:        battle_api.BuffMode(buff.Mode),
								Keep:        int8(buff.ShowKeep),
								Value:       int32(buff.Value),
								SkillId:     int16(buff.Skill),
								MaxOverride: int8(buff.MaxOverride),
								OverrideNum: int8(buff.OverrideNum),
								Uncleanable: buff.Uncleanable,
							},
						})
					}
				}
			}
		}

		rspAttackerGroups[groupIndex] = battle_api.StartBattle_Out_AttackerGroups{
			Attackers:  rspAttackers,
			SelfBuffs:  SelfBuffs,
			BuddyBuffs: BuddyBuffs,
		}
	}

	rspDefenderGroups := make([]battle_api.StartBattle_Out_DefenderGroups, len(b.Defenders.Groups))

	for groupIndex, fighters := range b.Defenders.Groups {
		rspDefenders := make([]battle_api.StartBattle_Out_DefenderGroups_Defenders, 0, len(fighters))

		for _, v := range fighters {
			if v == nil {
				continue
			}

			skills := make([]battle_api.StartBattle_Out_DefenderGroups_Defenders_Skills, len(v.Skills))

			for i, skill := range v.Skills {
				if skill == nil {
					skills[i] = battle_api.StartBattle_Out_DefenderGroups_Defenders_Skills{
						Skill: battle_api.SkillInfo{
							SkillId: -2,
						},
					}
					continue
				}

				skills[i] = battle_api.StartBattle_Out_DefenderGroups_Defenders_Skills{
					Skill: battle_api.SkillInfo{
						SkillId:  int16(skill.SkillId),
						IncPower: int8(skill.IncPower),
						DecPower: int8(skill.DecPower),
					},
					SkillId2: int16(v.SkillInfos[i].SkillId2),
				}
			}

			ghosts := []battle_api.BattleRole_Ghosts{}
			for _, roleGhost := range v.Ghosts {
				var relatedGhost int16
				if roleGhost.RelationGhost != nil {
					relatedGhost = roleGhost.RelationGhost.GhostId
				}
				ghosts = append(ghosts, battle_api.BattleRole_Ghosts{
					GhostId:      roleGhost.GhostId,
					GhostStar:    roleGhost.GhostStar,
					GhostLevel:   roleGhost.GhostLevel,
					GhostSkillId: int32(roleGhost.GhostSkillId),
					RelatedGhost: relatedGhost,
				})
				if roleGhost.RelationGhost != nil {
					ghosts = append(ghosts, battle_api.BattleRole_Ghosts{
						GhostId:      roleGhost.RelationGhost.GhostId,
						GhostStar:    roleGhost.RelationGhost.GhostStar,
						GhostLevel:   roleGhost.RelationGhost.GhostLevel,
						GhostSkillId: int32(roleGhost.RelationGhost.GhostSkillId),
					})
				}
			}

			rspDefenders = append(rspDefenders, battle_api.StartBattle_Out_DefenderGroups_Defenders{
				Role: battle_api.BattleRole{
					Kind:                battle_api.FighterKind(v.Kind),
					PlayerId:            v.PlayerId,
					RoleId:              int32(v.RoleId),
					RoleLevel:           int16(v.Level),
					Position:            int32(v.Position),
					FashionId:           v.FashionId,
					FriendshipLevel:     v.FriendshipLevel,
					Health:              int32(v.Health),
					MaxHealth:           int32(v.MaxHealth),
					Power:               int16(v.Power),
					MaxPower:            int16(v.MaxPower),
					SunderValue:         int16(v.GetSunderValue()),
					SunderMaxValue:      int16(v.SunderMaxValue),
					SunderMinHurtRate:   int16(v.SunderMinHurtRate),
					SunderEndHurtRate:   int16(v.SunderEndHurtRate),
					SunderEndDefendRate: int16(v.SunderEndDefendRate),

					Speed: int32(v.Speed),

					GhostPower:  int32(v.GetGhostPower()),
					Ghosts:      ghosts,
					CanUseGhost: true,
					//GhostSkillId:      int16(v.GhostSkillId),
					//GhostSkillRate:    int32(v.GhostSkillRate),
					GhostShieldValue: int32(v.GhostShieldValue),
					//MainGhostId:       int32(v.MainGhostId),
				},
				Skills: skills,
			})
		}

		rspDefenderGroups[groupIndex] = battle_api.StartBattle_Out_DefenderGroups{
			Defenders: rspDefenders,
		}
	}

	firstFighter := b.GetFirstFighter()

	rsp := &battle_api.StartBattle_Out{
		IsMainRoleFirst: firstFighter.Kind == battle.FK_PLAYER,
		IsAttackerFirst: firstFighter.Ftype == battle.FT_ATK,
		AttackerGroups:  rspAttackerGroups,
		DefenderGroups:  rspDefenderGroups,
	}

	attckerPlayerIds := b.GetAttackerPlayerIds()
	rsp.AttackerPlayerIds = make([]battle_api.StartBattle_Out_AttackerPlayerIds, len(attckerPlayerIds))
	for i, pid := range attckerPlayerIds {
		rsp.AttackerPlayerIds[i] = battle_api.StartBattle_Out_AttackerPlayerIds{
			PlayerId: pid,
		}
	}

	for i, totem := range b.Attackers.TotemInfo {
		if totem != nil {
			rsp.AttackerTotems = append(rsp.AttackerTotems, battle_api.StartBattle_Out_AttackerTotems{
				Round:   int16(i),
				TotemId: int16(totem.Id),
			})
		}

	}
	for i, totem := range b.Defenders.TotemInfo {
		if totem != nil {
			rsp.DefenderTotems = append(rsp.DefenderTotems, battle_api.StartBattle_Out_DefenderTotems{
				Round:   int16(i),
				TotemId: int16(totem.Id),
			})
		}
	}

	// 攻方职业信息
	rsp.AllAttackers = make([]battle_api.StartBattle_Out_AllAttackers, len(b.Attackers.Players))

	for i, player := range b.Attackers.Players {
		rsp.AllAttackers[i] = battle_api.StartBattle_Out_AllAttackers{
			PlayerId: player.PlayerId,
		}
	}

	// 守方职业信息
	rsp.AllDefenders = make([]battle_api.StartBattle_Out_AllDefenders, len(b.Defenders.Players))

	for i, player := range b.Defenders.Players {
		rsp.AllDefenders[i] = battle_api.StartBattle_Out_AllDefenders{
			PlayerId: player.PlayerId,
		}
	}

	//防守放分组上阵组数
	rsp.TotalGroup = 1

	return rsp
}
