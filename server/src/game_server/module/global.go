package module

import (
	"crypto/md5"
	"game_server/config"
	"sync"
)

var (
	g_sessionStateStore map[int64]*SessionState = make(map[int64]*SessionState)
	g_stateStoreMutex   sync.Mutex
)

func StoreSessionState(state *SessionState) {
	g_stateStoreMutex.Lock()
	defer g_stateStoreMutex.Unlock()

	g_sessionStateStore[state.PlayerId] = state
}

func ReStoreSessionState(playerId int64) *SessionState {
	g_stateStoreMutex.Lock()
	defer g_stateStoreMutex.Unlock()

	if state, ok := g_sessionStateStore[playerId]; ok {
		delete(g_sessionStateStore, playerId)
		return state
	}

	return nil
}

func GetServerIdWithPlayerId(pid int64) (int, bool) {
	sid := int(int32(pid >> 32))
	return sid, (sid == config.ServerCfg.ServerId)
}

func HashNow(delimiter []byte, sequence [][]byte) []byte {
	hash := md5.New()
	for i, unit := range sequence {
		if i > 0 {
			hash.Write(delimiter)
		}

		hash.Write(unit)
	}
	return hash.Sum(nil)
}

func PostContent(unit *PostUnit, delay int64) {
	postman.doPost(unit, delay)
}

func DeleteContent(unit *PostUnit) {
	postman.doDelete(unit)
}

func GetLatestWorldChannelMessage() []*Message {
	return worldChatChannel.GetLatest(50, 7*86400)
}

func AddWorldChannelMessage(msg *Message) {
	worldChatChannel.AddMsg(msg)
}
