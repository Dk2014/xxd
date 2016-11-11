package platform_server

import (
	"bytes"
	"core/fail"
	"core/log"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	MSDK_AUTH_URL      = "http://msdk.tencent-cloud.net/auth"
	MSDK_AUTH_URL_TEST = "http://msdktest.qq.com/auth"
	ANYSDK_AUTH_URL    = "http://asdk.xindong.com/api/User/LoginOauth/"
)

const (
	APPID_QQ  = 1000001161
	APPKEY_QQ = "CDFhyo38ffNllQ6L"

	APPID_WEIXIN  = "wxd4ade9d6689aff30"
	APPKEY_WEIXIN = "9203705497046c7657af816c5cfc118b"

	APPID_GUEST  = "G_1000001161"
	APPKEY_GUEST = "CDFhyo38ffNllQ6L"
)

func AuthURLByToken(token *string) string {
	if strings.HasPrefix(*token, "_TEST_") {
		*token = (*token)[6:]
		return MSDK_AUTH_URL_TEST
	}
	return MSDK_AUTH_URL
}

func CheckToken(iType uint8, openid, token, userip string) (*map[string]interface{}, error) {

	if token == "_SANDBOX_" {
		return &map[string]interface{}{"error": float64(0)}, nil
	}

	c := make(chan map[string]interface{})
	defer close(c)

	var resp map[string]interface{}

	var params string
	url := AuthURLByToken(&token) + "/%v/?timestamp=%v&appid=%v&sig=%v&openid=%v&encode=1"

	timestamp := time.Now().Unix()
	switch iType {

	case TYPE_MOBILE_IOS_QQ, TYPE_MOBILE_ANDROID_QQ:
		sig := makeSig(APPKEY_QQ + strconv.FormatInt(timestamp, 10))
		qq_template := `{"appid": %v, "openkey": "%v", "openid":"%v", "userip":"%v"}`
		params = fmt.Sprintf(qq_template, APPID_QQ, token, openid, userip)
		url = fmt.Sprintf(url, "verify_login", timestamp, APPID_QQ, sig, openid)

	case TYPE_MOBILE_IOS_WEIXIN, TYPE_MOBILE_ANDROID_WEIXIN:
		sig := makeSig(APPKEY_WEIXIN + strconv.FormatInt(timestamp, 10))
		weixin_template := `{"appid": "%v", "accessToken": "%v", "openid":"%v", "userip":"%v"}`
		params = fmt.Sprintf(weixin_template, APPID_WEIXIN, token, openid, userip)
		url = fmt.Sprintf(url, "check_token", timestamp, APPID_WEIXIN, sig, openid)

	case TYPE_MOBILE_IOS_GUEST, TYPE_MOBILE_ANDROID_GUEST:
		sig := makeSig(APPKEY_GUEST + strconv.FormatInt(timestamp, 10))
		guest_template := `{"guestid": "%v", "accessToken": "%v"}`
		params = fmt.Sprintf(guest_template, openid, token)
		url = fmt.Sprintf(url, "guest_check_token", timestamp, APPID_GUEST, sig, openid)

	default:
		log.Errorf("check token, unknown type:%v\n", iType)
		return &map[string]interface{}{"error": float64(1)}, nil
	}

	go requestForTencent(url, strings.NewReader(params), c)

	select {
	case resp = <-c:
		resp["error"] = resp["ret"]
		delete(resp, "ret")
		if resp["error"] != float64(0) {
			log.Errorf("auth/check resp:%v, url:%v, params:%v\n", resp, url, params)
		}
		return &resp, nil
	case <-time.After(15 * time.Second):
		resp = make(map[string]interface{})
		resp["error"] = 400
		resp["msg"] = "time out"
		log.Errorf("auth/check timeout, url:%v, params:%v\n", url, params)
		return &resp, nil
	}

	log.Errorf("auth/check 服务器出错, url:%v, params:%v\n", url, params)
	return nil, errors.New("服务器出错")
}

func CheckTokenForAnySDK() {

}

func makeSig(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

//http post request
func requestForTencent(uri string, body io.Reader, c chan map[string]interface{}) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Recovered in %v", r)
		}
	}()

	req, _ := http.NewRequest("POST", uri, body)
	defer func() {
		if req.Body != nil {
			req.Body.Close()
		}
	}()

	resp, e := http.DefaultClient.Do(req)
	if e != nil {
		c <- map[string]interface{}{"ret": 400, "msg": "服务器http请求出错"}
		return
	}
	if resp.StatusCode != 200 {
		c <- map[string]interface{}{"ret": resp.StatusCode, "msg": "服务器http请求出错"}
		return
	}
	result := make(map[string]interface{})
	j := json.NewDecoder(resp.Body)
	err := j.Decode(&result)
	if err != nil {
		log.Infof("json decode error, body:%v\n", resp.Body)
		c <- map[string]interface{}{"ret": 400, "msg": "解析http response body 出错"}
		return
	}
	c <- result
}

type ReqRefreshToken struct {
	ReqBase
	RefreshToken string
}

type procRefreshToken struct {
}

func (proc procRefreshToken) Req() interface{} {
	return &ReqRefreshToken{}
}

func (proc procRefreshToken) Validate(in interface{}) {
	req := in.(*ReqRefreshToken)

	fail.When(req.OpenId == "", "incorrect openid")
	fail.When(req.RefreshToken == "", "incorrect refreshToken")
	fail.When(req.Type == 2, "incorrect app type")
}

func (proc procRefreshToken) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqRefreshToken)

	c := make(chan map[string]interface{})
	defer close(c)

	var resp map[string]interface{}

	var params string
	url := AuthURLByToken(&(req.Token)) + "/%v/?timestamp=%v&appid=%v&sig=%v&openid=%v&encode=1"

	timestamp := time.Now().Unix()
	sig := makeSig(APPKEY_WEIXIN + strconv.FormatInt(timestamp, 10))
	params_template := `{"appid": "%v", "refreshToken": "%v"}`
	params = fmt.Sprintf(params_template, APPID_WEIXIN, req.RefreshToken)
	url = fmt.Sprintf(url, "refresh_token", timestamp, APPID_WEIXIN, sig, req.OpenId)

	go requestForTencent(url, strings.NewReader(params), c)

	select {
	case resp = <-c:
		resp["error"] = resp["ret"]
		delete(resp, "ret")
		if resp["error"] != 0 {
			log.Errorf("auth/refreshToken error:%v, url:%v, params:%v\n", resp["msg"], url, params)
		}
		return &resp, nil
	case <-time.After(15 * time.Second):
		resp = make(map[string]interface{})
		resp["error"] = 400
		resp["msg"] = "time out"
		log.Errorf("auth/refreshToken timeout,url:%v, params:%v\n", url, params)
		return &resp, nil
	}

	log.Errorf("auth/refreshToken 服务器出错, url:%v, params:%v\n", url, params)
	return nil, errors.New("服务器出错")
}

func TransportAnySDKOauth(body []byte, w io.Writer) {
	c := make(chan map[string]interface{})
	defer close(c)

	go requestForAnySdk(ANYSDK_AUTH_URL, body, c)
	var resp map[string]interface{}
	select {
	case resp = <-c:
		b, err := json.Marshal(resp)
		fail.When(err != nil, err)
		w.Write(b)
	case <-time.After(15 * time.Second):
		resp = make(map[string]interface{})
		resp["status"] = "fail"
		resp["data"] = `{"error":"time out"}`
		resp["common"] = "{}"
		log.Errorf("auth/refreshToken timeout,url:%v\n", ANYSDK_AUTH_URL)
		b, err := json.Marshal(resp)
		fail.When(err != nil, err)
		w.Write(b)
	}
}

//http post request
func requestForAnySdk(uri string, body []byte, c chan map[string]interface{}) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Recovered in %v", r)
		}
	}()

	req, _ := http.NewRequest("POST", uri, bytes.NewReader(body))
	defer func() {
		if req.Body != nil {
			req.Body.Close()
		}
	}()

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, e := http.DefaultClient.Do(req)
	if e != nil {
		respMap := make(map[string]interface{})
		respMap["status"] = "fail"
		respMap["data"] = `{"error":"服务器http请求出错"}`
		c <- respMap
		return
	}
	if resp.StatusCode != 200 {
		respMap := make(map[string]interface{})
		respMap["status"] = "fail"
		respMap["data"] = `{"error":"服务器http请求出错"}`
		c <- respMap
	}
	result := make(map[string]interface{})
	j := json.NewDecoder(resp.Body)
	err := j.Decode(&result)
	if err != nil {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Infof("json decode error, body:%s\n", string(body))
		respMap := make(map[string]interface{})
		respMap["status"] = "fail"
		respMap["data"] = `{"error":"解析http response body 出错"}`
		c <- respMap
		return
	}
	c <- result
}
