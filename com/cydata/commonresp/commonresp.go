package commonresp

import (
	"encoding/json"
	"net/http"
)

type JsonResult struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

const (
	okCode     = 200
	failedCode = 500
)

func CommonOK() JsonResult {
	return JsonResult{
		Code: okCode,
		Msg:  "",
	}
}

func OK(msg interface{}) JsonResult {
	return JsonResult{
		Code: okCode,
		Msg:  msg,
	}
}

func Failed(msg interface{}) JsonResult {
	return JsonResult{
		Code: failedCode,
		Msg:  msg,
	}
}

func CommonFailed() JsonResult {
	return JsonResult{
		Code: failedCode,
		Msg:  "",
	}
}

func JsonRespOK(w http.ResponseWriter, data interface{}) {
	w.Header().Set("content-type", "text/json")
	if dataString, ok := data.(string); ok {
		resp := OK(dataString)
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
	} else {
		resp := OK(data)
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
	}

}

func JsonRespFail(w http.ResponseWriter, msg string) {
	w.Header().Set("content-type", "text/json")
	resp := Failed(msg)
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
