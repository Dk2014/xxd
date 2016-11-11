package platform_server

import (
	"core/fail"
	"core/log"
	"core/redis"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"sort"
	"strconv"
	"sync"
	"time"
)

var (
	_server_list_lock   sync.RWMutex
	_server_list_hash   []byte
	_server_lists       = make(map[string]map[string]Server)
	_server_update_time = time.Now()
)

type GServer struct {
	GSId    int32
	Ip      string
	Port    string
	HD      bool
	RPCIp   string
	RPCPort string
}

type Server struct {
	Id        int32     // 游戏服ID
	Type      uint8     // 平台类型
	Name      string    // 名称
	Status    int8      // 状态[0-维护,1-通畅,2-繁忙,3-拥挤]
	StatusMsg string    // 状态信息
	IsNew     bool      // 新服标记[0-否,1-是]
	IsHot     bool      // 推荐服[0-否,1-是]
	OpenTime  int64     // 开服时间，时间戳，单位s
	GServers  []GServer // game_server 列表信息
}

type ServerType struct {
	AreaId uint8 // 大区（1微信，2手Q）
	PlatId uint8 // 平台（0ios，1安卓）
}

func (s *ServerType) GetType() uint8 {
	//广州测试服area不能和正式服一样，所以定义了98-》手q，99-》微信
	if s.AreaId == uint8(98) {
		s.AreaId = 2
	} else if s.AreaId == uint8(99) {
		s.AreaId = 1
	}
	return s.AreaId | (s.PlatId << 4)
}

func GetServerType(iType uint8) ServerType {
	return ServerType{AreaId: iType & 0x0f, PlatId: ((iType & 0xf0) >> 4)}
}

func MatchServerType(reqType uint8, svrType uint8, is_equal_audit bool) bool {
	//如果客户端版本等于审核版本，则只下发相应的审核的服
	if is_equal_audit {
		if (reqType == TYPE_MOBILE_IOS_WEIXIN && svrType == TYPE_MOBILE_AUDIT_WEIXIN) ||
			(reqType == TYPE_MOBILE_IOS_QQ && svrType == TYPE_MOBILE_AUDIT_QQ) ||
			(reqType == TYPE_MOBILE_IOS_GUEST && svrType == TYPE_MOBILE_AUDIT_GUEST) {
			return true
		} else {
			return false
		}
	}
	//适用于四个平台的测试服
	if svrType == TYPE_MOBILE_SANDBOX {
		return true
	}

	return reqType == svrType
}

type ReqServerList struct {
	ReqBase
}

func InitServerList() (err error) {
	err = updateServerList()
	if err != nil {
		return err
	}

	// start a ticker to check if there is any update about "gameserverlist" every 1 min.
	ticker := time.NewTicker(time.Second * 60)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("Recovered in %v: %v", r, string(debug.Stack()))
			}
		}()

		for _ = range ticker.C {
			err := updateServerList()
			if err != nil {
				log.Errorf("fetch gs list error inside ticker: %v", err)
			}
		}
	}()
	return
}

func updateServerList() (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Recovered in %v: %v", r, string(debug.Stack()))
		}
	}()

	dbmap := GetDBMap()

	_server_list_lock.Lock()
	defer _server_list_lock.Unlock()

	for app, _ := range dbmap {
		// fetch from redis by key "gameserverlist"
		c := GetDBConn(app)
		defer c.Close()

		v, err := redis.Values(c.Do("HGETALL", RedisKey_GameServerList(c)))
		fail.When(err != nil, "fail to get game server list")

		var list_map = make(map[string]Server)

		for len(v) > 0 {
			var k string
			var b []byte
			v, err = redis.Scan(v, &k, &b)
			fail.When(err != nil, fmt.Sprintf("fail to parse game server list: %v", v))

			var svr Server
			// validate serverlist is proper json format
			err = json.Unmarshal(b, &svr)

			fail.When(err != nil, fmt.Sprintf("fail to parse game server info: %v", string(b)))

			list_map[HashServerKey(svr.Id, svr.Type)] = svr
		}

		log.Infof("updating game server list: %s", app)

		_server_lists[app] = list_map
		_server_update_time = time.Now()
	}

	return
}

func ServerList(app string) map[string]Server {
	_server_list_lock.RLock()
	defer _server_list_lock.RUnlock()

	_server_list, ok := _server_lists[app]
	if ok {
		return _server_list
	}
	return nil
}

// iType/AreaId|PlatID<<4 uint8
// Sid/Partition uint32 // 小区ID
// PlatId uint8 // 平台（0ios，1安卓）
func GetGServerInfoByOpenIdSid(OpenId string, iType uint8, Sid int32, app string) (gs GServer, ok bool) {
	c := GetDBConn(app)
	defer c.Close()

	key_rolelist_by_openid := RedisKey_RoleListByOpenidType(OpenId, iType)

	b, err := redis.Bytes(c.Do("HGET", key_rolelist_by_openid, Sid))
	if err != nil {
		log.Errorf("user info not found")
		return gs, false
	}

	var role RoleInfo
	// validate serverlist is proper json format
	err = json.Unmarshal(b, &role)
	if err != nil {
		log.Errorf("fail to parse role info %v", err)
		return gs, false
	}

	return GetGServerInfo(Sid, iType, role.Gsid, app)
}

func GetGServerInfo(Sid int32, iType uint8, GSId int32, app string) (gs GServer, ok bool) {
	var s Server
	s, ok = GetServerInfo(Sid, iType, app)
	if !ok {
		return
	}

	for _, v := range s.GServers {
		if v.GSId == GSId {
			return v, true
		}
	}
	ok = false
	return
}

func GetServerInfo(Sid int32, iType uint8, app string) (s Server, ok bool) {
	_server_list_lock.RLock()
	defer _server_list_lock.RUnlock()

	_server_list, b := _server_lists[app]
	if !b {
		return Server{}, false
	}

	s, ok = _server_list[HashServerKey(Sid, iType)]

	if !ok {
		for _, v := range _server_list {
			if v.Id == Sid && (v.Type == TYPE_MOBILE_SANDBOX || v.Type == TYPE_MOBILE_AUDIT_WEIXIN || v.Type == TYPE_MOBILE_AUDIT_QQ || v.Type == TYPE_MOBILE_AUDIT_GUEST) {
				s = v
				ok = true
				break
			}
		}
	}
	return
}

func NextServerListRevsion(app string) int {
	c := GetDBConn(app)
	defer c.Close()

	v, err := redis.Int(c.Do("INCR", RedisKey_GameServerListRevCounter()))
	if err != nil {
		c.Do("SET", RedisKey_GameServerListRevCounter(), 1)
		return 1
	}
	return v
}

func HashServerKey(Sid int32, iType uint8) string {
	return fmt.Sprintf("%08v_%03v", Sid, iType)
}

type procServerList struct {
}

func (proc procServerList) Req() interface{} {
	return &ReqServerList{}
}

func (proc procServerList) Validate(in interface{}) {
	req := in.(*ReqServerList)
	fail.When(req.OpenId == "", "incorrect openid")
	ValidateMobileType(req.Type)
}

func (proc procServerList) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqServerList)
	if req.App == "" {
		req.App = "xxd_qq"
	}

	is_in_openid_white := isInOpenidWhiteTable(req.OpenId, req.App)
	is_in_version_black := isInVersionBlackTable(req.Version, req.Type, req.App)
	is_less_min_version := isLessThanMinVersion(req.Type, req.Version, req.App)

	var list struct {
		List        []Server
		NeedUpgrade string
	}

	//小于最小版本，则下发空的服务器列表
	if is_less_min_version || is_in_version_black {
		upgrade_url := getClientUpgradeUrl(req.Type, req.App)
		list.NeedUpgrade = "1"
		if upgrade_url != "" {
			list.NeedUpgrade = upgrade_url
		}
		return list, nil
	}

	_server_list_lock.RLock()
	defer _server_list_lock.RUnlock()

	_server_list, _ := _server_lists[req.App]

	var sorted_keys []string
	for k, _ := range _server_list {
		sorted_keys = append(sorted_keys, k)
	}

	sort.Strings(sorted_keys)

	for i, j := 0, len(sorted_keys)-1; i < j; i, j = i+1, j-1 {
		sorted_keys[i], sorted_keys[j] = sorted_keys[j], sorted_keys[i]
	}

	now := time.Now().Unix()

	isEqualAudit := isEqualAuditForApp(req.Type, req.Version, req.App)
	is_close_type := isCloseTypeServers(req.Type, req.App)

	//过滤符合要求的服务器
	for _, k := range sorted_keys {
		v := _server_list[k]
		if (is_in_openid_white || v.OpenTime <= now) && MatchServerType(req.Type, v.Type, isEqualAudit) {
			list.List = append(list.List, v)
		}
	}

	//如果开启了此平台的维护，给不在白名单的玩家下面维护的服务器列表
	if len(list.List) > 0 && is_close_type && !is_in_openid_white {
		for i, _ := range list.List {
			list.List[i].Status = 0
		}
	}
	return list, nil
}

type ReqGServerList struct {
	ReqBase
}

type procGServerList struct {
}

func (proc procGServerList) Req() interface{} {
	return &ReqGServerList{}
}

func (proc procGServerList) Validate(in interface{}) {

}

type RPCAddr struct {
	GSId    int32
	HD      bool
	RPCIp   string
	RPCPort string
	Type    uint8
}

func (proc procGServerList) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqGServerList)
	if req.App == "" {
		req.App = "xxd_qq"
	}

	_server_list_lock.RLock()
	defer _server_list_lock.RUnlock()

	_server_list, _ := _server_lists[req.App]

	gservers := make(map[string]RPCAddr)

	for _, server := range _server_list {
		_gservers := server.GServers
		for _, gserver := range _gservers {
			gsid := strconv.Itoa(int(gserver.GSId))
			gservers[gsid] = RPCAddr{
				GSId:    gserver.GSId,
				HD:      gserver.HD,
				RPCIp:   gserver.RPCIp,
				RPCPort: gserver.RPCPort,
				Type:    server.Type,
			}
		}
	}

	return gservers, nil
}

func isEqualAudit(itype uint8, client_version int32) bool {
	c := redisPool_.Get()
	defer c.Close()

	audit_version_key := RedisKey_ClientAuditVersionByType(itype)
	audit_version, err := redis.Int(c.Do("GET", audit_version_key))

	if err != nil || audit_version <= 0 {
		return false
	}
	return client_version == int32(audit_version)
}

func isEqualAuditForApp(itype uint8, client_version int32, app string) bool {
	c := GetDBConn(app)
	defer c.Close()

	audit_version_key := RedisKey_ClientAuditVersionByType(itype)
	audit_version, err := redis.Int(c.Do("GET", audit_version_key))

	if err != nil || audit_version <= 0 {
		return false
	}
	return client_version == int32(audit_version)
}

func isLessThanMinVersion(itype uint8, client_version int32, app string) bool {
	c := GetDBConn(app)
	defer c.Close()

	min_version_key := RedisKey_ClientMinVersionByType(itype)
	min_version, err := redis.Int(c.Do("GET", min_version_key))

	if err != nil || min_version <= 0 {
		return false
	}
	return client_version < int32(min_version)
}

func isInOpenidWhiteTable(openid string, app string) bool {
	c := GetDBConn(app)
	defer c.Close()

	openid_white_table_key := RedisKey_OpenidWhiteTable()
	isIn, err := redis.Bool(c.Do("SISMEMBER", openid_white_table_key, openid))

	if err != nil {
		log.Infof("run sismember error, openid_white_table_key:%v\n", openid_white_table_key)
		return false
	}

	if isIn {
		return true
	}

	return false
}

func isInVersionBlackTable(version int32, iType uint8, app string) bool {
	c := GetDBConn(app)
	defer c.Close()

	sVersion := strconv.Itoa(int(version))
	version_black_table_key := RedisKey_ClientVersionBlackTable(iType)
	isIn, err := redis.Bool(c.Do("SISMEMBER", version_black_table_key, sVersion))
	if err != nil {
		log.Infof("run sismember error, version_black_table_key:%v\n", version_black_table_key)
		return false
	}

	if isIn {
		return true
	}
	return false
}

func getClientUpgradeUrl(iType uint8, app string) string {
	c := GetDBConn(app)
	defer c.Close()

	upgrade_key := RedisKey_ClientUpgradeUrl(iType)
	url, err := redis.String(c.Do("get", upgrade_key))
	if err != nil {
		log.Infof("get upgradeUrl error, key is %v\n", upgrade_key)
		return ""
	}

	return url
}

func isCloseTypeServers(itype uint8, app string) bool {
	c := GetDBConn(app)
	defer c.Close()

	key := RedisKey_CloseTypeServers(itype)
	iClose, err := redis.Int(c.Do("GET", key))

	if err == nil && iClose == 1 {
		return true
	}
	return false
}
