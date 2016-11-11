package platform_server

import (
	"core/fail"
	"core/log"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"time"
)

type IAPI interface {
	Req() interface{}
	Validate(interface{})
	Process(interface{}) (interface{}, error)
}

type ReqBase struct {
	OpenId   string
	Type     uint8
	AreaId   uint8 // 大区（1微信，2手Q, 5游客）
	PlatId   uint8 // 平台（0ios，1安卓）
	Token    string
	PayToken string
	ClientIp string
	Version  int32
	App      string
}

func route(action string, reqBody io.Reader, w io.Writer, clientIp string) (err error) {
	body, err := ioutil.ReadAll(reqBody)
	fail.When(err != nil, err)
	logAccessInfo(clientIp, action, body)

	var prc IAPI

	switch action {
	case "/server/list", "/serverall":
		prc = procServerList{}
	case "/user/server", "/user/list":
		prc = procUserList{}
	case "/user/logon":
		prc = procUserLogon{}
	case "/user/create":
		prc = procUserCreate{}
	case "/user/gsinfo":
		prc = procUserGSInfo{}
	case "/user/update":
		prc = procUserUpdate{}
	case "/patch":
		prc = procClientPatch{}
	case "/totalResource":
		prc = procTotalResource{}
	case "/announce":
		prc = procAnnounce{}
	case "/systemInfo":
		prc = procSystemInfo{}
	case "/disableActionPic":
		prc = procDisableActionPic{}
	case "/stat", "/status":
		prc = procStatus{}
	case "/productInfo":
		prc = procProductInfo{}

	//xdgm api
	case "/gserverall":
		prc = procGServerList{}
	case "/announce/update":
		prc = procSetAnnounce{}
	case "/xdgm/addBlackIp":
		prc = procAddBlackIp{}
	case "/xdgm/removeBlackIp":
		prc = procRemoveBlackIp{}
	case "/xdgm/getIpBlackTable":
		prc = procGetIpBlackTable{}

	//wegame api
	case "/wegame/register":
		prc = procWGRegister{}
	case "/wegame/login":
		prc = procWGLogin{}
	case "/wegame/modifyPasswd":
		prc = procWGModifyPasswd{}
	case "/wegame/bindThird":
		prc = procWGBindThird{}
	case "/wegame/bindAccount":
		prc = procWGBindAccount{}

	case "/anysdk/loginOauth":
		TransportAnySDKOauth(body, w)
		return
	default:
		w.Write([]byte(`{"error":404,"msg":"the action is not found."}`))
		log.Debugf("unhandled request: %v", action)
		return
	}

	startTime := time.Now().UnixNano()

	// parse reqBody as json, throw error on error
	in := prc.Req()
	err = json.Unmarshal(body, in)
	fail.When(err != nil, "Unmarshal request-data error")

	prc.Validate(in)

	v := reflect.ValueOf(in).Elem() // .FieldByName("ClientIp").SetString(clientIp)

	if v.Kind() == reflect.Struct {
		fClientIp := v.FieldByName("ClientIp")

		if fClientIp.IsValid() && fClientIp.CanSet() && fClientIp.Kind() == reflect.String {
			fClientIp.SetString(clientIp)
		}
		fApp := v.FieldByName("App")
		app := fApp.String()
		if app == "" && fApp.IsValid() && fApp.CanSet() && fApp.Kind() == reflect.String {
			fApp.SetString("xxd_qq")
		}

		fType := v.FieldByName("Type")
		if app == "xxd_vn" && fType.IsValid() && fType.CanSet() && fType.Kind() == reflect.Uint8 {
			fType.SetUint(1)
		}

		//check ip
		if isInIpBlackTable(app, clientIp) {
			w.Write([]byte(`{"error":400,"msg":"your ip is forbidden."}`))
			return
		}
	}

	out, err := prc.Process(in)

	if err != nil {
		log.Errorf("Process error %v", err)
	}

	if out != nil {
		// write marshaled json of rsp
		b, err := json.Marshal(out)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)
	}

	endTime := time.Now().UnixNano()
	log.Infof("[access]%s - - %s used time: %dms\n", clientIp, action, (endTime-startTime)/1000000)
	return
}

func logAccessInfo(ip, action string, body []byte) {
	log.Infof(fmt.Sprintf("[access]%v - - [%v] \"%v %s\"\n", ip, time.Now(), action, body))
}
