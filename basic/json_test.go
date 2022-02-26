package basic

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestJsonDeserialize(t *testing.T) {
	var paramA []string
	jsonStr1 := "{\"inbound_id\":\"INSGL0002107190004\",\"creator\":\"zhiwei.yang@shopee.com\",\"print_at\":\"2021-08-17 18:24:53\",\"__test\":\"test\"}"
	//jsonStr2 := "[\"aaa\",\"bbb\",\"cc\"]"
	var jsonStr2 string

	if jsonStr2 != "" {
		json.Unmarshal([]byte(jsonStr2), &paramA)
	} else {
		paramA = extractValidationParams(jsonStr1)
	}
	fmt.Println(paramA)

}

func extractValidationParams(paramsJsonStr string) []string {
	var params map[string]interface{}
	if err := json.Unmarshal([]byte(paramsJsonStr), &params); err != nil {
		return nil
	}
	validationParams := []string{}
	for key := range params {
		if strings.HasPrefix(key, "__") {
			continue
		}
		validationParams = append(validationParams, key)
	}
	return validationParams
}
