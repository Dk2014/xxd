package channel_dat

import (
	"fmt"
	"strconv"
)

type MessageTpl interface {
	GetTplId() int16
	GetParameters() []byte
}

type ParamCliqueBoat struct {
	BoatId   int64
	OwnerPid int64
	BoatName string
}

func (param ParamCliqueBoat) GetType() int8 {
	return PARAM_TYPE_CLIQUE_BOAT
}
func (param ParamCliqueBoat) ToString() string {
	return fmt.Sprintf(`{"_type":%d, "boat_name": %s, "boat_id":"%d", "owner_pid": "%d"}`, param.GetType(), strconv.Quote(param.BoatName), param.BoatId, param.OwnerPid)
}

type ParamString struct {
	Content string
}

func (param ParamString) GetType() int8 {
	return PARAM_TYPE_STRING
}

func (param ParamString) ToString() string {
	return fmt.Sprintf(`{"_type":%d, "content":%s}`, param.GetType(), strconv.Quote(param.Content))
}

type ParamPlayer struct {
	Nick []byte
	Pid  int64
}

func (param ParamPlayer) GetType() int8 {
	return PARAM_TYPE_PLAYER
}

func (param ParamPlayer) ToString() string {
	return fmt.Sprintf(`{"_type":%d, "nick":%s, "pid": %d}`, param.GetType(), strconv.Quote(string(param.Nick)), param.Pid)
}

type ParamItem struct {
	ItemType int8
	ItemId   int16
}

func (param ParamItem) GetType() int8 {
	return PARAM_TYPE_ITEM
}

func (param ParamItem) ToString() string {
	return fmt.Sprintf(`{"_type":%d, "item_type":%d, "item_id": %d}`, param.GetType(), param.ItemType, param.ItemId)
}

type ParamClique struct {
	Name string
	Id   int64
}

func (param ParamClique) GetType() int8 {
	return PARAM_TYPE_CLIQUE
}

func (param ParamClique) ToString() string {
	return fmt.Sprintf(`{"_type": %d, "name":%s, "clique_id": %d}`, param.GetType(), strconv.Quote(param.Name), param.Id)
}
