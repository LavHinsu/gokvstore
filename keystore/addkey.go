package keystore

import (
	"log"
	"time"
)

// function used to add a key to our store
func Addkey(Keyname string, Value string, ipaddr string) int {
	mutex.Lock()
	defer mutex.Unlock()
	_, KeyExists := keyStore[Keyname]
	if KeyExists {
		log.Println(ipaddr + " key already exists: " + Keyname)
		return 409
	} else {
		keyStore[Keyname] = &keyStoreMap{
			Value: Value,             //initalize with a value
			C_AT:  time.Now().Unix(), // intialize with an empty time object
			U_AT:  0,                 // intialize with an empty time object
		}
		log.Println(ipaddr+" added key:", Keyname)
		return 200
	}
}
