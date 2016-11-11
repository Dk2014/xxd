package battle

import (
	"game_server/api/protocol/battle_api"
	"game_server/battle"
)

func GetNewFighterGroupResponse(playerId int64, ftype int8, fighters []*battle.Fighter, skillUsage map[int64]map[int]int) *battle_api.NewFighterGroup_Out {
	rspNewFighterGroup := &battle_api.NewFighterGroup_Out{
		PlayerId: playerId,
		Ftype:    ftype,
		//GhostSkillIndex: 0, //暂时无需求
	}

	var rspFighters []battle_api.NewFighterGroup_Out_Fighters
	var restReleaseNum int16
	for _, fighter := range fighters {
		if fighter == nil {
			continue
		}
		skills := make([]battle_api.NewFighterGroup_Out_Fighters_Skills, len(fighter.Skills))
		for i, skill := range fighter.Skills {
			if skill == nil {
				skills[i] = battle_api.NewFighterGroup_Out_Fighters_Skills{
					Skill: battle_api.SkillInfo{
						SkillId: -2,
					},
				}
				continue
			}
			restReleaseNum = -1
			if skillUsage != nil {
				if _, isCurrPlayer := skillUsage[fighter.PlayerId]; isCurrPlayer {
					if num, check := skillUsage[fighter.PlayerId][skill.SkillId]; check {
						restReleaseNum = int16(num)
					}
				}
			}
			skills[i] = battle_api.NewFighterGroup_Out_Fighters_Skills{
				Skill: battle_api.SkillInfo{
					SkillId:  int16(skill.SkillId),
					IncPower: int8(skill.IncPower),
					DecPower: int8(skill.DecPower),
				},
				RestReleaseNum: restReleaseNum,
			}
		}

		ghosts := []battle_api.BattleRole_Ghosts{}
		for _, roleGhost := range fighter.Ghosts {
			if roleGhost != nil {
				ghosts = append(ghosts, battle_api.BattleRole_Ghosts{
					GhostId:      roleGhost.GhostId,
					GhostStar:    roleGhost.GhostStar,
					GhostSkillId: int32(roleGhost.GhostSkillId),
				})
			}

		}
		rspFighters = append(rspFighters, battle_api.NewFighterGroup_Out_Fighters{
			Skills: skills,
			Role: battle_api.BattleRole{
				Kind:                battle_api.FighterKind(fighter.Kind),
				PlayerId:            fighter.PlayerId,
				RoleId:              int32(fighter.RoleId),
				RoleLevel:           int16(fighter.Level),
				Position:            int32(fighter.Position),
				FashionId:           fighter.FashionId,
				Health:              int32(fighter.Health),
				MaxHealth:           int32(fighter.MaxHealth),
				Power:               int16(fighter.Power),
				MaxPower:            int16(fighter.MaxPower),
				SunderValue:         int16(fighter.GetSunderValue()),
				SunderMaxValue:      int16(fighter.SunderMaxValue),
				SunderMinHurtRate:   int16(fighter.SunderMinHurtRate),
				SunderEndHurtRate:   int16(fighter.SunderEndHurtRate),
				SunderEndDefendRate: int16(fighter.SunderEndDefendRate),
				Speed:               int32(fighter.Speed),
				GhostPower:          int32(fighter.GetGhostPower()),
				Ghosts:              ghosts,
				GhostShieldValue:    int32(fighter.GhostShieldValue),
			},
		})

	}
	rspNewFighterGroup.Fighters = rspFighters
	return rspNewFighterGroup
}
