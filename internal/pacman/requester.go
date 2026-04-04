package pacman

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)


func DownloadData(urlStr string) ([]byte,error) {
	// Validate URL format and scheme
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, errors.New("invalid URL format: " + err.Error())
	}
	
	// Only allow HTTP and HTTPS schemes
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return nil, errors.New("invalid URL scheme: only http and https are allowed")
	}
	
	// Prevent SSRF by blocking private IP ranges
	if parsedURL.Hostname() != "" {
		host := parsedURL.Hostname()
		// Block localhost and private IPs
		if strings.HasPrefix(host, "127.") || 
		   strings.HasPrefix(host, "localhost") ||
		   strings.HasPrefix(host, "10.") ||
		   strings.HasPrefix(host, "172.16.") ||
		   strings.HasPrefix(host, "192.168.") ||
		   host == "0.0.0.0" ||
		   strings.HasPrefix(host, "169.254.") { // link-local
			return nil, errors.New("access to private/local network addresses is not allowed")
		}
	}
	
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
	
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent","Venera Package Manager")
	
	r, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return nil, errors.New("status code different from 200 calling for "+urlStr)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil,err
	}
	return body, nil
}