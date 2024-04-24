package keystore

import (
	"sync"
	"time"
)

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
