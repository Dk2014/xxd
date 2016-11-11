package idip_server

import "errors"

func (req *{{.Req}}) Process() ({{.Rsp}}, error) {
	return {{.Rsp}}{}, errors.New("{{.Req}} Process sample")
}
