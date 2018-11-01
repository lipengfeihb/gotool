package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

//生成32位md5字串
func Md5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

//
func Sha1(data []byte) string {
	hash := sha1.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}
