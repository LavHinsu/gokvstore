package keystore

import (
	"log"
	"time"
)

/*
function used to update a key in our store
returns a int. 200 if successfully, 409 if the key already exists.
use the UpdateKey to update a key
*/
func Addkey(Keyname string, Value string, E_AT int64, ipaddr string) int {
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
			E_AT:  E_AT,              // update expires at
			U_AT:  0,                 // intialize with an empty time object
		}
		log.Println(ipaddr+" added key:", Keyname)
		return 200
	}
}
