package database

import (
	"fmt"
	"time"
)

func ProcessRecord(id int64, table int) {
	tx, err := g_db.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			if tx != nil {
				tx.Rollback()
			}
			panic(err)
		}
	}()
	_, err = tx.Exec(fmt.Sprintf("update `%s` set `try`=`try`+1, `try_timestamp`=%d where id= %d", TableId2TableName(table), time.Now().Unix(), id))
	if err != nil {
		panic(err)
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}

//通知游戏服次数增加
func BatchProcessRecord(ids []int64, table int) {
	tx, err := g_db.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			if tx != nil {
				tx.Rollback()
			}
			panic(err)
		}
	}()
	now := time.Now().Unix()
	for _, id := range ids {
		_, err := tx.Exec(fmt.Sprintf("update `%s` set `try`=`try`+1,`try_timestamp`=%d where id= %d", TableId2TableName(table), now, id))
		if err != nil {
			panic(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}

func BatchDeliverRecord(ids []int64, table int) {
	tx, err := g_db.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			if tx != nil {
				tx.Rollback()
			}
			panic(err)
		}
	}()
	for _, id := range ids {
		_, err = tx.Exec(fmt.Sprintf("update `%s`set `is_delivered`=1 where id= %d", TableId2TableName(table), id))
		if err != nil {
			panic(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}

//拿到游戏服处理完成确认
func DeleteRecordById(id int64, table int) {
	tx, err := g_db.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			if tx != nil {
				tx.Rollback()
			}
			panic(err)
		}
	}()
	_, err = tx.Exec(fmt.Sprintf("delete from `%s` where id= %d", TableId2TableName(table), id))
	if err != nil {
		panic(err)
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}

func BatchDeleteRecordById(ids []int64, table int) {
	tx, err := g_db.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			if tx != nil {
				tx.Rollback()
			}
			panic(err)
		}
	}()
	for _, id := range ids {
		_, err = tx.Exec(fmt.Sprintf("delete from `%s` where id= %d", TableId2TableName(table), id))
		if err != nil {
			panic(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
