package player_common

import (
	"core/fail"
	"core/skiplist"
	"fmt"
	"game_server/api/protocol/player_api"
	"sync"
)

type PlayerRankData struct {
	Pid       int64
	Nick      string
	Values    []player_api.GetRanks_Out_Ranks_Values
	Rank      int64
	Timestamp int64 // 数据更新时间
}

func NewPlayerRankData() *PlayerRankData {
	obj := &PlayerRankData{}
	obj.Values = []player_api.GetRanks_Out_Ranks_Values{}
	return obj
}

func (this *PlayerRankData) Clone() *PlayerRankData {
	obj := NewPlayerRankData()
	obj.Pid = this.Pid
	obj.Nick = this.Nick
	obj.Values = nil
	obj.Values = append(obj.Values, this.Values...)
	obj.Rank = this.Rank
	obj.Timestamp = this.Timestamp
	return obj
}

func (this *PlayerRankData) Less(lesser skiplist.Lesser) bool {
	data := lesser.(*PlayerRankData)

	// comparing the supposed rank number
	if this.Values[0].Value > data.Values[0].Value {
		return true
	} else if this.Values[0].Value == data.Values[0].Value {
		if this.Timestamp < data.Timestamp {
			return true
		} else if this.Timestamp == data.Timestamp {
			return this.Pid < data.Pid
		} else {
			return false
		}
	} else {
		return false
	}
}

func (this *PlayerRankData) Equal(lesser skiplist.Lesser) bool {
	data := lesser.(*PlayerRankData)

	return this.Pid == data.Pid
}

type PlayerRankTable struct {
	sync.Mutex
	list  *skiplist.SL
	datas map[int64]*PlayerRankData
}

func NewPlayerRankTable() *PlayerRankTable {
	obj := &PlayerRankTable{}
	obj.list = skiplist.NewSkipList(10000)
	obj.datas = make(map[int64]*PlayerRankData)
	return obj
}

func (this *PlayerRankTable) Update(datas []*PlayerRankData) {
	this.Lock()
	defer this.Unlock()

	for i := 0; i < len(datas); i++ {
		data := datas[i]
		oldData, ok := this.datas[data.Pid]
		if ok {
			this.list.Remove(oldData)
		}
		this.list.Insert(data)
		this.datas[data.Pid] = data
	}
}

func (this *PlayerRankTable) GetByPid(pid int64) *PlayerRankData {
	this.Lock()
	defer this.Unlock()

	data := this.datas[pid]
	fail.When(data == nil, fmt.Sprintf("Fail cause cant find pid %d", pid))

	_, index := this.list.Find(data)
	data.Rank = int64(index + 1)
	return data.Clone()
}

func (this *PlayerRankTable) IterRank(begin int64, callback func(*PlayerRankData) bool) {
	this.Lock()
	defer this.Unlock()

	node := this.list.Index(int(begin - 1))
	if node == nil {
		return
	}

	rank := begin

	this.list.Iter(node, func(node *skiplist.Node) bool {
		data := node.Data.(*PlayerRankData)
		data.Rank = rank
		rank++
		return callback(data)
	})
}

func (this *PlayerRankTable) ViewKeyValue(pid int64) player_api.GetRanks_Out_Ranks_Values {
	this.Lock()
	defer this.Unlock()

	return this.datas[pid].Values[0]
}
