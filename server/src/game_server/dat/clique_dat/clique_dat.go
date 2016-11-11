package clique_dat

import (
	"core/mysql"
)

func Load(db *mysql.Connection) {
	// load所有的帮派相关数据
	loadCliqueCenterBuildingLevelInfo(db)
	loadCliqueBuildingBankDat(db)
	loadCliqueTemple(db)
	loadCliqueTempleUpgrade(db)
	loadCliqueKongfuLevel(db)
	loadCliqueBuildingKongfuLevelInfo(db)
	loadCliqueKongfu(db)
	loadCliqueBuilding(db)
	loadCliqueDailyQuest(db)
	loadCliqueBuildingQuest(db)
	loadCliqueBoat(db)
}
