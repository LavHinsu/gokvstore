package keystore

import (
	"fmt"
	"log"
)

var (
	keyStoreMap = make(map[string]string)
)

func init() {
	keyStoreMap["foo"] = "bar"
	keyStoreMap["good"] = "stuff"
}

func Getkey(key string) string {
	log.Print("getting key", key)
	value, ok := keyStoreMap[key]
	if ok {
		return value
	} else {
		log.Printf("key not found")
		return "key not found"
	}
}

func Addkey() {
	fmt.Println("Inside addkey")
}
