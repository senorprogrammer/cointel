package modules

import (
	"encoding/json"
)

func Json(cont *CurrencyContainer) string {
	json, err := json.Marshal(cont)
	if err != nil {
		return ""
	}

	return string(json)
}
