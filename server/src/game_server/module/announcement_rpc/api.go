package announcement_rpc

import (
	"core/time"

	"core/net"
	"game_server/api/protocol/announcement_api"
	"game_server/mdb"
	"game_server/module"
)

type AnnouncementAPI struct{}

func init() {
	announcement_api.SetInHandler(AnnouncementAPI{})
}

func (m AnnouncementAPI) GetList(session *net.Session, in *announcement_api.GetList_In) {
	timestamp := time.GetNowTime()
	out := &announcement_api.GetList_Out{}

	module.State(session).Database.Select.GlobalAnnouncement(func(row *mdb.GlobalAnnouncementRow) {
		if row.ExpireTime() < timestamp {
			return
		}

		out.Announcements = append(out.Announcements, announcement_api.GetList_Out_Announcements{
			Id:          row.Id(),
			TplId:       row.TplId(),
			ExpireTime:  row.ExpireTime(),
			Parameters:  []byte(row.Parameters()),
			Content:     []byte(row.Content()),
			SpacingTime: int32(row.SpacingTime()),
		})
	})

	session.Send(out)
}
