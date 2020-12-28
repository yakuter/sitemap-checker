package checker

import (
	"net/http"
	"net/url"
)

type Status struct {
	status   int
	redirect *url.URL
}

func CheckRedirect(url string) (Status, error) {
	s := Status{}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	response, err := client.Get(url)
	if err != nil {
		return s, err
	}
	redirect, err := response.Location()
	if err != nil {
		return s, err
	}
	defer response.Body.Close()
	s.status = response.StatusCode
	s.redirect = redirect
	return s, err
}
