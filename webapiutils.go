package steam4go

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func perror(err error) {
	if err != nil {
		panic(err)
	}
}

//Navigate to url
func Navigate(address string) (response string) {
	_, err := url.Parse(address)
	perror(err)

	resp, err := http.Get(address)
	perror(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}
