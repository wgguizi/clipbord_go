package request

import (
	"context"
	"net/http"
)

const (
	extTempleteDir = 1
)

// 设置上下文变量
func setContexData(r *http.Request, key any, val any) *http.Request {
	c2 := context.WithValue(r.Context(), key, val)
	r2 := r.WithContext(c2)
	return r2
}

// 设置模板目录
func SetTemplateDir(r *http.Request, val string) *http.Request {
	return setContexData(r, extTempleteDir, val)
}

// 获取模板目录
func GetTemplateDir(r *http.Request) string {
	return r.Context().Value(extTempleteDir).(string)
}

// 获取Ip
func GetIp(r *http.Request) string {
	ip := r.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
}
