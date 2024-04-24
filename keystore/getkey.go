package keystore

import "log"

// function used to get a key from our store
func Getkey(key string, ipaddr string) *keyStoreMap {
	mutex.RLock()
	defer mutex.RUnlock()
	value, ok := keyStore[key]
	if ok {
		log.Println(ipaddr + " get key: " + key)
		return value
	} else {
		log.Println(ipaddr + " key not found: " + key)
		return nil
	}
}
