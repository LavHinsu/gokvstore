package keystore

import (
	"log"
	"sync"
	"time"
)

var (
	keyStore    = make(map[string]*keyStoreMap)
	healthCheck = make(map[string]*HealthCheckStruct)
	mutex       = &sync.RWMutex{}
)

// initialize our store with a value (that expires)
func init() {
	mutex.Lock()
	defer mutex.Unlock()
	defer Reaper() // start our reaper routine after the store has initialized
	keyStore["foo"] = &keyStoreMap{
		Value: "bar",                  //initalize with a value
		C_AT:  time.Now().Unix(),      // intialize with an empty time object
		U_AT:  0,                      // intialize with an empty time object
		E_AT:  time.Now().Unix() + 10, // expire foo in 10 seconds
	}
	log.Println("intitalized kvstore")
}
