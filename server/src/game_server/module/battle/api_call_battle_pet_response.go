package battle

import (
	"game_server/api/protocol/battle_api"
	"game_server/battle"
)

func GetCallBattlePetResponse(playerFighter, petFighter *battle.Fighter) *battle_api.CallBattlePet_Out {
	rsp := &battle_api.CallBattlePet_Out{}
	rsp.Success = true
	rsp.PlayerPower = int16(playerFighter.Power)
	rsp.Skills = make([]battle_api.CallBattlePet_Out_Skills, len(petFighter.Skills))

	ghosts := []battle_api.BattleRole_Ghosts{}
	for _, roleGhost := range petFighter.Ghosts {
		ghosts = append(ghosts, battle_api.BattleRole_Ghosts{
			GhostId:      roleGhost.GhostId,
			GhostStar:    roleGhost.GhostStar,
			GhostSkillId: int32(roleGhost.GhostSkillId),
		})
	}

	rsp.Role = battle_api.BattleRole{
		Kind:                battle_api.FighterKind(petFighter.Kind),
		PlayerId:            petFighter.PlayerId,
		RoleId:              int32(petFighter.RoleId),
		RoleLevel:           int16(petFighter.Level),
		Position:            int32(petFighter.Position),
		FashionId:           petFighter.FashionId,
		Health:              int32(petFighter.Health),
		MaxHealth:           int32(petFighter.MaxHealth),
		Power:               int16(petFighter.Power),
		MaxPower:            int16(petFighter.MaxPower),
		SunderValue:         int16(petFighter.GetSunderValue()),
		SunderMaxValue:      int16(petFighter.SunderMaxValue),
		SunderMinHurtRate:   int16(petFighter.SunderMinHurtRate),
		SunderEndHurtRate:   int16(petFighter.SunderEndHurtRate),
		SunderEndDefendRate: int16(petFighter.SunderEndDefendRate),

		Speed: int32(petFighter.Speed),

		//GhostPower:        int32(petFighter.GhostPower),
		Ghosts:            ghosts,
		GhostShieldValue:  int32(petFighter.GhostShieldValue),
		CouldUseSwordSoul: petFighter.SwordSoulValue > 0,
	}

	for i, skill := range petFighter.Skills {
		if skill == nil {
			rsp.Skills[i] = battle_api.CallBattlePet_Out_Skills{
				Skill: battle_api.SkillInfo{
					SkillId: -2,
				},
			}
			continue
		}
		rsp.Skills[i] = battle_api.CallBattlePet_Out_Skills{
			Skill: battle_api.SkillInfo{
				SkillId:  int16(skill.SkillId),
				IncPower: int8(skill.IncPower),
				DecPower: int8(skill.DecPower),
			},
		}
	}
	return rsp
}
