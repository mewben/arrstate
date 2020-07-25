package utils

import "strings"

// GetSubdomain -
// example url: http://melvin.arrstate.lh:8000
func GetSubdomain(url string) string {
	sp := strings.Split(url, ".")
	if len(sp) < 0 {
		return ""
	}
	sp = strings.Split(sp[0], "/")
	if len(sp) < 0 {
		return sp[0]
	}

	return sp[len(sp)-1]

}
