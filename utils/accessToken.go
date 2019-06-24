package utils

import (
	"encoding/json"
	"net/http"

	"io/ioutil"
)

// GetAccessToken query access_token from wechat
func GetAccessToken() (string, error) {
	req, err := http.NewRequest("GET", "https://api.weixin.qq.com/cgi-bin/token", nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("appid", "wxf2b43b00a21af0c5")
	q.Add("secret", "1dce117e15c630b0a11c343c49163132")
	q.Add("grant_type", "client_credential")

	req.URL.RawQuery = q.Encode()

	var resp *http.Response
	if resp, err = http.DefaultClient.Do(req); err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(body), &dat); err != nil {
		return "", err
	}

	accessToken := dat["access_token"].(string)
	return accessToken, nil
}
