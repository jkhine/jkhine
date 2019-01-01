package pager_duty

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func  jsonToMap(contents []byte) {
	var f interface{}
	json.Unmarshal(contents, &f)
	m := f.(map[string]interface{})
	onCallMap := m["oncalls"]
	fmt.Println(onCallMap)
}



func main() {
	var pgdOncallUrl string = "https://api.pagerduty.com/oncalls?time_zone=EST&escalation_policy_ids%5B%5D="
	var networkEid string = "P2XZH4B"
	var pgUrl = pgdOncallUrl + networkEid

	client := &http.Client{}
	req, err := http.NewRequest("GET", pgUrl, nil)
	req.Header.Add("Accept", "application/vnd.pagerduty+json;version=2")
	req.Header.Add("Authorization", "Token token=N4JzsHJNzhdxrU_drUEo")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	} else {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {

		} else {
			jsonToMap(contents)
		}

	}
}

