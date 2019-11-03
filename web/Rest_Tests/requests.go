package main

// Reference web pages
// https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type HTTPBinResponse struct {
	Args    map[string]string `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	URL     string            `json:"url"`
}

func jsonUnmarshal() {
	APIURL := "https://httpbin.org/get?name=roger&country=switzerland"

	req, err := http.NewRequest(http.MethodGet, APIURL, nil)
	if err != nil { log.Fatal("http.NewRequest() Error : ", err)}
	c := http.DefaultClient
	resp, err := c.Do(req)
	if err != nil { log.Fatal( "client.Do() : ", err)}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { log.Fatal("ioutil.ReadAll() Error : ", err)}

	var r HTTPBinResponse
	if err = json.Unmarshal(body, &r)  ; err !=nil { log.Fatal("Unmarshal() Error : ", err)}
	fmt.Println(r)
	fmt.Println(r.Headers)
	fmt.Println(r.Args)
}

func netHTTPPost() {
	requestBody, err := json.Marshal(map[string] string {
		"name" : "noval djokovic",
		"country" : "serbia",
	})
	if err != nil { log.Fatal("Marshal() Error : ", err)}

	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(requestBody))
	if err != nil { log.Fatal("Post() Error : ", err) }
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { log.Fatal("ReadAll() Error  : ", err)}
	fmt.Println(string(body))
}

func netHTTPPostForm() {
	formData := url.Values{
		"players" : {"roger federer", "novak djokovic"},
	}

	resp, err := http.PostForm("https://httpbin.org/post", formData)
	if err != nil { log.Fatal( "PostForm() Error : ", err)}

	var result map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&result) ; err != nil { log.Fatal("NewDecoder().Decode() Error : ", err)}
	fmt.Println(result["form"])
}

func netHTTPGet() {
	// JSON Get using built-in net/http Get
	APIURL := "http://httpbin.org/get"

	req, err := http.NewRequest(http.MethodGet, APIURL, nil)  // req, err := http.NewRequest("GET", APIURL, nil)
	if err != nil { log.Fatal(err)}
	c := http.DefaultClient // c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil { log.Fatal(err)}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { log.Fatal(err)}
	fmt.Println(string(body))
}


func gRequestsSessionPost() {
	// JSON Post using grequests session object, Post
	auth := map[string]string{
		"user_name": "admin",
		"password": "admin",}

	session := grequests.NewSession(nil)
	res, err := session.Post("http://httpbin.org/post", &grequests.RequestOptions{Data: auth})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.StatusCode)
	fmt.Println(res.String())
	fmt.Println(session.RequestOptions.Cookies)
}

func main() {
	netHTTPPostForm()
	// netHTTPPost()
	// netHTTPGet()
	// jsonUnmarshal()
	// gRequestsSessionPost()
}
