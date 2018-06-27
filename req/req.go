package req

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var Client = &http.Client{}

func Get(url string) (int, string, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Travis-API-Version", "3")
	resp, err := Client.Do(req)
	if err != nil {
		fmt.Printf("Error fetching url %s", url)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(body), err
}
