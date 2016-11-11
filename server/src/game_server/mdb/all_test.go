package mdb

// import "time"
// import "testing"
// import "fmt"
// import "math/rand"
// import "core/log"

// func Test_MDB(t *testing.T) {
// 	mysqlInfo := make(map[string]interface{})

// 	mysqlInfo["unix_socket"] = "tcp"
// 	mysqlInfo["charset"] = "utf8"
// 	mysqlInfo["host"] = "127.0.0.1"
// 	mysqlInfo["port"] = 3306
// 	mysqlInfo["uname"] = "lms"
// 	mysqlInfo["pass"] = "123456"
// 	mysqlInfo["dbname"] = "test"

// 	log.Setup("./log", false)

// 	Start(false, 1, mysqlInfo, "./log", "./log")

// 	db := g_Database

// 	g_Database.Transaction(9999, func() {
// 		for i := 0; i < 50000; i++ {
// 			user := fmt.Sprintf("testMDB_%d", i)
// 			player := &Player{
// 				User:       user,
// 				Nick:       user,
// 				MainRoleId: 0,
// 			}
// 			db.Init()
// 			db.Insert.Player(player)

// 			insertItem := []int64{}
// 			selectItem := []int64{}
// 			for i := 0; i < 50; i++ {
// 				item := &PlayerItem{
// 					Pid:          player.Id,
// 					ItemId:       int16(i),
// 					Num:          int16(rand.Intn(50) + 1),
// 					IsDressed:    0,
// 					BuyBackState: 0,
// 					AppendixId:   0,
// 					RefineLevel:  0,
// 				}
// 				db.Insert.PlayerItem(item)
// 				insertItem = append(insertItem, item.Id)
// 			}

// 			db.Select.PlayerItem(func(row *PlayerItemRow) {
// 				item := row.GoObject()
// 				selectItem = append(selectItem, item.Id)
// 			})

// 			fmt.Printf("new player: id %d, name %s, lookup player: id %d, name %s \n", player.Id, player.User, db.Lookup.Player(player.Id).Id, db.Lookup.Player(player.Id).User)
// 			fmt.Printf("insert item %v \n", insertItem)
// 			fmt.Printf("look item %v \n", selectItem)
// 		}
// 	})

// 	time.Sleep(time.Second * 10)
// }

// func Test_String(t *testing.T) {
// 	gPlayer := &Player{
// 		User: "test",
// 		Nick: "nickname",
// 	}

// 	cPlayer := gPlayer.CObject()

// 	rPlayer := &PlayerRow{row: cPlayer}

// 	if rPlayer.User() != gPlayer.User {
// 		t.Fatal("username not match")
// 	}

// 	if rPlayer.Nick() != gPlayer.Nick {
// 		t.Fatal("nickname not match")
// 	}
// }

/*
func Test_GlobalTable(t *testing.T) {
	g_transController = newTransController(nil)

	db := NewDatabase()

	// insert
	db.Insert.GlobalArenaData(&GlobalArenaData{
		Num: 1,
		Pid: 10,
	})
	db.Insert.GlobalArenaData(&GlobalArenaData{
		Num: 2,
		Pid: 20,
	})
	db.Insert.GlobalArenaData(&GlobalArenaData{
		Num: 3,
		Pid: 30,
	})

	// lookup
	data1 := db.Lookup.GlobalArenaData(2)
	if data1.Num != 2 || data1.Pid != 20 {
		t.Fatal("data1 not match")
	}

	// update
	data1.Pid = 21
	db.Update.GlobalArenaData(data1)
	data2 := db.Lookup.GlobalArenaData(2)
	if data1.Num != 2 || data1.Pid != 21 {
		t.Fatal("data2 not match")
	}

	// delete
	db.Delete.GlobalArenaData(data2)

	// select
	count := 0
	match1 := false
	match2 := false
	db.Select.GlobalArenaData(func(row *GlobalArenaDataRow) {
		t.Log(row.Num(), row.Pid())

		if row.Num() == 1 {
			if row.Pid() != 10 {
				t.Fatal("not match")
			} else {
				match1 = true
				count += 1
			}
		}

		if row.Num() == 3 {
			if row.Pid() != 30 {
				t.Fatal("not match")
			} else {
				match2 = true
				count += 1
			}
		}
	})
	if count != 2 || match1 == false || match2 == false {
		t.Fatal("select not match")
	}
}

func Test_PlayerTable(t *testing.T) {
	g_transController = newTransController(nil)

	// insert
	db := GetOrNewDatabase(1)
	db.Insert.PlayerFriend(&PlayerFriend{
		Id:        10,
		Pid:       123,
		FriendPid: 456,
	})
	db.Insert.PlayerFriend(&PlayerFriend{
		Id:        11,
		Pid:       321,
		FriendPid: 654,
	})
	db.Insert.PlayerFriend(&PlayerFriend{
		Id:        12,
		Pid:       456,
		FriendPid: 789,
	})

	// lookup
	db = GetOrNewDatabase(1)
	data1 := db.Lookup.PlayerFriend(11)
	if data1.Id != 11 || data1.Pid != 321 || data1.FriendPid != 654 {
		t.Fatal("data1 not match")
	}

	// update
	db = GetOrNewDatabase(1)
	data1.FriendPid = 890
	db.Update.PlayerFriend(data1)
	data2 := db.Lookup.PlayerFriend(11)
	if data2.Id != 11 || data2.Pid != 321 || data2.FriendPid != 890 {
		t.Fatal("data2 not match")
	}

	// delete
	db = GetOrNewDatabase(1)
	db.Delete.PlayerFriend(data2)

	// select
	db = GetOrNewDatabase(1)
	count := 0
	match1 := false
	match2 := false
	db.Select.PlayerFriend(func(row *PlayerFriendRow) {
		t.Log(row.Id(), row.Pid(), row.FriendPid())

		if row.Id() == 10 {
			if row.Pid() != 123 || row.FriendPid() != 456 {
				t.Fatal("not match")
			} else {
				match1 = true
				count += 1
			}
		}

		if row.Id() == 12 {
			if row.Pid() != 456 || row.FriendPid() != 789 {
				t.Fatal("not match")
			} else {
				match2 = true
				count += 1
			}
		}
	})
	if count != 2 || match1 == false || match2 == false {
		t.Fatal("select not match")
	}
}

func Test_Rollback(t *testing.T) {
	g_transController = newTransController(nil)

	// rollback insert
	db := NewDatabase()
	db.Insert.GlobalArenaData(&GlobalArenaData{
		Num: 123,
		Pid: 10,
	})
	if db.Lookup.GlobalArenaData(123) == nil {
		t.Fatal("insert failed")
	}
	g_transController.Rollback()
	if db.Lookup.GlobalArenaData(123) != nil {
		t.Fatal("rollback insert failed")
	}

	// rollback delete
	db = NewDatabase()
	db.Insert.GlobalArenaData(&GlobalArenaData{
		Num: 123,
		Pid: 10,
	})
	data1 := db.Lookup.GlobalArenaData(123)
	if data1 == nil {
		t.Fatal("insert failed")
	}
	g_transController.Commit()
	db.Delete.GlobalArenaData(data1)
	if db.Lookup.GlobalArenaData(123) != nil {
		t.Fatal("delete failed")
	}
	g_transController.Rollback()
	if db.Lookup.GlobalArenaData(123) == nil {
		t.Fatal("rollback delete failed")
	}

	// rollback update
	db = NewDatabase()
	data1.Pid = 11
	db.Update.GlobalArenaData(data1)
	data2 := db.Lookup.GlobalArenaData(123)
	if data2.Pid != 11 {
		t.Fatal("update failed")
	}
	g_transController.Rollback()
	data3 := db.Lookup.GlobalArenaData(123)
	if data3.Pid != 10 {
		t.Fatal("rollback update failed")
	}
}

func Test_Start(t *testing.T) {
	info := make(map[string]interface{})
	info["host"] = "127.0.0.1"
	info["uname"] = "lms"
	info["pass"] = "123456"
	info["port"] = 3306
	info["unix_socket"] = "tcp"
	info["dbname"] = "lms_trunk"
	info["charset"] = "utf8"

	log.Setup()

	Start(info)

	log.Flush()
}

func Test_NestedTrans(t *testing.T) {
	defer func() {
		if err := recover(); err.(TransError).Message.(string) != "nested transaction" {
			t.Fail()
		}
	}()

	trans := newTransController(nil)

	trans.Transaction(func() {
		trans.Transaction(func() {
			t.Log("nested trans")
		})
	})
}

type TestGlobalArenaDataHook struct {
	InsertCallback func()
	DeleteCallback func()
	UpdateCallback func()
}

func (hook *TestGlobalArenaDataHook) Insert(_ *GlobalArenaDataRow) {
	hook.InsertCallback()
}

func (hook *TestGlobalArenaDataHook) Delete(_ *GlobalArenaDataRow) {
	hook.DeleteCallback()
}

func (hook *TestGlobalArenaDataHook) Update(_, _ *GlobalArenaDataRow) {
	hook.UpdateCallback()
}

func Test_Hook(t *testing.T) {
	g_transController = newTransController(nil)

	db := NewDatabase()

	insertOk := false
	deleteOk := false
	updateOk := false

	Hook.GlobalArenaData(&TestGlobalArenaDataHook{
		InsertCallback: func() {
			insertOk = true
		},
		DeleteCallback: func() {
			deleteOk = true
		},
		UpdateCallback: func() {
			updateOk = true
		},
	})

	db.Insert.GlobalArenaData(&GlobalArenaData{
		Num: 234,
		Pid: 10,
	})
	g_transController.Commit()
	if !insertOk {
		t.Fatal("hook insert failed")
	}

	data1 := db.Lookup.GlobalArenaData(234)
	db.Update.GlobalArenaData(data1)
	g_transController.Commit()
	if !updateOk {
		t.Fatal("hook update failed")
	}

	db.Delete.GlobalArenaData(data1)
	g_transController.Commit()
	if !deleteOk {
		t.Fatal("hook delete failed")
	}
}

*/
