package main

import "time"
import "fmt"

import "math/rand"
import "core/log"
import "game_server/mdb"

/*
	测试说明：

	> 测试mdb对数据的操作（增删改查）
	> 测试一对多查询，一对一查询

	> 最终测试结果只会保留一半player和player_item各一半的数据(要测试删除)


	PS: 测试数据不符合游戏正常业务逻辑数据格式。请手动删除
*/

const (
	PlayerMaxNum     = 5000
	PlayerItemMaxNum = 5
)

func main() {
	mysqlInfo := make(map[string]interface{})

	mysqlInfo["unix_socket"] = "tcp"
	mysqlInfo["charset"] = "utf8"
	mysqlInfo["host"] = "127.0.0.1"
	mysqlInfo["port"] = 3306
	mysqlInfo["uname"] = "lms"
	mysqlInfo["pass"] = "123456"
	mysqlInfo["dbname"] = "test"

	log.Setup("./log", true)

	mdb.Start(false, 1, mysqlInfo, "./log", "./log")

	pids := []int64{}

	// insert
	for i := 0; i < PlayerMaxNum; i++ {
		user := fmt.Sprintf("testMDB_%d", i)
		player := &mdb.Player{
			User:       user,
			Nick:       user,
			MainRoleId: 0,
		}
		db := mdb.NewDatabase()
		db.Transaction(11111, func() {
			db.Init()
			db.Insert.Player(player)
			pids = append(pids, player.Id)
			fmt.Printf("new player: id %d, name %s, lookup player: id %d, name %s \n", player.Id, player.User, db.Lookup.Player(player.Id).Id, db.Lookup.Player(player.Id).User)
		})
	}

	// lookup|select|insert
	for _, pid := range pids {
		go func(pid int64) {
			mdb.Transaction(9999, func() {
				mdb.GlobalExecute(func(db *mdb.Database) {
					db.AgentExecute(pid, func(pdb *mdb.Database) {
						insertItem := []int64{}
						selectItem := []int64{}
						for i := 0; i < PlayerItemMaxNum; i++ {
							item := &mdb.PlayerItem{
								Pid:          pid,
								ItemId:       int16(i),
								Num:          int16(rand.Intn(20000) + 1),
								IsDressed:    0,
								BuyBackState: 0,
								AppendixId:   0,
								RefineLevel:  0,
							}
							pdb.Insert.PlayerItem(item)
							insertItem = append(insertItem, item.Id)
						}

						pdb.Select.PlayerItem(func(row *mdb.PlayerItemRow) {
							item := row.GoObject()
							selectItem = append(selectItem, item.Id)
						})
						fmt.Printf("insert item %v \n", insertItem)
						fmt.Printf("look item %v \n", selectItem)
					})
				})
			})
		}(pid)
	}

	// delete
	mdb.Transaction(9999, func() {
		mdb.GlobalExecute(func(db *mdb.Database) {
			i := 0
			for _, pid := range pids {
				if len(pids)/2 == i {
					break
				}

				i++
				db.AgentExecute(pid, func(pdb *mdb.Database) {
					pdb.Delete.Player(pdb.Lookup.Player(pid))
					pdb.Select.PlayerItem(func(row *mdb.PlayerItemRow) {
						pdb.Delete.PlayerItem(row.GoObject())
					})
				})
			}
		})
	})

	time.Sleep(5 * time.Minute)
}
