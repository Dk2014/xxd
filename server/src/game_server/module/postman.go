package module

import (
	"core/debug"
	"core/log"
	"core/net"
	util "core/time"
	"game_server/api/protocol/notify_api"
	"game_server/mdb"
	"sync"
	"time"
)

var postman *postMan

const (
	GLOBAL_POST_UNIT_TYPE_MAIL = iota
	GLOBAL_POST_UNIT_TYPE_ANNOUNCEMENT
	GLOBAL_DELETE_UNIT_TYPE_ANNOUNCEMENT
)

type PostUnit struct {
	Type      int8         // 投递类型
	StartTime int64        // 生效时间
	EndTime   int64        // 失效时间
	Response  net.Response // 投递的消息
	Used      bool         //是否使用
	Id        int64        //关联数据库Id
}

type postMan struct {
	sync.Mutex
	postChan chan *PostUnit // 投递队列（包含了走马灯公告和全局邮件）
	cylinder []*PostUnit    // (邮件、公告)筒
}

// 在主线程main中调用
func PostManStart() {
	postman = &postMan{
		postChan: make(chan *PostUnit, 1000),
		cylinder: make([]*PostUnit, 0, 100),
	}

	postman.loopWork()

	mdb.Transaction(mdb.TRANS_TAG_GlobalPostMan, func() {
		mdb.GlobalExecute(func(db *mdb.Database) {

			nowTime := util.GetNowTime()
			mailResponse := &notify_api.SendGlobalMail_Out{}

			var unit *PostUnit

			db.Select.GlobalMail(func(row *mdb.GlobalMailRow) {
				startTime := row.SendTime()
				expireTime := row.ExpireTime()

				// 丢弃过期邮件
				if nowTime >= expireTime {
					db.Select.GlobalMailAttachments(func(attacRow *mdb.GlobalMailAttachmentsRow) {
						if attacRow.GlobalMailId() == row.Id() {
							db.Delete.GlobalMailAttachments(attacRow.GoObject())
						}

					})
					db.Delete.GlobalMail(row.GoObject())
					return
				}

				unit = &PostUnit{
					Id:        row.Id(),
					Type:      GLOBAL_POST_UNIT_TYPE_MAIL,
					StartTime: startTime,
					EndTime:   expireTime,
					Response:  mailResponse,
					Used:      true,
				}

				if startTime > nowTime {
					postman.doPost(unit, startTime-nowTime)
				} else {
					postman.addUnit(unit)
				}
			})

			db.Select.GlobalAnnouncement(func(row *mdb.GlobalAnnouncementRow) {
				startTime := row.SendTime()
				expireTime := row.ExpireTime()

				// 丢弃过期公告
				if nowTime >= expireTime {
					db.Delete.GlobalAnnouncement(row.GoObject())
					return
				}

				unit = &PostUnit{
					Id:        row.Id(),
					Type:      GLOBAL_POST_UNIT_TYPE_ANNOUNCEMENT,
					StartTime: startTime,
					EndTime:   expireTime,
					Used:      true,
					Response: &notify_api.SendAnnouncement_Out{
						Id:          row.Id(),
						TplId:       row.TplId(),
						ExpireTime:  expireTime,
						Parameters:  []byte(row.Parameters()),
						Content:     []byte(row.Content()),
						SpacingTime: int32(row.SpacingTime()),
					},
				}

				if startTime > nowTime {
					postman.doPost(unit, startTime-nowTime)
				} else {
					postman.addUnit(unit)
				}
			})

		})
	})

}

func (this *postMan) addUnit(unit *PostUnit) {
	this.Lock()
	defer this.Unlock()

	this.cylinder = append(this.cylinder, unit)
}

func (this *postMan) doPost(unit *PostUnit, delay int64) {
	if delay < 0 {
		return
	}

	time.AfterFunc(time.Duration(delay)*time.Second, func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf(`PostMan::doPost error
Error = %v
Stack =
%s`,
					err,
					debug.Stack(1, "    "),
				)
			}
		}()
		unit.Used = true
		this.addUnit(unit)
		this.doPush(unit)
	})
}

func (this *postMan) doPush(unit *PostUnit) {
	select {
	case this.postChan <- unit:
	default:
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Errorf("[postMan] doPush failed: %v", err)
				}
			}()
			this.postChan <- unit
		}()
	}
}

func (this *postMan) doDelete(unit *PostUnit) {
	switch unit.Type {
	case GLOBAL_DELETE_UNIT_TYPE_ANNOUNCEMENT:
		for _, v := range this.cylinder {
			if v.Type == GLOBAL_POST_UNIT_TYPE_ANNOUNCEMENT && v.Id == unit.Id {
				v.Used = false
			}
		}
	}
	//post给客户端，通知消除
	unit.Used = true
	this.addUnit(unit)
	this.doPush(unit)
}

func (this *postMan) loopWork() {
	go func() {

		defer func() {
			if err := recover(); err != nil {
				log.Errorf(`PostMan::loop error
Error = %v
Stack =
%s`,
					err,
					debug.Stack(1, "    "),
				)
			}
		}()

	L:
		for {
			select {
			case unit, ok := <-this.postChan:
				if !ok {
					break L
				}
				if unit.Used == true {
					API.Broadcast(Player, unit.Response)
				} else {
					break L
				}

			case _, ok := <-time.After((7 * 86400) * time.Second): // 7day后定时清理失效的信息
				if !ok {
					break L
				}

				this.Lock()
				//defer this.Unlok()

				nowTime := util.GetNowTime()
				newCylinder := make([]*PostUnit, 0, len(this.cylinder))

				//var expireMail []int64
				var expireAnn []int64
				for _, unit := range this.cylinder {
					if nowTime >= unit.EndTime {
						switch unit.Type {
						case GLOBAL_POST_UNIT_TYPE_MAIL:
							//expireMail = append(expireMail, unit.Id)
						case GLOBAL_POST_UNIT_TYPE_ANNOUNCEMENT:
							expireAnn = append(expireAnn, unit.Id)
						}
						continue
					}

					newCylinder = append(newCylinder, unit)
				}

				mdb.Transaction(mdb.TRANS_TAG_GlobalPostMan, func() {
					mdb.GlobalExecute(func(db *mdb.Database) {
						/*
							for _, id := range expireMail {
								mail := db.Lookup.GlobalMail(id)
								if mail != nil {
									db.Delete.GlobalMail(mail)
								}
							}
						*/

						for _, id := range expireAnn {
							ann := db.Lookup.GlobalAnnouncement(id)
							if ann != nil {
								db.Delete.GlobalAnnouncement(ann)
							}
						}
					})
				})

				this.cylinder = newCylinder
				this.Unlock()
			}
		}
	}()
}
