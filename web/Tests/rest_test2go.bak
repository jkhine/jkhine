// Json post - use

package main

import (
	"net/url"
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
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

func jsonUnmarshal(apiUrl string) {
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil { log.Fatal(err) }

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil { log.Fatal(err) }

	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil { log.Fatal(err) }

	phoneMap := make(map[string]interface{})

	if err := json.Unmarshal(result, &phoneMap); err != nil { log.Fatal(err) }
	for k,v := range phoneMap {
		fmt.Printf("%s = %s\n", k,v)
	}
}

func jsonDecodeType(apiUrl string) {
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil { log.Fatal(err) }

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil { log.Fatal(err) }

	defer resp.Body.Close()

	var record Numverify
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil { log.Fatal(err) }
	fmt.Println(record)
}

func main() {
	phone := "17188969720"
	key := "1af9e2ba48e1e24dbf7f71ae28a4918b"
	baseUrl := "http://apilayer.net/api/validate?access_key=%s&number=%s"

	safeKey := url.QueryEscape(key)
	safePhone := url.QueryEscape(phone)

	apiUrl := fmt.Sprintf(baseUrl, safeKey, safePhone)
	jsonUnmarshal(apiUrl)
	jsonDecodeType(apiUrl)

}
