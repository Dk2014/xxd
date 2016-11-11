package dat

import (
	"core/mysql"

	"game_server/dat/arena_award_box_dat"
	"game_server/dat/arena_buy_cost_config_dat"
	"game_server/dat/awaken_dat"
	"game_server/dat/battle_pet_dat"
	"game_server/dat/buy_boss_level_times_config_dat"
	"game_server/dat/buy_hard_level_times_config_dat"
	"game_server/dat/buy_resource_times_config_dat"
	"game_server/dat/chest_dat"
	"game_server/dat/clique_dat"
	"game_server/dat/coins_exchange_dat"
	"game_server/dat/daily_sign_in_dat"
	"game_server/dat/driving_sword_dat"
	"game_server/dat/enemy_deploy_dat"
	"game_server/dat/event_dat"
	"game_server/dat/fashion_dat"
	"game_server/dat/friend_dat"
	"game_server/dat/ghost_dat"
	"game_server/dat/heart_draw_dat"
	"game_server/dat/item_dat"
	"game_server/dat/mail_dat"
	"game_server/dat/mission_dat"
	"game_server/dat/monster_property_addition_dat"
	"game_server/dat/multi_level_dat"
	"game_server/dat/payments_rule_dat"
	"game_server/dat/physical_buy_cost_config_dat"
	"game_server/dat/player_dat"
	"game_server/dat/push_notify_dat"
	"game_server/dat/quest_dat"
	"game_server/dat/rainbow_buy_cost_config_dat"
	"game_server/dat/rainbow_dat"
	"game_server/dat/role_dat"
	"game_server/dat/skill_dat"
	"game_server/dat/sword_soul_dat"
	"game_server/dat/team_dat"
	"game_server/dat/totem_dat"
	"game_server/dat/tower_level_dat"
	"game_server/dat/town_dat"
	"game_server/dat/trader_dat"
	"game_server/dat/vip_dat"
)

//func GlobalLoad(mysqlInfo map[string]interface{}) {
//	db, err1 := mysql.Connect(mysqlInfo)
//	if err1 != nil {
//		panic(err1)
//	}
//	defer db.Close()
//	clique_dat.GlobalLoad(db)
//}

func Load(mysqlInfo map[string]interface{}) {
	db, err1 := mysql.Connect(mysqlInfo)
	if err1 != nil {
		panic(err1)
	}
	defer db.Close()

	role_dat.Load(db)
	skill_dat.Load(db)
	mission_dat.Load(db)
	item_dat.Load(db)
	mail_dat.Load(db)
	enemy_deploy_dat.Load(db)
	ghost_dat.Load(db)
	player_dat.Load(db)
	sword_soul_dat.Load(db)
	quest_dat.Load(db)
	team_dat.Load(db)
	town_dat.Load(db)
	tower_level_dat.Load(db)
	multi_level_dat.Load(db)
	battle_pet_dat.Load(db)
	heart_draw_dat.Load(db)
	chest_dat.Load(db)
	coins_exchange_dat.Load(db)
	arena_award_box_dat.Load(db)
	vip_dat.Load(db)
	trader_dat.Load(db)
	daily_sign_in_dat.Load(db)
	friend_dat.Load(db)
	event_dat.Load(db)
	rainbow_dat.Load(db)
	fashion_dat.Load(db)
	physical_buy_cost_config_dat.Load(db)
	arena_buy_cost_config_dat.Load(db)
	rainbow_buy_cost_config_dat.Load(db)
	totem_dat.Load(db)
	driving_sword_dat.Load(db)
	clique_dat.Load(db)
	awaken_dat.Load(db)
	push_notify_dat.Load()
	payments_rule_dat.Load(db)
	buy_resource_times_config_dat.Load(db)
	monster_property_addition_dat.Load(db)
	buy_hard_level_times_config_dat.Load(db)
	buy_boss_level_times_config_dat.Load(db)
}
