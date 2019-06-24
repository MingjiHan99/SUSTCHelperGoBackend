package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Jeffail/gabs"
)

func Login(code string) string {
	jsonObj := gabs.New()
	jsonObj.Set(false, "success")

	req, err := http.NewRequest("GET", "https://api.weixin.qq.com/sns/jscode2session", nil)
	if err != nil {
		return jsonObj.String()
	}

	q := req.URL.Query()
	fmt.Println(code)
	q.Add("appid", "wxf2b43b00a21af0c5")
	q.Add("secret", "1dce117e15c630b0a11c343c49163132")
	q.Add("js_code", code)
	q.Add("grant_type", "authorization_code")

	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())

	var resp *http.Response
	if resp, err = http.DefaultClient.Do(req); err != nil {
		return jsonObj.String()
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return jsonObj.String()
	}

	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(body), &dat); err != nil {
		return jsonObj.String()
	}

	openid := dat["openid"].(string)
	time := time.Now().Format("2006-01-02 15:04:05")
	str := openid + "SusTechHelper" + time
	sessionKey := Sha256Encryption(&str)

	if err := SetSession(openid, sessionKey); err != nil {
		fmt.Errorf("Set session failed" + err.Error())
		return jsonObj.String()
	}
	jsonObj.Set(true, "success")
	jsonObj.Set(sessionKey, "session_key")
	return jsonObj.String()
}
