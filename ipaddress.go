package main

import (
	"encoding/json"
	"net"
	"net/http"
)

func GetIpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	respBody, _ := json.Marshal(map[string]string{
		"ip": getIpFrom(r),
	})
	w.Write(respBody)
}

func getIpFrom(r *http.Request) string {
	realIp := r.Header.Get("X-Real-Ip")
	if realIp == "" {
		realIp = r.Header.Get("X-Forwarded-For")
		if realIp == "" {
			realIp, _, _ = net.SplitHostPort(r.RemoteAddr)
		}
	}
	return realIp
}
