package common

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
	"net/url"
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseInt(value string) int {
	if value == "" {
		return 0
	}
	val, _ := strconv.Atoi(value)
	return val
}

func IntString(value int) string {
	return strconv.Itoa(value)
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Atoa(str string) string {
	var result string
	for i := 0; i < len(str); i++ {
		c := rune(str[i])
		if 'A' <= c && c <= 'Z' && i > 0 {
			result = result + "_" + strings.ToLower(string(str[i]))
		} else {
			result = result + string(str[i])
		}
	}
	return result
}

func Urlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

func GetRemoteIp(r *http.Request) (ip string) {
	ip = r.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = r.RemoteAddr
	}
	ip = strings.Split(ip, ":")[0]
	if len(ip) < 7 || ip == "127.0.0.1" {
		ip = "localhost"
	}
	return
}

func GetToken() string {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	rndNum := rand.Int63()
	uuid := Md5(Md5(strconv.FormatInt(nano, 10)) + Md5(strconv.FormatInt(rndNum, 10)))
	return uuid
}
