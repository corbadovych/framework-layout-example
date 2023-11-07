package stdlib

import (
	"github.com/valyala/fasthttp"
)

func GetRemoteAddress(req *fasthttp.RequestCtx) (string, error) {
	ip := string(req.Request.Header.Peek("X-Forwarded-For"))
	if ip == "" {
		ip = req.RemoteIP().String()
	}

	return ip, nil
}
