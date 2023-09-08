package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/stanleyh24/clientmanager/models"
)

func SendRequest(req *http.Request) string {

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Timeout: time.Duration(1) * time.Second, Transport: transCfg}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	if resp.StatusCode == 204 {
		return ""
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return ""
	}
	formattedData := formatJSON(body)

	return formattedData
}

func Querymaker(query models.QueryParams) *http.Request {
	var req *http.Request

	switch query.Method {

	case "GET":
		req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("https://%s/rest/%s", query.Ip, query.Resource), nil)
		req.SetBasicAuth(query.Username, query.Password)
		req.Header.Set("Content-Type", "application/json")
		return req
	case "PUT":
		a, _ := json.Marshal(query.Body)
		req, _ = http.NewRequest(http.MethodPut, fmt.Sprintf("https://%s/rest/%s", query.Ip, query.Resource), bytes.NewBuffer(a))
		req.SetBasicAuth(query.Username, query.Password)
		req.Header.Set("Content-Type", "application/json")
		return req
	case "DELETE":
		req, _ = http.NewRequest(http.MethodDelete, fmt.Sprintf("https://%s/rest/%s", query.Ip, query.Resource), nil)
		req.SetBasicAuth(query.Username, query.Password)
		req.Header.Set("Content-Type", "application/json")
		return req
	}

	return req
}

// function to format JSON data
func formatJSON(data []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	d := out.Bytes()
	return string(d)
}
