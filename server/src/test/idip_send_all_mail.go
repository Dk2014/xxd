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

// 封号请求
type IDIP_DO_SEND_MAIL_ALL_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
		AreaId      uint32 // 服务器：微信（1），手Q（2）
		Partition   uint32 // 小区ID
		PlatId      uint8  // 平台：IOS（0），安卓（1）
		MailTitle   string // 邮件标题
		MailContent string // 邮件内容
		SendTime    uint64 // 邮件发送时间
		EndTime     uint64 // 邮件自动删除时间
		ItemId1     uint64 //  邮件赠送道具1ID
		ItemNum1    uint32 // 邮件赠送道具1数量
		ItemId2     uint64 //  邮件赠送道具2ID
		ItemNum2    uint32 // 邮件赠送道具2数量
		ItemId3     uint64 //  邮件赠送道具3ID
		ItemNum3    uint32 // 邮件赠送道具3数量
		ItemId4     uint64 //  邮件赠送道具4ID
		ItemNum4    uint32 // 邮件赠送道具4数量
		Source      uint32 // 渠道号，由前端生成，不需要填写
		Serial      string // 流水号，由前端生成，不需要填写
	} `json:"body"`
}

// 封号应答
type IDIP_DO_SEND_MAIL_ALL_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
		Result int32  // 结果：（0）成功，（其他）失败
		RetMsg string // 返回消息

	} `json:"body"`
}

func main() {
	client := &http.Client{}

	var request_par IDIP_DO_SEND_MAIL_ALL_REQ

	request_par.Head.PacketLen = 0
	request_par.Head.Cmdid = 4161
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
	request_par.Body.Source = 1
	request_par.Body.Serial = "11"
	request_par.Body.MailTitle = "Boy"
	request_par.Body.MailContent = "Hello"
	request_par.Body.SendTime = 1414074000
	request_par.Body.EndTime = 1414074600
	request_par.Body.ItemId1 = 31
	request_par.Body.ItemNum1 = 1
	request_par.Body.ItemId2 = 350
	request_par.Body.ItemNum2 = 2
	request_par.Body.ItemId3 = 242
	request_par.Body.ItemNum3 = 50
	request_par.Body.ItemId4 = 0
	request_par.Body.ItemNum4 = 0
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
	req2 := new(IDIP_DO_SEND_MAIL_ALL_RSP)
	err = json.Unmarshal(body, req2)
	fmt.Println(req2)

}
