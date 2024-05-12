package keystore

import (
	"log"
	"time"
)

// function used to update a key in our store
func UpdateKey(Keyname string, Value string, E_AT int64, ipaddr string) int {
	mutex.Lock()
	defer mutex.Unlock()
	_, KeyExists := keyStore[Keyname]
	if KeyExists {
		if E_AT != 0 { // only update the E_AT value if one is sent
			keyStore[Keyname].Value = Value
			keyStore[Keyname].U_AT = time.Now().Unix()
			keyStore[Keyname].E_AT = E_AT

		} else {
			keyStore[Keyname].Value = Value
			keyStore[Keyname].U_AT = time.Now().Unix()
		}
		log.Println(ipaddr+" updated key:", Keyname)
		return 200
	} else {
		log.Println(ipaddr+" couldn't find key to update:", Keyname)
		return 404
	}
}
