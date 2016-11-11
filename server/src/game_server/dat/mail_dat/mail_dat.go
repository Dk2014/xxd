package mail_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapMail map[int32]*Mail
)

// 系统邮件接口
type Mailer interface {
	GetMailId() int32
	GetParameters() string
	GetTitle() string
	GetContent() string
	GetAttachments() []*Attachment
	GetSendTime() int64
	GetExpireTime() int64
	GetPriority() int8
	GetMinLevel() int16
	GetMaxLevel() int16
	GetMinVIPLevel() int16
	GetMaxVIPLevel() int16
}

func Load(db *mysql.Connection) {
	LoadMail(db)
}

type Mail struct {
	Id      int32
	Title   string
	Content string
}

func LoadMail(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM mail ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iTitle := res.Map("title")
	iContent := res.Map("content")

	var pri_id int32
	mapMail = map[int32]*Mail{}
	for _, row := range res.Rows {
		pri_id = row.Int32(iId)
		mail := &Mail{
			Id:      pri_id,
			Title:   row.Str(iTitle),
			Content: row.Str(iContent),
		}
		mapMail[pri_id] = mail
	}
}

func GetMail(mailId int32) (title, content string) {
	mail := mapMail[mailId]
	fail.When(mail == nil, "wrong mailId")
	return mail.Title, mail.Content
}

// 附件
type Attachment struct {
	AttachmentType int8
	ItemId         int16
	ItemNum        int64
}

type EmptyMail struct {
	MailId      int32
	Title       string
	Content     string
	Parameters  string
	Attachments []*Attachment
	SendTime    int64
	ExpireTime  int64
	Priority    int8
	MinLevel    int16 //0不限制
	MaxLevel    int16 //0不限制
	MinVIPLevel int16 //0不限制
	MaxVIPLevel int16 //0不限制
}

func (this *EmptyMail) GetMailId() int32 {
	return this.MailId
}

func (this *EmptyMail) GetSendTime() int64 {
	return this.SendTime
}

func (this *EmptyMail) GetExpireTime() int64 {
	return this.ExpireTime
}

func (this *EmptyMail) GetPriority() int8 {
	return this.Priority
}

func (this *EmptyMail) GetTitle() string {
	return this.Title
}

func (this *EmptyMail) GetContent() string {
	return this.Content
}

func (this *EmptyMail) GetParameters() string {
	return this.Parameters
}

func (this *EmptyMail) GetAttachments() []*Attachment {
	return this.Attachments
}

func (this *EmptyMail) GetMinLevel() int16 {
	return this.MinLevel
}

func (this *EmptyMail) GetMaxLevel() int16 {
	return this.MaxLevel
}

func (this *EmptyMail) GetMinVIPLevel() int16 {
	return this.MinVIPLevel
}

func (this *EmptyMail) GetMaxVIPLevel() int16 {
	return this.MaxVIPLevel
}
