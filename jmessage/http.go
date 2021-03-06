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
	return request("DELETE", url, nil)
}

func delete(url string, data *strings.Reader) (rest *bytes.Buffer) {
	return request("DELETE", url, data)
}

// 构造请求
func request(method string, url string, data *strings.Reader) (rest *bytes.Buffer) {
	var req *http.Request
	var err error
	if data == nil {
		req, err = http.NewRequest(method, baseApi+url, nil)
	} else {
		req, err = http.NewRequest(method, baseApi+url, data)
	}

	if err != nil {
		return bytes.NewBuffer([]byte("{}"))
	}

	req.Header.Set("Authorization", "Basic "+basic)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	response, _ := (&http.Client{}).Do(req)

	defer response.Body.Close()

	res, _ := ioutil.ReadAll(response.Body)

	return bytes.NewBuffer(res)
}
