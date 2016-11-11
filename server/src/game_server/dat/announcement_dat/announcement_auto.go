package announcement_dat

import (
	"strconv"
	"bytes"
	
)

const( 
	MODULE_ANNOUNCEMENT = 0
	BACKSTAGE_ANNOUNCEMENT = 1
)

type AnnouncementConfig struct {
	TplId  int64
	Params  []string
	ShowTiming  int64
}

var StaticAnnouncements = []AnnouncementConfig{
}





// 测试公告
type AnnounceTestAnnounce struct {
	P1 string // 参数一
}

func (this AnnounceTestAnnounce) GetAnnouncementTplId() int32 {
	return int32(1)
}

func (this AnnounceTestAnnounce) GetDuration() int64 {
	return int64(0)
}

func (this AnnounceTestAnnounce) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.P1))
	b.WriteString("]")
	return string(b.Bytes())
}

// 测试公告2
type AnnounceTestAnnounce2 struct {
}

func (this AnnounceTestAnnounce2) GetAnnouncementTplId() int32 {
	return int32(2)
}

func (this AnnounceTestAnnounce2) GetDuration() int64 {
	return int64(0)
}

func (this AnnounceTestAnnounce2) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

// 测试公告3
type AnnounceTestAnnounce3 struct {
}

func (this AnnounceTestAnnounce3) GetAnnouncementTplId() int32 {
	return int32(3)
}

func (this AnnounceTestAnnounce3) GetDuration() int64 {
	return int64(0)
}

func (this AnnounceTestAnnounce3) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

// 巡游商人出现公告
type AnnounceTraderShowupAnnounce struct {
	Timing string // 出现时间
	Disappear string // 消失时间
}

func (this AnnounceTraderShowupAnnounce) GetAnnouncementTplId() int32 {
	return int32(4)
}

func (this AnnounceTraderShowupAnnounce) GetDuration() int64 {
	return int64(0)
}

func (this AnnounceTraderShowupAnnounce) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Timing))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Disappear))
	b.WriteString("]")
	return string(b.Bytes())
}


