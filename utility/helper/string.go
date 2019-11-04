package helper

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// 获取MD5
func MD5(str string) string {
	_md5 := md5.New()
	_md5.Write([]byte(str))
	return hex.EncodeToString(_md5.Sum([]byte(nil)))
}

func MD5FILE(filepath string) string {
	f, _ := os.Open(filepath)
	defer f.Close()

	_md5 := md5.New()
	io.Copy(_md5, f)
	return hex.EncodeToString(_md5.Sum([]byte(nil)))
}

// 获取SHA1
func SHA1(str string) string {
	_sha1 := sha1.New()
	_sha1.Write([]byte(str))
	return hex.EncodeToString(_sha1.Sum([]byte(nil)))
}

// 获取SHA256
func SHA256(str string) string {
	_sha256 := sha256.New()
	_sha256.Write([]byte(str))
	return hex.EncodeToString(_sha256.Sum([]byte(nil)))
}

// 获取HMAC
func HMAC(key, data string) string {
	_hmac := hmac.New(md5.New, []byte(key))
	_hmac.Write([]byte(data))
	return hex.EncodeToString(_hmac.Sum([]byte(nil)))
}

func EmptyValue(v interface{}) bool{
	switch t := v.(type) {
	case string:
		return t == ""
	case int:
		return t == 0
	case int64:
		return t == 0
	default:
		return true
	}
}

func StrToInt(s string, def int) int{
	i, err := strconv.Atoi(s)

	if err != nil {
		return def
	}

	return i
}

// 合并字符串
func StrJoin(sep string, e... string) string {
	return strings.Join(e, sep)
}

func RandNumber(width int) string {
	rand.Seed(time.Now().UnixNano())

	var res string
	for i := 0; i < width; i++ {
		res += strconv.Itoa(rand.Intn(10))
	}
	return res
}

// 统一切掉账号的0
func RepairPhone(cc, phone string) (string, string) {
	if cc != "" && cc != "241" {
		return cc, strings.TrimLeft(phone, "0")
	}

	return cc, phone
}
