package keystore

import (
	"time"
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
