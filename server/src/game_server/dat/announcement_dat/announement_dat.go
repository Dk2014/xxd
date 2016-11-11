package announcement_dat

// 公告接口
type Announcer interface {
	GetAnnouncementTplId() int32
	GetParameters() string
	GetDuration() int64
}
