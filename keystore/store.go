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
	_, KeyExists := keyStoreMap[Keyname]
	if KeyExists {
		log.Println("key already exists:" + Keyname)
		return 409
	} else {
		keyStoreMap[Keyname] = Value
		log.Println("added key:", Keyname)
		return 200
	}
}

func UpdateKey(Keyname string, Value string) int {
	_, KeyExists := keyStoreMap[Keyname]
	if KeyExists {
		keyStoreMap[Keyname] = Value
		log.Println("updated key:", Keyname, Value)
		return 200
	} else {
		log.Println("couldn't find key in update key call:", Keyname)
		return 404
	}
}

func DeleteKey(Keyname string) int {
	_, KeyExists := keyStoreMap[Keyname]
	if KeyExists {
		delete(keyStoreMap, Keyname)
		log.Println("deleted key:", Keyname)
		return 200
	} else {
		return 404
	}
}
