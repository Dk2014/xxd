// this is auto-genrated file,
// Don't modify this file manually
package idip_server

import (
	"bytes"
	"core/fail"
	"core/log"
	"encoding/json"
	"io"
	"io/ioutil"
)

func route(reqBody io.Reader, w io.Writer) error {

	body, err := ioutil.ReadAll(reqBody)
	fail.When(err != nil, err)

	// remove data_packet=
	if bytes.HasPrefix(body, []byte("data_packet=")) {
		body = body[12:]
	}

	log.Infof("idip request parameter: %v", body)
	// parse for cmdid
	cmd := new(IDIP_COMMON_REQ)
	err = json.Unmarshal(body, cmd)
	if err != nil {
		log.Errorf("Unmarshal request-command error %v", err)
		return err
	}

	log.Debugf("process request: %v", cmd.Head)

	switch cmd.Head.Cmdid {
	{{with $mg:=.NET_CMD_ID}}
	{{range $m:=$mg.Macro}}
	case {{$m.Val}}: //{{$m.Desc}}
		// parse reqBody as json, throw error on error
		req:= new({{$m.Name}})
		err = json.Unmarshal(body, req)
		if err != nil {
			log.Errorf("Unmarshal request-data error %v", err)
			return err
		}

		rsp, err := req.Process()
		if err != nil {
			log.Errorf("process error %v", err)
			return err
		}

		// write marshaled json of rsp
		b, err := json.Marshal(rsp)
		if err != nil {
			log.Errorf("rsp marshal error: %v", err)
			return err
		}

		w.Write(b)
	{{end}}
	{{end}}
	default:
		log.Debugf("unhandled request: %v", cmd.Head.Cmdid)
		return nil
	}

	return nil
}

