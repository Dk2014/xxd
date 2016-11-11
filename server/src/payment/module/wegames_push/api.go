package wegames_push

import (
	"bytes"
	"core/debug"
	"core/log"
	"core/time"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"payment/config"
	"payment/database"
	"payment/module"
	"sort"
	"strings"
)

func init() {
	module.Wegames = WegamesPushMod{}
}

type WegamesPushMod struct{}

func (mod WegamesPushMod) PushGooglePlay(paymentQueue []*database.GooglePlayPendingQueue) {
	payData := make([]*wegamesPayData, 0, len(paymentQueue))
	otherItem := ""
	payIds := []int64{}
	for _, pay := range paymentQueue {
		payIds = append(payIds, pay.Id)
		productCfg, ok := config.CFG.ProductInfo[pay.ProductId]
		if !ok {
			log.Errorf("[PushGooglePlay] Unknow productId [%s] id[%d]", pay.ProductId, pay.Id)
			continue
		}
		if productCfg.IsMonthCard {
			otherItem = "month"
		} else {
			otherItem = ""
		}
		payData = append(payData, &wegamesPayData{
			ServerCode:   fmt.Sprintf("%d", pay.GameUserId>>32/10),
			RoleName:     pay.Nickname,
			PlatformUid:  pay.OpenId,
			PaySource:    "google",
			ThirdOrderId: pay.TransactionId,
			PayCurrCode:  "TWD",
			PayAmount:    productCfg.Cost,
			GameMoney:    productCfg.GameMoney,
			OtherItem:    otherItem,
			PaymentTime:  fmt.Sprintf("%d", pay.PurchaseDateMs/1000),
		})
	}
	payDataByte, err := json.Marshal(payData)
	if err != nil {
		log.Errorf("[PushGooglePlay] marshal error [%v]", err)
	}
	pushRsp, err := post(fmt.Sprintf("%s", payDataByte))
	if err != nil {
		log.Errorf("[PushGooglePlay] post error [%v]", err)
	}
	log.Debugf("[PushGooglePlay] resp %s", debug.Print(0, false, true, "    ", nil, pushRsp))
	database.BatchDeleteRecordById(payIds, database.TABLE_WEGAMES_GOOGLE_PLAY_DELIVER)
}

func (mod WegamesPushMod) PushAppStore(paymentQueue []*database.AppStorePendingQueue) {
	payData := make([]*wegamesPayData, 0, len(paymentQueue))
	otherItem := ""
	payIds := []int64{}
	for _, pay := range paymentQueue {
		payIds = append(payIds, pay.Id)
		productCfg, ok := config.CFG.ProductInfo[pay.ProductId]
		if !ok {
			log.Errorf("[PushGooglePlay] Unknow productId [%s] id[%d]", pay.ProductId, pay.Id)
			continue
		}
		if productCfg.IsMonthCard {
			otherItem = "month"
		} else {
			otherItem = ""
		}
		payData = append(payData, &wegamesPayData{
			ServerCode:   fmt.Sprintf("%d", pay.GameUserId>>32/10),
			RoleName:     pay.Nickname,
			PlatformUid:  pay.OpenId,
			PaySource:    "app",
			ThirdOrderId: pay.TransactionId,
			PayCurrCode:  "TWD",
			PayAmount:    productCfg.Cost,
			GameMoney:    productCfg.GameMoney,
			OtherItem:    otherItem,
			PaymentTime:  fmt.Sprintf("%d", pay.PurchaseDateMs/1000),
		})
	}
	payDataByte, err := json.Marshal(payData)
	if err != nil {
		log.Errorf("[PushAppStore] marshal error [%v]", err)
	}
	pushRsp, err := post(fmt.Sprintf("%s", payDataByte))
	if err != nil {
		log.Errorf("[PushAppStore] post error [%v]", err)
	}
	log.Debugf("[PushAppStore] resp %v", debug.Print(0, false, true, "    ", nil, pushRsp))
	database.BatchDeleteRecordById(payIds, database.TABLE_WEGAMES_APP_STORE_DELIVER)
}

type wegamesPayData struct {
	ServerCode   string  `json:"wg_server_code"`
	RoleName     string  `json:"wg_role_name"`
	PlatformUid  string  `json:"wg_platform_uid"`
	PaySource    string  `json:"wg_pay_source"`
	ThirdOrderId string  `json:"wg_third_order_id"`
	PayCurrCode  string  `json:"wg_pay_curr_code"`
	PayAmount    float64 `json:"wg_pay_amount"`
	GameMoney    int64   `json:"wg_game_money"`
	OtherItem    string  `json:"wg_other_item"`
	PaymentTime  string  `json:"wg_payment_time"`
}

type WegamesPushRespError struct {
	Status    int32  `json:"status"`
	ErrorCode int32  `json:"errorcode"`
	Msg       string `json:"msg"`
}

type WegamesPushRespData struct {
	OrderSuccessNum int32                           `json:"order_success_num"`
	OrderErrorNum   int32                           `json:"order_error_num"`
	OrderSuccess    string                          `json:"order_success"`
	OrderError      string                          `json:"order_error"`
	ErrorList       map[string]WegamesPushRespError `json:"error_list"`
}

type WegamesPushResp struct {
	Status int32                `json:"status"`
	Msg    string               `json:"msg"`
	Data   *WegamesPushRespData `json:"data"`
}

func post(payData string) (*WegamesPushResp, error) {
	data := url.Values{}
	postUrl := ""
	data["wg_method"] = []string{"pay.back_log_batch"}
	if config.CFG.Wegames.IsProduction {
		postUrl = config.CFG.Wegames.ProductUrl
	} else {
		postUrl = config.CFG.Wegames.DevUrl
	}
	data["wg_game_code"] = []string{config.CFG.Wegames.GameCode}
	data["wg_pay_data"] = []string{payData}
	data["wg_version"] = []string{"1"}
	data["wg_time"] = []string{fmt.Sprintf("%d", time.GetNowTime())}
	_, sign := postSign(data)
	data["wg_sign"] = []string{sign}
	log.Debugf("PostPayLogToWegames sign [%s]", sign)
	resp, err := http.PostForm(postUrl, data)
	if err != nil {
		return nil, err
	}
	if config.CFG.Debug {
		respBytes, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			log.Debugf("wegames post require args  %s", debug.Print(0, false, true, "    ", nil, data))
			log.Debugf("wegames post response content  %s", respBytes)
		} else {
			log.Debugf("wegames post response error %v", err)
		}
	}

	respStruct := WegamesPushResp{}
	err = json.NewDecoder(resp.Body).Decode(&respStruct)
	if err != nil {
		return nil, err
	}
	return &respStruct, nil
}

func postSign(postData url.Values) (urlSign string, computedSign string) {
	var paramKeys sort.StringSlice
	var sign string
	for key, vals := range postData {
		if len(vals) == 0 {
			continue
		}
		if key == "wg_sign" {
			sign = vals[0]
			continue
		}
		paramKeys = append(paramKeys, key)
	}
	paramKeys.Sort()
	buff := &bytes.Buffer{}
	for i, key := range paramKeys {
		buff.WriteString(key)
		buff.WriteByte('=')
		buff.WriteString(postData[key][0])
		if i+1 < len(paramKeys) {
			buff.WriteByte('&')
		}
	}
	buff.WriteString(config.CFG.Wegames.PrivateKey)
	return sign, strings.ToLower(fmt.Sprintf("%x", md5.Sum(buff.Bytes())))
}
