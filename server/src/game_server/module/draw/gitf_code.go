package draw

import (
	"core/i18l"
	"core/net"
	"core/time"
	"game_server/api/protocol/draw_api"
	//"game_server/config"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
	//"game_server/module/mail"
	"game_server/dat/mail_dat"
	"game_server/module/rpc"
	"game_server/xdlog"
)

func exchangeGiftCode(session *net.Session, code string) {
	state := module.State(session)
	dupExchange := false
	state.Database.Select.PlayerGlobalGiftCardRecord(func(row *mdb.PlayerGlobalGiftCardRecordRow) {
		if row.Code() == code {
			dupExchange = true
			row.Break()
		}
	})
	if dupExchange {
		session.Send(&draw_api.ExchangeGiftCode_Out{
			Result: draw_api.EXCHANGE_GIFT_RESULT_DUP_EXCHANGE,
		})
		return
	}
	now := time.GetNowTime()
	info, ok := global.TakeGift(code)
	if !ok || info.EffectTimestamp > now || info.ExpireTimestamp < now {
		session.Send(&draw_api.ExchangeGiftCode_Out{
			Result: draw_api.EXCHANGE_GIFT_RESULT_EXPIRE,
		})
		return
	}
	var attach []*mail_dat.Attachment
	if info.ItemId > 0 {
		attach = append(attach, &mail_dat.Attachment{mail_dat.ATTACHMENT_ITEM, info.ItemId, 1})
	}
	for _, cfg := range info.Config {
		attach = append(attach, &mail_dat.Attachment{cfg.ItemType, cfg.ItemId, cfg.Num})
	}
	mail := &mail_dat.EmptyMail{
		Title:       i18l.T.Tran("兑换码领取成功"),
		Content:     i18l.T.Tran(info.Content),
		Attachments: attach,
	}
	session.Send(&draw_api.ExchangeGiftCode_Out{
		Result: draw_api.EXCHANGE_GIFT_RESULT_SUCCESS,
	})
	xdlog.GiftCodeLog(state.Database, code, int32(info.Type))
	state.Database.Insert.PlayerGlobalGiftCardRecord(&mdb.PlayerGlobalGiftCardRecord{
		Pid:  state.PlayerId,
		Code: code,
	})
	if info.Type == global.GIFT_CODE_MONOP_TYPE {
		state.Database.Insert.GlobalGiftCardRecord(&mdb.GlobalGiftCardRecord{
			Pid:       state.PlayerId,
			Code:      code,
			Timestamp: now,
		})
	}
	rpc.RemoteMailSend(state.PlayerId, mail)
}
