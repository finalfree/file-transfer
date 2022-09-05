package service

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func GetIpFromRequest(r *http.Request) string {
	realIp := r.Header.Get("X-Real-Ip")
	if realIp == "" {
		realIp = r.Header.Get("X-Forwarded-For")
		if realIp == "" {
			realIp, _, _ = net.SplitHostPort(r.RemoteAddr)
		}
	}
	return realIp
}

func PrintSenderAndTime(r *http.Request) {
	fmt.Println(GetIpFromRequest(r), time.Now())
}
