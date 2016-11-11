package idip_server

func (req *IDIP_QUERY_TASK_INFO_REQ) Process() (IDIP_QUERY_TASK_INFO_RSP, error) {
	rsp := IDIP_QUERY_TASK_INFO_RSP{}
	rsp.Head.Cmdid = req.Head.Cmdid + 1
	rsp.Body.TaskProgressList_count = 0
	rsp.Body.TaskProgressList = append(rsp.Body.TaskProgressList, STaskProgressInfo{})
	return rsp, nil
}
