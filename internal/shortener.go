package internal

import (
	"crypto/md5"
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
	hash := string(md5.New().Sum([]byte(link)))
	key := hash[:16]
	for reserved.includes(key) {
		key = key[1:] + string(hash[rand.Intn(len(hash))])
	}
	return key
}
