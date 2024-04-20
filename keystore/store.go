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

func Getkey(key string, ipaddr string) string {
	value, ok := keyStoreMap[key]
	if ok {
		log.Println(ipaddr + " get key: " + key)
		return value
	} else {
		log.Println(ipaddr + " key not found: " + key)
		return "404"
	}
}

func Addkey(Keyname string, Value string, ipaddr string) int {
	_, KeyExists := keyStoreMap[Keyname]
	if KeyExists {
		log.Println(ipaddr + " key already exists: " + Keyname)
		return 409
	} else {
		keyStoreMap[Keyname] = Value
		log.Println(ipaddr+" added key:", Keyname)
		return 200
	}
}

func UpdateKey(Keyname string, Value string, ipaddr string) int {
	_, KeyExists := keyStoreMap[Keyname]
	if KeyExists {
		keyStoreMap[Keyname] = Value
		log.Println(ipaddr+" updated key:", Keyname)
		return 200
	} else {
		log.Println(ipaddr+" couldn't find key to update:", Keyname)
		return 404
	}
}

func DeleteKey(Keyname string, ipaddr string) int {
	_, KeyExists := keyStoreMap[Keyname]
	if KeyExists {
		delete(keyStoreMap, Keyname)
		log.Println(ipaddr+" deleted key:", Keyname)
		return 200
	} else {
		log.Println(ipaddr+" couldn't find key to delete:", Keyname)
		return 404
	}
}
