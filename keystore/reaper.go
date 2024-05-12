package keystore

import (
	"log"
	"time"
)

func deleteExpiredKeys() {
	keysToExpire := getKeysThatExpire() //
	if len(keysToExpire) > 0 {          // only call the delete if there are keys to delete. saves unnecessary mutex locks
		removeKeys(keysToExpire)
	}

}

/*
returns keys from the map that have expired in a slice.
*/
func getKeysThatExpire() []string {
	defer mutex.RUnlock()
	mutex.RLock()
	/*
		only calculate it once. if done in the loopchecker, we add some overhead of calculating the stamp every loop iteration
	*/
	timenow := time.Now().Unix()

	keysToExpire := []string{}

	for key, value := range keyStore {
		if value.E_AT != 0 { // 0 = E_AT not set
			if value.E_AT < timenow { // ie, if the stamp is in the past
				keysToExpire = append(keysToExpire, key)
			}
		}
	}
	return keysToExpire
}

/*
This function deletes all keys it is given in the argument.
*/
func removeKeys(keys []string) {
	defer mutex.Unlock()
	mutex.Lock()
	for _, key := range keys {
		log.Println("reaper delete expired key:", key)
		/*
			delete the key from our map. we don't need to worry about concurrent delete (ie, if someone sent a delete request)
			because in golang a delete operation over a non-existent key is just a no-op.
		*/
		delete(keyStore, key)
	}
}

// this is our reaper function. it will run a goroutine that checks and deletes keys if present
func Reaper() {
	/*
		set period to run our function. we can call it every second because the iterator in the for loop will not let this run concurrently.
		this is because there is nothing listening on the channel when it is updated when the iterator is running
		if there was one, it would cause heavy RWMutex locks. fortunately for us, this is handled internally by golang.
	*/
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			<-ticker.C
			deleteExpiredKeys()
		}
	}()
}
