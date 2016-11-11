package module

import (
	"game_server/global"
)

type PlayerArenaState struct {
	PlayerRank       int32
	PlayerFightNum   int32
	TargetPlayerRank int32

	TargetInfo *global.PlayerInfo
}
