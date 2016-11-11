package event_dat

import (
	"core/time"
	"game_server/dat/mail_dat"
)

/*

  "TotalPhysicalCost":
    {
      "StartUnixTime": 0,
      "EndUnixTime": 0,
      "MailTitle": "累计消耗体力奖励",
      "MailContent": "累计消耗体力{0}点，特此奖励，再接再厉!",
      "MailParams": "cost,100",
      "List": [
        {
          "TotalCost": 0,
          "Ingot ": 0,
          "Coin ": 0,
          "Item1Id ": 0,
          "Item1Num ": 0,
          "Item2Id ": 0,
          "Item2Num ": 0,
          "Item3Id  ": 0,
          "Item3Num ": 0
        }
      ]
    }


*/

var (
	totalPhysicalCost    *TotalPhysicalCost
	mapTotalPhysicalCost map[int32][]*mail_dat.Attachment = make(map[int32][]*mail_dat.Attachment)
)

type PhysicalTotalAward struct {
	TotalCost int32 // 累计消耗体力
	EventDefaultAward
}

type TotalPhysicalCost struct {
	StartUnixTime int64
	EndUnixTime   int64

	MailTitle   string
	MailContent string
	MailParams  string

	List []*PhysicalTotalAward
}

func EnableTotalPhysicalCost() (ok bool) {
	if totalPhysicalCost != nil {
		nowTime := time.GetNowTime()
		ok = nowTime >= totalPhysicalCost.StartUnixTime && nowTime < totalPhysicalCost.EndUnixTime
	}
	return
}

func GetTotalPhysicalCostAward(costTotal int64) (attachs []*mail_dat.Attachment, ok bool) {
	// if attachs, ok=mapTotalPhysicalCost[costTotal]
	return
}

// 累计体力消耗运营活动数据配置
func LoadTotalPhysicalCost(list *TotalPhysicalCost) {
	totalPhysicalCost = list
	if list != nil {
		for _, item := range list.List {
			mapTotalPhysicalCost[item.TotalCost] = makeAttachment(int64(item.Coin), item.Ingot, item.Item1Id, item.Item1Num, item.Item2Id, item.Item2Num,
				item.Item3Id, item.Item3Num, item.Item4Id, item.Item4Num, item.Item5Id, item.Item5Num)
		}

		totalPhysicalCost.List = nil
	}
}
