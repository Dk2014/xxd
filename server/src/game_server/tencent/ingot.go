package tencent

/*
	蛋疼的腾讯游戏币相关接口
*/

import (
	"core/fail"
	"core/log"
	"core/time"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"game_server/api/protocol/player_api"
	. "game_server/config"
	"game_server/module"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

const (
	GET_BALANCE = "/mpay/get_balance_m"
	PAY         = "/mpay/pay_m"
	PRESENT     = "/mpay/present_m"
)

type ParameterMapSort []ParameterStr

type ParameterStr struct {
	Key string
	Val string
}

func GetIngotTest(moneyState *module.MoneyState, callback func(*GET_BALANCE_RSP)) {
	parameter := getpar(moneyState)
	parameter["sig"] = getsig(parameter, GET_BALANCE)

	rsp := new(GET_BALANCE_RSP)
	rsp.Ret = -12345 // 默认设置一个错误的ret
	Send(ServerCfg.MoneySdk.MoneySdkAddr, parameter, GET_BALANCE, moneyState.PlatformType, rsp, func(result string) {
		if err := json.Unmarshal([]byte(result), rsp); err != nil {
			log.Errorf("response result: %v, 解析json失败: %v, err: %v, parameter: %v\n", result, GET_BALANCE, err, parameter)
		}

		callback(rsp)
	})
}

func GetIngot(moneyState *module.MoneyState) (ingot int64) {
	parameter := getpar(moneyState)
	parameter["sig"] = getsig(parameter, GET_BALANCE)

	rsp := getIngot(parameter, GET_BALANCE, moneyState.PlatformType)

	fail.When(rsp.Ret != 0, fmt.Sprintf("%v", rsp))

	return rsp.Balance
}

//获取玩家总充值金额，不包括赠送游戏币个数
func GetSaveAmt(moneyState *module.MoneyState) (saveamt int64) {
	parameter := getpar(moneyState)
	parameter["sig"] = getsig(parameter, GET_BALANCE)

	rsp := getIngot(parameter, GET_BALANCE, moneyState.PlatformType)

	fail.When(rsp.Ret != 0, fmt.Sprintf("%v", rsp))

	return rsp.SaveAmt
}

//获取玩家月卡信息
func GetMonthCardInfo(moneyState *module.MoneyState) []TssList {
	if moneyState == nil {
		return nil
	}
	parameter := getpar(moneyState)
	parameter["sig"] = getsig(parameter, GET_BALANCE)
	rsp := getIngot(parameter, GET_BALANCE, moneyState.PlatformType)

	fail.When(rsp.Ret != 0, fmt.Sprintf("%v", rsp))

	return rsp.Tss_list
}

//检查货币是否足够,mtype:COINS,INGOT
func CheckIngot(moneyState *module.MoneyState, num int64) (ret bool) {
	ingot := GetIngot(moneyState)
	if ingot >= num {
		return true
	}
	return false
}

//增加货币,mtype:COINS(0),INGOT(1)
func IncIngot(moneyState *module.MoneyState, discountid, giftid string, num int64) int64 {
	parameter := getpar(moneyState)
	parameter["discountid"] = discountid
	parameter["giftid"] = giftid
	parameter["presenttimes"] = strconv.FormatInt(num, 10)
	parameter["sig"] = getsig(parameter, PRESENT)

	rsp := incIngot(parameter, PRESENT, moneyState.PlatformType)

	fail.When(rsp.Ret != 0, fmt.Sprintf("%v", rsp))

	return rsp.Balance
}

//减少货币,mtype:COINS(0),INGOT(1)
func DecIngot(moneyState *module.MoneyState, num int64) int64 {
	parameter := getpar(moneyState)
	parameter["amt"] = strconv.FormatInt(num, 10)
	parameter["sig"] = getsig(parameter, PAY)

	rsp := decIngot(parameter, PAY, moneyState.PlatformType)

	fail.When(rsp.Ret != 0, fmt.Sprintf("%v", rsp))

	return rsp.Balance
}

func getIngot(reqvalue map[string]string, orgloc string, platformtype int8) *GET_BALANCE_RSP {
	rsp := new(GET_BALANCE_RSP)
	reqURL := Send(ServerCfg.MoneySdk.MoneySdkAddr, reqvalue, orgloc, platformtype, rsp, nil)
	if rsp.Ret != 0 {
		log.Errorf("tencent money api[getIngot]: %s", reqURL)
	}
	return rsp
}

func decIngot(reqvalue map[string]string, orgloc string, platformtype int8) *PAY_RSP {
	rsp := new(PAY_RSP)
	reqURL := Send(ServerCfg.MoneySdk.MoneySdkAddr, reqvalue, orgloc, platformtype, rsp, nil)
	if rsp.Ret != 0 {
		log.Errorf("tencent money api[decIngot]: %s", reqURL)
	}
	return rsp
}

func incIngot(reqvalue map[string]string, orgloc string, platformtype int8) *PRESENT_RSP {
	rsp := new(PRESENT_RSP)
	reqURL := Send(ServerCfg.MoneySdk.MoneySdkAddr, reqvalue, orgloc, platformtype, rsp, nil)
	if rsp.Ret != 0 {
		log.Errorf("tencent money api[incIngot]: %s", reqURL)
	}
	return rsp
}

//生成sig
func getsig(parameter map[string]string, path string) string {
	var Parameter = make(map[string]string, 0)
	for k, v := range parameter {
		if k != "sig" {
			Parameter[k] = url.QueryEscape(v)
		}
	}
	ParameterSort := NewMapSorter(Parameter)
	sort.Sort(ParameterSort)
	var strmap []string
	for _, v2 := range ParameterSort {
		strmap = append(strmap, v2.Key+url.QueryEscape("=")+v2.Val)
	}
	str := strings.Join(strmap, url.QueryEscape("&"))
	str = "GET&" + url.QueryEscape(path) + "&" + str
	appkey := ServerCfg.MoneySdk.MoneySdkAppKey + "&"
	mac := hmac.New(sha1.New, []byte(appkey))
	mac.Write([]byte(str))
	base64str := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return base64str
}

func (ps ParameterMapSort) Len() int {
	return len(ps)
}

func (ps ParameterMapSort) Less(i, j int) bool {
	return ps[i].Key < ps[j].Key // 按值排序
}

func (ps ParameterMapSort) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func NewMapSorter(p map[string]string) ParameterMapSort {
	ParameterSort := make(ParameterMapSort, 0, len(p))

	for k, v := range p {
		ParameterSort = append(ParameterSort, ParameterStr{k, v})
	}

	return ParameterSort
}

func getpar(moneyState *module.MoneyState) map[string]string {
	parameter := make(map[string]string, 0)
	parameter["openid"] = moneyState.Openid
	if moneyState.PlatformType == int8(player_api.PLATFORM_TYPE_MOBILE_IOS_GUEST) {
		parameter["openkey"] = "openkey"
	} else {
		parameter["openkey"] = moneyState.Openkey
	}
	parameter["pay_token"] = moneyState.PayToken
	parameter["appid"] = ServerCfg.MoneySdk.MoneySdkAppId
	parameter["ts"] = strconv.FormatInt(time.GetNowTime(), 10)
	parameter["pf"] = moneyState.Pf
	parameter["pfkey"] = moneyState.Pfkey
	parameter["zoneid"] = strconv.Itoa(moneyState.Zoneid)
	return parameter
}
