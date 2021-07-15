package jmessage

import (
	"bytes"
	// "encoding/base64"
	// "encoding/json"
	// "fmt"
	"io/ioutil"
	"net/http"
	"strings"
	// _ "time"
)

func post(url string, data *strings.Reader) (rest *bytes.Buffer) {
	return request("POST", url, data)
}

func put(url string, data *strings.Reader) (rest *bytes.Buffer) {
	return request("PUT", url, data)
}

func get(url string) (rest *bytes.Buffer) {
	req, _ := http.NewRequest("GET", baseApi+url, nil)
	req.Header.Set("Authorization", "Basic "+basic)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	response, _ := (&http.Client{}).Do(req)
	defer response.Body.Close()
	res, _ := ioutil.ReadAll(response.Body)
	return bytes.NewBuffer(res)
}

// 构造请求
func request(method string, url string, data *strings.Reader) (rest *bytes.Buffer) {
	req, _ := http.NewRequest(method, baseApi+url, data)
	req.Header.Set("Authorization", "Basic "+basic)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	response, _ := (&http.Client{}).Do(req)

	defer response.Body.Close()

	res, _ := ioutil.ReadAll(response.Body)

	return bytes.NewBuffer(res)
}
