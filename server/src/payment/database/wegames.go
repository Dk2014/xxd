package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

type WegamesPaymentLog struct {
	Id               int64  `column:"-"`
	OrderId          string `column:"order_id"`
	PlatformUid      string `column:"platform_uid"`
	PayAmount        string `column:"pay_amount"`
	TwdMoney         string `column:"twd_money"`
	GameCode         string `column:"game_code"`
	ServerCode       int    `column:"server_code"`
	GameMoney        int64  `column:"game_money"`
	PresentGameMoney int64  `column:"present_game_money"`
	OtherItem        string `column:"other_item"`
	VirtualItems     string `column:"virtual_items"`
	Time             int64  `column:"time"`
	Sign             string `column:"sign"`
}

func (record *WegamesPaymentLog) InsertSQL() string {
	return g_Hepler.InsertSQL("wegames_payment_log", record)
}

type WegamesPendingQueue struct {
	Id               int64  `column:"id"`
	TryTimestamp     int64  `column:"try_timestamp"`
	Try              int8   `column:"try"`
	OrderId          string `column:"order_id"`
	PlatformUid      string `column:"platform_uid"`
	ServerCode       int    `column:"server_code"`
	TwdMoney         string `column:"twd_money"`
	GameMoney        int64  `column:"game_money"`
	PresentGameMoney int64  `column:"present_game_money"`
	OtherItem        string `column:"other_item"`
	VirtualItems     string `column:"virtual_items"`
	Time             int64  `column:"time"`
}

func (record *WegamesPendingQueue) InsertSQL() string {
	return g_Hepler.InsertSQL("wegames_pending_queue", record)
}

//开启事务插入 payment_log 插入pending_queue
func WegamesNewPaymentLog(log *WegamesPaymentLog) (dup bool) {
	var tx *sql.Tx
	var err error
	defer func() {
		if err := recover(); err != nil {
			if tx != nil {
				tx.Rollback()
			}
			panic(err)
		}
	}()
	tx, err = g_db.Begin()
	if err != nil {
		panic(err)
	}
	result, err := tx.Exec(log.InsertSQL(),
		log.OrderId, log.PlatformUid, log.PayAmount, log.TwdMoney, log.GameCode, log.ServerCode, log.GameMoney, log.PresentGameMoney, log.OtherItem, log.VirtualItems, log.Time, log.Sign)

	if err != nil {
		//超级黑科技
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			tx.Rollback()
			return true
		}
		panic(err)
	}

	latestId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	log.Id = latestId
	pending := &WegamesPendingQueue{
		Id:               latestId,
		TryTimestamp:     0,
		Try:              0,
		OrderId:          log.OrderId,
		PlatformUid:      log.PlatformUid,
		TwdMoney:         log.TwdMoney,
		ServerCode:       log.ServerCode,
		GameMoney:        log.GameMoney,
		PresentGameMoney: log.PresentGameMoney,
		OtherItem:        log.OtherItem,
		VirtualItems:     log.VirtualItems,
		Time:             log.Time,
	}
	_, err = tx.Exec(pending.InsertSQL(), pending.Id, pending.TryTimestamp, pending.Try, pending.OrderId, pending.PlatformUid, pending.ServerCode, pending.TwdMoney, pending.GameMoney, pending.PresentGameMoney, pending.OtherItem, pending.VirtualItems, pending.Time)
	if err != nil {
		panic(err)
	}
	tx.Commit()
	return false
}

func WegamesFetchPendingById(id int64) (*WegamesPendingQueue, bool) {
	rows, err := g_db.Query(fmt.Sprintf("select id, try_timestamp, try, order_id, platform_uid, server_code, twd_money , game_money, present_game_money, other_item, virtual_items, time from pending_queue where id=%d", id))
	if err != nil {
		panic(err)
	}
	defer func() {
		rows.Close()
	}()
	var (
		ok bool
		pq WegamesPendingQueue
	)

	for rows.Next() {
		ok = true
		rows.Scan(&pq.Id, &pq.TryTimestamp, &pq.Try, &pq.OrderId, &pq.PlatformUid, &pq.ServerCode, &pq.TwdMoney, &pq.GameMoney, &pq.PresentGameMoney, &pq.OtherItem, &pq.VirtualItems, &pq.Time)
	}
	if ok {
		return &pq, ok
	}
	return nil, ok
}

//启动时开启处理掉所有调用
func WegamesFetchPending(triedTime int8, limit int32) (results []*WegamesPendingQueue) {
	rows, err := g_db.Query(fmt.Sprintf("select id, try_timestamp, try, order_id, platform_uid, server_code, twd_money,  game_money, present_game_money, other_item, virtual_items, time from wegames_pending_queue where try=%d limit %d", triedTime, limit))
	if err != nil {
		panic(err)
	}
	defer func() {
		rows.Close()
	}()
	for rows.Next() {
		var pq WegamesPendingQueue
		rows.Scan(&pq.Id, &pq.TryTimestamp, &pq.Try, &pq.OrderId, &pq.PlatformUid, &pq.ServerCode, &pq.TwdMoney, &pq.GameMoney, &pq.PresentGameMoney, &pq.OtherItem, &pq.VirtualItems, &pq.Time)
		results = append(results, &pq)
	}
	return results
}
