package main

import (
	"encoding/json"
	"fmt"
)

type InputJson struct {
	AbilityDict    AbilityDict `json:"ability_dict"`
	DetectInterval int         `json:"detect_interval"`
	AlarmInterval  int         `json:"alarm_interval"`
}

type AbilityDict struct {
	AbilityId AbilityId `json:"ability_id,omitempty"`
}

type AbilityId struct {
	Conf float64 `json:"conf,omitempty"`
}

func main() {
	var abilityId1 AbilityId
	abilityId1.Conf = 0.5

	var abilityDict1 AbilityDict
	abilityDict1.AbilityId = abilityId1

	var inputJson1 InputJson
	inputJson1.AbilityDict = abilityDict1
	inputJson1.AlarmInterval = 180
	inputJson1.DetectInterval = 30

	marshal, _ := json.Marshal(inputJson1)
	fmt.Println(string(marshal))
	//{"ability_dict":{"ability_id":{}},"detect_interval":30,"alarm_interval":180}

	var str string = "{\"ability_dict\":{\"ability_id\":{}},\"detect_interval\":30,\"alarm_interval\":180}"

	var inputJson2 InputJson
	json.Unmarshal([]byte(str), &inputJson2)
	fmt.Println(inputJson2.AbilityDict.AbilityId)
}
