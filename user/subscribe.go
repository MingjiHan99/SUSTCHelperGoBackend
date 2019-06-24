package user

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

func Subsribe(username, password, formID, course, session string) string {
	jsonObj := gabs.New()
	jsonObj.Set(false, "success")
	if openid, err := GetSession(session); err != nil {
		return jsonObj.String()
	} else {
		fmt.Println(openid, course, formID)
		Send(openid, course, formID)
	}
	return jsonObj.String()
}
