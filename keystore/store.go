package keystore

import (
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
	log.Println("get key:", key)
	value, ok := keyStoreMap[key]
	if ok {
		return value
	} else {
		log.Println("key: " + key + ", not found")
		return "404"
	}
}

func Addkey(Keyname string, Value string) int {
	if _, isKeyExists := keyStoreMap[Keyname]; isKeyExists {
		log.Println("key already exists:" + Keyname)
		return 409
	} else {
		keyStoreMap[Keyname] = Value
		log.Println("add key:", Keyname)
		return 200
	}
}
