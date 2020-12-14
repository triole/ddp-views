package req

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get sends a get request
func (s Req) Get(url string) (req *http.Request) {
	req = s.initRequest("GET", s.BaseURL+url+"/", []byte{})
	s.printResponse(req)
	return
}

// Put sends a put request
func (s Req) Put(url string, json []byte) (req *http.Request) {
	req = s.initRequest("PUT", s.BaseURL+url+"/", json)
	s.printResponse(req)
	return
}

func (s Req) printResponse(req *http.Request) {
	client := &http.Client{}
	resp, err := client.Do(req)
	check(err)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	fmt.Printf("Request %s\t%v\n", req.URL, resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("%v\n", bodyString)
	}
}

func (s Req) initRequest(method string, url string, json []byte) (req *http.Request) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(json))
	check(err)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Token "+s.Token)
	return req
}
