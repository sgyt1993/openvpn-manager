package commonresp

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JsonResult
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

func JudgeError(c *gin.Context, data interface{}, err error) {
	if err != nil {
		if len(err.Error()) == 0 {
			c.JSON(http.StatusOK, Failed("system is error"))
		} else {
			c.JSON(http.StatusOK, Failed(err.Error()))
		}
	} else {
		c.JSON(http.StatusOK, OK(data))
	}
}
