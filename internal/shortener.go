package internal

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

type strslice []string

func (s strslice) includes(value string) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

func shorten(link string, reserved strslice) string {
	hashBytes := md5.Sum([]byte(link))
	hash := hex.EncodeToString(hashBytes[:])
	key := hash[:16]
	for reserved.includes(key) {
		key = key[1:] + string(hash[rand.Intn(len(hash))])
	}
	return key
}
