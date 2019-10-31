package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://private-anon-cf6e9edfbc-restapi3.apiary-mock.com/notes"
	fmt.Println("URL:>", url)

	values := make(map[string]string)
	values["title"] = "Buy cheese and bread for breakfast."
	jsonVal, _ := json.Marshal(values)

	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonVal))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonVal))
	if err != nil {
		log.Fatal(err)
	}
	//c := &http.Client{}
	//resp, err := c.Do(req)
	//if err != nil { log.Fatal(resp) }
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
