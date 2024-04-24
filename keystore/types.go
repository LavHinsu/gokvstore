package keystore

type keyStoreMap struct {
	Value string // value of the key
	C_AT  int64  // created_at field (int64 because unix)
	U_AT  int64  // updated_at field (int64 because unix)
}

type HealthCheckStruct struct {
	TotalKeys  string // total keys in the store
	ServerTime int64  // current server time in unix
}
