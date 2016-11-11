package module

import (
	"bytes"
	"core/fail"
	"core/net"
	"encoding/gob"
	"game_server/api/protocol/item_api"
	"game_server/dat/item_dat"
	"game_server/mdb"
)

type SealedBookAttribute struct {
	ItemId int16 // 物品id
	Status int8  // 状态 1：拥有,2，激活
}

type SealedBookRecord map[int8][]*SealedBookAttribute

func (record *SealedBookRecord) Encode() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(record)
	if err != nil {
		fail.When(true, err.Error())
	}
	return buffer.Bytes()
}

func (record *SealedBookRecord) Decode(buf []byte) {
	if len(buf) > 0 {
		buffer := bytes.NewBuffer(buf)
		decoder := gob.NewDecoder(buffer)
		err := decoder.Decode(record)
		if err != nil {
			fail.When(true, err.Error())
		}
	}
}

func (record *SealedBookRecord) Load(db *mdb.Database) {
	sealedBook := db.Lookup.PlayerSealedbook(db.PlayerId())
	if sealedBook == nil {
		sealedBook = &mdb.PlayerSealedbook{
			Pid: db.PlayerId(),
		}
		db.Insert.PlayerSealedbook(sealedBook)
	}
	buffer := sealedBook.SealedRecord
	record.Decode(buffer)
}

func (record *SealedBookRecord) FindRecord(itemType int8, itemID int16) (status int8, exist bool) {

	if values, ok := (*record)[itemType]; ok {
		for _, item := range values {
			if itemID == item.ItemId {
				status = item.Status
				exist = true
				break
			}
		}
	}

	return
}

func addBookItems(in *item_api.GetSealedbooks_Out, status int8, listout *[]*item_api.GetSealedbooks_Out_Items) {
	if in != nil {
		for _, value := range in.Items {
			if value.Status == status {
				*listout = append(*listout, &item_api.GetSealedbooks_Out_Items{
					ItemType: value.ItemType,
					ItemId:   int64(value.ItemId),
					Status:   value.Status,
				})

			}
		}
	}

	return
}

func (record *SealedBookRecord) GetRecordsByStatus(status int8) []*item_api.GetSealedbooks_Out_Items {
	outItemList := make([]*item_api.GetSealedbooks_Out_Items, 0)

	roles := record.GetRecordsByType(item_dat.STEALDBOOK_TYPE_ROLES)
	addBookItems(roles, status, &outItemList)

	ghost := record.GetRecordsByType(item_dat.STEALDBOOK_TYPE_GHOSTS)
	addBookItems(ghost, status, &outItemList)

	swordsouls := record.GetRecordsByType(item_dat.STEALDBOOK_TYPE_SWORDSOULS)
	addBookItems(swordsouls, status, &outItemList)

	pets := record.GetRecordsByType(item_dat.STEALDBOOK_TYPE_PETS)
	addBookItems(pets, status, &outItemList)

	equipments := record.GetRecordsByType(item_dat.STEALDBOOK_TYPE_WEAPON)
	addBookItems(equipments, status, &outItemList)

	equipments = record.GetRecordsByType(item_dat.STEALDBOOK_TYPE_CLOTHES)
	addBookItems(equipments, status, &outItemList)

	equipments = record.GetRecordsByType(item_dat.STEALDBOOK_TYPE_SHOE)
	addBookItems(equipments, status, &outItemList)

	equipments = record.GetRecordsByType(item_dat.STEALDBOOK_TYPE_ACCESSORIES)
	addBookItems(equipments, status, &outItemList)

	totems := record.GetRecordsByType(item_dat.STEALDBOOK_TYPE_TOTEMS)
	addBookItems(totems, status, &outItemList)

	noequipments := record.GetRecordsByType(item_dat.STEALDBOOK_TYPE_NOEQUIPMENTS)
	addBookItems(noequipments, status, &outItemList)

	battletools := record.GetRecordsByType(item_dat.STEALDBOOK_TYPE_BATTLETOOLS)
	addBookItems(battletools, status, &outItemList)

	return outItemList
}

func (record *SealedBookRecord) GetRecordsByType(itemType int8) *item_api.GetSealedbooks_Out {
	list := &item_api.GetSealedbooks_Out{}
	if values, ok := (*record)[itemType]; ok {
		for _, item := range values {
			list.Items = append(list.Items, item_api.GetSealedbooks_Out_Items{
				ItemType: itemType,
				ItemId:   int64(item.ItemId),
				Status:   item.Status,
			})
		}
	}

	return list
}

func (record *SealedBookRecord) Update(db *mdb.Database) {
	recordBytes := record.Encode()
	oldItem := db.Lookup.PlayerSealedbook(db.PlayerId())
	oldItem.SealedRecord = recordBytes
	db.Update.PlayerSealedbook(oldItem)
}

func (record *SealedBookRecord) AddRecord(itemType int8, itemID int16, status int8, db *mdb.Database) bool {

	if _, ok := record.FindRecord(itemType, itemID); ok {
		return false
	}

	(*record)[itemType] = append((*record)[itemType], &SealedBookAttribute{
		ItemId: itemID, // 物品id
		Status: status, // 状态 1：拥有,2，激活
	})
	record.Update(db)

	return true
}

func (record *SealedBookRecord) ChangeStatus(itemType int8, itemID int16, status int8, db *mdb.Database) {
	if value, ok := (*record)[itemType]; ok {
		var index int
		for i, itemlist := range value {
			if itemlist.ItemId == itemID {
				index = i
			}
		}
		(*record)[itemType][index].Status = status
		record.Update(db)
	}
}

func InitPlayerSealedBookRecord(session *net.Session) {
	state := State(session)
	state.GetSealedBookRecord()
}
