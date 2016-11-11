package platform_server

import (
	"core/fail"
	"core/log"
	"core/redis"
	"encoding/json"
	"errors"
	"fmt"
	"runtime/debug"
	"sort"
	"strings"
	"time"
)

const (
	APP_WEGAME = "xxd_tw"
)

type RoleList []RoleInfo

func (v RoleList) Len() int           { return len(v) }
func (v RoleList) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v RoleList) Less(i, j int) bool { return v[i].LoginTime > v[j].LoginTime }

func (proc procUserList) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqRoleList)

	isEqualAudit := isEqualAudit(req.Type, req.Version)
	var roles RoleList
	var sType uint8
	if isEqualAudit {
		switch req.Type {
		case 1:
			sType = TYPE_MOBILE_AUDIT_WEIXIN
		case 2:
			sType = TYPE_MOBILE_AUDIT_QQ
		case 5:
			sType = TYPE_MOBILE_AUDIT_GUEST
		}

		roles = getRoleListByOpenIdType(req.App, req.OpenId, sType)
	} else {
		roles = getRoleListByOpenIdType(req.App, req.OpenId, req.Type)
		sandboxRoles := getRoleListByOpenIdType(req.App, req.OpenId, TYPE_MOBILE_SANDBOX)

		if len(sandboxRoles) > 0 {
			for _, v := range sandboxRoles {
				roles = append(roles, v)
			}
		}
	}

	sort.Sort(roles)
	rsp := &RspRoleList{List: roles}

	return rsp, nil
}

type RoleInfo struct {
	Nick      string // 游戏角色名
	RoleId    int8   // 角色模板ID
	LoginTime int64  // 最近一次登录时间
	Hash      []byte
	Gsid      int32  `json:",omitempty"` // Game Server ID
	IP        string `json:",omitempty"` // Game Server IP
	Port      string `json:",omitempty"` // Game Server Port
	RoleLevel int16  // 角色等级
	Sid       int32  // 游戏逻辑服ID
	GuestId   string
}

// "/user/gsinfo"
type ReqUserGSInfo struct {
	ReqBase
	Sid int32
}

type procUserGSInfo struct {
}

func (proc procUserGSInfo) Req() interface{} {
	return &ReqUserGSInfo{}
}

func (proc procUserGSInfo) Validate(in interface{}) {
	req := in.(*ReqUserGSInfo)

	fail.When(req.OpenId == "", "incorrect openid")
	fail.When(req.Sid <= 0, "incorrect sid")

	ValidateMobileType(req.Type)
}

func (proc procUserGSInfo) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqUserGSInfo)

	// 验证提交的服务器ID和平台类型是否正常
	s, ok := GetServerInfo(req.Sid, req.Type, req.App)
	fail.When(!ok, fmt.Sprintf("incorrect sid:%v or type:%v", req.Sid, req.Type))

	c := GetDBConn(req.App)
	defer c.Close()

	key_rolelist_by_openid := RedisKey_RoleListByOpenidType(req.OpenId, s.Type)

	// 验证openid是否已有角色
	b, err := redis.Bytes(c.Do("HGET", key_rolelist_by_openid, req.Sid))
	if err == nil {
		roleinfo := &RoleInfo{}
		err = json.Unmarshal(b, roleinfo)
		if err == nil {
			gs, ok := GetGServerInfo(req.Sid, req.Type, roleinfo.Gsid, req.App)
			fail.When(!ok, "game server not found for user")

			// 已有角色
			return &map[string]interface{}{"Gsid": gs.GSId, "IP": gs.Ip, "Port": gs.Port}, nil
		}
		log.Errorf("User roleinfo unmarshal error: %v of %v @ %v", err, req.OpenId, req.Sid)
	}

	elected_idx, elected_ksgc := electServer(s, c)

	if elected_idx < 0 {
		return nil, errors.New("game server election failed!")
	}

	rsp := s.GServers[elected_idx]
	log.Infof("send user %v of %v to %v, %v:%v", req.OpenId, req.Sid, rsp.GSId, rsp.Ip, rsp.Port)

	// update RoleInfo of OpenidSid - redisKey_RoleInfoByOpenidSid
	b, err = json.Marshal(rsp)
	if err != nil {
		log.Errorf("rsp marshal error: %v", err)
		return rsp, err
	}

	c.Send("HSET", key_rolelist_by_openid, req.Sid, b)
	c.Send("INCR", elected_ksgc)
	c.Flush()

	//计算真实的type，因为越南android的type被当作ios处理了
	rType := (&ServerType{AreaId: req.AreaId, PlatId: req.PlatId}).GetType()
	go logFirstCreateUser(req.App, key_rolelist_by_openid, req.Sid, rType, req.OpenId, req.ClientIp)
	return &map[string]interface{}{"Gsid": rsp.GSId, "IP": rsp.Ip, "Port": rsp.Port}, nil
}

func logFirstCreateUser(app, hkey string, sid int32, iType uint8, openid, ip string) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Recovered in %v: %v", r, string(debug.Stack()))
		}
	}()

	c := GetDBConn(app)
	defer c.Close()

	len, _ := redis.Int(c.Do("HLEN", hkey))
	if len == 1 {
		cid := getCidOfApp(app, iType)
		logObj := UserCreateLog{
			Type:    "account",
			Cid:     cid,
			Sid:     sid,
			Account: openid,
			Time:    time.Now().Unix(),
			Ip:      ip,
		}
		LogAccountCreate(&logObj)
	}
}

type ReqUserCreate struct {
	ReqBase
	Sid    int32
	RoleId int8
	Nick   string
}

type procUserCreate struct {
}

func (proc procUserCreate) Req() interface{} {
	return &ReqUserCreate{}
}

func (proc procUserCreate) Validate(in interface{}) {
	req := in.(*ReqUserCreate)

	fail.When(req.OpenId == "", "incorrect openid")
	fail.When(req.Nick == "", "incorrect nick")

	ValidateRoleId(req.RoleId)
}

func (proc procUserCreate) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqUserCreate)

	switch {
	case len(req.Nick) < 3:
		return &map[string]int{"NickTooShort": 1, "error": 1325}, errors.New("nick name too short: " + req.Nick)
	case (req.App != "xxd_vn" && len([]rune(req.Nick)) > 8) || (req.App == "xxd_vn" && len([]rune(req.Nick)) > 16):
		return &map[string]int{"NickTooLong": 1, "error": 413}, errors.New("nick name too long: " + req.Nick)
	case ValidateNickname(req.Nick) != nil:
		return &map[string]int{"NickIllegal": 1, "error": 2925}, errors.New("nick name illegal: " + req.Nick)
	}

	if req.App == "" || req.App == "xxd_qq" {
		out, err := CheckToken(req.Type, req.OpenId, req.Token, req.ClientIp)
		if err != nil || (*out)["error"] != float64(0) {
			return &map[string]int{"TokenUnavailable": 1, "error": 405}, errors.New("token is not avaiable")
		}
	}

	// 验证提交的服务器ID和平台类型是否正常
	s, ok := GetServerInfo(req.Sid, req.Type, req.App)
	fail.When(!ok, fmt.Sprintf("incorrect sid:%v or type:%v", req.Sid, req.Type))

	c := GetDBConn(req.App)
	defer c.Close()

	key_nick_exist := RedisKey_ExistenceBySidTypeNick(s.Id, s.Type, req.Nick)

	// 验证昵称是否重复
	exists, _ := redis.Bool(c.Do("EXISTS", key_nick_exist))
	if exists {
		return &map[string]int{"NickExist": 1, "error": 304}, errors.New("nick already exist: " + req.Nick)
	}

	key_rolelist_by_openid := RedisKey_RoleListByOpenidType(req.OpenId, s.Type)
	rsp := &RoleInfo{}

	// 验证openid是否已有角色
	b, err := redis.Bytes(c.Do("HGET", key_rolelist_by_openid, req.Sid))
	if err == nil {
		err = json.Unmarshal(b, rsp)
		if err == nil {

			gs, ok := GetGServerInfo(req.Sid, req.Type, rsp.Gsid, req.App)
			fail.When(!ok, "game server not found for user")

			// 已有角色
			rsp.IP = gs.Ip
			rsp.Port = gs.Port
			if rsp.Sid <= 0 {
				rsp.Sid = s.Id
			}

			rsp.Nick = req.Nick

			if req.RoleId > 0 {
				rsp.RoleId = req.RoleId
			}

			if rsp.RoleLevel <= 0 {
				rsp.RoleLevel = 1
			}

			rsp.LoginTime = time.Now().Unix()
			//计算真实的type，因为越南android的type被当作ios处理了
			rType := (&ServerType{AreaId: req.AreaId, PlatId: req.PlatId}).GetType()
			rsp.Hash = hashResponse(rType, rsp.RoleId, req.OpenId, rsp.Nick, rsp.LoginTime)

			// update RoleInfo of OpenidSid - redisKey_RoleInfoByOpenidSid
			b, err = json.Marshal(rsp)
			if err != nil {
				log.Errorf("rsp marshal error: %v", err)
				return rsp, err
			}

			c.Do("HSET", key_rolelist_by_openid, req.Sid, b)
			// update Existence by Sid Nick
			c.Send("SET", key_nick_exist, 1)
			c.Flush()

			//针对wegame渠道，传递角色信息给wegame平台
			if req.App == APP_WEGAME {
				if !strings.HasPrefix(req.OpenId, "guest_") {
					roleList := make(RoleList, 0)
					roleList = append(roleList, *rsp)
					err := setWGRole(req.App, req.OpenId, req.ClientIp, getWGSDKUrl(req.Token), roleList)
					if err != nil {
						log.Error(err.Error())
					}
				}
			}

			return rsp, nil
		}

		log.Errorf("User roleinfo unmarshal error: %v of %v @ %v", err, req.OpenId, req.Sid)
	} else {
		panic(fmt.Sprintf("user roleinfo not found, openid:%v, itype:%v, sid:%v", req.OpenId, req.Type, req.Sid))
	}

	return rsp, nil
}

func electServer(s Server, c redis.Conn) (int, string) {
	// 直接选取最少用户的game_server
	elected_idx := -1
	elected_ksgc := ""
	last_count := ^uint64(0)
	for i, gs := range s.GServers {
		if gs.HD {
			continue
		}
		ksgc := RedisKey_CounterBySidGsid(s.Id, s.Type, gs.GSId)
		gs_total_user_count, _ := redis.Uint64(c.Do("GET", ksgc))
		// ignore err since the key-value may not exist
		if gs_total_user_count < last_count {
			last_count = gs_total_user_count
			elected_ksgc = ksgc
			elected_idx = i
		}
	}
	return elected_idx, elected_ksgc
}

// /user/logon
type ReqUserLogon struct {
	ReqBase
	Sid    int32
	Nick   string // 游戏角色名
	RoleId int8   // 角色模板ID
}

type procUserLogon struct {
}

func (proc procUserLogon) Req() interface{} {
	return &ReqUserLogon{}
}

func (proc procUserLogon) Validate(in interface{}) {
	req := in.(*ReqUserLogon)

	fail.When(req.OpenId == "", "incorrect openid")
	fail.When(req.Sid <= 0, "incorrect sid")

	ValidateMobileType(req.Type)
}

func (proc procUserLogon) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqUserLogon)

	if req.App == "xxd_qq" {
		out, err := CheckToken(req.Type, req.OpenId, req.Token, req.ClientIp)
		if err != nil || (*out)["error"] != float64(0) {
			return &map[string]int{"TokenUnavailable": 1, "error": 405}, errors.New("token is not avaiable")
		}
	}

	rsp := &RoleInfo{}

	// 验证提交的服务器ID和平台类型是否正常
	s, ok := GetServerInfo(req.Sid, req.Type, req.App)
	fail.When(!ok, fmt.Sprintf("incorrect sid:%v or type:%v", req.Sid, req.Type))

	c := GetDBConn(req.App)
	defer c.Close()

	not_found := &map[string]int{"NeedCreate": 1, "error": 404}

	key_rolelist_by_openid := RedisKey_RoleListByOpenidType(req.OpenId, s.Type)

	b, err := redis.Bytes(c.Do("HGET", key_rolelist_by_openid, req.Sid))
	if err != nil {
		if strings.Contains(err.Error(), "redigo: nil returned") {
			log.Errorf("user roleinfo not found, openid:%v, itype:%v, sid:%v\n", req.OpenId, req.Type, req.Sid)
			return not_found, err
		}
		panic(err)
	}

	err = json.Unmarshal(b, rsp)
	if err != nil {
		log.Errorf("User roleinfo unmarshal error: %v", err)
		return not_found, err
	}

	if rsp.Nick == "" {
		log.Errorf("user nick not found, openid:%v, itype:%v, sid:%v\n", req.OpenId, req.Type, req.Sid)
		return not_found, err
	}

	gs, ok := GetGServerInfo(req.Sid, req.Type, rsp.Gsid, req.App)
	fail.When(!ok, fmt.Sprintf("incorrect sid:%v or type:%v or Gsid:%v", req.Sid, req.Type, rsp.Gsid))

	// 需要更新部分字段
	rsp.IP = gs.Ip
	rsp.Port = gs.Port
	rsp.LoginTime = time.Now().Unix()
	//计算真实的type，因为越南android的type被当作ios处理了
	rType := (&ServerType{AreaId: req.AreaId, PlatId: req.PlatId}).GetType()
	rsp.Hash = hashResponse(rType, rsp.RoleId, req.OpenId, rsp.Nick, rsp.LoginTime)

	// update RoleInfoByOpenidSid - redisKey_RoleInfoByOpenidSid
	b, err = json.Marshal(rsp)
	if err != nil {
		log.Errorf("rsp marshal error: %v", err)
	} else {
		c.Do("HSET", key_rolelist_by_openid, req.Sid, b)
	}

	return rsp, nil
}

type ReqRoleList struct {
	ReqBase
}

type RspRoleList struct {
	List []RoleInfo
}

type procUserList struct {
}

func (proc procUserList) Req() interface{} {
	return &ReqRoleList{}
}

func (proc procUserList) Validate(in interface{}) {
	req := in.(*ReqRoleList)

	fail.When(req.OpenId == "", "incorrect openid")

	ValidateMobileType(req.Type)
}

func getRoleListByOpenIdType(app string, openid string, iType uint8) RoleList {
	c := GetDBConn(app)
	defer c.Close()

	roleList := make(RoleList, 0)

	key_rolelist_by_openid := RedisKey_RoleListByOpenidType(openid, iType)
	v, err := redis.Values(c.Do("HGETALL", key_rolelist_by_openid))
	if err != nil {
		log.Errorf("user roleinfo not found, openid:%v, itype:%v\n", openid, iType)
		return roleList
	}

	for len(v) > 0 {
		var sid string
		var b []byte
		v, err = redis.Scan(v, &sid, &b)
		if err != nil {
			log.Errorf("fail to scan role info %v", err)
			continue
		}

		var u RoleInfo
		// validate serverlist is proper json format
		err = json.Unmarshal(b, &u)
		if err != nil {
			log.Errorf("fail to parse role info %v", err)
			continue
		}

		roleList = append(roleList, u)
	}

	return roleList
}

type ReqUserUpdate struct {
	ReqBase
	Nick      string // 游戏角色名
	Sid       int32
	RoleLevel int16
}

type procUserUpdate struct {
}

func (proc procUserUpdate) Req() interface{} {
	return &ReqUserUpdate{}
}

func (proc procUserUpdate) Validate(in interface{}) {
	req := in.(*ReqUserUpdate)

	fail.When(req.OpenId == "", "incorrect openid")
	ValidateMobileType(req.Type)
	fail.When(req.RoleLevel < 1, "incorrect role level")
	fail.When(req.Sid < 1, "incorrect server id")
}

func (proc procUserUpdate) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqUserUpdate)

	if req.App == "xxd_qq" {
		out, err := CheckToken(req.Type, req.OpenId, req.Token, req.ClientIp)
		if err != nil || (*out)["error"] != float64(0) {
			return &map[string]int{"TokenUnavailable": 1, "error": 405}, errors.New("token is not avaiable")
		}
	}

	// 验证提交的服务器ID和平台类型是否正常
	s, ok := GetServerInfo(req.Sid, req.Type, req.App)
	fail.When(!ok, fmt.Sprintf("incorrect sid:%v or type:%v", req.Sid, req.Type))

	c := GetDBConn(req.App)
	defer c.Close()

	key_rolelist_by_openid := RedisKey_RoleListByOpenidType(req.OpenId, s.Type)

	b, err := redis.Bytes(c.Do("HGET", key_rolelist_by_openid, req.Sid))
	if err != nil {
		log.Errorf("user roleinfo not found, openid:%v, itype:%v, sid:%v\n", req.OpenId, req.Type, req.Sid)
		return nil, err
	}

	var role RoleInfo
	err = json.Unmarshal(b, &role)
	if err != nil {
		log.Errorf("role info unmarshal error: %v", err)
		return nil, err
	}

	// update rolelevel by openid sid
	if role.Sid == req.Sid {
		role.RoleLevel = req.RoleLevel

		if len(req.Nick) > 1 {
			role.Nick = req.Nick
			// update nick exist
			key_nick_exist := RedisKey_ExistenceBySidTypeNick(s.Id, s.Type, req.Nick)
			c.Send("SET", key_nick_exist, 1)
			c.Flush()
		}

		b, err := json.Marshal(role)
		if err != nil {
			log.Errorf("role info marshal error: %v", err)
		} else {
			c.Do("HSET", key_rolelist_by_openid, req.Sid, b)
		}
	}

	return role, nil
}
