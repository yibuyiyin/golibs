package http

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

// Get 发起get请求
func Get(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic("http get fail." + err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("read body fail." + err.Error())
	}
	return body
}

// PostJson 发起post请求
// params, _ := json.Marshal(map[string]string{"name": "Test"})
func PostJson(url string, params []byte) []byte {
	responseBody := bytes.NewBuffer(params)
	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}
