package platform_server

import (
	"core/fail"
	"core/log"
	"core/redis"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	WEGAME_SDK_TEST_URL       = "http://test-api.wegames.com.tw/api/"
	WEGAME_SDK_PRODUCTION_URL = "http://api.wegames.com.tw/api/"

	WEGAME_GAME_CODE   = "AAACAA"
	WEGAME_GAME_SECRET = "c5a7b0b41b3774863d31e88f221ba42b"

	HTTP_REQUEST_TIME_OUT = time.Second * 15
)

func handleWGResponse(c chan map[string]interface{}) map[string]interface{} {
	pushResp := make(map[string]interface{})
	select {
	case resp := <-c:
		pullResp := make(map[string]interface{})
		j := json.NewDecoder(resp["data"].(io.ReadCloser))
		err := j.Decode(&pullResp)
		if err != nil {
			log.Error(err.Error())
			pushResp["status"] = 0
			pushResp["msg"] = "inner server error"
			pushResp["errcode"] = 1
		} else {
			if resp["error"] != 0 {
				log.Error("wg register error")
				pushResp["msg"] = resp["msg"]
				pushResp["errcode"] = 1
				pushResp["status"] = 0
			} else {
				//succeed
				for key, val := range pullResp {
					pushResp[key] = val
				}
			}
		}
	case <-time.After(HTTP_REQUEST_TIME_OUT):
		pushResp["msg"] = "request timeout"
		pushResp["errcode"] = 1
		pushResp["status"] = 0
		log.Error("timeout")
	}

	return pushResp
}

type ReqWGRegister struct {
	ReqBase
	UserName  string
	Password  string
	Email     string
	Telephone string
}

type procWGRegister struct {
}

func (proc procWGRegister) Req() interface{} {
	return &ReqWGRegister{}
}

func (proc procWGRegister) Validate(in interface{}) {
	req := in.(*ReqWGRegister)
	fail.When(req.OpenId == "", "incorrect openid")
}

func (proc procWGRegister) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqWGRegister)

	c := make(chan map[string]interface{})
	defer close(c)

	pushResp := make(map[string]interface{})

	var url string
	if req.Token == "_SANDBOX_" {
		url = WEGAME_SDK_TEST_URL
	} else {
		url = WEGAME_SDK_PRODUCTION_URL
	}

	if req.UserName == "" || req.Password == "" {
		pushResp["msg"] = "username or password is empty"
		pushResp["errcode"] = 1
		pushResp["status"] = 0

		return &pushResp, nil
	}
	requestMap := make(map[string]string)
	requestMap["wg_method"] = "user.register"
	requestMap["wg_game_code"] = WEGAME_GAME_CODE
	requestMap["wg_username"] = req.UserName
	requestMap["wg_password"] = req.Password
	if req.Email != "" {
		requestMap["wg_email"] = req.Email
	}
	if req.Telephone != "" {
		requestMap["wg_telephone"] = req.Telephone
	}
	requestMap["wg_ip"] = req.ClientIp
	requestMap["wg_version"] = "1"
	requestMap["wg_time"] = strconv.FormatInt(time.Now().Unix(), 10)

	keys := make([]string, 0)
	for key, _ := range requestMap {
		keys = append(keys, key)
	}

	sort.StringSlice(keys).Sort()
	for i, key := range keys {
		keys[i] = key + "=" + requestMap[key]
	}
	params := strings.Join(keys, "&")
	sign := makeWGSign(params)
	params = params + "&wg_sign=" + sign

	go post(url, []byte(params), c)
	pushResp = handleWGResponse(c)

	var err error

	if pushResp["status"] == float64(1) {
		data := pushResp["data"].(map[string]interface{})
		platform_uid := data["platform_uid"].(string)

		//move guest_id roles to platform_uid roles
		if isEqualAudit(req.Type, req.Version) {
			err = moveRoles(req.App, req.OpenId, platform_uid, TYPE_MOBILE_AUDIT_WEIXIN)
		} else {
			err = moveRoles(req.App, req.OpenId, platform_uid, req.Type)
			err = moveRoles(req.App, req.OpenId, platform_uid, TYPE_MOBILE_SANDBOX)
		}

		guest_roles := getRoles(req.App, req.OpenId, req.Type, req.Version)
		//set wegame roles
		err = setWGRole(req.App, platform_uid, req.ClientIp, getWGSDKUrl(req.Token), guest_roles)
	}

	return &pushResp, err
}

func makeWGSign(params string) string {
	source := params + WEGAME_GAME_SECRET

	MD5 := md5.New()
	MD5.Write([]byte(source))

	return fmt.Sprintf("%x", MD5.Sum(nil))
}

type ReqWGLogin struct {
	ReqBase
	UserName string
	Password string
}

type procWGLogin struct {
}

func (proc procWGLogin) Req() interface{} {
	return &ReqWGLogin{}
}

func (proc procWGLogin) Validate(in interface{}) {
	req := in.(*ReqWGLogin)
	fail.When(req.OpenId == "", "incorrect openid")
}

func (proc procWGLogin) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqWGLogin)

	c := make(chan map[string]interface{})
	defer close(c)

	pushResp := make(map[string]interface{})

	var url string
	if req.Token == "_SANDBOX_" {
		url = WEGAME_SDK_TEST_URL
	} else {
		url = WEGAME_SDK_PRODUCTION_URL
	}

	if req.UserName == "" || req.Password == "" {
		pushResp["msg"] = "username or password is empty"
		pushResp["errcode"] = 1
		pushResp["status"] = 0
		return &pushResp, nil
	}
	requestMap := make(map[string]string)
	requestMap["wg_method"] = "user.login"
	requestMap["wg_game_code"] = WEGAME_GAME_CODE
	requestMap["wg_username"] = req.UserName
	requestMap["wg_password"] = req.Password
	requestMap["wg_ip"] = req.ClientIp
	requestMap["wg_version"] = "1"
	requestMap["wg_time"] = strconv.FormatInt(time.Now().Unix(), 10)

	keys := make([]string, 0)
	for key, _ := range requestMap {
		keys = append(keys, key)
	}

	sort.StringSlice(keys).Sort()
	for i, key := range keys {
		keys[i] = key + "=" + requestMap[key]
	}
	params := strings.Join(keys, "&")
	sign := makeWGSign(params)
	params = params + "&wg_sign=" + sign

	go post(url, []byte(params), c)
	pushResp = handleWGResponse(c)

	return &pushResp, nil
}

type ReqWGModifyPasswd struct {
	ReqBase
	UserName  string
	OldPasswd string
	NewPasswd string
}

type procWGModifyPasswd struct {
}

func (proc procWGModifyPasswd) Req() interface{} {
	return &ReqWGModifyPasswd{}
}

func (proc procWGModifyPasswd) Validate(in interface{}) {
	req := in.(*ReqWGModifyPasswd)
	fail.When(req.OpenId == "", "incorrect openid")
}

func (proc procWGModifyPasswd) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqWGModifyPasswd)

	c := make(chan map[string]interface{})
	defer close(c)

	pushResp := make(map[string]interface{})

	var url string
	if req.Token == "_SANDBOX_" {
		url = WEGAME_SDK_TEST_URL
	} else {
		url = WEGAME_SDK_PRODUCTION_URL
	}

	if req.UserName == "" || req.OldPasswd == "" || req.NewPasswd == "" {
		pushResp["msg"] = "username or password is empty"
		pushResp["errcode"] = 1
		pushResp["status"] = 0
		return &pushResp, nil
	}
	requestMap := make(map[string]string)
	requestMap["wg_method"] = "user.modify_passwd"
	requestMap["wg_game_code"] = WEGAME_GAME_CODE
	requestMap["wg_username"] = req.UserName
	requestMap["wg_old_passwd"] = req.OldPasswd
	requestMap["wg_new_passwd"] = req.NewPasswd
	requestMap["wg_ip"] = req.ClientIp
	requestMap["wg_version"] = "1"
	requestMap["wg_time"] = strconv.FormatInt(time.Now().Unix(), 10)

	keys := make([]string, 0)
	for key, _ := range requestMap {
		keys = append(keys, key)
	}

	sort.StringSlice(keys).Sort()
	for i, key := range keys {
		keys[i] = key + "=" + requestMap[key]
	}
	params := strings.Join(keys, "&")
	sign := makeWGSign(params)
	params = params + "&wg_sign=" + sign

	go post(url, []byte(params), c)
	pushResp = handleWGResponse(c)

	return &pushResp, nil
}

type ReqWGBindThird struct {
	ReqBase
	Source   string
	UniqueId string
	OpenId2  string
	Email    string
}

type procWGBindThird struct {
}

func (proc procWGBindThird) Req() interface{} {
	return &ReqWGBindThird{}
}

func (proc procWGBindThird) Validate(in interface{}) {
	req := in.(*ReqWGBindThird)
	fail.When(req.OpenId == "", "incorrect openid")
}

func (proc procWGBindThird) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqWGBindThird)

	cli := GetDBConn(req.App)
	defer cli.Close()

	pushResp := make(map[string]interface{})
	if req.Source == "" || req.UniqueId == "" || req.OpenId2 == "" {
		pushResp["msg"] = "source or unique_id or open_id2 is empty"
		pushResp["errcode"] = 1
		pushResp["status"] = 0
		return &pushResp, nil
	}

	c := make(chan map[string]interface{})
	defer close(c)
	third_key := RedisKey_WG_THIRD_BIND_KEY(req.Source, req.UniqueId)
	platform_uid, _ := redis.String(cli.Do("GET", third_key))
	if platform_uid != "" {
		pushResp["msg"] = "ok"
		pushResp["errcode"] = 0
		pushResp["status"] = 1
		_data := make(map[string]string)
		_data["platform_uid"] = platform_uid
		pushResp["data"] = _data
		return &pushResp, nil
	}

	requestMap := make(map[string]string)
	requestMap["wg_method"] = "user.third"
	requestMap["wg_game_code"] = WEGAME_GAME_CODE
	requestMap["wg_source"] = req.Source
	requestMap["wg_unique_id"] = req.UniqueId
	requestMap["wg_open_id"] = req.OpenId2
	if req.Email != "" {
		requestMap["wg_email"] = req.Email
	}
	requestMap["wg_ip"] = req.ClientIp
	requestMap["wg_version"] = "1"
	requestMap["wg_time"] = strconv.FormatInt(time.Now().Unix(), 10)

	keys := make([]string, 0)
	for key, _ := range requestMap {
		keys = append(keys, key)
	}

	sort.StringSlice(keys).Sort()
	for i, key := range keys {
		keys[i] = key + "=" + requestMap[key]
	}
	params := strings.Join(keys, "&")
	sign := makeWGSign(params)
	params = params + "&wg_sign=" + sign

	var url string
	if req.Token == "_SANDBOX_" {
		url = WEGAME_SDK_TEST_URL
	} else {
		url = WEGAME_SDK_PRODUCTION_URL
	}

	go post(url, []byte(params), c)
	pushResp = handleWGResponse(c)

	if pushResp["status"] == float64(1) {
		data := pushResp["data"].(map[string]interface{})
		_, err := cli.Do("SET", third_key, data["platform_uid"].(string))
		if err != nil {
			log.Error(err.Error())
		}
	}

	return &pushResp, nil
}

type ReqWGBindAccount struct {
	ReqBase
	UserName string
	Password string
	Sign     string
}

/*
	绑定分两种，一种是客户端用用户名和密码绑定，
	一种是客服用用户名来绑定，后者需要验证来保证请求的合法性
*/
type procWGBindAccount struct {
}

func (proc procWGBindAccount) Req() interface{} {
	return &ReqWGBindAccount{}
}

func (proc procWGBindAccount) Validate(in interface{}) {
	req := in.(*ReqWGBindAccount)
	fail.When(req.OpenId == "", "incorrect openid")
}

func (proc procWGBindAccount) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqWGBindAccount)

	c := make(chan map[string]interface{})
	defer close(c)

	pushResp := make(map[string]interface{})

	var url string
	if req.Token == "_SANDBOX_" {
		url = WEGAME_SDK_TEST_URL
	} else {
		url = WEGAME_SDK_PRODUCTION_URL
	}

	if req.UserName == "" || (req.Password == "" && req.Sign == "") {
		pushResp["msg"] = "username, password or sign error"
		pushResp["errcode"] = 1
		pushResp["status"] = 0
		return &pushResp, nil
	}

	requestMap := make(map[string]string)
	if req.Sign != "" {
		//客服手动绑定，check sign
		hash := md5.New()
		hash.Write([]byte(req.UserName + "_" + XDGM_REQUEST_SECRET))
		my_sign := fmt.Sprintf("%x", hash.Sum(nil))
		if my_sign != req.Sign {
			pushResp["msg"] = "sign error"
			pushResp["errcode"] = 1
			pushResp["status"] = 0
			return &pushResp, nil
		}

		requestMap["wg_method"] = "user.account_bind2"
	} else {
		requestMap["wg_method"] = "user.account_bind"
	}

	requestMap["wg_game_code"] = WEGAME_GAME_CODE
	requestMap["wg_guest_id"] = req.OpenId
	requestMap["wg_username"] = req.UserName
	if req.Password != "" {
		requestMap["wg_password"] = req.Password
	}
	requestMap["wg_ip"] = req.ClientIp
	requestMap["wg_version"] = "1"
	requestMap["wg_time"] = strconv.FormatInt(time.Now().Unix(), 10)

	keys := make([]string, 0)
	for key, _ := range requestMap {
		keys = append(keys, key)
	}

	sort.StringSlice(keys).Sort()
	for i, key := range keys {
		keys[i] = key + "=" + requestMap[key]
	}
	params := strings.Join(keys, "&")
	sign := makeWGSign(params)
	params = params + "&wg_sign=" + sign

	go post(url, []byte(params), c)
	pushResp = handleWGResponse(c)

	var err error
	if pushResp["status"] == float64(1) {
		data := pushResp["data"].(map[string]interface{})
		platform_uid := data["platform_uid"].(string)

		//move guest_id roles to platform_uid roles
		if isEqualAudit(req.Type, req.Version) {
			err = moveRoles(req.App, req.OpenId, platform_uid, TYPE_MOBILE_AUDIT_WEIXIN)
		} else {
			err = moveRoles(req.App, req.OpenId, platform_uid, req.Type)
			err = moveRoles(req.App, req.OpenId, platform_uid, TYPE_MOBILE_SANDBOX)
		}

		guest_roles := getRoles(req.App, req.OpenId, req.Type, req.Version)
		//set wegame roles
		err = setWGRole(req.App, platform_uid, req.ClientIp, getWGSDKUrl(req.Token), guest_roles)
	}

	return &pushResp, err
}

func moveRoles(app, fromOpenId, toOpenId string, iType uint8) error {
	c := GetDBConn(app)
	defer c.Close()

	from_key := RedisKey_RoleListByOpenidType(fromOpenId, iType)
	to_key := RedisKey_RoleListByOpenidType(toOpenId, iType)
	v, err := redis.Values(c.Do("HGETALL", from_key))
	if err != nil {
		log.Errorf("user roleinfo not found, openid:%v, itype:%v\n", from_key, iType)
		return err
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

		//注入guest_id
		u.GuestId = fromOpenId
		b, err = json.Marshal(u)
		if err != nil {
			log.Errorf("fail to marshal role info: %v", u)
			continue
		}

		_, err = c.Do("HSET", to_key, sid, string(b))
		if err != nil {
			log.Error(err.Error())
			return err
		}
	}

	_, err = c.Do("DEL", from_key)
	return err
}

func getRoles(app string, openid string, iType uint8, version int32) RoleList {
	isEqualAudit := isEqualAudit(iType, version)
	var roles RoleList
	if isEqualAudit {
		roles = getRoleListByOpenIdType(app, openid, TYPE_MOBILE_AUDIT_WEIXIN)
	} else {
		roles = getRoleListByOpenIdType(app, openid, iType)
		sandboxRoles := getRoleListByOpenIdType(app, openid, TYPE_MOBILE_SANDBOX)

		if len(sandboxRoles) > 0 {
			for _, v := range sandboxRoles {
				roles = append(roles, v)
			}
		}
	}

	return roles
}

func getWGSDKUrl(token string) string {
	if token == "_SANDBOX_" {
		return WEGAME_SDK_TEST_URL
	}

	return WEGAME_SDK_PRODUCTION_URL
}

func setWGRole(app, platform_uid, ip, url string, roles RoleList) (err error) {
	if len(roles) == 0 {
		return nil
	}

	c := make(chan map[string]interface{})
	defer close(c)

	for _, role := range roles {
		requestMap := make(map[string]string)
		requestMap["wg_method"] = "user.setrole"
		requestMap["wg_game_code"] = WEGAME_GAME_CODE
		requestMap["wg_server_code"] = strconv.Itoa(int(role.Sid))
		requestMap["wg_platform_uid"] = platform_uid
		requestMap["wg_role_name"] = role.Nick
		requestMap["wg_game_uid"] = "0"
		requestMap["wg_ip"] = ip
		requestMap["wg_version"] = "1"
		requestMap["wg_time"] = strconv.FormatInt(time.Now().Unix(), 10)

		keys := make([]string, 0)
		for key, _ := range requestMap {
			keys = append(keys, key)
		}

		sort.StringSlice(keys).Sort()
		for i, key := range keys {
			keys[i] = key + "=" + requestMap[key]
		}
		params := strings.Join(keys, "&")
		sign := makeWGSign(params)
		params = params + "&wg_sign=" + sign

		go post(url, []byte(params), c)
		pushResp := handleWGResponse(c)

		if pushResp["status"] != float64(1) {
			msg := pushResp["msg"].(string)
			log.Info(msg)
			err = errors.New(msg)
		}
	}

	return err
}
