package myHTTP

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

//HTTPRequestJSON :
func HTTPRequestJSON(urlStr string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

//HTTPRequestJSONProxy :
func HTTPRequestJSONProxy(urlStr string, proxyHost string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	proxyURL, err := url.Parse(proxyHost)
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

//HTTPRequestGET :
func HTTPRequestGET(client *http.Client, urlStr string, params map[string]string) ([]byte, error) {
	param := ""
	if len(params) > 0 {
		for idx, value := range params {
			param = param + string(idx) + "=" + string(value) + "&"
		}
		param = "?" + param[:len(param)-1]
	}

	resp, err := client.Get(urlStr + param)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

//HTTPAuthPOST :
func HTTPAuthPOST(client *http.Client, urlLoc string, auth string) ([]byte, error) {
	data := []byte{}

	req, err := http.NewRequest("POST", urlLoc, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", auth)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	u, _ := url.Parse("https://www.oyoos.com/")
	client.Jar.SetCookies(u, resp.Cookies())
	return bodyBytes, nil
}
