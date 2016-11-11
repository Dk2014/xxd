package module

import (
	"core/log"
	"game_server/battle"
	"game_server/dat/awaken_dat"
	"game_server/dat/fashion_dat"
	"game_server/dat/fight_num_dat"
	"game_server/dat/ghost_dat"
	"game_server/dat/item_dat"
	"game_server/dat/player_dat"
	"game_server/dat/role_dat"
	"game_server/dat/skill_dat"
	"game_server/dat/sword_soul_dat"
	"game_server/dat/team_dat"
	"game_server/dat/totem_dat"
	"game_server/mdb"
	"math"
)

const (
	ALL_FIGHTER_POS_NUM = 15
	NONE_ROLE_ID        = -1 // 空角色ID
	POS_TYPE_MAIN_FORM  = 0  // 主阵位置
	POS_TYPE_SUPPORT    = 1  // 支援伙伴位置

	FIGHT_FOR_EMPTY         = -1 // 不做数据加成
	FIGHT_FOR_ALL           = 0  // 战斗属性加成标识
	FIGHT_FOR_ROLE_LEVEL    = 1  // 角色等级属性标识
	FIGHT_FOR_SKILL         = 2  // 绝招加成标识
	FIGHT_FOR_EQUIP         = 3  // 装备加成标识
	FIGHT_FOR_GHOST         = 4  // 魂侍加成标识
	FIGHT_FOR_REALM         = 5  // 境界加成标识
	FIGHT_FOR_SWORD_SOUL    = 6  // 剑心加成标识
	FIGHT_FOR_GHOST_SKILL   = 7  // 魂侍技能
	FIGHT_FOR_FASHION       = 8  // 时装加成标识
	FIGHT_FOR_FRIENDSHIP    = 9  // 羁绊加成标识
	FIGHT_FOR_TEAMSHIP      = 10 // 伙伴配合加成标识
	FIGHT_FOR_TOTEM         = 11 // 阵印加成
	FIGHT_FOR_SEALDEBOOK    = 12 // 天书加成
	FIGHT_FOR_BUDDY_COOP    = 13 // 伙伴协力加成
	FIGHT_FOR_CLIQUE_KONGFU = 14 // 帮派武功
	FIGHT_FOR_AWAKEN        = 15 // 觉醒属性
)

type InFormRoleInfo struct {
	RoleId            int8 // id
	Pos               int  // 站位
	NotGetAllAddition bool // 不获取全体加成
	FormType          int8 // 站位类型 主阵 支援
}

type BattleBiz struct {
}

// 战斗数据初始化的时候使用
func (this BattleBiz) GetInFormRoleForBattle(db *mdb.Database, playerForm *mdb.PlayerFormation, mainRoleOnly bool) (inFormRoleInfos []*InFormRoleInfo, jobs [][]*battle.JobInfo) {
	inFormRoleInfos = make([]*InFormRoleInfo, 1, ALL_FIGHTER_POS_NUM)

	/* 主阵部分 */
	inFormRoleInfos = this.setMainFormRoleForBattle(db, playerForm, mainRoleOnly, inFormRoleInfos)
	//职业已经废弃
	//jobs = this.GetJobsForBattle(inFormRoleInfos, mainRoleOnly)

	return
}

// 构造主阵的InFormRoleInfo
func (this BattleBiz) setMainFormRoleForBattle(db *mdb.Database, playerForm *mdb.PlayerFormation, mainRoleOnly bool, inFormRoleInfos []*InFormRoleInfo) []*InFormRoleInfo {
	/*
		如果mainRoleOnly为真
		inFormRoleInfos中只有一个元素 就是主角

		如果mainRoleOnly为假
		inFormRoleInfos中有2个元素 0为主角 1为伙伴
		inFormRoleInfos中有3个元素 0为主角 1为小伙伴 2为大伙伴
	*/
	if playerForm.Pos0 != NONE_ROLE_ID {
		inFormRoleInfos = this.setFormRoleInfo(playerForm.Pos0, POS_TYPE_MAIN_FORM, 1, mainRoleOnly, inFormRoleInfos)
	}
	if playerForm.Pos1 != NONE_ROLE_ID {
		inFormRoleInfos = this.setFormRoleInfo(playerForm.Pos1, POS_TYPE_MAIN_FORM, 2, mainRoleOnly, inFormRoleInfos)
	}
	if playerForm.Pos2 != NONE_ROLE_ID {
		inFormRoleInfos = this.setFormRoleInfo(playerForm.Pos2, POS_TYPE_MAIN_FORM, 3, mainRoleOnly, inFormRoleInfos)
	}
	if playerForm.Pos3 != NONE_ROLE_ID {
		inFormRoleInfos = this.setFormRoleInfo(playerForm.Pos3, POS_TYPE_MAIN_FORM, 6, mainRoleOnly, inFormRoleInfos)
	}
	if playerForm.Pos4 != NONE_ROLE_ID {
		inFormRoleInfos = this.setFormRoleInfo(playerForm.Pos4, POS_TYPE_MAIN_FORM, 7, mainRoleOnly, inFormRoleInfos)
	}
	if playerForm.Pos5 != NONE_ROLE_ID {
		inFormRoleInfos = this.setFormRoleInfo(playerForm.Pos5, POS_TYPE_MAIN_FORM, 8, mainRoleOnly, inFormRoleInfos)
	}
	if playerForm.Pos6 != NONE_ROLE_ID {
		inFormRoleInfos = this.setFormRoleInfo(playerForm.Pos6, POS_TYPE_MAIN_FORM, 11, mainRoleOnly, inFormRoleInfos)
	}
	if playerForm.Pos7 != NONE_ROLE_ID {
		inFormRoleInfos = this.setFormRoleInfo(playerForm.Pos7, POS_TYPE_MAIN_FORM, 12, mainRoleOnly, inFormRoleInfos)
	}
	if playerForm.Pos8 != NONE_ROLE_ID {
		inFormRoleInfos = this.setFormRoleInfo(playerForm.Pos8, POS_TYPE_MAIN_FORM, 13, mainRoleOnly, inFormRoleInfos)
	}

	//TODO 临时取消排序
	//// 伙伴排序
	//if !mainRoleOnly && len(inFormRoleInfos) == 3 && inFormRoleInfos[1].RoleId > inFormRoleInfos[2].RoleId {
	//	inFormRoleInfos[1], inFormRoleInfos[2] = inFormRoleInfos[2], inFormRoleInfos[1]
	//}

	return inFormRoleInfos
}

// 构造InFormRoleInfo对象
func (this BattleBiz) setFormRoleInfo(roleId, formType int8, pos int, mainRoleOnly bool, inFormRoleInfos []*InFormRoleInfo) []*InFormRoleInfo {
	if role_dat.IsMainRole(roleId) {
		inFormRoleInfos[0] = &InFormRoleInfo{
			RoleId:   roleId,
			Pos:      pos,
			FormType: POS_TYPE_MAIN_FORM,
		}
	} else if !mainRoleOnly {
		inFormRoleInfos = append(inFormRoleInfos, &InFormRoleInfo{
			RoleId:   roleId,
			Pos:      pos,
			FormType: formType,
		})
	}
	return inFormRoleInfos
}

func (this BattleBiz) SetFighters(db *mdb.Database, inFormRoleInfos []*InFormRoleInfo, fighters []*battle.Fighter,
	autoFight bool, isCountFightersNum bool, fightersNumType int8) (fighterNums map[int16]int32, totalFighterNum int32) {
	/*
		1.战斗 isCountFightersNum = false
		2.战力变化 isCountFightersNum = true fightersNumType = 具体模块参数
		3.单个角色详细战力 isCountFightersNum = true fightersNumType = 0
	*/
	if isCountFightersNum {
		fighterNums = make(map[int16]int32)
	}

	playerForm := db.Lookup.PlayerFormation(db.PlayerId())

	playerRole := make(map[int8]*mdb.PlayerRole)
	db.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
		playerRole[row.RoleId()] = row.GoObject()
	})

	/*
		创建各角色的Fighter对象
		fighters 上阵的fighter
		inFormRoleInfos 上阵的InFormRoleInfo
	*/
	for _, inFormRoleInfo := range inFormRoleInfos {
		fighter := newFighterForPlayerRole(db.PlayerId(), int(inFormRoleInfo.Pos+1), playerRole[inFormRoleInfo.RoleId], autoFight, role_dat.IsMainRole(inFormRoleInfo.RoleId))
		fighters[inFormRoleInfo.Pos] = fighter
	}

	// 角色等级数据战力
	totalFighterNum = this.countEachModuleFighterNum(isCountFightersNum, FIGHT_FOR_ROLE_LEVEL, fighterNums, 0, this.getFightNumForFighters(fighters))

	// ###### 各种属性加成 coding here ######

	// 绝招加成
	if fightersNumType == FIGHT_FOR_ALL || fightersNumType == FIGHT_FOR_SKILL {
		this.fighterSetSkillInfo(db, fighters, inFormRoleInfos, playerForm)
		totalFighterNum = this.countEachModuleFighterNum(isCountFightersNum, FIGHT_FOR_SKILL, fighterNums, totalFighterNum, this.getFightNumForFighters(fighters))
	}

	// 设置初次魂侍技能
	if fightersNumType == FIGHT_FOR_ALL || fightersNumType == FIGHT_FOR_GHOST_SKILL {
		this.fighterSetGhostSkillInfo(db, fighters, inFormRoleInfos, playerForm)
		totalFighterNum = this.countEachModuleFighterNum(isCountFightersNum, FIGHT_FOR_GHOST_SKILL, fighterNums, totalFighterNum, this.getFightNumForFighters(fighters))
	}

	// 装备加成
	if fightersNumType == FIGHT_FOR_ALL || fightersNumType == FIGHT_FOR_EQUIP {
		this.fighterSetEquipInfo(db, fighters, inFormRoleInfos)
		totalFighterNum = this.countEachModuleFighterNum(isCountFightersNum, FIGHT_FOR_EQUIP, fighterNums, totalFighterNum, this.getFightNumForFighters(fighters))
	}

	// 羁绊加成
	if fightersNumType == FIGHT_FOR_ALL || fightersNumType == FIGHT_FOR_FRIENDSHIP {
		this.fighterSetFriShpInfo(db, fighters, inFormRoleInfos)
		totalFighterNum = this.countEachModuleFighterNum(isCountFightersNum, FIGHT_FOR_FRIENDSHIP, fighterNums, totalFighterNum, this.getFightNumForFighters(fighters))
	}

	// 魂侍加成
	if Player.IsOpenFunc(db, player_dat.FUNC_GHOST) {
		if fightersNumType == FIGHT_FOR_ALL || fightersNumType == FIGHT_FOR_GHOST {
			this.fighterGhostAddData(db, fighters, inFormRoleInfos)
			totalFighterNum = this.countEachModuleFighterNum(isCountFightersNum, FIGHT_FOR_GHOST, fighterNums, totalFighterNum, this.getFightNumForFighters(fighters))
		}
	}

	// 剑心加成
	if Player.IsOpenFunc(db, player_dat.FUNC_SWORD_SOUL) {
		if fightersNumType == FIGHT_FOR_ALL || fightersNumType == FIGHT_FOR_SWORD_SOUL {
			this.fighterSwordSoulAddData(db, fighters, inFormRoleInfos)
			totalFighterNum = this.countEachModuleFighterNum(isCountFightersNum, FIGHT_FOR_SWORD_SOUL, fighterNums, totalFighterNum, this.getFightNumForFighters(fighters))
		}
	}

	//时装加成
	if Fashion.FashionEnable(db) {
		if fightersNumType == FIGHT_FOR_ALL || fightersNumType == FIGHT_FOR_FASHION {
			this.fighterSetFashionInfo(db, fighters, inFormRoleInfos)
			totalFighterNum = this.countEachModuleFighterNum(isCountFightersNum, FIGHT_FOR_FASHION, fighterNums, totalFighterNum, this.getFightNumForFighters(fighters))
		}
	}

	//伙伴配合加成
	if fightersNumType == FIGHT_FOR_ALL || fightersNumType == FIGHT_FOR_TEAMSHIP {
		this.fighterSetTeamship(db, fighters, inFormRoleInfos)
		totalFighterNum = this.countEachModuleFighterNum(isCountFightersNum, FIGHT_FOR_TEAMSHIP, fighterNums, totalFighterNum, this.getFightNumForFighters(fighters))
	}

	//阵印加成
	if Player.IsOpenFunc(db, player_dat.FUNC_TOTEM) {
		if fightersNumType == FIGHT_FOR_ALL || fightersNumType == FIGHT_FOR_TOTEM {
			this.fighterSetTotem(db, fighters, inFormRoleInfos)
			totalFighterNum = this.countEachModuleFighterNum(isCountFightersNum, FIGHT_FOR_TOTEM, fighterNums, totalFighterNum, this.getFightNumForFighters(fighters))
		}
	}

	//天书加成
	if fightersNumType == FIGHT_FOR_ALL || fightersNumType == FIGHT_FOR_SEALDEBOOK {
		this.fighterSetSealedBook(db, fighters, inFormRoleInfos)
		totalFighterNum = this.countEachModuleFighterNum(isCountFightersNum, FIGHT_FOR_SEALDEBOOK, fighterNums, totalFighterNum, this.getFightNumForFighters(fighters))
	}

	//伙伴协力
	if fightersNumType == FIGHT_FOR_ALL || fightersNumType == FIGHT_FOR_BUDDY_COOP {
		this.fighterSetBuddyCoop(db, fighters, inFormRoleInfos)
		totalFighterNum = this.countEachModuleFighterNum(isCountFightersNum, FIGHT_FOR_BUDDY_COOP, fighterNums, totalFighterNum, this.getFightNumForFighters(fighters))
	}

	//帮派武功
	if fightersNumType == FIGHT_FOR_ALL || fightersNumType == FIGHT_FOR_CLIQUE_KONGFU {
		this.fighterSetCliqueKongfu(db, fighters, inFormRoleInfos)
		totalFighterNum = this.countEachModuleFighterNum(isCountFightersNum, FIGHT_FOR_CLIQUE_KONGFU, fighterNums, totalFighterNum, this.getFightNumForFighters(fighters))
	}

	//觉醒属性
	if fightersNumType == FIGHT_FOR_ALL || fightersNumType == FIGHT_FOR_AWAKEN {
		this.fighterSetAwaken(db, fighters, inFormRoleInfos)
		totalFighterNum = this.countEachModuleFighterNum(isCountFightersNum, FIGHT_FOR_AWAKEN, fighterNums, totalFighterNum, this.getFightNumForFighters(fighters))
	}

	// 全部加成完成之后再设置满血
	for _, f := range fighters {
		if f != nil {
			f.Health = f.MaxHealth
		}
	}

	return
}

// 加载战斗角色的招式信息
func (this BattleBiz) fighterSetSkillInfo(db *mdb.Database, fighters []*battle.Fighter, inFormRoleInfos []*InFormRoleInfo, playerForm *mdb.PlayerFormation) {
	db.Select.PlayerUseSkill(func(playerUseSkill *mdb.PlayerUseSkillRow) {
		for _, inFormRoleInfo := range inFormRoleInfos {
			if inFormRoleInfo.RoleId == playerUseSkill.RoleId() {
				fighter := fighters[inFormRoleInfo.Pos]
				switch role_dat.IsMainRole(playerUseSkill.RoleId()) {
				case true:
					// 主角
					result := make([]*battle.SkillInfo, 0, 4)
					if playerUseSkill.SkillId1() > skill_dat.SKILL_IS_NULL {
						result = append(result, this.getMainRoleSkillInfo(db, playerUseSkill.SkillId1(), playerUseSkill.RoleId()))
					}
					if playerUseSkill.SkillId2() > skill_dat.SKILL_IS_NULL {
						result = append(result, this.getMainRoleSkillInfo(db, playerUseSkill.SkillId2(), playerUseSkill.RoleId()))
					}
					if playerUseSkill.SkillId3() > skill_dat.SKILL_IS_NULL {
						result = append(result, this.getMainRoleSkillInfo(db, playerUseSkill.SkillId3(), playerUseSkill.RoleId()))
					}
					if playerUseSkill.SkillId4() > skill_dat.SKILL_IS_NULL {
						result = append(result, this.getMainRoleSkillInfo(db, playerUseSkill.SkillId4(), playerUseSkill.RoleId()))
					}
					fighter.SkillInfos = result
				case false:
					result := make([]*battle.SkillInfo, 0, 2)

					if playerUseSkill.SkillId1() > skill_dat.SKILL_IS_NULL {
						skillModelInfo1 := skill_dat.GetSkillInfo(playerUseSkill.SkillId1())
						result = append(result, this.getBuddySkillInfo(db, skillModelInfo1, playerUseSkill.RoleId()))
					} else {
						result = append(result, &battle.SkillInfo{})
					}

					//FIXME 伙伴技能1、4槽有效（逻辑不同意，需要修复）
					if playerUseSkill.SkillId4() > skill_dat.SKILL_IS_NULL {
						skillModelInfo2 := skill_dat.GetSkillInfo(playerUseSkill.SkillId4())
						result = append(result, this.getBuddySkillInfo(db, skillModelInfo2, playerUseSkill.RoleId()))
					}

					fighter.SkillInfos = result
				}
				break
			}
		}
	})
	return
}

func (this BattleBiz) fighterSetGhostSkillInfo(db *mdb.Database, fighters []*battle.Fighter, inFormRoleInfos []*InFormRoleInfo, playerForm *mdb.PlayerFormation) {
	db.Select.PlayerGhostEquipment(func(row *mdb.PlayerGhostEquipmentRow) {
		for _, inFormRoleInfo := range inFormRoleInfos {
			if inFormRoleInfo.RoleId == row.RoleId() {
				fighter := fighters[inFormRoleInfo.Pos]

				//只有主魂侍需要加载，因为起他魂侍不能放技能
				if row.Pos1() > 0 {
					playerGhost := db.Lookup.PlayerGhost(row.Pos1())
					ghostSkill := skill_dat.GetGhostSkillByGhostId(int8(playerGhost.GhostId))

					//fighter.InitGhostPower += Ghost.GetRoleInitGhostPower(playerGhost)

					ghost := &battle.FightGhost{
						GhostId:      playerGhost.GhostId,
						GhostStar:    playerGhost.Star,
						GhostLevel:   playerGhost.Level,
						GhostSkillId: int(ghostSkill.Id),
						GhostSkillLv: playerGhost.SkillLevel,
					}
					if relationGhost := db.Lookup.PlayerGhost(playerGhost.RelationId); relationGhost != nil {
						// 连锁魂侍
						relationGhostSkill := skill_dat.GetGhostSkillByGhostId(int8(relationGhost.GhostId))
						ghost.RelationGhost = &battle.FightGhost{
							GhostId:      relationGhost.GhostId,
							GhostStar:    relationGhost.Star,
							GhostLevel:   relationGhost.Level,
							GhostSkillId: int(relationGhostSkill.Id),
							GhostSkillLv: relationGhost.SkillLevel,
						}
					}
					fighter.Ghosts = append(fighter.Ghosts, ghost)
				}

				//fighter.GhostPower = fighter.InitGhostPower
			}
		}
	})
}

// 构造伙伴的战斗招式对象
func (this BattleBiz) getBuddySkillInfo(db *mdb.Database, skillModelInfo *skill_dat.SkillInfo, roleId int8) *battle.SkillInfo {
	var playerSkill *mdb.PlayerSkill
	db.Select.PlayerSkill(func(row *mdb.PlayerSkillRow) {
		if row.SkillId() == skillModelInfo.Id && row.RoleId() == roleId {
			playerSkill = row.GoObject()
			row.Break()
		}
	})
	skillInfo := &battle.SkillInfo{
		SkillId:    int(skillModelInfo.Id),
		SkillTrnLv: int16(playerSkill.SkillTrnlv),
		//SkillJob:      int(skillModelInfo.ChildType),
		//Rhythm:        int(skillContent.RecoverRoundNum + 1),
		//UseRhythm:     int(skillContent.RecoverRoundNum + skillContent.ReleaseNum),
		//RecoverRhythm: int(skillContent.RecoverRoundNum),
	}
	if skillModelInfo.ChildKind == skill_dat.SKILL_KIND_ULTIMATE {
		skillContent := skill_dat.GetSkillContent(skillModelInfo.Id)
		skillInfo.ReleaseNum = int(skillContent.ReleaseNum)
		skillInfo.MaxReleaseNum = int(skillContent.ReleaseNum)
	}
	return skillInfo
}

// 获取主角的战斗招式对象
func (this BattleBiz) getMainRoleSkillInfo(db *mdb.Database, skillId int16, roleId int8) *battle.SkillInfo {
	var playerSkill *mdb.PlayerSkill
	db.Select.PlayerSkill(func(row *mdb.PlayerSkillRow) {
		if row.SkillId() == skillId && row.RoleId() == roleId {
			playerSkill = row.GoObject()
			row.Break()
		}
	})
	return &battle.SkillInfo{
		SkillId:    int(skillId),
		SkillJob:   int(skill_dat.GetSkillInfo(skillId).ChildType),
		SkillTrnLv: int16(playerSkill.SkillTrnlv),
	}
}

func (this BattleBiz) countEachModuleFighterNum(isCountFightersNum bool, fighterNumType int16, fighterNums map[int16]int32, oldFightNum, newFightNum int32) int32 {
	if !isCountFightersNum {
		return 0
	}

	fighterNums[fighterNumType] = newFightNum - oldFightNum
	return newFightNum
}

// 计算多个fighter的战力
func (this BattleBiz) getFightNumForFighters(fighters []*battle.Fighter) int32 {
	var fightNum int32
	for _, fighter := range fighters {
		if fighter != nil {
			fightNum += this.getFightNumForFighter(fighter)
		}
	}
	return fightNum
}

// 计算单个fighter的战力
func (this BattleBiz) getFightNumForFighter(fighter *battle.Fighter) int32 {
	var fightNum float64
	fightNum += float64(fighter.Level * fight_num_dat.LEVEL_RATE)                   // 等级
	fightNum += float64(fighter.MaxHealth) * float64(fight_num_dat.MAX_HEALTH_RATE) // 生命
	fightNum += fighter.Cultivation * fight_num_dat.CULTIVATION_RATE                // 内力
	fightNum += fighter.Speed * fight_num_dat.SPEED_RATE                            // 速度
	fightNum += fighter.Attack * fight_num_dat.ATTACK_RATE                          // 攻击
	fightNum += fighter.Defend * fight_num_dat.DEFEND_RATE                          // 防御
	fightNum += float64(fighter.SunderMaxValue * fight_num_dat.SUNDER_RATE)         // 破甲
	fightNum += fighter.CritialLevel * fight_num_dat.CRITIAL_LEVEL_RATE             // 暴击等级
	fightNum += fighter.TenacityLevel * fight_num_dat.TENACITY_LEVEL_RATE           // 韧性等级
	fightNum += fighter.HitLevel * fight_num_dat.HIT_LEVEL_RATE                     // 命中等级
	fightNum += fighter.DodgeLevel * fight_num_dat.DODGE_LEVEL_RATE                 // 闪避等级
	fightNum += fighter.DestroyLevel * fight_num_dat.DESTROY_LEVEL_RATE             // 破击等级
	fightNum += fighter.BlockLevel * fight_num_dat.BLOCK_LEVEL_RATE                 // 格挡等级
	fightNum += fighter.SpiritHurt * fight_num_dat.SPIRIT_HURT_RATE                 // 仙灵伤害
	fightNum += fighter.HumanHurt * fight_num_dat.HUMAN_HURT_RATE                   // 人畜伤害
	fightNum += fighter.DevilHurt * fight_num_dat.DEVIL_HURT_RATE                   // 妖魔伤害

	fightNum += fighter.SleepLevel * fight_num_dat.SLEEP_LEVEL_RATE
	fightNum += fighter.DizzinessLevel * fight_num_dat.DIZZINESS_LEVEL_RATE
	fightNum += fighter.RandomLevel * fight_num_dat.RANDOM_LEVEL_RATE
	fightNum += fighter.DisableSkillLevel * fight_num_dat.DISABLE_SKILL_LEVEL_RATE
	fightNum += fighter.PoisoningLevel * fight_num_dat.POISONING_LEVEL_RATE

	fightNum += fighter.CritialHurtLevel * fight_num_dat.CRITIAL_HURT_LEVEL_RATE // 必杀等级
	fightNum += fighter.Critial * fight_num_dat.CRITIAL_RATE                     // 暴击百分比
	fightNum += fighter.Tenacity * fight_num_dat.TENACITY_RATE                   // 韧性百分比
	fightNum += fighter.Hit * fight_num_dat.HIT_RATE                             // 命中百分比
	fightNum += fighter.Dodge * fight_num_dat.DODGE_RATE                         // 闪避百分比
	fightNum += fighter.Destroy * fight_num_dat.DESTROY_RATE                     // 破击百分比
	fightNum += fighter.Block * fight_num_dat.BLOCK_RATE                         // 格挡百分比

	fightNum += fighter.Sleep * fight_num_dat.SLEEP_RATE
	fightNum += fighter.Dizziness * fight_num_dat.DIZZINESS_RATE
	fightNum += fighter.Random * fight_num_dat.RANDOM_RATE
	fightNum += fighter.DisableSkill * fight_num_dat.DISABLE_SKILL_RATE
	fightNum += fighter.Poisoning * fight_num_dat.POISONING_RATE

	fightNum += fighter.CritialHurt * fight_num_dat.CRITIAL_HURT_RATE // 必杀百分比
	fightNum -= 1000

	/*
		fmt.Println("======================")
		fmt.Println(fighter.RoleId)
		fmt.Println("等级=", float64(fighter.Level))
		fmt.Println("生命=", float64(fighter.MaxHealth))
		fmt.Println("内力=", fighter.Cultivation)
		fmt.Println("速度=", fighter.Speed)
		fmt.Println("攻击=", fighter.Attack)
		fmt.Println("防御=", fighter.Defend)
		fmt.Println("破甲=", float64(fighter.SunderMaxValue))
		fmt.Println("暴击等级=", fighter.CritialLevel)
		fmt.Println("韧性等级=", fighter.TenacityLevel)
		fmt.Println("命中等级=", fighter.HitLevel)
		fmt.Println("闪避等级=", fighter.DodgeLevel)
		fmt.Println("破击等级=", fighter.DestroyLevel)
		fmt.Println("格挡等级=", fighter.BlockLevel)
		fmt.Println("意志等级=", fighter.WillLevel)
		fmt.Println("暴击等级=", fighter.CritialHurtLevel)
		fmt.Println("暴击百分比=", fighter.Critial)
		fmt.Println("韧性百分比=", fighter.Tenacity)
		fmt.Println("命中百分比=", fighter.Hit)
		fmt.Println("闪避百分比=", fighter.Dodge)
		fmt.Println("破击百分比=", fighter.Destroy)
		fmt.Println("格挡百分比=", fighter.Block)
		fmt.Println("意志百分比=", fighter.Will)
		fmt.Println("暴击百分比=", fighter.CritialHurt)
		fmt.Println("======================")
	*/

	return int32(math.Ceil(fightNum))
}

//加载时装信息
func (this BattleBiz) fighterSetFashionInfo(db *mdb.Database, fighters []*battle.Fighter, inFormRoleInfos []*InFormRoleInfo) {
	mainRole := Role.GetMainRole(db)
	mainRoleLv := int32(mainRole.Level)
	for _, inFormRoleInfo := range inFormRoleInfos {
		if inFormRoleInfo.RoleId == mainRole.RoleId {
			fighter := fighters[inFormRoleInfo.Pos]
			fashionState := db.Lookup.PlayerFashionState(db.PlayerId())
			db.Select.PlayerFashion(func(row *mdb.PlayerFashionRow) {
				fashionDat := fashion_dat.GetFashionData(row.FashionId())
				if fashionDat.Level <= mainRoleLv {
					fighter.MaxHealth += int(fashionDat.Health)
					fighter.Speed += float64(fashionDat.Speed)
					fighter.Attack += float64(fashionDat.Attack)
					fighter.Defend += float64(fashionDat.Defence)
					fighter.Cultivation += float64(fashionDat.Cultivation)
					fighter.FashionId = fashionState.DressedFashionId
				}
			})
		}
	}
}

// 加载战斗角色的羁绊信息
func (this BattleBiz) fighterSetFriShpInfo(db *mdb.Database, fighters []*battle.Fighter, inFormRoleInfos []*InFormRoleInfo) {
	for _, inFormRoleInfo := range inFormRoleInfos {
		fighter := fighters[inFormRoleInfo.Pos]
		friendshipStuff := role_dat.GetRoleFriendship(inFormRoleInfo.RoleId, int16(fighter.FriendshipLevel))
		fighter.MaxHealth += int(friendshipStuff.Health)
		fighter.Attack += float64(friendshipStuff.Attack)
		fighter.Defend += float64(friendshipStuff.Defend)
		fighter.Cultivation += float64(friendshipStuff.Cultivation)
	}
}

// 加载伙伴配合信息
func (this BattleBiz) fighterSetTeamship(db *mdb.Database, fighters []*battle.Fighter, inFormRoleInfos []*InFormRoleInfo) {

	// get player team info
	teamInfo := db.Lookup.PlayerTeamInfo(db.PlayerId())
	if teamInfo == nil {
		return
	}

	// convert each attribute level to value
	ts_health := team_dat.GetTeamshipStuff(teamInfo.HealthLv).Health
	ts_attack := team_dat.GetTeamshipStuff(teamInfo.AttackLv).Attack
	ts_defence := team_dat.GetTeamshipStuff(teamInfo.DefenceLv).Defence

	for _, inFormRoleInfo := range inFormRoleInfos {
		fighter := fighters[inFormRoleInfo.Pos]
		fighter.MaxHealth += int(ts_health)
		fighter.Attack += float64(ts_attack)
		fighter.Defend += float64(ts_defence)
	}
}

func (this BattleBiz) fighterSetTotem(db *mdb.Database, fighters []*battle.Fighter, inFormRoleInfos []*InFormRoleInfo) {
	playerTotemInfo := db.Lookup.PlayerTotemInfo(db.PlayerId())
	if playerTotemInfo == nil {
		return
	}
	for _, inFormRoleInfo := range inFormRoleInfos {
		totemEquipmentPosDataPtr := []*int64{
			&playerTotemInfo.Pos1,
			&playerTotemInfo.Pos2,
			&playerTotemInfo.Pos3,
			&playerTotemInfo.Pos4,
			&playerTotemInfo.Pos5,
		}
		for _, ptr := range totemEquipmentPosDataPtr {
			if *ptr != totem_dat.TOTEM_POS_EMPTY {
				playerTotem := db.Lookup.PlayerTotem(*ptr)
				quality := totem_dat.GetTotemQualiyById(playerTotem.TotemId)
				totemLevelInfo := totem_dat.GetTotemLevelInfo(quality, playerTotem.Level)
				fighter := fighters[inFormRoleInfo.Pos]
				fighter.Health += int(totemLevelInfo.Health)
				fighter.MaxHealth += int(totemLevelInfo.Health)
				fighter.Attack += totemLevelInfo.Attack
				fighter.Defend += totemLevelInfo.Defence
				fighter.Cultivation += totemLevelInfo.Cultivation
			}
		}
	}
}

// 加载战斗角色的装备信息
func (this BattleBiz) fighterSetEquipInfo(db *mdb.Database, fighters []*battle.Fighter, inFormRoleInfos []*InFormRoleInfo) {
	db.Select.PlayerEquipment(func(row *mdb.PlayerEquipmentRow) {
		for _, inFormRoleInfo := range inFormRoleInfos {
			if inFormRoleInfo.RoleId == row.RoleId() {
				var minRefineLevel int16
				var playerEquipment *mdb.PlayerItem
				fighter := fighters[inFormRoleInfo.Pos]

				if row.WeaponId() > 0 {
					setFighterEquipData(db, fighter, row.WeaponId())
					playerEquipment = db.Lookup.PlayerItem(row.WeaponId())
					minRefineLevel = playerEquipment.RefineLevel
				}

				if row.ClothesId() > 0 {
					setFighterEquipData(db, fighter, row.ClothesId())
					playerEquipment = db.Lookup.PlayerItem(row.ClothesId())
					if minRefineLevel > playerEquipment.RefineLevel {
						minRefineLevel = playerEquipment.RefineLevel
					}
				} else {
					minRefineLevel = 0
				}

				if row.AccessoriesId() > 0 {
					setFighterEquipData(db, fighter, row.AccessoriesId())
					playerEquipment = db.Lookup.PlayerItem(row.AccessoriesId())
					if minRefineLevel > playerEquipment.RefineLevel {
						minRefineLevel = playerEquipment.RefineLevel
					}
				} else {
					minRefineLevel = 0
				}

				if row.ShoeId() > 0 {
					setFighterEquipData(db, fighter, row.ShoeId())
					playerEquipment = db.Lookup.PlayerItem(row.ShoeId())
					if minRefineLevel > playerEquipment.RefineLevel {
						minRefineLevel = playerEquipment.RefineLevel
					}
				} else {
					minRefineLevel = 0
				}
				equipmentResonanceDat := item_dat.GetEquipmentResonance(minRefineLevel)
				fighter.Health += equipmentResonanceDat.Health
				fighter.MaxHealth += equipmentResonanceDat.Health
				fighter.Attack += equipmentResonanceDat.Attack
				fighter.Defend += equipmentResonanceDat.Defence

				break
			}
		}
	})
}

// 装备属性加成
func setFighterEquipData(db *mdb.Database, fighter *battle.Fighter, playerItemId int64) {
	playerItem := db.Lookup.PlayerItem(playerItemId)

	if playerItem == nil {
		log.Errorf("[setFighterEquipData]playerItem不应该为nil. playerItemId = %d", playerItemId)
		return
	}

	item_eq := item_dat.GetItem(playerItem.ItemId)

	fighter.MaxHealth += int(item_eq.Health)
	fighter.Speed += float64(item_eq.Speed)
	fighter.Attack += float64(item_eq.Attack)
	fighter.Defend += float64(item_eq.Defence)

	// 装备精炼等级加成
	if playerItem.RefineLevel > 0 {
		addEquipmentRefineAddDate(fighter, item_eq, playerItem.RefineLevel)
	}

	if playerItem.AppendixId != 0 {
		playerItemAppendix := db.Lookup.PlayerItemAppendix(playerItem.AppendixId)
		fighter.MaxHealth += int(playerItemAppendix.Health)
		fighter.Speed += float64(playerItemAppendix.Speed)
		fighter.Attack += float64(playerItemAppendix.Attack)
		fighter.Defend += float64(playerItemAppendix.Defence)
		fighter.Cultivation += float64(playerItemAppendix.Cultivation)

		fighter.DodgeLevel += float64(playerItemAppendix.DodgeLevel)
		fighter.HitLevel += float64(playerItemAppendix.HitLevel)
		fighter.BlockLevel += float64(playerItemAppendix.BlockLevel)
		fighter.CritialLevel += float64(playerItemAppendix.CriticalLevel)
		fighter.TenacityLevel += float64(playerItemAppendix.TenacityLevel)
		fighter.DestroyLevel += float64(playerItemAppendix.DestroyLevel)
	}
}

// 装备精炼等级加成
func addEquipmentRefineAddDate(fighter *battle.Fighter, item *item_dat.Item, refineLevel int16) {
	//equipmentRefine := item_dat.GetEquipmentRefineLevel(item.Quality, refineLevel)
	equipmentRefine := item_dat.GetEquipmentRefineDat(int8(item.TypeId), int8(item.EquipTypeId))
	switch item.TypeId {
	case item_dat.TYPE_WEAPON:
		fighter.Attack += refineGainAttr(equipmentRefine.BaseVal, equipmentRefine.IncreVal, refineLevel)
	case item_dat.TYPE_CLOTHES:
		fighter.Defend += refineGainAttr(equipmentRefine.BaseVal, equipmentRefine.IncreVal, refineLevel)
	case item_dat.TYPE_SHOE:
		fighter.Defend += refineGainAttr(equipmentRefine.BaseVal, equipmentRefine.IncreVal, refineLevel)
	case item_dat.TYPE_ACCESSORIES:
		fighter.MaxHealth += int(refineGainAttr(equipmentRefine.BaseVal, equipmentRefine.IncreVal, refineLevel))
	}
}

func refineGainAttr(baseVal int32, increVal int32, refineLevel int16) float64 {
	// 向上取整
	//return math.Ceil(float64(refine_base) * float64(gainPct) / 100)
	return float64(increVal * int32(refineLevel))
}

// 魂侍加成和主魂侍绝招信息
func (this BattleBiz) fighterGhostAddData(db *mdb.Database, fighters []*battle.Fighter, inFormRoleInfos []*InFormRoleInfo) {
	db.Select.PlayerGhostEquipment(func(row *mdb.PlayerGhostEquipmentRow) {
		for _, inFormRoleInfo := range inFormRoleInfos {

			if row.RoleId() == inFormRoleInfo.RoleId {
				fighter := fighters[inFormRoleInfo.Pos]

				var totalGhostHth int32
				var shieldGhost *mdb.PlayerGhost

				if row.Pos1() > 0 {
					playerGhost := db.Lookup.PlayerGhost(row.Pos1())
					ghostAddData := addEquipmentGhostAddDate(db, fighter, playerGhost)
					totalGhostHth += ghostAddData.Health

					if playerGhost.Star >= 2 {
						shieldGhost = playerGhost
					}
				}

				if row.Pos2() > 0 {
					playerGhost := db.Lookup.PlayerGhost(row.Pos2())
					ghostAddData := addEquipmentGhostAddDate(db, fighter, playerGhost)
					totalGhostHth += ghostAddData.Health
				}

				if row.Pos3() > 0 {
					playerGhost := db.Lookup.PlayerGhost(row.Pos3())
					ghostAddData := addEquipmentGhostAddDate(db, fighter, playerGhost)
					totalGhostHth += ghostAddData.Health
				}

				if row.Pos4() > 0 {
					playerGhost := db.Lookup.PlayerGhost(row.Pos4())
					ghostAddData := addEquipmentGhostAddDate(db, fighter, playerGhost)
					totalGhostHth += ghostAddData.Health
				}

				if shieldGhost != nil {

					// 护盾值比例
					shieldDiscount := ghost_dat.GetGhostShieldDiscount(shieldGhost.Star)

					fighter.EnableGhostShield = true
					fighter.GhostShieldValue = int(float64(totalGhostHth) * shieldDiscount)
					fighter.ShieldGhostId = shieldGhost.GhostId
				}
			}
		}
	})
}

// 增加魂侍加成
func addEquipmentGhostAddDate(db *mdb.Database, fighter *battle.Fighter, playerGhost *mdb.PlayerGhost) (ghostAddData *ghost_dat.GhostAddData) {

	ghostAddData = Ghost.GetGhostAddData(playerGhost)

	// 连锁魂侍加成
	var change float64 = 1.0
	if playerGhost.RelationId > 0 {
		change += ghost_dat.RELATION_GHOST_ADD_PER
	}
	fighter.MaxHealth += int(float64(ghostAddData.Health) * change)
	fighter.Attack += float64(ghostAddData.Attack) * change
	fighter.Defend += float64(ghostAddData.Defence) * change

	return ghostAddData
}

// 剑心加成
func (this BattleBiz) fighterSwordSoulAddData(db *mdb.Database, fighters []*battle.Fighter, inFormRoleInfos []*InFormRoleInfo) {
	db.Select.PlayerSwordSoulEquipment(func(row *mdb.PlayerSwordSoulEquipmentRow) {
		for _, inFormRoleInfo := range inFormRoleInfos {
			if inFormRoleInfo.RoleId == row.RoleId() {

				fighter := fighters[inFormRoleInfo.Pos]

				// 获取角色剑心装备信息
				if row.Pos1() > 0 {
					playerSwordSoul := db.Lookup.PlayerSwordSoul(row.Pos1())
					realAddSwordSoulAddData(fighter, playerSwordSoul.SwordSoulId, playerSwordSoul.Level)
				}

				if row.Pos2() > 0 {
					playerSwordSoul := db.Lookup.PlayerSwordSoul(row.Pos2())
					realAddSwordSoulAddData(fighter, playerSwordSoul.SwordSoulId, playerSwordSoul.Level)
				}

				if row.Pos3() > 0 {
					playerSwordSoul := db.Lookup.PlayerSwordSoul(row.Pos3())
					realAddSwordSoulAddData(fighter, playerSwordSoul.SwordSoulId, playerSwordSoul.Level)
				}

				if row.Pos4() > 0 {
					playerSwordSoul := db.Lookup.PlayerSwordSoul(row.Pos4())
					realAddSwordSoulAddData(fighter, playerSwordSoul.SwordSoulId, playerSwordSoul.Level)
				}

				if row.Pos5() > 0 {
					playerSwordSoul := db.Lookup.PlayerSwordSoul(row.Pos5())
					realAddSwordSoulAddData(fighter, playerSwordSoul.SwordSoulId, playerSwordSoul.Level)
				}

				if row.Pos6() > 0 {
					playerSwordSoul := db.Lookup.PlayerSwordSoul(row.Pos6())
					realAddSwordSoulAddData(fighter, playerSwordSoul.SwordSoulId, playerSwordSoul.Level)
				}

				if row.Pos7() > 0 {
					playerSwordSoul := db.Lookup.PlayerSwordSoul(row.Pos7())
					realAddSwordSoulAddData(fighter, playerSwordSoul.SwordSoulId, playerSwordSoul.Level)
				}

				if row.Pos8() > 0 {
					playerSwordSoul := db.Lookup.PlayerSwordSoul(row.Pos8())
					realAddSwordSoulAddData(fighter, playerSwordSoul.SwordSoulId, playerSwordSoul.Level)
				}

				if row.Pos9() > 0 {
					playerSwordSoul := db.Lookup.PlayerSwordSoul(row.Pos9())
					realAddSwordSoulAddData(fighter, playerSwordSoul.SwordSoulId, playerSwordSoul.Level)
				}

				break
			}
		}
	})
}

func realAddSwordSoulAddData(fighter *battle.Fighter, swordSoulId int16, level int8) {
	swordSoul := sword_soul_dat.GetSwordSoul(swordSoulId)

	var addValue int32

	// 判断是否满级
	if l := len(swordSoul.LevelValue); level > int8(l) {
		addValue = swordSoul.LevelValue[l-1]
	} else {
		addValue = swordSoul.LevelValue[level-1]
	}

	// 根据剑心类型进行加成
	//ATTACK           // 攻击
	//DEFENCE          // 防御
	//HEALTH           // 生命
	//SPEED            // 速度
	//CULTIVATION      // 内力
	//HIT_LEVEL        // 命中
	//CRITICAL_LEVEL   // 暴击
	//BLOCK_LEVEL      // 格挡
	//DESTROY_LEVEL    // 破击
	//TENACITY_LEVEL   // 韧性
	//DODGE_LEVEL      // 闪避
	//SUNDER_MAX_VALUE // 护甲

	switch swordSoul.TypeId {
	case sword_soul_dat.TYPE_ATTACK:
		fighter.Attack += float64(addValue)
	case sword_soul_dat.TYPE_DEFENCE:
		fighter.Defend += float64(addValue)
	case sword_soul_dat.TYPE_HEALTH:
		fighter.MaxHealth += int(addValue)
	case sword_soul_dat.TYPE_SPEED:
		fighter.Speed += float64(addValue)
	case sword_soul_dat.TYPE_CULTIVATION:
		fighter.Cultivation += float64(addValue)
	case sword_soul_dat.TYPE_HIT_LEVEL:
		fighter.HitLevel += float64(addValue)
	case sword_soul_dat.TYPE_CRITICAL_LEVEL:
		fighter.CritialLevel += float64(addValue)
	case sword_soul_dat.TYPE_BLOCK_LEVEL:
		fighter.BlockLevel += float64(addValue)
	case sword_soul_dat.TYPE_DESTROY_LEVEL:
		fighter.DestroyLevel += float64(addValue)
	case sword_soul_dat.TYPE_TENACITY_LEVEL:
		fighter.TenacityLevel += float64(addValue)
	case sword_soul_dat.TYPE_DODGE_LEVEL:
		fighter.DodgeLevel += float64(addValue)
	case sword_soul_dat.TYPE_SUNDER_MAX_VALUE:
		fighter.SunderMaxValue += int(addValue)
	}
}

func (this BattleBiz) fighterSetSealedBook(db *mdb.Database, fighters []*battle.Fighter, inFormRoleInfos []*InFormRoleInfo) {

	if session, ok := Player.GetPlayerOnline(db.PlayerId()); ok {
		state := State(session)
		stealdbook := state.GetSealedBookRecord().GetRecordsByStatus(item_dat.STEALDBOOK_ACTIVATION)
		var (
			health                       int
			attack, defence, cultivation float64
		)
		for _, item := range stealdbook {
			if bookinfo := item_dat.GetSealedBookInfo(item.ItemType, int16(item.ItemId)); bookinfo != nil {
				health += int(bookinfo.Health)
				attack += float64(bookinfo.Attack)
				defence += float64(bookinfo.Defense)
				cultivation += float64(bookinfo.Cultivation)
			}
		}
		for _, inFormRoleInfo := range inFormRoleInfos {
			fighter := fighters[inFormRoleInfo.Pos]
			fighter.MaxHealth += health
			fighter.Attack += attack
			fighter.Defend += defence
			fighter.Cultivation += cultivation
		}
	}
}

//伙伴协力
func (this BattleBiz) fighterSetBuddyCoop(db *mdb.Database, fighters []*battle.Fighter, inFormRoleInfos []*InFormRoleInfo) {
	inFormRoleSet := map[int8]int8{} // roleId -> _
	for _, inFormRoleInfo := range inFormRoleInfos {
		fighter := fighters[inFormRoleInfo.Pos]
		inFormRoleSet[int8(fighter.RoleId)] = 1
	}
	for _, inFormRoleInfo := range inFormRoleInfos {
		fighter := fighters[inFormRoleInfo.Pos]
		if role_dat.IsMainRole(int8(fighter.RoleId)) {
			var buddyCount int8
			db.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
				if !role_dat.IsMainRole(row.RoleId()) {
					buddyCount++
				}
			})
			coopDat := role_dat.GetMainRoleCooperation(buddyCount)
			fighter.MaxHealth += int(coopDat.Health)
			fighter.Speed += float64(coopDat.Speed)
			fighter.Attack += float64(coopDat.Attack)
			fighter.Defend += float64(coopDat.Defence)
			fighter.Cultivation += float64(coopDat.Cultivation)
		} else {
			db.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
				if row.RoleId() != int8(fighter.RoleId) || row.CoopId() <= 0 {
					return
				}
				coopDat := role_dat.GetBuddyCooperation(row.CoopId())
				_, ok1 := inFormRoleSet[coopDat.RoleId1]
				_, ok2 := inFormRoleSet[coopDat.RoleId2]
				if ok1 && ok2 {
					fighter.MaxHealth += int(coopDat.Health)
					fighter.Speed += float64(coopDat.Speed)
					fighter.Attack += float64(coopDat.Attack)
					fighter.Defend += float64(coopDat.Defence)
					fighter.Cultivation += float64(coopDat.Cultivation)
					fighter.SunderMaxValue += int(coopDat.Sunder)

					fighter.DodgeLevel += float64(coopDat.DodgeLevel)
					fighter.HitLevel += float64(coopDat.HitLevel)
					fighter.BlockLevel += float64(coopDat.BlockLevel)
					fighter.CritialLevel += float64(coopDat.CriticalLevel)
					fighter.CritialHurtLevel += float64(coopDat.CriticalHurtLevel)
					fighter.TenacityLevel += float64(coopDat.TenacityLevel)
					fighter.DestroyLevel += float64(coopDat.DestroyLevel)

					fighter.InitGhostPower += int(coopDat.GhostPower)
					fighter.SkillLevel += int16(coopDat.SkillLevel)
				}
			})
		}
	}
}

//帮派武学
func (this BattleBiz) fighterSetCliqueKongfu(db *mdb.Database, fighters []*battle.Fighter, inFormRoleInfos []*InFormRoleInfo) {
	playerCliqueKongfuAttrib := db.Lookup.PlayerCliqueKongfuAttrib(db.PlayerId())
	if playerCliqueKongfuAttrib == nil {
		return
	}
	for _, inFormRoleInfo := range inFormRoleInfos {
		fighter := fighters[inFormRoleInfo.Pos]
		fighter.MaxHealth += int(playerCliqueKongfuAttrib.Health)
		fighter.Attack += float64(playerCliqueKongfuAttrib.Attack)
		fighter.Defend += float64(playerCliqueKongfuAttrib.Defence)
	}
}

//觉醒属性
func (this BattleBiz) fighterSetAwaken(db *mdb.Database, fighters []*battle.Fighter, inFormRoleInfos []*InFormRoleInfo) {
	// role id mapping to fighter
	fighterById := make(map[int8]*battle.Fighter)
	for _, inFormRoleInfo := range inFormRoleInfos {
		fighterById[inFormRoleInfo.RoleId] = fighters[inFormRoleInfo.Pos]
	}
	// now iterate the database and upgrade fighters
	db.Select.PlayerAwakenGraphic(func(row *mdb.PlayerAwakenGraphicRow) {
		attr := awaken_dat.GetRoleAttr(row.RoleId(), row.AttrImpl())
		if !attr.IsSkill {
			attr_v := float64(attr.Attr) * float64(row.Level())
			fighter := fighterById[row.RoleId()]
			if fighter != nil {
				switch attr.Type {
				case awaken_dat.AWAKE_HEALTH:
					fighter.MaxHealth += int(attr_v)
				case awaken_dat.AWAKE_ATTACK:
					fighter.Attack += attr_v
				case awaken_dat.AWAKE_DEFEND:
					fighter.Defend += attr_v
				case awaken_dat.AWAKE_CULTIVATION:
					fighter.Cultivation += attr_v
				case awaken_dat.AWAKE_SPEED:
					fighter.Speed += attr_v
				case awaken_dat.AWAKE_SUNDER:
					fighter.SunderMaxValue += int(attr_v)
				case awaken_dat.AWAKE_HITLV:
					fighter.HitLevel += attr_v
				case awaken_dat.AWAKE_DODGELV:
					fighter.DodgeLevel += attr_v
				case awaken_dat.AWAKE_BLOCKLV:
					fighter.BlockLevel += attr_v
				case awaken_dat.AWAKE_DESTORYLV:
					fighter.DestroyLevel += attr_v
				case awaken_dat.AWAKE_CRITIALLV:
					fighter.CritialLevel += attr_v
				case awaken_dat.AWAKE_TENACITYLV:
					fighter.TenacityLevel += attr_v
				case awaken_dat.AWAKE_POISON:
					fighter.PoisoningLevel += attr_v
				case awaken_dat.AWAKE_DISSKILL:
					fighter.DisableSkillLevel += attr_v
				case awaken_dat.AWAKE_SLEEP:
					fighter.SleepLevel += attr_v
				case awaken_dat.AWAKE_DIZZINESS:
					fighter.DizzinessLevel += attr_v
				case awaken_dat.AWAKE_RANDOM:
					fighter.RandomLevel += attr_v
				case awaken_dat.AWAKE_SPIRIT:
					fighter.SpiritHurt += attr_v
				case awaken_dat.AWAKE_HUMAN:
					fighter.HumanHurt += attr_v
				case awaken_dat.AWAKE_DEVIL:
					fighter.DevilHurt += attr_v
				}
			}
		}
	})
}
