package modules

import (
	"encoding/json"
)

func Json(cc *CurrencyContainer) string {
	json, err := json.Marshal(cc)
	if err != nil {
		return ""
	}

	return string(json)
}
