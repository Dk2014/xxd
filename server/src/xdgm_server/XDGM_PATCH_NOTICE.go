package xdgm_server

import ()

func (this *XDGM_PATCH_NOTICE_REQ) Process(platid uint8) (XDGM_RSP, error) {
	rsp := XDGM_RSP{}
	errorCode, message := SetAnnounce(this.Content, this.Time, platid)
	if errorCode == 0 {
		rsp.Status = 1
		rsp.Message = "success"
	} else {
		rsp.Status = 0
		rsp.Message = message
	}
	return rsp, nil
}
