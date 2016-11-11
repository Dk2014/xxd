package clique_dat

import (
	"core/mysql"
)

var (
	mapCliqueTemple map[int8][]*CliqueTemple
	mapLeveUpFee    map[int16]*CliqueTempleUpgrade
)

type CliqueTemple struct {
	//Level   int16 // 宗祠等级
	Fame    int16 // 声望
	Contrib int32 // 帮贡
}

type CliqueTempleUpgrade struct {
	UpgradeFee int32  // 升级费用（铜钱）
	Desc       string // 对应等级描述
}

func loadCliqueTemple(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM clique_temple ORDER BY `worship_type`,`level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iWorshipType := res.Map("worship_type")
	iFame := res.Map("fame")
	iContrib := res.Map("contrib")
	//iLevel := res.Map("level")
	var worshiptype int8
	mapCliqueTemple = map[int8][]*CliqueTemple{}
	for _, row := range res.Rows {
		worshiptype = row.Int8(iWorshipType)
		mapCliqueTemple[worshiptype] = append(mapCliqueTemple[worshiptype], &CliqueTemple{
			//	Level:   row.Int16(iLevel),
			Fame:    row.Int16(iFame),
			Contrib: row.Int32(iContrib),
		})
	}
}

func loadCliqueTempleUpgrade(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM clique_temple_upgrade ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iUpgradeFee := res.Map("upgrade_fee")
	iLevel := res.Map("level")
	iDesc := res.Map("desc")

	mapLeveUpFee = map[int16]*CliqueTempleUpgrade{}
	for _, row := range res.Rows {
		mapLeveUpFee[row.Int16(iLevel)] = &CliqueTempleUpgrade{
			UpgradeFee: row.Int32(iUpgradeFee),
			Desc:       row.Str(iDesc),
		}
	}
}

// ############### 对外接口实现 coding here ###############

func GetCliqueTempleByType(worshipType int8, level int16) *CliqueTemple {
	cliqueTempleList := mapCliqueTemple[worshipType]
	if int(level) <= len(cliqueTempleList) {
		return cliqueTempleList[level-1]
	}

	return nil
}

func GetUpgradeFeeByLevel(level int16) *CliqueTempleUpgrade {
	if fee, ok := mapLeveUpFee[level]; ok {
		return fee
	}
	return nil
}
