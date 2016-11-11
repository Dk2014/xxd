package clique_rpc

import (
	//"game_server/mdb"
	"sync"
)

var (
	g_PlayerCliqueApplyTable *playerCliqueApplyTable = &playerCliqueApplyTable{}
)

func init() {
	g_PlayerCliqueApplyTable.applyMap = make(map[int64][3]int64)
}

type playerCliqueApplyTable struct {
	sync.RWMutex
	applyMap map[int64][3]int64
}

func AddApply(pid, cliqueId int64) {
	g_PlayerCliqueApplyTable.Lock() //避免同时有人操作‘
	defer g_PlayerCliqueApplyTable.Unlock()

	var (
		exist    bool
		newArray = [3]int64{}
		index    int
	)

	if oldclique, ok := g_PlayerCliqueApplyTable.applyMap[pid]; ok {
		for _, value := range oldclique {
			if cliqueId == value {
				exist = true
				break
			}
			if value != 0 {
				newArray[index] = value
				index++
			}
		}

	}

	//1,无此申请记录，且不足3条
	if !exist && index < 3 {
		newArray[index] = cliqueId
		g_PlayerCliqueApplyTable.applyMap[pid] = newArray
	}

}

func DeleteAllApply(pid int64) {
	g_PlayerCliqueApplyTable.Lock() //避免同时有人操作‘
	defer g_PlayerCliqueApplyTable.Unlock()

	if applies, ok := g_PlayerCliqueApplyTable.applyMap[pid]; ok {
		for _, cliqueId := range applies {
			g_CliqueTable.DeleteApply(pid, cliqueId)
		}

		g_PlayerCliqueApplyTable.applyMap[pid] = [3]int64{}
	}

}

func DeleteApply(pid, cliqueId int64) {
	g_PlayerCliqueApplyTable.Lock() //避免同时有人操作‘
	defer g_PlayerCliqueApplyTable.Unlock()
	var (
		newArray = [3]int64{}
		index    int
	)

	if oldclique, ok := g_PlayerCliqueApplyTable.applyMap[pid]; ok {
		for i := 0; i < len(oldclique); i++ {
			if cliqueId != oldclique[i] {
				newArray[index] = oldclique[i]
				index++
			}
		}

		g_PlayerCliqueApplyTable.applyMap[pid] = newArray
	}

}

func ApplyFull(pid int64) bool {
	g_PlayerCliqueApplyTable.Lock() //避免同时有人操作
	defer g_PlayerCliqueApplyTable.Unlock()

	var (
		count    int
		fullFlag bool
	)

	if oldclique, ok := g_PlayerCliqueApplyTable.applyMap[pid]; ok {
		for _, value := range oldclique {
			if value != 0 {
				count++
			}
		}
	}

	if count >= 3 {
		fullFlag = true
	}

	return fullFlag
}

func FetchPlayerCliqueApply(Pid int64, cb func(int64)) {
	g_PlayerCliqueApplyTable.Lock() //避免同时有人操作
	defer g_PlayerCliqueApplyTable.Unlock()

	if cliqueid, ok := g_PlayerCliqueApplyTable.applyMap[Pid]; ok {
		for _, values := range cliqueid {
			if values > 0 {
				cb(values)
			}
		}
	}
}
