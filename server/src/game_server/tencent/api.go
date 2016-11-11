package tencent

import (
	"core/fail"
	_ "core/log"
	"encoding/json"
	"fmt"
	"game_server/api/protocol/player_api"
	. "game_server/config"
	"game_server/module"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type GET_BALANCE_RSP struct {
	Ret         int64     `json:"ret"` //返回码 0：成功； 1001：参数错误1018：登陆校验失败。其它：失败
	Msg         string    `json:"msg"`
	Balance     int64     `json:"balance"`     //游戏币个数（包含了赠送游戏币）
	Gen_balance int64     `json:"gen_balance"` //赠送游戏币个数
	First_save  int       `json:"first_save"`  //是否满足首次充值，1：满足，0：不满足
	SaveAmt     int64     `json:"save_amt"`    //累计充值金额
	Tss_list    []TssList `json:"tss_list"`    //已经开通的订阅物品列表
}

type PAY_RSP struct {
	Ret        int64  `json:"ret"` //返回码 0：成功； 1001：参数错误1018：登陆校验失败。其它：失败
	Msg        string `json:"msg"`
	Billno     string `json:"billno"`       //预扣流水号
	Balance    int64  `json:"balance"`      //预扣后的余额
	UsedGenAmt int64  `json:"used_gen_amt"` //用掉的赠送游戏币个数
}

type PRESENT_RSP struct {
	Ret     int64  `json:"ret"` //返回码 0：成功； 1001：参数错误1018：登陆校验失败。其它：失败
	Msg     string `json:"msg"`
	Balance int64  `json:"balance"` //预扣后的余额
}

type TssList struct {
	InnerProductId string `json:"inner_productid"`     //用户开通的订阅物品id（注：该参数为计费分配的订阅物品的servicecode）
	BeginTime      string `json:"begintime"`           //用户订阅的开始时间
	EndTime        string `json:"endtime"`             //用户订阅的结束时间
	PayChan        string `json:"paychan"`             //用户订阅该物品id最后一次的支付渠道
	PaySubChan     int64  `json:"paysubchan"`          //用户订阅该物品id最后一次的支付子渠道id
	OpenDays       int64  `json:"grandtotal_opendays"` //用户订阅天数
}

func setPayReqCookie(req *http.Request, platformtype int8, orgloc string) {
	/*
		Cookie里面需要包含的参数：
		session_id    用户账户类型，session_id =“openid”
		session_type	session类型，session_type = “kp_actoken”
		org_loc     	需要填写: mpay/get_balance_m
		appip       	（可选）来源的第三方应用的服务IP
		注意：cookie里面参数的值，需要进行urlencode
	*/
	if platformtype == int8(player_api.PLATFORM_TYPE_MOBILE_ANDROID_WEIXIN) {
		req.AddCookie(&http.Cookie{Name: "session_id", Value: "hy_gameid"})
		req.AddCookie(&http.Cookie{Name: "session_type", Value: "wc_actoken"})
	} else if platformtype == int8(player_api.PLATFORM_TYPE_MOBILE_ANDROID_QQ) {
		req.AddCookie(&http.Cookie{Name: "session_id", Value: "openid"})
		req.AddCookie(&http.Cookie{Name: "session_type", Value: "kp_actoken"})
	} else if platformtype == int8(player_api.PLATFORM_TYPE_MOBILE_IOS_WEIXIN) {
		req.AddCookie(&http.Cookie{Name: "session_id", Value: "wechatid"})
		req.AddCookie(&http.Cookie{Name: "session_type", Value: "wc_actoken"})
	} else if platformtype == int8(player_api.PLATFORM_TYPE_MOBILE_IOS_QQ) {
		req.AddCookie(&http.Cookie{Name: "session_id", Value: "openid"})
		req.AddCookie(&http.Cookie{Name: "session_type", Value: "kp_actoken"})
	} else if platformtype == int8(player_api.PLATFORM_TYPE_MOBILE_IOS_GUEST) {
		req.AddCookie(&http.Cookie{Name: "session_id", Value: "hy_gameid"})
		req.AddCookie(&http.Cookie{Name: "session_type", Value: "st_dummy"})
	}
	req.AddCookie(&http.Cookie{Name: "org_loc", Value: orgloc})
}

func Send(addr string, val map[string]string, orgloc string, platformtype int8, rspStruct interface{}, asyncCallback func(response string)) string {
	var strtemp []string
	for k, v := range val {
		strtemp = append(strtemp, k+"="+url.QueryEscape(v))
	}

	reqURL := addr + orgloc + "?" + strings.Join(strtemp, "&")

	// 异步请求(目前仅用于玩家登陆服务器进行token验证)
	if asyncCallback != nil {
		req, err := module.UrlServer.NewRequestForGet(reqURL)
		if err != nil {
			return reqURL
		}
		setPayReqCookie(req, platformtype, orgloc)
		module.UrlServer.Push(module.UrlWork{
			Req:      req,
			Callback: asyncCallback,
		})
		return reqURL
	}

	req, _ := http.NewRequest("GET", reqURL, nil)
	DefaultClient := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(time.Duration(ServerCfg.MoneySdk.MoneySdkTimeOutSec) * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Duration(ServerCfg.MoneySdk.MoneySdkTimeOutSec)*time.Second)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}
	setPayReqCookie(req, platformtype, orgloc)
	resp, err := DefaultClient.Do(req)
	fail.When(err != nil, fmt.Sprintf("tencent money: error: %v, request: %s\n", err, reqURL))

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fail.When(err != nil, fmt.Sprintf("tencent money: error: %v, request: %s\n", err, reqURL))

	err = json.Unmarshal(body, rspStruct)
	fail.When(err != nil, fmt.Sprintf("解析json失败: %v [%v], %s\n", err, body, reqURL))

	return reqURL
}
