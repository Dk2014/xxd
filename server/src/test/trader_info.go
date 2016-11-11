package main

import (
	"fmt"
	"time"

	"client_test"
	"game_server/api/protocol/trader_api"
)

/*
var (
	HASH_DELIMITER = []byte("-")
	HASH_SALT      = []byte("Uh981-jsdO!%&^#kahdbp0;v/'@$1*UOysd")
)

func FromPlatformLogin(client *client_test.Client, platFormType int8, user string, nickname string, roleId int8, restore bool) {
	timestamp := time.Now().Unix()
	hash := module.HashNow(HASH_DELIMITER, [][]byte{
		[]byte(strconv.FormatInt(int64(platFormType), 10)),
		[]byte(user),
		HASH_SALT,
		[]byte(strconv.FormatInt(timestamp, 10)),
		[]byte(nickname),
		[]byte(strconv.FormatInt(int64(roleId), 10)),
	})
	encodedHash := base64.StdEncoding.EncodeToString(hash)
	client.Player_FromPlatformLogin(player_api.PlatformType(platFormType), user, nickname, roleId, encodedHash, timestamp, restore)
}
*/

func main() {
	client := client_test.NewClient("127.0.0.1:8080")
	//client.Player_FromPlatformLogin(1, "trader_info17", "trader_info17", 2, "", 0, false)
	client_test.FromPlatformLogin(client, 1, "trader_info17", "trader_info17", 2, false)

	//注册回调函数
	client.OutTrader_Info = func(out *trader_api.Info_Out) {
		fmt.Println("\n==============")
		fmt.Printf("trader_api.InfoOut:\n%#v\n", out)
		fmt.Println("==============\n")
	}

	client.Trader_Info()

	b := time.Tick(5 * time.Second)
	<-b
}
