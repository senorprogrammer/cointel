package modules

import (
	"fmt"
	"log"

	"encoding/json"
)

func Json(client *CoinbaseClient, persistFlag bool) {

	client.Refresh()

	json, err := json.Marshal(client.Container)
	if err != nil {
		log.Panic("JSON error")
	}

	fmt.Println(string(json))
}
