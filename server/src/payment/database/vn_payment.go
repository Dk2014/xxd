package database

import (
	corelog "core/log"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

type PrivateData struct {
	IP     string `json:"ip"`
	Public bool   `json:"public"`
}

type PaymentLog struct {
	Id            int64  `column:"-"`
	OrderId       string `column:"order_id"`
	ProductCount  int64  `column:"product_count"`
	Amount        string `column:"amount"`
	PayStatus     string `column:"pay_status"`
	PayTime       string `column:"pay_time"`
	UserId        string `column:"user_id"`
	OrderType     string `column:"order_type"`
	GameUserId    int64  `column:"game_user_id"`
	ServerId      int64  `column:"server_id"`
	ProductName   string `column:"product_name"`
	ProductId     string `column:"product_id"`
	PrivateData   string `column:"private_data"`
	ChannelNumber string `column:"channel_number"`
	Sign          string `column:"sign"`
	Source        string `column:"source"`
	EnchancedSign string `column:"enhanced_sign"`
	//OriginRequest string
}

func (record *PaymentLog) InsertSQL() string {
	return g_Hepler.InsertSQL("payment_log", record)
	//return "INSERT INTO `payment_log` (`order_id`, `product_count`, `amount`, `pay_status`, `pay_time`, `user_id`, `order_type`, `game_user_id`, `server_id`, `product_name`, `product_id`, `private_data`, `channel_number`, `sign`, `source`, `enhanced_sign` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ); "
}

type PendingQueue struct {
	Id           int64
	TryTimestamp int64  `column:"try_timestamp"`
	Try          int8   `column:"try"`
	ProductId    string `column:"product_id"`
	GameUserId   int64  `column:"game_user_id"`
	Amount       string `column:"amount"`
	IP           string `column:"ip"`
	Public       int8   `column:"public"`
}

func (record *PendingQueue) InsertSQL() string {
	return g_Hepler.InsertSQL("pending_queue", record)
	//return "INSERT INTO `pending_queue` (`id`, `try_timestamp`, `try`, `product_id`, `game_user_id`,`amount`, `ip`) VALUES (?, ?, ?, ?, ?, ?, ?);"
}

//开启事务插入 payment_log 插入pending_queue
func NewPaymentLog(log *PaymentLog) (dup bool) {
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
		log.OrderId, log.ProductCount, log.Amount, log.PayStatus, log.PayTime, log.UserId, log.OrderType, log.GameUserId, log.ServerId, log.ProductName, log.ProductId, log.PrivateData, log.ChannelNumber, log.Sign, log.Source, log.EnchancedSign)

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

	privateData := &PrivateData{}
	if err1 := json.Unmarshal([]byte(log.PrivateData), privateData); err1 != nil {
		corelog.Warnf("unmarsha private data error: %v. record id %d", err1, log.Id)
		_ = err1
	}

	pending := &PendingQueue{
		Id:         latestId,
		Try:        0,
		ProductId:  log.ProductId,
		GameUserId: log.GameUserId,
		Amount:     log.Amount,
		IP:         privateData.IP,
		// amount maybe
	}
	if privateData.Public {
		pending.Public = 1
	}
	_, err = tx.Exec(pending.InsertSQL(), pending.Id, 0, pending.Try, pending.ProductId, pending.GameUserId, pending.Amount, pending.IP, pending.Public)
	if err != nil {
		panic(err)
	}
	tx.Commit()
	return false
}

func FetchPendingById(id int64) (*PendingQueue, bool) {
	rows, err := g_db.Query(fmt.Sprintf("select id, try_timestamp, try, product_id, game_user_id, amount,ip from pending_queue where id=%d", id))
	if err != nil {
		panic(err)
	}
	defer func() {
		rows.Close()
	}()
	var (
		ok bool
		pq PendingQueue
	)

	for rows.Next() {
		ok = true
		rows.Scan(&pq.Id, &pq.TryTimestamp, &pq.Try, &pq.ProductId, &pq.GameUserId, &pq.Amount, &pq.IP)
	}
	if ok {
		return &pq, ok
	}
	return nil, ok
}

//启动时开启处理掉所有调用
func FetchPending(triedTime int8, limit int32) (results []*PendingQueue) {
	rows, err := g_db.Query(fmt.Sprintf("select id, try_timestamp, try, product_id, game_user_id, amount,ip from pending_queue where try= %d limit %d", triedTime, limit))
	if err != nil {
		panic(err)
	}
	defer func() {
		rows.Close()
	}()
	for rows.Next() {
		var pq PendingQueue
		rows.Scan(&pq.Id, &pq.TryTimestamp, &pq.Try, &pq.ProductId, &pq.GameUserId, &pq.Amount, &pq.IP)
		results = append(results, &pq)
	}
	return results
}
