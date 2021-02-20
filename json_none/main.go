package main

import (
	"encoding/json"
	"fmt"
)

// неизвестный json
func main() {
	j := []byte(`{["Name":"Vasya",
		"Job":{"Department":"Operations","Title":"Boss"},
		"Good":["lddasd", "jdalksdjla", 12]]}`)

	var p2 interface{}
	json.Unmarshal(j, &p2)
	fmt.Printf("p2: %v\n", p2)

	person, ok := p2.(map[string]interface{})

	if !ok {
		p := p2.([]map[string]interface{})
		person = p[0]
	}

	fmt.Printf("name=%s\n", person["Good"])

	for _, v := range person["Good"].([]interface{}) {
		switch v.(type) {
		case string:
			fmt.Printf("%s  это строка\n", v)
		case float64:
			fmt.Printf("%.2f  это float64\n", v)
		case int:
			fmt.Printf("%d  это int\n", v)

		}

	}
}
