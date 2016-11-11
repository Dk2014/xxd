package xdgm_server

import (
	"bytes"
	"core/log"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"platform_server"
	"strconv"
	"time"
	mconfig "xdgm_server/config"
)

const XDGM_REQUEST_SECRET = "xxd@2015"

func Pid2GameServerId(pid int64) int {
	return int(pid >> 32)
}

func Pid2GlobalServerId(pid int64) int {
	gsid := (pid >> 32)
	return int((gsid/10)*10 + 9)
}

func Sid2GlobalServerId(sid int) int {
	return int(sid*10 + 9)
}

func GetGsServer(app string, platid uint8, sid int) (platform_server.Server, bool) {
	st := &platform_server.ServerType{AreaId: uint8(1), PlatId: platid}
	gs, exists := platform_server.GetServerInfo(int32(sid), st.GetType(), app)
	return gs, exists
}

func GetAllGsServer(app string, platid uint8) map[string]platform_server.Server {
	var needRedis = make(map[string]platform_server.Server, 0)
	st := &platform_server.ServerType{AreaId: uint8(1), PlatId: platid}
	res := platform_server.ServerList(app)
	for k, v := range res {
		if st.GetType() == v.Type || v.Type == 255 {
			needRedis[k] = v
		}
	}
	return needRedis
}
func GetAllGsServerByHttp(platid uint8) []mconfig.RPCServerConf {
	serverInfoList := httpGetServerInfo()
	result := make([]mconfig.RPCServerConf, 0)
	for _, item := range serverInfoList {
		if platid == 0 || uint8(item.Type) == GetType(platid) || item.Type == 255 {
			result = append(result, item)
		}
	}
	return result
}

func GetGsServerByHttp(platid uint8, sid int) ([]mconfig.RPCServerConf, bool) {
	serverInfoList := httpGetServerInfo()
	result := make([]mconfig.RPCServerConf, 0)
	exists := false
	if serverInfoList != nil {
		for _, server := range serverInfoList {
			if server.GSID/10 == sid {
				exists = true
				result = append(result, server)
			}
		}
	}
	return result, exists
}

func SetAnnounce(content string, noticetime int64, platid uint8) (errorCode int, message string) {
	errorCode, message = httpSetAnnounce(content, noticetime, platid)
	return errorCode, message
}

func httpGetServerInfo() map[string]mconfig.RPCServerConf {
	reqURL := config.PlatformServerUrl + "/gserverall"
	// Add rpc server list from http response
	postValues := make(map[string]interface{})
	postValues["App"] = config.DefaultApp
	postDataBytes, _ := json.Marshal(postValues)
	postBytesReader := bytes.NewReader(postDataBytes)
	req, _ := http.NewRequest("POST", reqURL, postBytesReader)
	req.Header.Add("Content-Type", "application/json")

	DefaultClient := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(15 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*15)
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
		return nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(fmt.Sprintf("error: %v, request: %s\n", err, reqURL))
		return nil
	}
	list := make(map[string]mconfig.RPCServerConf)
	err = json.Unmarshal(body, &list)
	if err != nil {
		log.Error(fmt.Sprintf("解析json失败: %v [%v], %s\n", err, body, reqURL))
		return nil
	}
	return list
}

func httpSetAnnounce(content string, noticetime int64, platid uint8) (int, string) {
	reqURL := config.PlatformServerUrl + "/announce/update"
	// Add rpc server list from http response
	postValues := make(map[string]interface{})
	sign := md5.Sum([]byte(content + "_" + XDGM_REQUEST_SECRET))
	postValues["App"] = config.DefaultApp
	postValues["Type"] = GetType(platid)
	postValues["Announce"] = content
	postValues["Title"] = ""
	postValues["Date"] = time.Unix(int64(noticetime), 0).Format("2006-01-02 15:04:05")
	postValues["Sign"] = fmt.Sprintf("%x", sign)
	postDataBytes, _ := json.Marshal(postValues)
	postBytesReader := bytes.NewReader(postDataBytes)
	req, _ := http.NewRequest("POST", reqURL, postBytesReader)
	req.Header.Add("Content-Type", "application/json")

	DefaultClient := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(15 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*15)
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
		return 1, err.Error()
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(fmt.Sprintf("error: %v, request: %s\n", err, reqURL))
		return 1, err.Error()
	}
	returnMap := make(map[string]interface{})
	err = json.Unmarshal(body, &returnMap)
	if err != nil {
		log.Error(fmt.Sprintf("解析json失败: %v [%v], %s\n", err, body, reqURL))
		return 1, err.Error()
	}
	errorCode, err := strconv.Atoi(fmt.Sprintf("%v", returnMap["error"]))
	if err != nil {
		log.Error(fmt.Sprintf("解析error: %v [%v], %s\n", err, body, reqURL))
		return 1, err.Error()
	}
	return errorCode, fmt.Sprintf("%v", returnMap["msg"])
}

func BlackIp(ip string, mode string) (errorCode int, message string, list interface{}) {
	if mode == "0" {
		errorCode, message := addBlackIp(ip)
		return errorCode, message, nil
	} else if mode == "1" {
		errorCode, message := removeBlackIp(ip)
		return errorCode, message, nil
	} else if mode == "2" {
		errorCode, message, list := BlackIpList()
		return errorCode, message, list
	}
	return 3, "some other error", nil
}

func addBlackIp(ip string) (int, string) {
	reqURL := config.PlatformServerUrl + "/xdgm/addBlackIp"
	postValues := make(map[string]interface{})
	sign := md5.Sum([]byte(config.DefaultApp + "-" + ip + "_" + XDGM_REQUEST_SECRET))
	postValues["App"] = config.DefaultApp
	postValues["Ip"] = ip
	postValues["Sign"] = fmt.Sprintf("%x", sign)
	postDataBytes, _ := json.Marshal(postValues)
	postBytesReader := bytes.NewReader(postDataBytes)
	req, _ := http.NewRequest("POST", reqURL, postBytesReader)
	req.Header.Add("Content-Type", "application/json")

	DefaultClient := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(15 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*15)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}
	resp, err := DefaultClient.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(fmt.Sprintf("error: %v, request: %s\n", err, reqURL))
		return -1, err.Error()
	}
	returnMap := make(map[string]interface{})
	err = json.Unmarshal(body, &returnMap)
	if err != nil {
		log.Error(fmt.Sprintf("解析json失败: %v [%v], %s\n", err, body, reqURL))
		return -1, err.Error()
	}
	errorCode, err := strconv.Atoi(fmt.Sprintf("%v", returnMap["error"]))
	if err != nil {
		log.Error(fmt.Sprintf("解析error: %v [%v], %s\n", err, body, reqURL))
		return -1, err.Error()
	}
	return errorCode, ""
}

func removeBlackIp(ip string) (int, string) {
	reqURL := config.PlatformServerUrl + "/xdgm/removeBlackIp"
	postValues := make(map[string]interface{})
	sign := md5.Sum([]byte(config.DefaultApp + "-" + ip + "_" + XDGM_REQUEST_SECRET))
	postValues["App"] = config.DefaultApp
	postValues["Ip"] = ip
	postValues["Sign"] = fmt.Sprintf("%x", sign)
	postDataBytes, _ := json.Marshal(postValues)
	postBytesReader := bytes.NewReader(postDataBytes)
	req, _ := http.NewRequest("POST", reqURL, postBytesReader)
	req.Header.Add("Content-Type", "application/json")

	DefaultClient := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(15 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*15)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}
	resp, err := DefaultClient.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(fmt.Sprintf("error: %v, request: %s\n", err, reqURL))
		return -1, err.Error()
	}
	returnMap := make(map[string]interface{})
	err = json.Unmarshal(body, &returnMap)
	if err != nil {
		log.Error(fmt.Sprintf("解析json失败: %v [%v], %s\n", err, body, reqURL))
		return -1, err.Error()
	}
	errorCode, err := strconv.Atoi(fmt.Sprintf("%v", returnMap["error"]))
	if err != nil {
		log.Error(fmt.Sprintf("解析error: %v [%v], %s\n", err, body, reqURL))
		return -1, err.Error()
	}
	return errorCode, ""
}

func BlackIpList() (int, string, interface{}) {
	reqURL := config.PlatformServerUrl + "/xdgm/getIpBlackTable"
	postValues := make(map[string]interface{})
	sign := md5.Sum([]byte(config.DefaultApp + "_" + XDGM_REQUEST_SECRET))
	postValues["App"] = config.DefaultApp
	postValues["Sign"] = fmt.Sprintf("%x", sign)
	postDataBytes, _ := json.Marshal(postValues)
	postBytesReader := bytes.NewReader(postDataBytes)
	req, _ := http.NewRequest("POST", reqURL, postBytesReader)
	req.Header.Add("Content-Type", "application/json")

	DefaultClient := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(15 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*15)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}
	resp, err := DefaultClient.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(fmt.Sprintf("error: %v, request: %s\n", err, reqURL))
		return 1, err.Error(), nil
	}
	returnMap := make(map[string]interface{})
	err = json.Unmarshal(body, &returnMap)
	if err != nil {
		log.Error(fmt.Sprintf("解析json失败: %v [%v], %s\n", err, body, reqURL))
		return 1, err.Error(), nil
	}
	errorCode, err := strconv.Atoi(fmt.Sprintf("%v", returnMap["error"]))
	if err != nil {
		log.Error(fmt.Sprintf("解析error: %v [%v], %s\n", err, body, reqURL))
		return 1, err.Error(), nil
	}
	return errorCode, "", returnMap["blackTable"]
}

func GetType(platid uint8) uint8 {
	//广州测试服area不能和正式服一样，所以定义了98-》手q，99-》微信
	return 1 | (platid << 4)
}
