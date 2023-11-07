package stdlib

import "net/http"

func GetRemoteAddress(req *http.Request) (string, error) {
	ip := req.Header.Get("X-Forwarded-For")
	if ip != "" {
		return ip, nil
	}

	return req.RemoteAddr, nil
}
