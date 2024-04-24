package keystore

import "log"

// function used to delete a key in our store
func DeleteKey(Keyname string, ipaddr string) int {
	mutex.Lock()
	defer mutex.Unlock()
	_, KeyExists := keyStore[Keyname]
	if KeyExists {
		delete(keyStore, Keyname)
		log.Println(ipaddr+" deleted key:", Keyname)
		return 200
	} else {
		log.Println(ipaddr+" couldn't find key to delete:", Keyname)
		return 404
	}
}
