package dummy_handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func init() {
	http.HandleFunc("/debug/dummy", dummyHandler)
}

type DummyResponse struct {
	Status int   `json:"status"`
	Result []int `json:"result"`
}

func dummyHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	rsp := DummyResponse{
		Status: 0,
		Result: []int{1, 2, 3, 4},
	}
	rspBytes, err := json.Marshal(rsp)
	if err == nil {
		w.Write(rspBytes)
	} else {
		io.WriteString(w, fmt.Sprintf(`{status:1, msg:"%s"}`, err.Error()))
	}
}
