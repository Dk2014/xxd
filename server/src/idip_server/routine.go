// this is auto-genrated file,
// Don't modify this file manually
package idip_server

import (
	"bytes"
	"core/fail"
	"core/log"
	"encoding/json"
	"io"
	"io/ioutil"
)

func route(reqBody io.Reader, w io.Writer) error {

	body, err := ioutil.ReadAll(reqBody)
	fail.When(err != nil, err)

	// remove data_packet=
	if bytes.HasPrefix(body, []byte("data_packet=")) {
		body = body[12:]
	}

	log.Infof("idip request parameter: %s", body)
	// parse for cmdid
	cmd := new(IDIP_COMMON_REQ)
	err = json.Unmarshal(body, cmd)
	if err != nil {
		log.Errorf("Unmarshal request-command error %v", err)
		return err
	}

	log.Debugf("process request: %v", cmd.Head)

	switch cmd.Head.Cmdid {

	case 0x1003: //封号请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_BAN_USR_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1005: //解封请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_UNBAN_USR_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1007: //删除道具请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_DEL_ITEM_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1009: //修改角色经验请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_UPDATE_ROLE_EXP_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x100b: //修改用户元宝请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_UPDATE_ROLE_GOLD_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x100d: //修改用户体力请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_UPDATE_ROLE_PHYSICAL_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x100f: //修改角色铜钱请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_UPDATE_ROLE_COIN_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1013: //修改角色爱心请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_UPDATE_ROLE_HEART_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1017: //发放道具请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_SEND_ITEM_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1019: //查询当前个人信息请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_QUERY_USR_INFO_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x101d: //查询角色装备信息请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_QUERY_ROLE_EQUIP_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1021: //查询角色灵宠信息请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_QUERY_ROLE_PET_INFO_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1023: //查询剑心信息请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_QUERY_SWORD_INFO_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1025: //查询魂侍信息请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_QUERY_SOUL_INFO_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1027: //查询任务进度请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_QUERY_TASK_INFO_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1029: //查询背包存量信息请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_QUERY_BAG_INFO_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x102b: //发放邮件请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_SEND_MAIL_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x102d: //更新新闻公告请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_UPDATE_NOTICE_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x102f: //更新游戏内走马灯请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_UPDATE_GAME_LAMP_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1037: //查询比武场请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_QUERY_COMPETE_INFO_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x103f: //清零数据请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_CLEAR_DATA_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1039: //设置魂侍等级请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_SET_SOUL_LEVEL_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x103b: //设置剑心等级请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_SET_SWORD_LEVEL_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x103d: //设置宠物激活请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_SET_PET_ACTIVE_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1041: //发放全区全服邮件请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_SEND_MAIL_ALL_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1043: //查询走马灯公告请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_QUERY_GAME_LAMP_NOTICE_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1045: //删除公告请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_DEL_NOTICE_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1047: //修改vip等级请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_DO_UPDATE_VIP_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1049: //查询深渊挂关卡进度请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_QUERY_ABYSSPASS_STATE_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x104b: //查询战力请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_QUERY_FIGHT_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1031: //发送邮件（AQ）请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_AQ_DO_SEND_MAIL_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1033: //封号（AQ）请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_AQ_DO_BAN_USR_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	case 0x1035: //解除处罚（AQ）请求
		// parse reqBody as json, throw error on error
		req := new(IDIP_AQ_DO_RELIEVE_PUNISH_REQ)
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)

	default:
		log.Debugf("unhandled request: %v", cmd.Head.Cmdid)
		return nil
	}

	return nil
}
