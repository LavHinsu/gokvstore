package keystore

import "sync"

var (
	keyStore    = make(map[string]*keyStoreMap)
	healthCheck = make(map[string]*HealthCheckStruct)
	mutex       = &sync.RWMutex{}
)
