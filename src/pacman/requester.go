package pacman

import (
	"io/ioutil"
	"net/http"
)


func DownloadData(url string) ([]byte,error) {
	client := &http.Client{}
	req,_ := http.NewRequest("GET",url,nil)
	req.Header.Set("User-Agent","Venera Package Manager")
	
	r, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil,err
	}
	return body, nil
}