package user

import (
	utils "SUSTechHelperGoBackend/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unsafe"
)

const templateID = "kzj01w29QBY6VOikjcRjm-H7VVfcjjtVOcvDDMrr2UE"

func Send(openid, course, formID string) (err error) {
	var accessToken string
	if accessToken, err = utils.GetAccessToken(); err != nil {
		fmt.Errorf("Get access token failed:" + err.Error())
		return err
	}

	form := make(map[string]interface{})
	form["touser"] = openid
	form["template_id"] = templateID
	form["form_id"] = formID
	form["data"] = encode("11711918", "计组", "ntm", "gg", "gun")
	form["emphasis_keyword"] = "keyword3.DATA"
	bytesData, err := json.Marshal(form)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(bytesData[:]))
	reader := bytes.NewReader(bytesData)
	url := "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=" + accessToken
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
	return nil
}

func encode(sid, course, score, hint, now string) map[string]interface{} {
	data := make(map[string]interface{})
	data["keyword1"] = map[string]string{"value": sid}
	data["keyword2"] = map[string]string{"value": course}
	data["keyword3"] = map[string]string{"value": score}
	data["keyword4"] = map[string]string{"value": hint}
	data["keyword5"] = map[string]string{"value": now}
	return data
}
