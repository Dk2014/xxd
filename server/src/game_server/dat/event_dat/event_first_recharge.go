package event_dat

//首充返利活动
type EventFirstRecharge struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
}
