package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	map1 := make(map[string]string)
	map1["aa"] = "111"
	map1["bb"] = "222"
	// fmt.Println(map1)
	// for k, v := range map1 {
	// 	fmt.Printf("k: %s    v:%s\n", k, v)
	// }

	my_json , _ := json.Marshal(map1)
	fmt.Println(string(my_json))
}
