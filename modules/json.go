package modules

import (
	"fmt"
	"log"

	"encoding/json"
)

func Json(coinbaseClient *CoinbaseClient, persistFlag bool) {

	coinbaseClient.Refresh()

	json, err := json.Marshal(coinbaseClient.Container)
	if err != nil {
		log.Panic("JSON error")
	}

	fmt.Println(string(json))
}
