package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Numverify struct {
	Valid               bool   `json:"valid"`
	Number              string `json:"number"`
	LocalFormat         string `json:"local_format"`
	InternationalFormat string `json:"international_format"`
	CountryPrefix       string `json:"country_prefix"`
	CountryCode         string `json:"country_code"`
	CountryName         string `json:"country_name"`
	Location            string `json:"location"`
	Carrier             string `json:"carrier"`
	LineType            string `json:"line_type"`
}

func main() {
	phone := "14158586273"
	key := "1af9e2ba48e1e24dbf7f71ae28a4918b"
	// QueryEscape escapes the phone string so
	// it can be safely placed inside a URL query
	safePhone := url.QueryEscape(phone)
	accessKey := url.QueryEscape(key)

	url := fmt.Sprintf("http://apilayer.net/api/validate?access_key=%s&number=%s", accessKey, safePhone)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// panic(err)
	}
	phoneMap := make(map[string]string)
	if err := json.Unmarshal([]byte(result), &phoneMap); err != nil {
		for k, v := range phoneMap {
			fmt.Printf("(k)%s  = (v)%s\n", k, v)
		}
	}

	// Fill the record with the data from the JSON
	var record Numverify

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println("Json Decoder Error : ", err)
	}

	//fmt.Println("Phone No. = ", record.InternationalFormat)
	//fmt.Println("Country   = ", record.CountryName)
	//fmt.Println("Location  = ", record.Location)
	//fmt.Println("Carrier   = ", record.Carrier)
	//fmt.Println("LineType  = ", record.LineType)

}
