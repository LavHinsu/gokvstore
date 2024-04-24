package keystore

import (
	"strconv"
	"time"
)

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
