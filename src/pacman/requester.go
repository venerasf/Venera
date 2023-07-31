package pacman

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)


func DownloadData(url string) ([]byte,error) {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
					Timeout:   5 * time.Second,
					KeepAlive: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   3 * time.Second,
			ResponseHeaderTimeout: 3 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	
	req,_ := http.NewRequest("GET",url,nil)
	req.Header.Set("User-Agent","Venera Package Manager")
	
	r, err := client.Do(req)
	if err != nil {
		return nil,err
	}

	if r.StatusCode != 200 {
		return nil, errors.New("status code different from 200")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil,err
	}
	return body, nil
}