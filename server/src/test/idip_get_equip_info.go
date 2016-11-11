package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type IDIP_COMMON_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
}

// 角色装备信息对象
type SEquipInfo struct {
	EquipId   uint64 // 装备ID
	EquipName string // 装备名称
	Level     uint32 // 装备精炼等级
	IsBattle  uint8  // 上阵1/未上阵0

}

// IDIP消息头
type IDIP_REQ_HEAD struct {
	PacketLen    uint32 // 包长
	Cmdid        uint32 // 命令ID
	Seqid        uint32 // 流水号
	ServiceName  string // 服务名
	SendTime     uint32 // 发送时间YYYYMMDD对应的整数
	Version      uint32 // 版本号
	Authenticate string // 加密串
	Result       int32  // 错误码,返回码类型：0：处理成功，需要解开包体获得详细信息,1：处理成功，但包体返回为空，不需要处理包体（eg：查询用户角色，用户角色不存在等），-1: 网络通信异常,-2：超时,-3：数据库操作异常,-4：API返回异常,-5：服务器忙,-6：其他错误,小于-100 ：用户自定义错误，需要填写szRetErrMsg
	RetErrMsg    string // 错误信息

}

type IDIP_QUERY_ROLE_EQUIP_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
		AreaId    uint32 // 大区（1微信，2手Q）
		Partition uint32 // 小区ID
		PlatId    uint8  // 平台（0ios，1安卓）
		OpenId    string // openid
		RoleId    uint64 // 角色ID
	} `json:"body"`
}

type IDIP_QUERY_ROLE_EQUIP_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
		EquipList_count uint32       // 角色装备信息列表的最大数量
		EquipList       []SEquipInfo // 角色装备信息列表
	} `json:"body"`
}

func main() {
	client := &http.Client{}

	var request_par IDIP_QUERY_ROLE_EQUIP_REQ

	request_par.Head.PacketLen = 0
	request_par.Head.Cmdid = 4125
	request_par.Head.Seqid = 1234
	request_par.Head.ServiceName = "idip"
	request_par.Head.SendTime = 20140731
	request_par.Head.Version = 4
	request_par.Head.Authenticate = ""
	request_par.Head.Result = 0
	request_par.Head.RetErrMsg = "nil"
	request_par.Body.AreaId = 1
	request_par.Body.Partition = 1
	request_par.Body.PlatId = 0
	request_par.Body.OpenId = "_test_id1003022414061086775"
	request_par.Body.RoleId = 4
	json_val, _ := json.Marshal(request_par)
	form := url.Values{
		"data_packet": {
			string(json_val),
		},
	}

	val, _ := url.QueryUnescape(form.Encode())

	//获取到request对象
	req, _ := http.NewRequest("POST", "http://127.0.0.1:9999", strings.NewReader(val))
	fmt.Println(url.QueryUnescape(form.Encode()))
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	cmd := new(IDIP_COMMON_REQ)
	err := json.Unmarshal(body, cmd)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(cmd)
	req2 := new(IDIP_QUERY_ROLE_EQUIP_RSP)
	err = json.Unmarshal(body, req2)
	fmt.Println(req2)
	// requst_par := {
	// 	"Head": {
	// 		"PacketLen":    "",
	// 		"Cmdid":        4107,
	// 		"Seqid":        1234,
	// 		"ServiceName":  "idip",
	// 		"SendTime":     "20140731",
	// 		"Version":      4,
	// 		"Authenticate": "",
	// 		"Result":       0,
	// 		"RetErrMsg":    "",
	// 	},
	// 	"Body": {
	// 		"AreaId":    1,
	// 		"Partition": 1,
	// 		"PlatId":    1,
	// 		"OpenId":    "孙桃柳",
	// 	},
	// }

}
