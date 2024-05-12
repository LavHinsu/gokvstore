package keystore

// this is the core keystore map. it's a flat store of values.
// all of the time fields are in int64 because we only deal with unix timestamps
type keyStoreMap struct {
	Value string // value of the key
	C_AT  int64  // created_at field (int64 because unix)
	U_AT  int64  // updated_at field (int64 because unix)
	E_AT  int64
}

type HealthCheckStruct struct {
	TotalKeys  string // total keys in the store
	ServerTime int64  // current server time in unix
}
