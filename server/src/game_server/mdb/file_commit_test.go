package mdb

// import "time"
// import "testing"

// func Test_EmptySyncFile(t *testing.T) {
// 	file := NewSyncFile("./")
// 	defer file.Close()
// }

// func Test_DemoSyncFile(t *testing.T) {
// 	file := NewSyncFile("./")
// 	defer file.Close()

// 	oldData := PlayerInfo{
// 		Pid:             1,                                      // int64 玩家ID
// 		Ingot:           100,                                    // int64 元宝
// 		Coins:           200,                                    // int64 铜钱
// 		NewMailNum:      9,                                      // int16 新邮件数
// 		LastLoginTime:   time.Now().Unix(),                      // int64 上次登录时间
// 		LastOfflineTime: time.Now().Unix(),                      // int64 上次下线时间
// 		TotalOnlineTime: int64(time.Hour * 24),                  // int64 总在线时间
// 		InitGlobalSrv:   1,                                      // int8 是否在互动服已初始化. 0 - 没有
// 		FirstLoginTime:  time.Now().Add(-time.Hour * 48).Unix(), // int64 玩家注册时间
// 	}

// 	newData := oldData
// 	newData.Ingot += 30
// 	newData.Coins += 40

// 	log1 := PlayerInfoInsertLog{
// 		GoRow: &oldData,
// 	}

// 	log2 := PlayerInfoDeleteLog{
// 		GoOld: &oldData,
// 	}

// 	log3 := PlayerInfoUpdateLog{
// 		GoNew: &newData,
// 		GoOld: &oldData,
// 	}

// 	file.BeginTrans()
// 	file.WriteUint32(123)
// 	file.WriteInt64(time.Now().Unix())
// 	log1.CommitToFile(file)
// 	log2.CommitToFile(file)
// 	log3.CommitToFile(file)
// 	file.EndTrans()
// }
