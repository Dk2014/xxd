package platform_server

import (
	"core/log"
	"core/redis"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

var ip_black_tables = make(map[string][]string, 0)

const (
	XDGM_REQUEST_SECRET = "xxd@2015"
)

// "/patch"
type ReqClientPatch struct {
	ReqBase
	ServerVersion int32
}

type procClientPatch struct {
}

func (proc procClientPatch) Req() interface{} {
	return &ReqClientPatch{}
}

func (proc procClientPatch) Validate(in interface{}) {
	req := in.(*ReqClientPatch)

	ValidateMobileType(req.Type)
}

func (proc procClientPatch) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqClientPatch)

	c := GetDBConn(req.App)
	defer c.Close()

	key_patch := RedisKey_ClientPatchHash(req.Type)
	field := strconv.Itoa(int(req.Version)) + "_" + strconv.Itoa(int(req.ServerVersion))

	var patches = make([]string, 0)
	patch_file_str, err := redis.String(c.Do("HGET", key_patch, field))
	if err != nil || strings.EqualFold(strings.TrimSpace(patch_file_str), "") {
		return &map[string]interface{}{"patch": patches, "error": 0}, nil
	}

	patch_files := strings.Split(patch_file_str, ":")
	patch_url_prefix, err := redis.String(c.Do("GET", RedisKey_ClientPatchUrl()))
	if err != nil {
		return &map[string]interface{}{"patch": patches, "error": 0}, nil
	}

	for _, patch := range patch_files {
		patch_url := patch_url_prefix + patch
		patches = append(patches, patch_url)
	}

	return &map[string]interface{}{"patch": patches, "error": 0}, nil
}

// "/TotalResource"
type ReqTotalResource struct {
	ReqBase
	MaxTown int32
}

type procTotalResource struct {
}

func (proc procTotalResource) Req() interface{} {
	return &ReqTotalResource{}
}

func (proc procTotalResource) Validate(in interface{}) {
	req := in.(*ReqTotalResource)

	ValidateMobileType(req.Type)
}

type townResource struct {
	Id   int32  `json:"id"`
	Dat  string `json:"dat"`
	Size int32  `json:"size"`
}

const SANDBOX_RESOURCE_VERSION = 90000

func (proc procTotalResource) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqTotalResource)

	c := GetDBConn(req.App)
	defer c.Close()

	var dats = make([]townResource, 0)
	url_prefix, err := redis.String(c.Do("GET", RedisKey_ClientResourceUrl()))
	if err != nil {
		return &map[string]interface{}{"resource": dats, "error": 0}, nil
	}

	version := req.Version
	if strings.EqualFold(req.Token, "_SANDBOX_") || strings.HasPrefix(req.Token, "_TEST_") {
		version = SANDBOX_RESOURCE_VERSION
	}

	key_resource := RedisKey_TotalResourceHash(version)
	v, err := redis.Values(c.Do("HGETALL", key_resource))
	if err != nil {
		return &map[string]interface{}{"resource": dats, "error": 0}, nil
	}

	rspMap := make(map[string]interface{}, 0)
	maxTown := req.MaxTown
	for len(v) > 0 {
		var field string
		var val string
		v, err = redis.Scan(v, &field, &val)
		if err != nil {
			return &map[string]interface{}{"resource": dats, "error": 0}, nil
		}

		//val: url,size
		arrs := strings.Split(val, ",")
		if len(arrs) != 2 {
			continue
		}
		path := arrs[0]
		size, err := strconv.Atoi(arrs[1])
		if err != nil {
			continue
		}

		id, err := strconv.Atoi(field)
		if err == nil {
			if int32(id) > maxTown || strings.EqualFold(strings.TrimSpace(val), "") {
				continue
			}

			dats = append(dats, townResource{Id: int32(id), Dat: url_prefix + path, Size: int32(size)})
		} else {
			if strings.EqualFold(strings.TrimSpace(val), "") {
				continue
			}
			rspMap[field] = url_prefix + path
			rspMap[field+"_size"] = size
		}
	}

	rspMap["town"] = dats
	rspMap["error"] = 0
	return rspMap, nil
}

// "/announce"
type ReqAnnounce struct {
	ReqBase
	AnnounceRevision int32
}

type procAnnounce struct {
}

func (proc procAnnounce) Req() interface{} {
	return &ReqAnnounce{}
}

func (proc procAnnounce) Validate(in interface{}) {
	req := in.(*ReqAnnounce)

	ValidateMobileType(req.Type)
}

func (proc procAnnounce) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqAnnounce)

	c := GetDBConn(req.App)
	defer c.Close()

	key_announce_rev := RedisKey_AnnounceCurrentRevByType(req.Type)

	cur_rev, err := redis.Int(c.Do("GET", key_announce_rev))
	if err != nil {
		log.Infof("current revison not found, type:%v\n", req.Type)
		return &map[string]interface{}{"revision": -1}, nil
	}

	if req.AnnounceRevision == int32(cur_rev) {
		return &map[string]interface{}{"revision": cur_rev}, nil
	}

	key_announce := RedisKey_AnnounceByTypeRevision(req.Type, int32(cur_rev))
	announce, err := redis.String(c.Do("GET", key_announce))
	if err != nil {
		log.Infof("announce not found, type:%v, rev:%v\n", req.Type, cur_rev)
		return &map[string]interface{}{"revision": -1}, nil
	}
	return &map[string]interface{}{"revision": cur_rev, "content": announce}, nil
}

// "/announce"
type ReqSetAnnounce struct {
	ReqBase
	Announce string
	Title    string
	Date     string
	Sign     string
}

type procSetAnnounce struct {
}

func (proc procSetAnnounce) Req() interface{} {
	return &ReqSetAnnounce{}
}

func (proc procSetAnnounce) Validate(in interface{}) {
	req := in.(*ReqSetAnnounce)

	ValidateMobileType(req.Type)
}

func (proc procSetAnnounce) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqSetAnnounce)

	hash := md5.New()
	hash.Write([]byte(req.Announce + "_" + XDGM_REQUEST_SECRET))
	my_sign := fmt.Sprintf("%x", hash.Sum(nil))
	if my_sign != req.Sign {
		return &map[string]interface{}{"error": 1}, nil
	}

	b := SetAnnounce(req.Type, req.App, req.Announce, req.Title, req.Date)
	if !b {
		return &map[string]interface{}{"error": 2}, nil
	}
	return &map[string]interface{}{"error": 0}, nil
}

func SetAnnounce(iType uint8, app, announce, tittle, date string) bool {
	c := GetDBConn(app)
	defer c.Close()

	rev_key := RedisKey_AnnounceCurrentRevByType(iType)
	rev, err := redis.Int(c.Do("GET", rev_key))
	if err != nil {
		rev = 0
	}

	next_rev := rev + 1
	announce_key := RedisKey_AnnounceByTypeRevision(iType, int32(next_rev))
	c.Do("SET", rev_key, next_rev)
	c.Do("SET", announce_key, announce)

	return true
}

// "/systemInfo"
type ReqSystemInfo struct {
	ReqBase
}

type procSystemInfo struct {
}

func (proc procSystemInfo) Req() interface{} {
	return &ReqSystemInfo{}
}

func (proc procSystemInfo) Validate(in interface{}) {
	req := in.(*ReqSystemInfo)

	ValidateMobileType(req.Type)
}

func (proc procSystemInfo) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqSystemInfo)

	isAudit := isEqualAudit(req.Type, req.Version)
	isActionAudit := isActionAudit(req.Type, req.App)
	sids, version, iCompatible, url := getDarkLaunchInfo(req.Type, req.App)

	//darkSids长度为0，表示没有灰度服
	//如果进灰度服，版本号低于darkClientVersion，则从darkUpdateUrl下载最新客户端
	//如何客户端版本高于等于darkClientVersion，如果darkCompatible标明不兼容非灰度服，则不能进非灰度服
	return &map[string]interface{}{
		"isAudit":           isAudit,
		"isActionAudit":     isActionAudit,
		"darkSids":          sids,
		"darkClientVersion": version,
		"darkCompatible":    iCompatible,
		"darkUpdateUrl":     url,
	}, nil
}

func isActionAudit(iType uint8, app string) bool {
	if iType == TYPE_MOBILE_ANDROID_WEIXIN || iType == TYPE_MOBILE_ANDROID_QQ || iType == TYPE_MOBILE_ANDROID_GUEST {
		return false
	}

	c := GetDBConn(app)
	defer c.Close()

	auditKey := RedisKey_ActionAudit(iType)
	isActionAudit, err := redis.Int(c.Do("GET", auditKey))

	if err == nil && isActionAudit == 0 {
		return false
	} else {
		return true
	}
}

func getDarkLaunchInfo(iType uint8, app string) ([]int, int32, int8, string) {
	c := GetDBConn(app)
	defer c.Close()

	key := RedisKey_DarkLaunch(iType)
	val, err := redis.String(c.Do("GET", key))
	if err != nil {
		return make([]int, 0), 0, 0, ""
	}

	arrs := strings.Split(val, ",")
	if len(arrs) != 4 {
		return make([]int, 0), 0, 0, ""
	}

	s_sids := strings.Split(arrs[0], ":")
	i_sids := make([]int, 0)
	for _, v := range s_sids {
		i, err := strconv.Atoi(v)
		if err == nil {
			i_sids = append(i_sids, i)
		}
	}
	version, err := strconv.Atoi(arrs[1])
	isCompatible, err := strconv.Atoi(arrs[2])
	url := arrs[3]

	return i_sids, int32(version), int8(isCompatible), url
}

type ReqDisableActionPic struct {
	ReqBase
}

type procDisableActionPic struct {
}

func (proc procDisableActionPic) Req() interface{} {
	return &ReqDisableActionPic{}
}

func (proc procDisableActionPic) Validate(in interface{}) {
	req := in.(*ReqDisableActionPic)

	ValidateMobileType(req.Type)
}

func (proc procDisableActionPic) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqDisableActionPic)

	c := GetDBConn(req.App)
	defer c.Close()

	key := RedisKey_DisableActionPic(req.Type)

	pictures, err := redis.Strings(c.Do("SMEMBERS", key))

	if err != nil || len(pictures) == 0 {
		return &map[string]interface{}{"disabled": make([]string, 0)}, nil
	}

	return &map[string]interface{}{"disabled": pictures}, nil
}

type ReqAddBlackIp struct {
	ReqBase
	Ip   string
	Sign string
}

type procAddBlackIp struct {
}

func (proc procAddBlackIp) Req() interface{} {
	return &ReqAddBlackIp{}
}

func (proc procAddBlackIp) Validate(in interface{}) {
	req := in.(*ReqAddBlackIp)

	ValidateMobileType(req.Type)
}

func (proc procAddBlackIp) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqAddBlackIp)

	if !checkSignForGM(req.Ip, req.Sign) {
		return &map[string]interface{}{"error": 1}, nil
	}

	c := GetDBConn(req.App)
	defer c.Close()

	key := RedisKey_IpBlackTable()
	_, err := c.Do("SADD", key, req.Ip)
	if err != nil {
		return &map[string]interface{}{"error": 2}, err
	}

	return &map[string]interface{}{"error": 0}, nil
}

type ReqRemoveBlackIp struct {
	ReqBase
	Ip   string
	Sign string
}

type procRemoveBlackIp struct {
}

func (proc procRemoveBlackIp) Req() interface{} {
	return &ReqRemoveBlackIp{}
}

func (proc procRemoveBlackIp) Validate(in interface{}) {
	req := in.(*ReqRemoveBlackIp)

	ValidateMobileType(req.Type)
}

func (proc procRemoveBlackIp) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqRemoveBlackIp)

	if !checkSignForGM(req.Ip, req.Sign) {
		return &map[string]interface{}{"error": 1}, nil
	}

	c := GetDBConn(req.App)
	defer c.Close()

	key := RedisKey_IpBlackTable()
	_, err := c.Do("SREM", key, req.Ip)
	if err != nil {
		return &map[string]interface{}{"error": 2}, err
	}

	return &map[string]interface{}{"error": 0}, nil
}

type ReqGetIpBlackTable struct {
	ReqBase
	Sign string
}

type procGetIpBlackTable struct {
}

func (proc procGetIpBlackTable) Req() interface{} {
	return &ReqGetIpBlackTable{}
}

func (proc procGetIpBlackTable) Validate(in interface{}) {
	req := in.(*ReqGetIpBlackTable)

	ValidateMobileType(req.Type)
}

func (proc procGetIpBlackTable) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqGetIpBlackTable)

	if !checkSignForGM(req.OpenId, req.Sign) {
		return &map[string]interface{}{"error": 1}, nil
	}

	blackTable, b := ip_black_tables[req.App]
	if !b {
		return &map[string]interface{}{"blackTable": make([]string, 0), "error": 0}, errors.New("no ip black table")
	}
	return &map[string]interface{}{"blackTable": blackTable, "error": 0}, nil
}

type ReqProductInfo struct {
	ReqBase
}

type procProductInfo struct {
}

func (proc procProductInfo) Req() interface{} {
	return &ReqProductInfo{}
}

func (proc procProductInfo) Validate(in interface{}) {
	req := in.(*ReqProductInfo)

	ValidateMobileType(req.Type)
}

func (proc procProductInfo) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqProductInfo)

	fmt.Println(req.App)
	if req.App == "xxd_vn" {
		info := `
		{
			"error": 0,
			"products": [
			{
				"order_info":"sg70_sc_10",
				"title": "50 : 10,000 VND",
				"descrition": "50",
				"price": 10},
			{
				"order_info":"sg70_sc_20",
				"title": "100 : 20,000 VND",
				"descrition": "100",
				"price": 20},
			{
				"order_info":"sg70_sc_30",
				"title": "150 : 30,000 VND",
				"descrition": "150",
				"price": 30},
			{
				"order_info":"sg70_sc_50",
				"title": "250 : 50,000 VND",
				"descrition": "250",
				"price": 50},
			{
				"order_info":"sg70_sc_100",
				"title": "500 : 100,000 VND",
				"descrition": "500",
				"price": 100},
			{
				"order_info":"sg70_sc_200",
				"title": "1000 : 200,000 VND",
				"descrition": "1,000",
				"price": 200},
			{
				"order_info":"sg70_sc_300",
				"title": "1500 : 300,000 VND",
				"descrition": "1,500",
				"price": 300},
			{
				"order_info":"sg70_sc_500",
				"title": "2500 : 500,000 VND",
				"descrition": "2,500",
				"price": 500},
				{
				"order_info":"sg70_sc_1000",
				"title": "5000 : 1,000,000 VND",
				"descrition": "5,000",
				"price": 1000},
			{
				"order_info":"sg70_sc_2000",
				"title": "10,000 : 2,000,000 VND",
				"descrition": "10,000",
				"price": 2000},
			{
				"order_info":"sg70_sc_5000",
				"title": "25,000 : 5,000,000 VND",
				"descrition": "25,000",
				"price": 5000}
			]
		}
		`
		resp := make(map[string]interface{})
		err := json.Unmarshal([]byte(info), &resp)
		if err == nil {
			return resp, nil
		}
		fmt.Println(err)
		log.Error(err.Error())
	}

	return &map[string]interface{}{"error": 1}, nil
}

func initSetting() (err error) {
	err = updateSetting()
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
			err := updateSetting()
			if err != nil {
				log.Errorf("fetch gs list error inside ticker: %v", err)
			}
		}
	}()
	return
}
func updateSetting() (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Recovered in %v: %v", r, string(debug.Stack()))
		}
	}()

	dbmap := GetDBMap()

	_server_list_lock.Lock()
	defer _server_list_lock.Unlock()

	for app, _ := range dbmap {
		c := GetDBConn(app)
		defer c.Close()

		v, err := redis.Strings(c.Do("SMEMBERS", RedisKey_IpBlackTable()))
		if err != nil {
			log.Error(err.Error())
			ip_black_tables[app] = make([]string, 0)
			continue
		}

		log.Infof("updating game setting: %s, %v", app, v)
		ip_black_tables[app] = v
	}
	return nil
}

func isInIpBlackTable(app, ip string) bool {
	blackTable, b := ip_black_tables[app]
	if !b {
		return false
	}
	for _, val := range blackTable {
		if ip == val {
			return true
		}
	}

	return false
}
