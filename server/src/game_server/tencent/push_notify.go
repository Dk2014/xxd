package tencent

import (
	"bytes"
	"core/fail"
	"core/log"
	util "core/time"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"game_server/api/protocol/player_api"
	"game_server/config"
	"game_server/module"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	XG_HTTP_METHOD = "GET"
	XG_URL         = "/v2/push/single_account"

	XG_MESSAGE_TYPE = "1" //消息类型：通知

	XG_PLATFORM_TYPE_IOS     = 0
	XG_PLATFORM_TYPE_ANDROID = 1

	QQ_SCORE     = "/profile/qqscore"
	WEIXIN_SCORE = "/profile/wxscore"
)

type XD_Message struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	BuilderId string `json:"builder_id"`
}

type XD_Rsp struct {
	RetCode int    `json:"ret_code"` //返回码 非零返回码标识出错
	ErrMsg  string `json:"err_msg"`
	Result  string `json:"result"`
}

type Api_Rsp struct {
	Ret  int    `json:"ret"` //返回码 非零返回码标识出错
	Msg  string `json:"msg"`
	Type int    `json:"type"`
}

type QQ_SCORE_REQ struct {
	Appid       string `json:"appid"`
	AccessToken string `json:"accessToken"`
	Openid      string `json:"openid"`
	Data        string `json:"data"`
	Type        int    `json:"type"`
	Bcover      int    `json:"bcover"`
	Expires     string `json:"expires"`
}

type WX_SCORE_REQ struct {
	Appid     string `json:"appid"`
	GrantType string `json:"grantType"`
	Openid    string `json:"openid"`
	Score     string `json:"score"`
	Expires   string `json:"expires"`
}

func SendNotificationToSingleAccount(pid int64, title string, content string) {
	if !config.ServerCfg.XGSdk.Enable {
		return
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("SendNotificationToSingleAccount error: %v\n", err)
			}
		}()

		parameter := make(map[string]string, 0)
		//通用参数
		parameter["access_id"] = config.ServerCfg.XGSdk.AccessId
		parameter["timestamp"] = fmt.Sprintf("%d", util.GetNowTime())

		parameter["account"] = fmt.Sprintf("%d", pid)
		parameter["message_type"] = XG_MESSAGE_TYPE
		parameter["message"] = getMessageForXG(title, content)

		parameter["environment"] = getEnvForXG()

		//在最后计算sign
		parameter["sign"] = getSignForXG(parameter)

		rsp := &XD_Rsp{}
		SendHTTPReqForXG(config.ServerCfg.XGSdk.Addr, parameter, XG_URL, rsp)

		fail.When(rsp.RetCode != 0, fmt.Sprintf("%v", rsp))
	}()
}

func SendFightNumToTencent(moneyState *module.MoneyState, fightnum int32) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("SendFightNumToTencent error: %v\n", err)
			}
		}()
		var reqURL string
		if moneyState.PlatformType == int8(player_api.PLATFORM_TYPE_MOBILE_ANDROID_QQ) || moneyState.PlatformType == int8(player_api.PLATFORM_TYPE_MOBILE_IOS_QQ) {
			parameter := QQ_SCORE_REQ{}
			parameter.Appid = config.ServerCfg.GameAppID
			parameter.Openid = moneyState.Openid
			parameter.AccessToken = moneyState.Openkey
			parameter.Data = strconv.FormatInt(int64(fightnum), 10)
			parameter.Type = 3
			parameter.Bcover = 1
			parameter.Expires = "0"
			reqURL = config.ServerCfg.MSDKApiAddr + QQ_SCORE + getParInInterface(moneyState)
			rsp := &Api_Rsp{}
			SendHTTPReqForAsyncApi(reqURL, parameter, rsp)

			fail.When(rsp.Ret != 0, fmt.Sprintf("%v", rsp))
		} else if moneyState.PlatformType == int8(player_api.PLATFORM_TYPE_MOBILE_ANDROID_WEIXIN) || moneyState.PlatformType == int8(player_api.PLATFORM_TYPE_MOBILE_IOS_WEIXIN) {
			parameter := WX_SCORE_REQ{}
			parameter.Appid = config.ServerCfg.GameAppID
			parameter.Openid = moneyState.Openid
			parameter.GrantType = "client_credential"
			parameter.Score = strconv.FormatInt(int64(fightnum), 10)
			parameter.Expires = "0"
			reqURL = config.ServerCfg.MSDKApiAddr + WEIXIN_SCORE + getParInInterface(moneyState)
			rsp := &Api_Rsp{}
			SendHTTPReqForAsyncApi(reqURL, parameter, rsp)

			fail.When(rsp.Ret != 0, fmt.Sprintf("%v", rsp))
		} else {
			return
		}
	}()
}

//计算信鸽推送加密签名
// parameter不能以`sign`为key
func getSignForXG(parameter map[string]string) string {
	if _, invalidKey := parameter["sign"]; invalidKey {
		fail.When(true, "getSignForXG 的 parameter 不能含有 `sign` 这个key")
	}

	sortedParam := NewMapSorter(parameter)
	sort.Sort(sortedParam)
	var strmap []string
	for _, v2 := range sortedParam {
		strmap = append(strmap, v2.Key+"="+v2.Val)
	}
	rawSign := XG_HTTP_METHOD + config.ServerCfg.XGSdk.Addr + XG_URL + strings.Join(strmap, "") + config.ServerCfg.XGSdk.SecretKey
	return fmt.Sprintf("%x", md5.Sum([]byte(rawSign)))
}

func getMessageForXG(title string, content string) string {
	if config.ServerCfg.XGSdk.PlatformType == XG_PLATFORM_TYPE_ANDROID {
		return fmt.Sprintf(`{"title":"%s", "content": "%s", "builder_id":0}`, title, content)
	} else if config.ServerCfg.XGSdk.PlatformType == XG_PLATFORM_TYPE_IOS {
		return fmt.Sprintf(`{"aps":{"alert":"%s:%s"}}`, title, content)
	}
	panic("unsupport  platform type")
}

/*
	environment
	向iOS设备推送时必填，1表示推送生产环境；2表示推送开发环境。推送Android平台不填或填0
*/
func getEnvForXG() string {
	if config.ServerCfg.XGSdk.PlatformType == XG_PLATFORM_TYPE_ANDROID {
		return "0"
	} else if config.ServerCfg.XGSdk.PlatformType == XG_PLATFORM_TYPE_IOS {
		return "1"
	}
	panic("unsupport  platform type")
}

func SendHTTPReqForXG(addr string, val map[string]string, orgloc string, rspStruct interface{}) {
	var strtemp []string
	for k, v := range val {
		strtemp = append(strtemp, k+"="+url.QueryEscape(v))
	}

	reqURL := "http://" + addr + orgloc + "?" + strings.Join(strtemp, "&")
	req, _ := http.NewRequest("GET", reqURL, nil)

	DefaultClient := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(time.Duration(config.ServerCfg.XGSdk.Timeout) * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Duration(config.ServerCfg.XGSdk.Timeout)*time.Second)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}

	resp, err := DefaultClient.Do(req)
	fail.When(err != nil, fmt.Sprintf("error: %v, request: %s\n", err, reqURL))

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fail.When(err != nil, fmt.Sprintf("error: %v, request: %s\n", err, reqURL))

	err = json.Unmarshal(body, rspStruct)
	fail.When(err != nil, fmt.Sprintf("解析json失败: %v [%v], %s\n", err, body, reqURL))

	return
}

func SendHTTPReqForAsyncApi(reqURL string, val interface{}, rspStruct interface{}) {
	postDataBytes, _ := json.Marshal(val)
	postBytesReader := bytes.NewReader(postDataBytes)
	req, _ := http.NewRequest("POST", reqURL, postBytesReader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	DefaultClient := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(time.Duration(config.ServerCfg.MoneySdk.MoneySdkTimeOutSec)*time.Second + 5*time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*time.Duration(config.ServerCfg.MoneySdk.MoneySdkTimeOutSec))
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}

	resp, err := DefaultClient.Do(req)
	if err != nil {
		log.Error(fmt.Sprintf("error: %v, request: %s\n", err, reqURL))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(fmt.Sprintf("error: %v, request: %s\n", err, reqURL))
	}

	err = json.Unmarshal(body, rspStruct)
	if err != nil {
		log.Error(fmt.Sprintf("解析json失败: %v [%v], %s\n", err, body, reqURL))
	}

	return
}

func getParInInterface(moneyState *module.MoneyState) string {
	params := make(map[string]string)
	params["timestamp"] = fmt.Sprintf("%d", util.GetNowTime())
	params["appid"] = config.ServerCfg.GameAppID
	str := fmt.Sprintf("%v%v", config.ServerCfg.GameAppKey, params["timestamp"])
	m := md5.New()
	m.Write([]byte(str))
	params["sig"] = hex.EncodeToString(m.Sum(nil))
	params["openid"] = moneyState.Openid
	params["encode"] = "1"
	params["conn"] = "1" //启用长连接
	parStr := "?"
	for key, val := range params {
		parStr += fmt.Sprintf("%v=%v&", key, val)
	}
	parStr = parStr[:len(parStr)-1]
	return parStr
}
