package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	dataToJSON := make(map[string]interface{})
	fmt.Println(dataToJSON)
	dataToJSON["key1"] = "string data"
	dataToJSON["key2"] = 5.56
	fmt.Println(dataToJSON)

	jsonString, err := json.Marshal(dataToJSON)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonString)
	var jsonToMap map[string]interface{}
	json.Unmarshal(jsonString,&jsonToMap)

	fmt.Println(jsonToMap)
	
	str := jsonToMap["key1"].(string)
	num := jsonToMap["key2"].(float64)

	fmt.Println(str)
	fmt.Println(num)


}