package draw

import (
	"game_server/dat/chest_dat"
	"game_server/dat/ghost_dat"
	"game_server/dat/item_dat"
	"game_server/dat/mail_dat"
	"game_server/dat/player_dat"
	"game_server/dat/sword_soul_dat"
	"game_server/module"
	"game_server/module/rpc"
	"math/rand"
)

func addGhost(state *module.SessionState, ghostId int16, itemReason, xdEventType int32) {
	ghost := ghost_dat.GetGhost(ghostId)
	funcName := player_dat.GetFuncById(player_dat.FUNC_GHOST).Name
	if !module.Player.IsOpenFunc(state.Database, player_dat.FUNC_GHOST) {
		rpc.RemoteMailSend(state.PlayerId, mail_dat.MailDrawTips{
			ItemName: ghost.Name,
			Func:     funcName,
		})
	}
	module.Ghost.AddGhost(state, ghostId, itemReason, xdEventType)
}

func addSwordSoul(state *module.SessionState, swordSoulId int16, itemReason int32) {
	swordSoul := sword_soul_dat.GetSwordSoul(swordSoulId)
	funcName := player_dat.GetFuncById(player_dat.FUNC_SWORD_SOUL).Name
	if !module.Player.IsOpenFunc(state.Database, player_dat.FUNC_SWORD_SOUL) {
		rpc.RemoteMailSend(state.PlayerId, mail_dat.MailDrawTips{
			ItemName: swordSoul.Name,
			Func:     funcName,
		})
	}
	module.SwordSoul.AddSwordSoul(state, swordSoulId, itemReason)
}

func addPet(state *module.SessionState, petId int32, num, itemReason, xdEventType int32) {
	for i := int32(0); i < num; i++ {
		module.BattlePet.AddPetWihoutNotify(state.Database, petId, itemReason, xdEventType)
	}
}

func randomItem(items []*item_dat.Item) (item *item_dat.Item) {
	index := rand.Intn(len(items))
	item = items[index]
	return item
}

func drawChest(chests []*chest_dat.Chest) (chestItem *chest_dat.ChestItem) {

	// 开箱
	probability := int8(rand.Float64() * 100)

	var chestId int32
	for _, chest := range chests {
		if probability < chest.Probability {
			chestId = chest.Id

			break
		}
		probability -= chest.Probability
	}

	// 获取物品
	chestItems := chest_dat.GetChestItems(chestId)
	chestItem = randomChestItem(chestItems)
	return chestItem
}

func randomChestItem(chestItems []*chest_dat.ChestItem) (chestItem *chest_dat.ChestItem) {
	index := rand.Intn(len(chestItems))
	chestItem = chestItems[index]
	return chestItem
}
