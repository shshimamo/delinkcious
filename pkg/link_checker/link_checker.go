package link_checker

import (
	"errors"
	"net/http"
)

func CheckLink(url string) (err error) {
	resp, err := http.Head(url)
	if err != nil {
		return
	}
	if resp.StatusCode >= 400 {
		err = errors.New(resp.Status)
	}
	return
}
