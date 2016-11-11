package web

import (
	"bytes"
	"core/log"
	"crypto/md5"
	"fmt"
	"net/http"
	"payment/config"
	"sort"
	"strings"
)

func wegamesValidate(req *http.Request) (ok bool) {
	if config.CFG.DisableValidate {
		return true
	}
	var paramKeys sort.StringSlice
	var sign string
	for key, vals := range req.PostForm {
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
		buff.WriteString(req.PostForm[key][0])
		if i+1 < len(paramKeys) {
			buff.WriteByte('&')
		}
	}
	buff.WriteString(config.CFG.Wegames.PrivateKey)
	log.Debugf("req %s\n", buff.String())
	srvSign := strings.ToLower(fmt.Sprintf("%x", md5.Sum(buff.Bytes())))
	log.Debugf("wegames validation client  %s server %s", sign, srvSign)
	return srvSign == sign
}

func validate(req *http.Request) (ok bool) {
	if config.CFG.DisableValidate {
		return true
	}
	var paramKeys sort.StringSlice
	var sign, enhanced_sign string
	for key, vals := range req.PostForm {
		var val string
		if len(vals) > 0 {
			val = vals[0]
		} else {
			continue
		}
		if key == "sign" {
			sign = val
			continue
		}
		if key == "enhanced_sign" {
			enhanced_sign = val
			continue
		}
		paramKeys = append(paramKeys, key)
	}
	paramKeys.Sort()
	var paramStr string
	for _, key := range paramKeys {
		paramStr += req.PostForm[key][0]
	}
	//fmt.Println("排序后的字符串", paramStr)
	if !config.CFG.EnhancedValidate {

		str1 := strings.ToLower(fmt.Sprintf("%x", md5.Sum([]byte(paramStr))))
		str2 := strings.ToLower(fmt.Sprintf("%x", md5.Sum([]byte(str1+config.CFG.PrivateKey))))
		//fmt.Println("普通校验")
		//fmt.Println(str1)
		//fmt.Println(str2)
		return str2 == sign
	} else {
		str1 := strings.ToLower(fmt.Sprintf("%x", md5.Sum([]byte(paramStr))))
		str2 := strings.ToLower(fmt.Sprintf("%x", md5.Sum([]byte(str1+config.CFG.EnhancedPrivateKey))))
		//fmt.Println("增强校验")
		//fmt.Println(str1)
		//fmt.Println(str2)
		return str2 == enhanced_sign
	}
}
