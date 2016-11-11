/*
内存数据库钩子
*/
package player

import (
	"game_server/mdb"
	"sync"
)

func init() {
	mdb.Hook.Player(new(playerTableHook))
}

type playerTableHook struct {
	sync.Mutex
}

// 当玩家表载入时触发
func (hook *playerTableHook) Load(row *mdb.PlayerRow) {
	hook.Lock()
	defer hook.Unlock()

	if row.Nick() != "" {
		g_NicknameIndex[row.Nick()] = row.Id()
	}

	if row.User() != "" {
		g_UsernameIndex[row.User()] = row.Id()
	}
}

// 只有在player::Login初次登陆时触发
func (hook *playerTableHook) Insert(row *mdb.PlayerRow) {
	hook.Lock()
	defer hook.Unlock()

	g_UsernameIndex[row.User()] = row.Id()
}

// 只有在player::init时更新
func (hook *playerTableHook) Update(row, old *mdb.PlayerRow) {
	hook.Lock()
	defer hook.Unlock()

	g_NicknameIndex[row.Nick()] = row.Id()
}

// 当玩家表删除时触发
func (hook *playerTableHook) Delete(row *mdb.PlayerRow) {
}
