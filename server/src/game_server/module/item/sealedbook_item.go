package item

import (
	"game_server/api/protocol/item_api"
	"game_server/dat/item_dat"
	"game_server/module"
)

func addSealedBook(itemType int8, itemid int16, out *item_api.GetSealedbooks_Out, state *module.SessionState) {
	var exist bool
	sealedbook := item_dat.GetSealedBookInfo(itemType, itemid)
	if sealedbook != nil {
		for _, value := range out.Items {
			if value.ItemId == int64(itemid) {
				exist = true
			}
		}

		if !exist {
			out.Items = append(out.Items, item_api.GetSealedbooks_Out_Items{
				ItemType: itemType,
				ItemId:   int64(itemid),
				Status:   item_dat.STEALDBOOK_HAVING,
			})
			//既然存在，就增加这个record
			state.GetSealedBookRecord().AddRecord(itemType, itemid, item_dat.STEALDBOOK_HAVING, state.Database)
		}

	}
}
