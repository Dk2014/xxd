package battle

import (
	"game_server/api/protocol/battle_api"
	"game_server/battle"
	"game_server/dat/skill_dat"
)

func GetNextRoundResponse(results []*battle.FightResult, status, nowRound int, battleState *battle.BattleState, autoPids []int64, skillReleaseNum map[int]int) *battle_api.NextRound_Out {
	attackers := battleState.Attackers.Players
	defenders := battleState.Defenders.Players

	rspResults := make([]battle_api.NextRound_Out_Results, len(results))

	for i, result := range results {
		rspResults[i] = battle_api.NextRound_Out_Results{
			Ftype:       battle_api.FighterType(result.Type),
			Event:       battle_api.RoundEvent(result.FighterEv),
			Position:    int8(result.FighterPos),
			Power:       int16(result.Power),
			Health:      int32(result.Health),
			SunderValue: int16(result.SunderValue),

			UseGhostSkill: result.UseGhostSkill,
			//GhostSkillRate: int32(result.GhostSkillRate),
			GhostShieldOn: result.GhostShieldOn,

			ShieldGhostId: result.ShieldGhostId,
			TotemId:       result.TotemId,
			GhostId:       result.GhostId,

			GhostPower: int32(result.GhostPower),
			AddPower:   int32(result.AddPower),

			Attacks: make([]battle_api.NextRound_Out_Results_Attacks, 0, len(result.Attacks)),
			Item:    []battle_api.NextRound_Out_Results_Item{},
		}

		rspResult := &rspResults[i]

		if result.ItemResult != nil {
			rspResult.Item = append(rspResult.Item, battle_api.NextRound_Out_Results_Item{
				ItemId:  result.ItemResult.ItemId,
				Targets: make([]battle_api.NextRound_Out_Results_Item_Targets, len(result.ItemResult.Result)),
			})

			for resultI, rs := range result.ItemResult.Result {
				buffs := make([]battle_api.NextRound_Out_Results_Item_Targets_Buffs, len(rs.Buffs))
				for buffI, b := range rs.Buffs {
					buffs[buffI] = battle_api.NextRound_Out_Results_Item_Targets_Buffs{
						Buffer: battle_api.BufferInfo{
							Mode:        battle_api.BuffMode(b.Mode),
							Keep:        int8(b.Keep),
							Value:       int32(b.Value),
							SkillId:     int16(b.Skill),
							MaxOverride: int8(b.MaxOverride),
							OverrideNum: int8(b.OverrideNum),
						},
					}
				}

				rspResult.Item[0].Targets[resultI] = battle_api.NextRound_Out_Results_Item_Targets{
					Ftype:    battle_api.FighterType(rs.Type),
					Health:   int32(rs.Health),
					Power:    int16(rs.Power),
					Hurt:     int32(rs.Hurt),
					Position: int8(rs.TargetPos),
					Buffs:    buffs,
				}
			}
		}

		for attacksI := 0; attacksI < len(result.Attacks); attacksI++ {
			var r = &result.Attacks[attacksI]

			if r.SkillId == skill_dat.SKILL_IS_NULL {
				continue
			}

			//更新技能
			/*
				var restReaseNum int16 = -1
				if skillReleaseNum == nil {
					restReaseNum = 0
				} else if num, check := skillReleaseNum[r.SkillId]; check {
					skillReleaseNum[r.SkillId]--
					restReaseNum = int16(num - 1)
				}

			*/

			rspResult.Attacks = append(rspResult.Attacks, battle_api.NextRound_Out_Results_Attacks{
				SkillId:        int32(r.SkillId),
				RestReleaseNum: int16(r.RestReleaseNum),
				Targets:        make([]battle_api.NextRound_Out_Results_Attacks_Targets, len(r.Targets)),
				SelfBuffs:      make([]battle_api.NextRound_Out_Results_Attacks_SelfBuffs, len(r.SelfBuffs)),
				BuddyBuffs:     make([]battle_api.NextRound_Out_Results_Attacks_BuddyBuffs, len(r.BuddyBuffs)),
			})

			x := &rspResult.Attacks[len(rspResult.Attacks)-1]

			for j, h := range r.Targets {
				x.Targets[j] = battle_api.NextRound_Out_Results_Attacks_Targets{
					Ftype:            battle_api.FighterType(h.Type),
					Hurt:             int32(h.Hurt),
					Event:            battle_api.RoundEvent(h.TargetEv),
					Position:         int8(h.TargetPos),
					TakeSunder:       int16(h.TakeSunder),
					TakeGhostShield:  int32(h.TakeGhostShield),
					DirectReductHurt: int32(h.DirectReductHurt),
					GhostShieldOn:    h.GhostShieldOn,
					GhostPower:       int32(h.GhostPower),
					ShieldGhostId:    h.ShieldGhostId,
					Buffs:            make([]battle_api.NextRound_Out_Results_Attacks_Targets_Buffs, 0, len(h.Buffs)),
				}

				for _, buff := range h.Buffs {
					if buff == nil {
						continue
					}

					x.Targets[j].Buffs = append(x.Targets[j].Buffs, battle_api.NextRound_Out_Results_Attacks_Targets_Buffs{
						Buffer: battle_api.BufferInfo{
							Mode:        battle_api.BuffMode(buff.Mode),
							Keep:        int8(buff.ShowKeep),
							Value:       int32(buff.Value),
							SkillId:     int16(buff.Skill),
							MaxOverride: int8(buff.MaxOverride),
							OverrideNum: int8(buff.OverrideNum),
						},
					})
				}
			}

			for j, buff := range r.SelfBuffs {
				x.SelfBuffs[j] = battle_api.NextRound_Out_Results_Attacks_SelfBuffs{
					Buffer: battle_api.BufferInfo{
						Mode:        battle_api.BuffMode(buff.Mode),
						Keep:        int8(buff.ShowKeep),
						Value:       int32(buff.Value),
						SkillId:     int16(buff.Skill),
						MaxOverride: int8(buff.MaxOverride),
						OverrideNum: int8(buff.OverrideNum),
					},
				}
			}

			for j, buff := range r.BuddyBuffs {
				x.BuddyBuffs[j] = battle_api.NextRound_Out_Results_Attacks_BuddyBuffs{
					Pos: int8(buff.Owner.Position),
					Buffer: battle_api.BufferInfo{
						Mode:        battle_api.BuffMode(buff.Mode),
						Keep:        int8(buff.ShowKeep),
						Value:       int32(buff.Value),
						SkillId:     int16(buff.Skill),
						MaxOverride: int8(buff.MaxOverride),
						OverrideNum: int8(buff.OverrideNum),
					},
				}
			}
		}
	}

	rspAllAttackers := make([]battle_api.NextRound_Out_AllAttackers, len(attackers))
	for i, attacker := range attackers {
		rspAllAttackers[i] = battle_api.NextRound_Out_AllAttackers{
			PlayerId: attacker.PlayerId,
		}
	}

	rspAllDefenders := make([]battle_api.NextRound_Out_AllDefenders, len(defenders))
	for i, defender := range defenders {
		rspAllDefenders[i] = battle_api.NextRound_Out_AllDefenders{
			PlayerId: defender.PlayerId,
		}
	}

	autos := make([]battle_api.NextRound_Out_Autos, len(autoPids))
	for i, pid := range autoPids {
		autos[i] = battle_api.NextRound_Out_Autos{
			PlayerId: pid,
		}
	}

	return &battle_api.NextRound_Out{
		Status:       battle_api.RoundStatus(status),
		NowRound:     int16(nowRound),
		AllAttackers: rspAllAttackers,
		AllDefenders: rspAllDefenders,
		Results:      rspResults,
		Autos:        autos,
	}
}
