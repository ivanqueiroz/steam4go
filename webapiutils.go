package steam4go

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

//APIList type represent return of supported methods list
type APIList struct {
	Apilist struct {
		Interfaces []struct {
			Methods []struct {
				Httpmethod string        `json:"httpmethod"`
				Name       string        `json:"name"`
				Parameters []interface{} `json:"parameters"`
				Version    float64       `json:"version"`
			} `json:"methods"`
			Name string `json:"name"`
		} `json:"interfaces"`
	} `json:"apilist"`
}

func perror(err error) {
	if err != nil {
		panic(err)
	}
}

//GetSupportedAPIListObj returns allc interfaces and method .
func GetSupportedAPIListObj() (data APIList) {
	u, _ := url.Parse(ifSteam)
	u.Path = "/ISteamWebAPIUtil/GetSupportedAPIList/v0001/"
	var apiList APIList
	jsonSrc := NavigateToByte(u.String())
	json.Unmarshal(jsonSrc, &apiList)
	return apiList
}

//GetSupportedAPIListStr returns allc interfaces and method .
func GetSupportedAPIListStr() (data string) {
	u, _ := url.Parse(ifSteam)
	u.Path = "/ISteamWebAPIUtil/GetSupportedAPIList/v0001/"
	return NavigateToString(u.String())
}

//NavigateToString to url returns String
func NavigateToString(address string) (response string) {
	_, err := url.Parse(address)
	perror(err)

	resp, err := http.Get(address)
	perror(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}

//NavigateToByte to url returns []byte
func NavigateToByte(address string) (response []byte) {
	_, err := url.Parse(address)
	perror(err)

	resp, err := http.Get(address)
	perror(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body
}
