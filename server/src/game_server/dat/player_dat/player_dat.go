package player_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapFunc               map[int32]*Func
	g_LockFuncs           []*Func //根据j权值开的功能
	g_LevelFuncs          []*Func //根据等级开放的功能
	mapQuestAwardFuncLock []*QuestFuncLock
)

func Load(db *mysql.Connection) {
	loadFunc(db)
	loadFameSystem(db)
	loadFameLevel(db)
	loadQuestFuncLock(db)
}

type QuestFuncLock struct {
	Lock  int16 //权值
	Order int32 //任务顺序
}

func loadQuestFuncLock(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM quest WHERE award_func_key > 0 ORDER BY `order` ASC"), -1)
	if err != nil {
		panic(err)
	}
	iLock := res.Map("award_func_key")
	iOrder := res.Map("order")

	mapQuestAwardFuncLock = []*QuestFuncLock{}
	for _, row := range res.Rows {
		mapQuestAwardFuncLock = append(mapQuestAwardFuncLock, &QuestFuncLock{
			Lock:  row.Int16(iLock),
			Order: row.Int32(iOrder),
		})
	}
}

type Func struct {
	Id        int32  //功能id
	Name      string //功能权值名称
	Type      int8   //类型
	Lock      int16  // 功能权值
	Level     int16  // 开放等级
	UniqueKey int64  // 唯一权值
	NeedPlay  int8   // 是否需要播放 0不需要 1需要
}

func loadFunc(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM func  WHERE `lock` < 30000  ORDER BY `lock`, `level`  ASC"), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iName := res.Map("name")
	iType := res.Map("type")
	iLock := res.Map("lock")
	iLevel := res.Map("level")
	iUniqueKey := res.Map("unique_key")
	iNeedPlay := res.Map("need_play")

	mapFunc = map[int32]*Func{}
	g_LockFuncs = make([]*Func, 0)
	g_LevelFuncs = make([]*Func, 0)
	for _, row := range res.Rows {
		typ := row.Int8(iType)
		id := row.Int32(iId)

		funcByLock := &Func{
			Id:        row.Int32(iId),
			Name:      row.Str(iName),
			Type:      typ,
			Level:     row.Int16(iLevel),
			Lock:      row.Int16(iLock),
			UniqueKey: row.Int64(iUniqueKey),
			NeedPlay:  row.Int8(iNeedPlay),
		}

		mapFunc[id] = funcByLock
		switch typ {
		case FUNC_OPEN_TYPE_BY_LOCK:
			g_LockFuncs = append(g_LockFuncs, funcByLock)
		case FUNC_OPEN_TYPE_BY_LEVEL:
			g_LevelFuncs = append(g_LevelFuncs, funcByLock)
		default:
			panic("undefine function open type")
		}
	}

}

// ############### 对外接口实现 coding here ###############

//获取所有按照权值开放的功能
func GetFuncMap() []*Func {
	return g_LockFuncs
}

func GetFuncById(id int32) *Func {
	function := mapFunc[id]
	fail.When(function == nil, "[GetFunById] can't found Func")
	return function
}

func GetSumUniqueKeyByLock(lock int16) (sum int64) {
	ok := false
	for _, funcByLock := range g_LockFuncs {
		sum += funcByLock.UniqueKey
		if funcByLock.Lock == lock {
			ok = true
			break
		}
	}
	fail.When(!ok, "GetSumUniqueKeyByLock wrong lock")
	return sum
}

func GetLevelFuncs(level int16) (funcs []*Func) {
	for _, function := range g_LevelFuncs {
		if level >= function.Level {
			funcs = append(funcs, function)
		} else {
			break
		}
	}
	return
}

func GetLockFuncs(funcKey int16) (funcs []*Func) {
	for _, function := range g_LockFuncs {
		if funcKey >= function.Lock {
			funcs = append(funcs, function)
		} else {
			break
		}
	}
	return funcs
}

// questOrder 是玩家当前任务的order，找出玩家当前任务之前最大的功能权值
// 找的过程中不包括玩家当前任务order的原因是：玩家的当前任务一定是未奖励或者为完成的。所以无需对当前任务存在的功能权值进行处理
func GetMaxLockByQuestOrder(questOrder int32) (maxLock int16) {
	for _, quest := range mapQuestAwardFuncLock {
		if quest.Order < questOrder {
			maxLock = quest.Lock
		} else {
			break
		}
	}
	return
}
