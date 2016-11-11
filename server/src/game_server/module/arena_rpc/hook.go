/*
内存数据库钩子
*/
package arena_rpc

import (
	"game_server/mdb"
)

func init() {
	mdb.Hook.GlobalArenaRank(new(arenaRankTableHook))
}

type arenaRankTableHook struct {
}

func (hook *arenaRankTableHook) Load(row *mdb.GlobalArenaRankRow) {
	rankTable.loadPlayerRank(row.Pid(), row.Rank())
}

func (hook *arenaRankTableHook) Insert(row *mdb.GlobalArenaRankRow) {
	rankTable.updatePlayerRank(row.Pid(), row.Rank())
}

func (hook *arenaRankTableHook) Update(row, old *mdb.GlobalArenaRankRow) {
	rankTable.updatePlayerRank(row.Pid(), row.Rank())
}

func (hook *arenaRankTableHook) Delete(row *mdb.GlobalArenaRankRow) {
}
