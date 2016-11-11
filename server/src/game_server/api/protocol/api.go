package protocol

import "core/net"
import "core/fail"
import (
	"game_server/api/protocol/announcement_api"
	"game_server/api/protocol/arena_api"
	"game_server/api/protocol/awaken_api"
	"game_server/api/protocol/battle_api"
	"game_server/api/protocol/battle_pet_api"
	"game_server/api/protocol/channel_api"
	"game_server/api/protocol/clique_api"
	"game_server/api/protocol/clique_building_api"
	"game_server/api/protocol/clique_escort_api"
	"game_server/api/protocol/clique_quest_api"
	"game_server/api/protocol/daily_sign_in_api"
	"game_server/api/protocol/debug_api"
	"game_server/api/protocol/despair_land_api"
	"game_server/api/protocol/draw_api"
	"game_server/api/protocol/driving_sword_api"
	"game_server/api/protocol/event_api"
	"game_server/api/protocol/fashion_api"
	"game_server/api/protocol/friend_api"
	"game_server/api/protocol/ghost_api"
	"game_server/api/protocol/item_api"
	"game_server/api/protocol/mail_api"
	"game_server/api/protocol/meditation_api"
	"game_server/api/protocol/mission_api"
	"game_server/api/protocol/money_tree_api"
	"game_server/api/protocol/multi_level_api"
	"game_server/api/protocol/notify_api"
	"game_server/api/protocol/pet_virtual_env_api"
	"game_server/api/protocol/player_api"
	"game_server/api/protocol/push_notify_api"
	"game_server/api/protocol/quest_api"
	"game_server/api/protocol/rainbow_api"
	"game_server/api/protocol/role_api"
	"game_server/api/protocol/server_info_api"
	"game_server/api/protocol/skill_api"
	"game_server/api/protocol/sword_soul_api"
	"game_server/api/protocol/taoyuan_api"
	"game_server/api/protocol/team_api"
	"game_server/api/protocol/totem_api"
	"game_server/api/protocol/tower_api"
	"game_server/api/protocol/town_api"
	"game_server/api/protocol/trader_api"
	"game_server/api/protocol/vip_api"
)

var TheDebugConfig DebugConfig

type DebugConfig interface {
	IsEnableDebugAPI() bool
}

type ProtoError struct {
	Message interface{}
}

type Request interface {
	Process(*net.Session)
	TypeName() string
	GetModuleIdAndActionId() (int8, int8)
}

func DecodeIn(message []byte) Request {

	defer func() {
		if err := recover(); err != nil {
			panic(ProtoError{err})
		}
	}()

	var moduleId = message[0]
	switch moduleId {
	case 0:
		return player_api.DecodeIn(message[1:])
	case 1:
		return town_api.DecodeIn(message[1:])
	case 2:
		return team_api.DecodeIn(message[1:])
	case 3:
		return role_api.DecodeIn(message[1:])
	case 4:
		return mission_api.DecodeIn(message[1:])
	case 5:
		return skill_api.DecodeIn(message[1:])
	case 6:
		return battle_api.DecodeIn(message[1:])
	case 7:
		return item_api.DecodeIn(message[1:])
	case 8:
		return notify_api.DecodeIn(message[1:])
	case 9:
		return ghost_api.DecodeIn(message[1:])
	case 10:
		return sword_soul_api.DecodeIn(message[1:])
	case 12:
		return mail_api.DecodeIn(message[1:])
	case 13:
		return quest_api.DecodeIn(message[1:])
	case 14:
		return friend_api.DecodeIn(message[1:])
	case 15:
		return tower_api.DecodeIn(message[1:])
	case 16:
		return multi_level_api.DecodeIn(message[1:])
	case 17:
		return battle_pet_api.DecodeIn(message[1:])
	case 18:
		return announcement_api.DecodeIn(message[1:])
	case 19:
		return arena_api.DecodeIn(message[1:])
	case 20:
		return vip_api.DecodeIn(message[1:])
	case 21:
		return trader_api.DecodeIn(message[1:])
	case 22:
		return daily_sign_in_api.DecodeIn(message[1:])
	case 23:
		return rainbow_api.DecodeIn(message[1:])
	case 24:
		return event_api.DecodeIn(message[1:])
	case 25:
		return fashion_api.DecodeIn(message[1:])
	case 26:
		return push_notify_api.DecodeIn(message[1:])
	case 27:
		return meditation_api.DecodeIn(message[1:])
	case 28:
		return pet_virtual_env_api.DecodeIn(message[1:])
	case 29:
		return channel_api.DecodeIn(message[1:])
	case 30:
		return driving_sword_api.DecodeIn(message[1:])
	case 31:
		return totem_api.DecodeIn(message[1:])
	case 32:
		return money_tree_api.DecodeIn(message[1:])
	case 33:
		return clique_api.DecodeIn(message[1:])
	case 34:
		return clique_building_api.DecodeIn(message[1:])
	case 35:
		return clique_quest_api.DecodeIn(message[1:])
	case 36:
		return clique_escort_api.DecodeIn(message[1:])
	case 37:
		return despair_land_api.DecodeIn(message[1:])
	case 38:
		return awaken_api.DecodeIn(message[1:])
	case 39:
		return taoyuan_api.DecodeIn(message[1:])
	case 50:
		return draw_api.DecodeIn(message[1:])
	case 88:
		return server_info_api.DecodeIn(message[1:])
	case 99:
		fail.When(!TheDebugConfig.IsEnableDebugAPI(), "Disable Debug API")
		return debug_api.DecodeIn(message[1:])
	}
	panic("unsupported protocol")
}

func DecodeOut(message []byte) Request {

	defer func() {
		if err := recover(); err != nil {
			panic(ProtoError{err})
		}
	}()

	var moduleId = message[0]
	switch moduleId {
	case 0:
		return player_api.DecodeOut(message[1:])
	case 1:
		return town_api.DecodeOut(message[1:])
	case 2:
		return team_api.DecodeOut(message[1:])
	case 3:
		return role_api.DecodeOut(message[1:])
	case 4:
		return mission_api.DecodeOut(message[1:])
	case 5:
		return skill_api.DecodeOut(message[1:])
	case 6:
		return battle_api.DecodeOut(message[1:])
	case 7:
		return item_api.DecodeOut(message[1:])
	case 8:
		return notify_api.DecodeOut(message[1:])
	case 9:
		return ghost_api.DecodeOut(message[1:])
	case 10:
		return sword_soul_api.DecodeOut(message[1:])
	case 12:
		return mail_api.DecodeOut(message[1:])
	case 13:
		return quest_api.DecodeOut(message[1:])
	case 14:
		return friend_api.DecodeOut(message[1:])
	case 15:
		return tower_api.DecodeOut(message[1:])
	case 16:
		return multi_level_api.DecodeOut(message[1:])
	case 17:
		return battle_pet_api.DecodeOut(message[1:])
	case 18:
		return announcement_api.DecodeOut(message[1:])
	case 19:
		return arena_api.DecodeOut(message[1:])
	case 20:
		return vip_api.DecodeOut(message[1:])
	case 21:
		return trader_api.DecodeOut(message[1:])
	case 22:
		return daily_sign_in_api.DecodeOut(message[1:])
	case 23:
		return rainbow_api.DecodeOut(message[1:])
	case 24:
		return event_api.DecodeOut(message[1:])
	case 25:
		return fashion_api.DecodeOut(message[1:])
	case 26:
		return push_notify_api.DecodeOut(message[1:])
	case 27:
		return meditation_api.DecodeOut(message[1:])
	case 28:
		return pet_virtual_env_api.DecodeOut(message[1:])
	case 29:
		return channel_api.DecodeOut(message[1:])
	case 30:
		return driving_sword_api.DecodeOut(message[1:])
	case 31:
		return totem_api.DecodeOut(message[1:])
	case 32:
		return money_tree_api.DecodeOut(message[1:])
	case 33:
		return clique_api.DecodeOut(message[1:])
	case 34:
		return clique_building_api.DecodeOut(message[1:])
	case 35:
		return clique_quest_api.DecodeOut(message[1:])
	case 36:
		return clique_escort_api.DecodeOut(message[1:])
	case 37:
		return despair_land_api.DecodeOut(message[1:])
	case 38:
		return awaken_api.DecodeOut(message[1:])
	case 39:
		return taoyuan_api.DecodeOut(message[1:])
	case 50:
		return draw_api.DecodeOut(message[1:])
	case 88:
		return server_info_api.DecodeOut(message[1:])
	case 99:
		return debug_api.DecodeOut(message[1:])
	}
	panic("unsupported protocol")
}
