package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	values := map[string]string{"url": "https://stackoverflow.com/questions/2707434/how-to-access-command-line-arguments-passed-to-a-go-program"}

	jsonValue, _ := json.Marshal(values)

	fmt.Println(jsonValue)
}
