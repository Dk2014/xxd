package sword_soul_dat

import (
	"core/fail"
	"core/mysql"
)

type SwordSoul struct {
	Id          int16               // 剑心ID
	Name        string              //剑心名
	TypeId      int8                // 类型ID
	Quality     int8                // 品质
	FragmentNum int16               // 碎片数量
	FragmentId  int16               // 碎片ID
	LevelValue  []int32             // 等级加成
	Exchange    map[int16]*Material // 兑换数据
}

type SwordSoulQuality struct {
	Id           int8    // 剑心品质ID
	LevelFullExp []int32 // 该等级满了的经验 包括初始经验
}

// 兑换材料
type Material struct {
	Id    int16
	Level int8
	Num   int8
}

var (
	mapSwordSoul               map[int16]*SwordSoul       // 所有剑心
	mapSwordSoulQuality        map[int8]*SwordSoulQuality // 品质
	mapDrawSwordSoulsByQuality map[int8][]int16           // 可以拔到的剑心列表
)

func Load(db *mysql.Connection) {
	loadSwordSoul(db)
	loadSwordSoulLevel(db)
	loadSwordSoulQuality(db)
	loadSwordSoulQualityLevel(db)
}

func loadSwordSoul(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM sword_soul ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	var (
		iId           = res.Map("id")
		iName         = res.Map("name")
		iTypeId       = res.Map("type_id")
		iQuality      = res.Map("quality")
		iFragmentNum  = res.Map("fragment_num")
		iFragmentId   = res.Map("fragment_id")
		iOnlyExchange = res.Map("only_exchange")
		pri_id        int16
		quality       int8
	)

	mapSwordSoul = map[int16]*SwordSoul{}
	mapDrawSwordSoulsByQuality = map[int8][]int16{}

	for _, row := range res.Rows {

		pri_id = row.Int16(iId)
		quality = row.Int8(iQuality)

		mapSwordSoul[pri_id] = &SwordSoul{
			Id:          pri_id,
			Name:        row.Str(iName),
			TypeId:      row.Int8(iTypeId),
			Quality:     quality,
			FragmentNum: row.Int16(iFragmentNum),
			FragmentId:  row.Int16(iFragmentId),
		}

		if !row.Bool(iOnlyExchange) {
			if mapDrawSwordSoulsByQuality[quality] == nil {
				mapDrawSwordSoulsByQuality[quality] = []int16{}
			}
			mapDrawSwordSoulsByQuality[quality] = append(mapDrawSwordSoulsByQuality[quality], pri_id)
		}
	}
}

func loadSwordSoulLevel(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM sword_soul_level ORDER BY `sword_soul_id`, `level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	var (
		iSwordSoulId       = res.Map("sword_soul_id")
		iLevel             = res.Map("level")
		iValue             = res.Map("value")
		li           int8  = 1
		lastId       int16 = -1
	)

	for _, row := range res.Rows {
		var (
			id        = row.Int16(iSwordSoulId)
			swordSoul = mapSwordSoul[id]
		)

		// 判断是否有没有填写的剑心等级
		if lastId != id {
			lastId = id
			li = 1
		} else {
			li++
		}
		fail.When(li != row.Int8(iLevel), "sword soul level mismatch")

		// 每个等级的附加属性值
		swordSoul.LevelValue = append(swordSoul.LevelValue, row.Int32(iValue))
	}
}

func loadSwordSoulQuality(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM sword_soul_quality ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")

	var pri_id int8
	mapSwordSoulQuality = map[int8]*SwordSoulQuality{}
	for _, row := range res.Rows {
		pri_id = row.Int8(iId)

		mapSwordSoulQuality[pri_id] = &SwordSoulQuality{
			Id: pri_id,
		}
	}
}

func loadSwordSoulQualityLevel(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM sword_soul_quality_level ORDER BY `quality_id`, `level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	var (
		iLevel              = res.Map("level")
		iQuality            = res.Map("quality_id")
		iExp                = res.Map("exp")
		li            int8  = 1
		lastQualityId int8  = -1
		fullExp       int32 = 0 // 前面等级的经验总和
	)

	for _, row := range res.Rows {
		var (
			exp       = row.Int32(iExp)
			qualityId = row.Int8(iQuality)
			quality   = mapSwordSoulQuality[qualityId]
		)

		// 判断是否有没有填写的剑心品质等级
		if lastQualityId != qualityId {
			lastQualityId = qualityId
			li = 1
			fullExp = 0
		} else {
			li++
		}
		fail.When(li != row.Int8(iLevel), "sword soul qualit level mismatch")

		// 每个等级品质的升级经验，包括当前等级
		fullExp += exp
		quality.LevelFullExp = append(quality.LevelFullExp, fullExp)
	}
}

// ############### 对外接口实现 coding here ###############

func GetSwordSoul(id int16) *SwordSoul {
	swordSoul, exist := mapSwordSoul[id]
	fail.When(!exist, "sword soul id is not exist")
	return swordSoul
}

func GetExchangeMaterial(id int16) (materials map[int16]*Material) {
	swordSoul, exist := mapSwordSoul[id]
	fail.When(!exist, "sword soul id is not exist")
	materials = swordSoul.Exchange
	fail.When(materials == nil, "sword soul materials is not exist")
	return materials
}

func GetQuality(id int8) *SwordSoulQuality {
	return mapSwordSoulQuality[id]
}

// 根据拔剑品质获取剑心列表
func GetDrawSwordSoulsByQuality(quality int8) (ids []int16) {
	ids = mapDrawSwordSoulsByQuality[quality]
	fail.When(len(ids) == 0, "this quality SwordSouls are not exist")
	return ids
}
