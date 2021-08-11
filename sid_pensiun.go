package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type init_hehe struct {
	urlapi string
}

func main() {
	init := init_hehe{
		"https://home.s.id/api/public/link/shorten",
	}
	hasil := os.Args
	if len(hasil) < 2 {
		fmt.Println("maaf, gunakan sid {url}")
	} else {
		fmt.Printf("url : %s\n", hasil[1])
	}

	values := map[string]string{"url": "https://stackoverflow.com/questions/2707434/how-to-access-command-line-arguments-passed-to-a-go-program"}

	jsonValue, _ := json.Marshal(values)

	respon, mengerror := http.NewRequest("POST", init.urlapi, bytes.NewReader(jsonValue))
	client := &http.Client{}
	resp, mengerror := client.Do(respon)
	if mengerror != nil {
		panic(mengerror)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
