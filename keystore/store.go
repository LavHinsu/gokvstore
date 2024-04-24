package keystore

import (
	"log"
	"strconv"
	"sync"
	"time"
)

type keyStoreMap struct {
	Value string // value of the key
	C_AT  int64  // created_at field (int64 because unix)
	U_AT  int64  // updated_at field (int64 because unix)
}

type HealthCheckStruct struct {
	TotalKeys  string // total keys in the store
	ServerTime int64  // current server time in unix
}

var (
	keyStore    = make(map[string]*keyStoreMap)
	healthCheck = make(map[string]*HealthCheckStruct)
	mutex       = &sync.RWMutex{}
)

// initialize our store with a value
func init() {
	mutex.Lock()
	defer mutex.Unlock()
	keyStore["foo"] = &keyStoreMap{
		Value: "bar",             //initalize with a value
		C_AT:  time.Now().Unix(), // intialize with an empty time object
		U_AT:  0,                 // intialize with an empty time object
	}
}

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

// function used to update a key in our store
func UpdateKey(Keyname string, Value string, ipaddr string) int {
	mutex.Lock()
	defer mutex.Unlock()
	_, KeyExists := keyStore[Keyname]
	if KeyExists {
		keyStore[Keyname].Value = Value
		keyStore[Keyname].U_AT = time.Now().Unix()
		log.Println(ipaddr+" updated key:", Keyname)
		return 200
	} else {
		log.Println(ipaddr+" couldn't find key to update:", Keyname)
		return 404
	}
}

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

// returns the number of keys in the store
func GetKeyCount() *HealthCheckStruct {
	mutex.Lock()
	defer mutex.Unlock()
	healthCheck["healthcheck"] = &HealthCheckStruct{
		TotalKeys:  strconv.Itoa(len(keyStore)),
		ServerTime: time.Now().Unix(),
	}
	return healthCheck["healthcheck"]
}
