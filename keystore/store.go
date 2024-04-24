package keystore

import (
	"log"
	"time"
)

type keyStoreMap struct {
	Value string
	C_AT  time.Time
	U_AT  time.Time
}

var (
	keyStore = make(map[string]*keyStoreMap)
)

// initialize our store with a value
func init() {
	keyStore["foo"] = &keyStoreMap{
		Value: "bar",            //initalize with a value
		C_AT:  time.Now().UTC(), // intialize with an empty time object
		U_AT:  time.Time{},      // intialize with an empty time object
	}
}

// function used to get a key from our store
func Getkey(key string, ipaddr string) *keyStoreMap {
	value, ok := keyStore[key]
	if ok {
		log.Println(ipaddr + " get key: " + key)
		return value
	} else {
		log.Println(ipaddr + " key not found: " + key)
		return nil
	}
}

// function used to add a key to our store
func Addkey(Keyname string, Value string, ipaddr string) int {
	_, KeyExists := keyStore[Keyname]
	if KeyExists {
		log.Println(ipaddr + " key already exists: " + Keyname)
		return 409
	} else {
		keyStore[Keyname] = &keyStoreMap{
			Value: Value,            //initalize with a value
			C_AT:  time.Now().UTC(), // intialize with an empty time object
			U_AT:  time.Time{},      // intialize with an empty time object
		}
		log.Println(ipaddr+" added key:", Keyname)
		return 200
	}
}

// function used to update a key in our store
func UpdateKey(Keyname string, Value string, ipaddr string) int {
	_, KeyExists := keyStore[Keyname]
	if KeyExists {
		keyStore[Keyname].Value = Value
		keyStore[Keyname].U_AT = time.Now().UTC()
		log.Println(ipaddr+" updated key:", Keyname)
		return 200
	} else {
		log.Println(ipaddr+" couldn't find key to update:", Keyname)
		return 404
	}
}

// function used to delete a key in our store
func DeleteKey(Keyname string, ipaddr string) int {
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

// returns the number of keys in the store
func GetKeyCount() int {
	return len(keyStore)
}
