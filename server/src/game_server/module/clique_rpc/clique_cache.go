package clique_rpc

import (
	"game_server/mdb"
	//"game_server/module"
	"core/time"
	"game_server/api/protocol/channel_api"
	"game_server/dat/clique_dat"
	"sort"
	"sync"
)

var (
	g_CliqueTable *cliqueTable = &cliqueTable{
		nameSet:   map[string]bool{},
		cliqueMap: map[int64]*CliqueInfo{},
		orderInfo: &orderCliqueList{},
	}
)

type cliqueTable struct {
	sync.RWMutex
	nameSet   map[string]bool //已占用的帮派名
	cliqueMap map[int64]*CliqueInfo
	orderInfo *orderCliqueList
}

type sortCliqueList []int64

func (lists sortCliqueList) Len() int {
	return len(lists)
}
func (lists sortCliqueList) Less(i, j int) bool {
	cliqueA := g_CliqueTable.cliqueMap[lists[i]]
	cliqueB := g_CliqueTable.cliqueMap[lists[j]]

	if cliqueA == nil {
		return false
	} else if cliqueB == nil {
		return true
	} else {
		if cliqueA.Level != cliqueB.Level {
			// 优先帮派等级
			return cliqueA.Level > cliqueB.Level
		} else {
			// 其次其他建筑物等级总和
			totalA := cliqueA.BankLevel + cliqueA.TempleLevel + cliqueA.JGLevel + cliqueA.HCLevel + cliqueA.SBLevel
			totalB := cliqueB.BankLevel + cliqueB.TempleLevel + cliqueB.JGLevel + cliqueB.HCLevel + cliqueB.SBLevel
			return totalA > totalB
		}
	}
}

func (lists sortCliqueList) Swap(i, j int) {
	lists[i], lists[j] = lists[j], lists[i]
}

type orderCliqueList struct {
	lists      sortCliqueList
	lastUpdate int64
}

func (table *cliqueTable) addCliqueName(row *mdb.GlobalCliqueRow) {
	table.Lock()
	defer table.Unlock()
	table.nameSet[row.Name()] = true
}

func (table *cliqueTable) deleteCliqueName(row *mdb.GlobalCliqueRow) {
	table.Lock()
	defer table.Unlock()
	delete(table.nameSet, row.Name())
}

func (table *cliqueTable) orderClique() {
	if time.GetNowTime()-table.orderInfo.lastUpdate > 300 /* 5分钟刷新间隔 */ {
		// 每隔一段时间执行排序
		sort.Sort(table.orderInfo.lists)
		table.orderInfo.lastUpdate = time.GetNowTime()
	}
}

func (table *cliqueTable) updateContrib(row *mdb.PlayerGlobalCliqueInfoRow) {
	if row.CliqueId() > 0 {
		table.Lock()
		defer table.Unlock()
		cliqueInfo := table.cliqueMap[row.CliqueId()]
		for idx := range cliqueInfo.Members {
			if cliqueInfo.Members[idx] != nil && cliqueInfo.Members[idx].Pid == row.Pid() {
				cliqueInfo.Members[idx].Contrib = row.Contrib()
				cliqueInfo.Members[idx].TotalContrib = row.TotalContrib()
			}
		}
	}
}

//TODO 缓存帮派成员信息 可能考虑用 SQL 加载
func (table *cliqueTable) addMember(row *mdb.PlayerGlobalCliqueInfoRow) {
	if row.CliqueId() > 0 {
		table.Lock()
		defer table.Unlock()
		if table.cliqueMap == nil {
			table.cliqueMap = make(map[int64]*CliqueInfo)
		}
		cliqueInfo, ok := table.cliqueMap[row.CliqueId()]
		if ok {
			for idx := range cliqueInfo.Members {
				if cliqueInfo.Members[idx] == nil {
					cliqueInfo.Members[idx] = &CliqueMember{
						Pid:          row.Pid(),
						Contrib:      row.Contrib(),
						TotalContrib: row.TotalContrib(),
					}
					cliqueInfo.MemberNum++
					break
				}
			}
		} else {
			cliqueInfo = &CliqueInfo{}
			cliqueInfo.MemberNum = 1
			cliqueInfo.Members[0] = &CliqueMember{
				Pid:          row.Pid(),
				Contrib:      row.Contrib(),
				TotalContrib: row.TotalContrib(),
			}
			table.cliqueMap[row.CliqueId()] = cliqueInfo
		}
	}
}

func (table *cliqueTable) deleteMember(cliqueId, pid int64) {
	table.Lock()
	defer table.Unlock()
	if cliqueInfo, ok := table.cliqueMap[cliqueId]; ok {
		for idx, _ := range cliqueInfo.Members {
			if cliqueInfo.Members[idx] != nil && cliqueInfo.Members[idx].Pid == pid {
				cliqueInfo.Members[idx] = nil
				cliqueInfo.MemberNum--
				break
			}
		}
	}
}

type CliqueInfo struct {
	Name              string
	Contrib           int64                //帮派贡献(不实时)
	AutoAudit         bool                 //开启自动审核
	AutoAuditMinLevel int16                //自动审核要求最低等级
	Level             int16                //帮派等级
	Members           [40]*CliqueMember    // 成员列表存储 pid
	MemberNum         int16                //当前帮派成员数量
	JoinApplies       map[int64]*JoinApply //加入请求
	Messages          []channel_api.CliqueMessage
	BankLevel         int16 //钱庄等级
	TempleLevel       int16 //宗祠等级
	HCLevel           int16 //回春堂等级
	JGLevel           int16 //金刚堂等级
	SBLevel           int16 //神兵堂等级
}

type CliqueMember struct {
	Pid int64
	//Nick    string //暂不支持 可能考虑用 sql 加载缓存信息
	Contrib      int64
	TotalContrib int64 // 玩家总共的帮派贡献
}

type JoinApply struct {
	Pid       int64
	Level     int16
	Timestamp int64
	Nick      string
	ArenaRank int32
}

//func (cache cliqueTable) addClique(cliqueId int64, name string) {
func (cache cliqueTable) addClique(row *mdb.GlobalCliqueRow) {
	g_CliqueTable.Lock()
	defer g_CliqueTable.Unlock()
	g_CliqueTable.nameSet[row.Name()] = true
	if clique, ok := g_CliqueTable.cliqueMap[row.Id()]; ok {
		clique.AutoAudit = row.AutoAudit() > 0
		clique.AutoAuditMinLevel = row.AutoAuditLevel()
		clique.Level = row.CenterBuildingLevel()
		clique.Contrib = row.Contrib()
		clique.Name = row.Name()
		clique.BankLevel = row.BankBuildingLevel()
		clique.TempleLevel = row.TempleBuildingLevel()
		clique.HCLevel = row.HealthBuildingLevel()
		clique.JGLevel = row.DefenseBuildingLevel()
		clique.SBLevel = row.AttackBuildingLevel()
	} else {
		g_CliqueTable.cliqueMap[row.Id()] = &CliqueInfo{
			JoinApplies:       map[int64]*JoinApply{},
			AutoAudit:         row.AutoAudit() > 0,
			AutoAuditMinLevel: row.AutoAuditLevel(),
			Level:             row.CenterBuildingLevel(),
			Name:              row.Name(),
			BankLevel:         row.BankBuildingLevel(),
			TempleLevel:       row.TempleBuildingLevel(),
			HCLevel:           row.HealthBuildingLevel(),
			JGLevel:           row.DefenseBuildingLevel(),
			SBLevel:           row.AttackBuildingLevel(),
			//Members 成员数在 addMember 时设置
		}
		g_CliqueTable.orderInfo.lists = append(g_CliqueTable.orderInfo.lists, row.Id())
	}
}

func (cache cliqueTable) AddApply(pid, cliqueId int64, apply *JoinApply) {
	g_CliqueTable.Lock()
	defer g_CliqueTable.Unlock()

	clique := g_CliqueTable.cliqueMap[cliqueId]
	//FIXME 需要统一处理初始化
	if clique.JoinApplies == nil {
		clique.JoinApplies = map[int64]*JoinApply{}
	}
	clique.JoinApplies[pid] = apply
}

func (cache cliqueTable) DeleteApply(pid, cliqueId int64) {
	g_CliqueTable.Lock()
	defer g_CliqueTable.Unlock()
	clique := g_CliqueTable.cliqueMap[cliqueId]
	if clique != nil {
		delete(clique.JoinApplies, pid)
	}
}

func (cache cliqueTable) AddMember(cliqueId, pid int64) {
	g_CliqueTable.Lock()
	defer g_CliqueTable.Unlock()

	clique, _ := g_CliqueTable.cliqueMap[cliqueId]
	for idx, _ := range clique.Members {
		if clique.Members[idx] == nil {
			clique.Members[idx] = &CliqueMember{
				Pid:          pid,
				Contrib:      0,
				TotalContrib: 0,
			}
			clique.MemberNum++
			return
		}
	}
	panic("成员已满")
}

func (cache cliqueTable) DeleteMember(cliqueId, pid int64) {
	g_CliqueTable.Lock()
	defer g_CliqueTable.Unlock()

	clique, _ := g_CliqueTable.cliqueMap[cliqueId]
	if clique == nil {
		return
	}
	for idx, _ := range clique.Members {
		if clique.Members[idx] != nil && clique.Members[idx].Pid == pid {
			clique.Members[idx] = nil
			clique.MemberNum--
		}
	}
}

func (cache cliqueTable) DeleteClique(cliqueId int64) {
	g_CliqueTable.Lock()
	defer g_CliqueTable.Unlock()
	delete(g_CliqueTable.cliqueMap, cliqueId)
	for index, clique_id := range g_CliqueTable.orderInfo.lists {
		if clique_id == cliqueId {
			g_CliqueTable.orderInfo.lists = append(g_CliqueTable.orderInfo.lists[:index], (g_CliqueTable.orderInfo.lists[index+1:])...)
		}
	}
}

func CacheGetCliqueInfo(cliqueId int64) *CliqueInfo {
	g_CliqueTable.RLock()
	defer g_CliqueTable.RUnlock()
	return g_CliqueTable.cliqueMap[cliqueId]
}

func CacheCheckCliqueByName(name string) (exist bool) {
	g_CliqueTable.RLock()
	defer g_CliqueTable.RUnlock()
	_, exist = g_CliqueTable.nameSet[name]
	return exist
}

func CacheAddMessage(cliqueId int64, msg channel_api.CliqueMessage) {
	g_CliqueTable.Lock()
	defer g_CliqueTable.Unlock()
	if clique, ok := g_CliqueTable.cliqueMap[cliqueId]; ok {
		clique.Messages = append(clique.Messages, msg)
		if len(clique.Messages) > clique_dat.MAX_RESORE_MESSAGE_NUM {
			clique.Messages = clique.Messages[len(clique.Messages)-clique_dat.MAX_RESORE_MESSAGE_NUM:]
		}
	}
}

func CacheGetLatestMessage(cliqueId int64) []channel_api.CliqueMessage {
	g_CliqueTable.RLock()
	defer g_CliqueTable.RUnlock()
	if clique, ok := g_CliqueTable.cliqueMap[cliqueId]; ok {
		return clique.Messages
	}
	return nil
}

func DeleteTimeOutApply(cliqueId int64) {
	g_CliqueTable.Lock()
	defer g_CliqueTable.Unlock()

	if clique, ok := g_CliqueTable.cliqueMap[cliqueId]; ok {
		for pid, join := range clique.JoinApplies {
			if time.GetNowTime()-join.Timestamp >= clique_dat.CLIQUE_TIME_OUT_APPLY {
				delete(clique.JoinApplies, pid)
				DeleteApply(pid, cliqueId)
			}
		}
	}
}

func FetchClique(cb func(int64, *CliqueInfo) bool) {
	g_CliqueTable.Lock()
	defer g_CliqueTable.Unlock()
	for cliqueId, clique := range g_CliqueTable.cliqueMap {
		if !cb(cliqueId, clique) {
			break
		}
	}
}

func GetCliquePid(cliuqeid int64) []int64 {
	g_CliqueTable.Lock()
	defer g_CliqueTable.Unlock()
	pidList := make([]int64, 0)
	if clique, ok := g_CliqueTable.cliqueMap[cliuqeid]; ok {
		for _, member := range clique.Members {
			if member != nil {
				pidList = append(pidList, member.Pid)
			}
		}
		return pidList
	}
	return nil
}

func CacheIsCliqueMember(cliqueId, pid int64) bool {
	g_CliqueTable.RLock()
	defer g_CliqueTable.RUnlock()
	if clique, ok := g_CliqueTable.cliqueMap[cliqueId]; ok {
		for _, member := range clique.Members {
			if member != nil && member.Pid == pid {
				return true
			}
		}
	}
	return false
}

func CacheGetMembeNum(cliqueId int64) int16 {
	g_CliqueTable.RLock()
	defer g_CliqueTable.RUnlock()
	if clique, ok := g_CliqueTable.cliqueMap[cliqueId]; ok {
		return clique.MemberNum
	}
	return 0
}

func CacheGetPageCliqueInfo(page int) (int, int, []int64) {
	g_CliqueTable.Lock()
	defer g_CliqueTable.Unlock()

	if g_CliqueTable.orderInfo.lastUpdate == 0 || page == 1 {
		// 初次或者新请求列表时排序
		g_CliqueTable.orderClique()
	}

	total := len(g_CliqueTable.orderInfo.lists)
	pageNum := total / clique_dat.CLIQUE_LIST_PAGE_SIZE
	if len(g_CliqueTable.orderInfo.lists)%clique_dat.CLIQUE_LIST_PAGE_SIZE != 0 {
		pageNum += 1
	}

	// 页数容错处理
	if page <= 0 {
		page = 1
	}
	if page > pageNum {
		page = pageNum
	}

	start := (page - 1) * clique_dat.CLIQUE_LIST_PAGE_SIZE
	end := page * clique_dat.CLIQUE_LIST_PAGE_SIZE
	if end > total {
		end = total
	}
	// 返回结果3个依次为当前页开始索引，总共的记录数，当前列表数据
	return start, total, g_CliqueTable.orderInfo.lists[start:end]
}
