package trader

import (
	"game_server/dat/trader_dat"
	"math/rand"
)

//输入格子id，随机返回配置中的(货物类型ID，货物ID, 数量, 价格) 元组
func randomGoodsByGridId(gridId int32, level int16) (goodsType int8, itemId int16, num int16, cost int64) {
	goodsList := trader_dat.GridConfig(gridId)
	rnd := int8(rand.Int31n(100)) //[0-100)
	for _, goods := range goodsList {
		if rnd < goods.Probability {
			goodsType = goods.GoodsType
			itemId = goods.ItemId
			num = goods.Num
			cost = goods.Cost
			if goods.GoodsType == trader_dat.EQUIPMENT {
				equiement := trader_dat.GetGridConfigEquiement(goods.Id, level)
				itemId = equiement.ItemId
				cost = equiement.Cost
				goodsType = trader_dat.ITEM
			}
			return
		}
		rnd -= goods.Probability
	}
	panic("unreachable")
}
