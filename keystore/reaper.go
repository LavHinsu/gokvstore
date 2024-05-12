package keystore

import (
	"fmt"
	"time"
)

func deleteExpiredKeys() {

}

func Reaper() {

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		t := <-ticker.C
		fmt.Println("running reaper", t)
		deleteExpiredKeys()
	}()
}
