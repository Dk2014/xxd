package platform_server

import (
	"bytes"
	"core/log"
	"crypto/md5"
	"fmt"
	"net/http"
)

//http post request
func post(uri string, body []byte, c chan map[string]interface{}) {
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
		c <- map[string]interface{}{"error": 1, "msg": "服务器http请求出错"}
		return
	}
	if resp.StatusCode != 200 {
		c <- map[string]interface{}{"error": 2, "msg": "http请求返回异常"}
		return
	}

	c <- map[string]interface{}{"error": 0, "data": resp.Body, "msg": "http请求返回异常"}
}

func checkSignForGM(source, sign string) bool {
	hash := md5.New()
	hash.Write([]byte(source + "_" + XDGM_REQUEST_SECRET))
	my_sign := fmt.Sprintf("%x", hash.Sum(nil))
	if my_sign == sign {
		return true
	}
	return false
}
