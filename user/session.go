package user

import (
	"SUSTechHelperGoBackend/utils"
)

// SetSession save map from session to openid
func SetSession(openid, session string) error {
	if err := utils.Set("SH:session:"+session+":openid", openid); err != nil {
		return err
	}
	return nil
}

// GetSession maps session to openid
func GetSession(session string) (string, error) {
	return utils.Get("SH:session:" + session + ":openid")
}
