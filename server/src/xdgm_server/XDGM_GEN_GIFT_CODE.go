package xdgm_server

import (
	"core/fail"
	"core/log"
	"encoding/json"
	"fmt"
	"game_server/dat/mail_dat"
	"net/http"
	"xdgm_server/gift_code"
	"xdgm_server/rpc"
)

type DataConfig struct {
	ItemType int16 `json:"item_type"` //道具类别
	ItemId   int16 `json:"item_id"`   // 赠送道具ID
	ItemNum  int64 `json:"num"`       // 赠送道具数量
}

func (this *XDGM_GEN_GIFT_CODE_REQ) Process() (XDGM_RSP, error) {
	rsp := XDGM_RSP{}
	var allGen bool
	if len(this.Servers) == 0 {
		allGen = true
		this.Servers = append(this.Servers, 0)
	}
	codes := make([]gift_code.GiftCode, 0, len(this.Servers)*int(this.Num))
	version := gift_code.GetVersion()
	config := JsonToConfig(this.Config)
	var err error
	for _, serverId := range this.Servers {
		newCodes, err := gift_code.Gen(gift_code.CodeType(this.Type), config, this.EffectTimestamp, this.ExpireTimestamp, serverId, this.Num, version, this.Content)
		if err != nil {
			log.Errorf("XDGM_GEN_GIFT_CODE_REQ Gen error:", err)
			break
		}
		codes = append(codes, newCodes...)
	}
	rsp.Data = codes
	if err != nil {
		rsp.Status = 0
		rsp.Message = fmt.Sprintf("%v", err)
		return rsp, err
	}
	go func() {
		// 生成完之后rpc通知相应服务器重新加载激活码数据
		if !allGen {
			for _, serverId := range this.Servers {
				hdId := Sid2GlobalServerId(serverId)
				rpc.RemoteXdgmReloadGiftCode(hdId, func(*rpc.Reply_XdgmReloadGiftCode, error) {
					if err != nil {
						log.Errorf("XDGM_GEN_GIFT_CODE_REQ rpc error:", err)
					}
				})
			}
		} else {
			servers := GetAllGsServerByHttp(0)
			for _, server := range servers {
				if server.HD == true {
					rpc.RemoteXdgmReloadGiftCode(server.GSID, func(*rpc.Reply_XdgmReloadGiftCode, error) {
						if err != nil {
							log.Errorf("XDGM_GEN_GIFT_CODE_REQ rpc error:", err)
						}
					})
				}
			}
		}
	}()
	rsp.Status = 1
	rsp.Message = "success"

	//rsp.Data = failpids
	return rsp, nil
}

func (this *XDGM_GEN_GIFT_CODE_REQ) Load(req *http.Request) (err error) {
	err = json.Unmarshal([]byte(req.PostFormValue("data")), this)
	if err != nil {
		return err
	}
	return nil
}

func JsonToConfig(reqconfig []GIFT_CODE_CONFIG) string {
	dataConfig := make([]DataConfig, 0)
	for _, v := range reqconfig {
		switch v.ItemId {
		//爱心
		case 247:
			dataConfig = append(dataConfig, DataConfig{
				ItemType: mail_dat.ATTACHMENT_HEART,
				ItemId:   v.ItemId,
				ItemNum:  v.ItemNum,
			})
		//铜钱
		case 244:
			dataConfig = append(dataConfig, DataConfig{
				ItemType: mail_dat.ATTACHMENT_COINS,
				ItemId:   v.ItemId,
				ItemNum:  v.ItemNum,
			})
		case 242:
			dataConfig = append(dataConfig, DataConfig{
				ItemType: mail_dat.ATTACHMENT_INGOT,
				ItemId:   v.ItemId,
				ItemNum:  v.ItemNum,
			})
		case 243:
			break
		case 245:
			break
		case 246:
			break
		case 248:
			break
		default:
			dataConfig = append(dataConfig, DataConfig{
				ItemType: mail_dat.ATTACHMENT_ITEM,
				ItemId:   v.ItemId,
				ItemNum:  v.ItemNum,
			})
		}
	}
	data, err := json.Marshal(dataConfig)
	fail.When(err != nil, fmt.Sprint("json to config string wrong"))
	return string(data)
}
