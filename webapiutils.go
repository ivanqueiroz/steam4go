package steam4go

import (
	"net/http"
	"net/url"
)

//Navigate to url
func Navigate(address string) (response string) {
	u, err := url.Parse(address)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
}
