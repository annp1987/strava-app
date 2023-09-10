package cronjob

import (
	"io/ioutil"
	"net/http"
)

func RequestFutureFunction(url string) func() ([]byte, error) {
	var body []byte
	var err error
	c := make(chan struct{}, 1)
	go func() {
		defer close(c)
		var resp *http.Response
		resp, err = http.Get(url)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
	}()
	return func() ([]byte, error) {
		<-c
		return body, err
	}
}
