package controller

import (
	"encoding/json"
	"net/http"
	"viry_sun/lib/dao/clipboard"
	"viry_sun/lib/request"
	"viry_sun/lib/response"

	"github.com/julienschmidt/httprouter"
)

type OutObj struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
type OutGet struct {
	Content string `json:"content"`
}

// 首页控制器
// ps:name
func GetAct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	code := ps.ByName("name") //
	cData := clipboard.GetByCode(code)

	var outObj OutObj
	outObj.Code = 0
	if cData != nil {
		var outGet OutGet
		outGet.Content = cData.Content
		outObj.Code = 1000
		outObj.Data = outGet
	}
	bt, _ := json.Marshal(outObj)

	response.WriteJson(w, bt)
}

func SaveAct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	code := ps.ByName("name") //
	content := r.PostFormValue("content")

	//优先获取新任上游转发的IP，如nginx
	ip := request.GetIp(r)

	var cData clipboard.CData
	cData.Code = code
	cData.Content = content
	cData.Ip = ip

	var outObj OutObj
	err := clipboard.Save(&cData)
	if err != "" {
		outObj.Code = 0
		outObj.Msg = err
	} else {
		outObj.Code = 1000
	}

	bt, _ := json.Marshal(outObj)
	response.WriteJson(w, bt)
}
