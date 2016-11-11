package battle

import (
	"fmt"
	"testing"
)

func TestBuffer(t *testing.T) {
	var f = new(Fighter)

	f.Attack = 10
	f.raw.Attack = 10

	f.addBuff(nil, BUFF_ATTACK, 20, 3, 10, 1)

	fmt.Println(f.Buffs.Keep)
	fmt.Println(f.Attack)
	fmt.Println("\n")

	f.Buffs.Update(f)
	fmt.Println(f.Buffs.Keep)
	f.addBuff(nil, BUFF_ATTACK, 20, 3, 10, 1)
	fmt.Println(f.Attack)
	fmt.Println("\n")

	f.Buffs.Update(f)
	f.Buffs.Update(f)
	fmt.Println(f.Buffs.Keep)
	f.addBuff(nil, BUFF_ATTACK, 20, 3, 10, 1)
	fmt.Println(f.Attack)
	fmt.Println("\n")

	f.Buffs.Update(f)
	f.Buffs.Update(f)
	f.Buffs.Update(f)
	f.addBuff(nil, BUFF_ATTACK, 20, 3, 10, 1)
	fmt.Println(f.Attack)
	fmt.Println("\n")

	t.Fail()
}

func TestRecord(t *testing.T) {
	StartDefaultRecordDb("battle_record.db")
	db := GetDefaultRecordDb()
	db.GetRecord(7, func(batType int, details, v []byte) {
		fmt.Println(len(v), batType)
	})
	StopDefaultRecordDb()
}
