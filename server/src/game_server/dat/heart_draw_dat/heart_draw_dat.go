package heart_draw_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapHeartDraw map[int16]*HeartDraw
)

func Load(db *mysql.Connection) {
	loadHeartDraw(db)
	loadHeartDrawAward(db)
}

type HeartDraw struct {
	DrawType  int8              // 抽奖类型（1-大转盘；2-刮刮卡）
	DailyNum  int8              // 每日可抽奖次数
	CostHeart int8              // 每次抽奖消耗爱心数
	Awards    []*HeartDrawAward // 奖品列表
}

func loadHeartDraw(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM heart_draw ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iDrawType := res.Map("draw_type")
	iDailyNum := res.Map("daily_num")
	iCostHeart := res.Map("cost_heart")

	var pri_id int16
	mapHeartDraw = map[int16]*HeartDraw{}
	for _, row := range res.Rows {
		pri_id = row.Int16(iId)
		mapHeartDraw[pri_id] = &HeartDraw{
			DrawType:  row.Int8(iDrawType),
			DailyNum:  row.Int8(iDailyNum),
			CostHeart: row.Int8(iCostHeart),
		}
	}
}

type HeartDrawAward struct {
	AwardType int8  // 奖品类型（1-铜钱；2-元宝；3-道具）
	AwardNum  int16 // 奖品数量
	ItemId    int16 // 道具奖品ID
	Chance    int8  // 抽奖概率%
}

func loadHeartDrawAward(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM heart_draw_award ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iHeartDrawId := res.Map("heart_draw_id")
	iAwardType := res.Map("award_type")
	iAwardNum := res.Map("award_num")
	iItemId := res.Map("item_id")
	iChance := res.Map("chance")

	var heartDrawId int16
	for _, row := range res.Rows {
		heartDrawId = row.Int16(iHeartDrawId)
		mapHeartDraw[heartDrawId].Awards = append(mapHeartDraw[heartDrawId].Awards, &HeartDrawAward{
			AwardType: row.Int8(iAwardType),
			AwardNum:  row.Int16(iAwardNum),
			ItemId:    row.Int16(iItemId),
			Chance:    row.Int8(iChance),
		})
	}
}

// ############### 对外接口实现 coding here ###############

func GetDrawInfo(drawType int8) (heartDraw *HeartDraw) {
	for _, heartDraw = range mapHeartDraw {
		if heartDraw.DrawType == drawType {
			return
		}
	}
	fail.When(heartDraw == nil, "not found heart draw")
	return
}
