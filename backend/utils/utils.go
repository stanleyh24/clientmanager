package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/stanleyh24/clientmanager/models"
)

type Response struct {
	Status  uint16
	Message string
	Body    map[string]any
}

func SendRequest(req *http.Request) (Response, error) {

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Timeout: time.Duration(1) * time.Second, Transport: transCfg}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return Response{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return Response{}, err
	}

	formattedData := formatJSON(body)

	var result map[string]any
	json.Unmarshal([]byte(formattedData), &result)

	var response Response
	if resp.StatusCode == http.StatusBadRequest {
		response = Response{
			Status:  uint16(resp.StatusCode),
			Message: result["detail"].(string),
		}
		return response, fmt.Errorf("router service add %s", result["detail"].(string))
	}
	response = Response{
		Status:  uint16(resp.StatusCode),
		Message: "",
		Body:    result,
	}
	return response, nil
}

func Querymaker(query models.QueryParams) *http.Request {
	var req *http.Request

	switch query.Method {

	case "GET":
		req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("https://%s/rest/%s", os.Getenv("ROUTER_IP"), query.Resource), nil)
		req.SetBasicAuth(os.Getenv("ROUTER_USERNAME"), os.Getenv("ROUTER_PASSWORD"))
		return req
	case "PUT":
		a, _ := json.Marshal(query.Body)
		req, _ = http.NewRequest(http.MethodPut, fmt.Sprintf("https://%s/rest/%s", os.Getenv("ROUTER_IP"), query.Resource), bytes.NewBuffer(a))
		req.SetBasicAuth(os.Getenv("ROUTER_USERNAME"), os.Getenv("ROUTER_PASSWORD"))
		return req
	case "PATCH":
		a, _ := json.Marshal(query.Body)
		req, _ = http.NewRequest(http.MethodPatch, fmt.Sprintf("https://%s/rest/%s", os.Getenv("ROUTER_IP"), query.Resource), bytes.NewBuffer(a))
		req.SetBasicAuth(os.Getenv("ROUTER_USERNAME"), os.Getenv("ROUTER_PASSWORD"))
		return req
	case "DELETE":
		req, _ = http.NewRequest(http.MethodDelete, fmt.Sprintf("https://%s/rest/%s", os.Getenv("ROUTER_IP"), query.Resource), nil)
		req.SetBasicAuth(os.Getenv("ROUTER_USERNAME"), os.Getenv("ROUTER_PASSWORD"))
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
