package keystore

import (
	"log"
	"time"
)

func deleteExpiredKeys() {
	defer log.Println("reaper complete after 10 seconds")
	log.Println("started reaper")
	time.Sleep(10 * time.Second)
}

func Reaper() {
	/*
		set period to run our function. we can call it every second because the ticker will not let this run concurrently.
		if it does, it would cause heavy RWMutex locks. fortunately for us, this is handled internally by golang.
	*/
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			<-ticker.C
			deleteExpiredKeys()
		}
	}()
}
