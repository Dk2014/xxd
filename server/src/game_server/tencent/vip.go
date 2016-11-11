package tencent

/*
	蛋疼的腾讯会员相关接口
*/

import (
	"bytes"
	"core/log"
	util "core/time"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"game_server/api/protocol/player_api"
	. "game_server/config"
	"game_server/module"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	LOAD_VIP_URL = "/profile/load_vip/"
)

type VipInfoList struct {
	Flag   int16 `json:"flag"`
	IsVip  int16 `json:"isvip"`
	Year   int16 `json:"year"`
	Level  int16 `json:"level"`
	Luxury int16 `json:"luxury"`
	IsPay  int16 `json:"ispay"`
}

type VipInfo struct {
	Ret   int64          `json:"ret"` //返回码 0：成功；其它：失败
	Msg   string         `json:"msg"`
	Lists []*VipInfoList `json:"lists"`
}

const (
	VIP_NOMAL  = 1   //qq会员
	VIP_BLUE   = 4   //蓝专
	VIP_RED    = 8   //红砖
	VIP_SUPER  = 16  //超级会员
	VIP_GAME   = 32  //游戏会员
	VIP_XINYUE = 64  //心悦俱乐部会员
	VIP_YELLOW = 128 //黄砖会员
)

const (
	IS_VIP  = 1
	NOT_VIP = 0
)

func QQVipStatus(moneyState *module.MoneyState) (status int16) {
	if moneyState.PlatformType == int8(player_api.PLATFORM_TYPE_MOBILE_ANDROID_QQ) || moneyState.PlatformType == int8(player_api.PLATFORM_TYPE_MOBILE_IOS_QQ) { //QQ平台
		rspStruct := new(VipInfo)
		SendByPostLoadVIP(ServerCfg.MSDKApiAddr, LOAD_VIP_URL, moneyState, rspStruct)
		if rspStruct.Ret == 0 && len(rspStruct.Lists) > 0 {
			// 有返回结果集
			for _, result := range rspStruct.Lists {
				if result.Flag == VIP_NOMAL && result.IsVip == 1 {
					status |= 1
				}
				if result.Flag == VIP_SUPER && result.IsVip == 1 {
					status |= 2
				}
			}

		}
	}
	return
}

func SendByPostLoadVIP(addr, orgloc string, moneyState *module.MoneyState, rspStruct interface{}) {
	reqURL := addr + orgloc + getParInVip(moneyState)

	postValues := make(map[string]interface{})
	postValues["appid"] = ServerCfg.GameAppID
	postValues["login"] = 2
	postValues["uin"] = 0
	postValues["openid"] = moneyState.Openid
	postValues["vip"] = VIP_NOMAL + VIP_SUPER // 目前要查询的有QQ会员和超级会员
	postDataBytes, _ := json.Marshal(postValues)
	postBytesReader := bytes.NewReader(postDataBytes)
	req, _ := http.NewRequest("POST", reqURL, postBytesReader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	DefaultClient := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(time.Duration(ServerCfg.MoneySdk.MoneySdkTimeOutSec)*time.Second + 5*time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*time.Duration(ServerCfg.MoneySdk.MoneySdkTimeOutSec))
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
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(fmt.Sprintf("error: %v, request: %s\n", err, reqURL))
		return
	}
	err = json.Unmarshal(body, rspStruct)
	if err != nil {
		log.Error(fmt.Sprintf("解析json失败: %v [%v], %s\n", err, body, reqURL))
		return
	}

	return

}

func getParInVip(moneyState *module.MoneyState) string {
	params := make(map[string]string)
	params["timestamp"] = fmt.Sprintf("%d", util.GetNowTime())
	params["appid"] = ServerCfg.GameAppID
	str := fmt.Sprintf("%v%v", ServerCfg.GameAppKey, params["timestamp"])
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
