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

//GetSupportedAPIListStruct returns all interfaces and method .
func GetSupportedAPIList() (data APIList) {
	u, _ := url.Parse(ifSteam)
	u.Path = "/ISteamWebAPIUtil/GetSupportedAPIList/v0001/"
	var apiList APIList
	jsonSrc := navigateToByte(u.String())
	json.Unmarshal(jsonSrc, &apiList)
	return apiList
}

//GetSupportedAPIListStr returns all interfaces and method as string .
func GetSupportedAPIListStr() (data string) {
	u, _ := url.Parse(ifSteam)
	u.Path = "/ISteamWebAPIUtil/GetSupportedAPIList/v0001/"
	return navigateToString(u.String())
}

func navigateToString(address string) (response string) {
	_, err := url.Parse(address)
	perror(err)

	resp, err := http.Get(address)
	perror(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}

func navigateToByte(address string) (response []byte) {
	_, err := url.Parse(address)
	perror(err)

	resp, err := http.Get(address)
	perror(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body
}
