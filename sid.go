package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gookit/color"
)

type init_hehe struct {
	urlapi string
}

type detail_shorturl struct {
	Short    string `json:"short"`
	Date     string `json:"date"`
	Timezone string `json:"timezone"`
	Long_url string `json:"long_url"`
}
type init_tampilan struct {
	Hello string
	Gris  string
}

func (hh *init_tampilan) initcli() {
	hh.Hello = "CommandLine tools shorturl untuk s.id shortener\n\n"
}
func (hh *init_tampilan) initgaris() {
	g := "="
	for a := 1; a < 6; a++ {
		g = g + g
	}
	hh.Gris = g + "\n"
}

func main() {
	menginit := init_tampilan{}
	menginit.initcli()
	menginit.initgaris()
	color.Cyan.Printf(menginit.Gris)
	color.Cyan.Printf(menginit.Hello)

	init := init_hehe{
		"https://home.s.id/api/public/link/shorten",
	}
	hasil := os.Args
	if len(hasil) < 2 {
		color.Green.Print("maaf, gunakan sid {url}")
	} else {
		color.Yellow.Printf("url : %s\n", hasil[1])
		data := url.Values{
			"url": {string(hasil[1])},
		}

		resp, err := http.PostForm(init.urlapi, data)

		HTTPSTATUSCODE := resp.StatusCode
		if HTTPSTATUSCODE == 429 {
			color.Red.Println("maaf, server memberikan nilai HTTP 429, silahkan coba lagi nanti")
		} else {
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			body, _ := ioutil.ReadAll(resp.Body)
			// fmt.Println("response Body:", string(body))

			var hasilapi detail_shorturl
			if error := json.Unmarshal(body, &hasilapi); error != nil {
				color.Red.Println("sori, respon server memberikan nilai balik engga valid")
			}

			color.Green.Println("URL      : http://s.id/" + hasilapi.Short)
			color.Green.Println("URL asli : " + hasilapi.Long_url)

		}

	}

}
