package arena_award_box_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapArenaAwardBox []*ArenaAwardBox
)

func Load(db *mysql.Connection) {
	loadArenaAwardBox(db)
}

type ArenaAwardBox struct {
	MaxRank  int32 // 排名
	Fame     int32 //声望
	Ingot    int32 // 元宝
	Coins    int32 // 铜钱
	ItemId   int16 // 物品ID
	ItemNum  int16 // 物品数量
	Item2Id  int16 // 物品2
	Item2Num int16 // 物品2数量
}

func loadArenaAwardBox(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM arena_award_box ORDER BY `max_rank` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iMaxRank := res.Map("max_rank")
	iFame := res.Map("fame")
	iIngot := res.Map("ingot")
	iCoins := res.Map("coins")
	iItemId := res.Map("item_id")
	iItemNum := res.Map("item_num")
	iItem2Id := res.Map("item2_id")
	iItem2Num := res.Map("item2_num")
	mapArenaAwardBox = []*ArenaAwardBox{}
	for _, row := range res.Rows {
		mapArenaAwardBox = append(mapArenaAwardBox, &ArenaAwardBox{
			MaxRank:  row.Int32(iMaxRank),
			Fame:     row.Int32(iFame),
			Ingot:    row.Int32(iIngot),
			Coins:    row.Int32(iCoins),
			ItemId:   row.Int16(iItemId),
			ItemNum:  row.Int16(iItemNum),
			Item2Id:  row.Int16(iItem2Id),
			Item2Num: row.Int16(iItem2Num),
		})
	}
}

// ############### 对外接口实现 coding here ###############

func GetAwardWithRank(rank int32) *ArenaAwardBox {
	for _, box := range mapArenaAwardBox {
		if rank <= box.MaxRank {
			return box
		}
	}

	fail.When(true, "can't found rank box")
	return nil
}
