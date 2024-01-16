package domain

import (
	"encoding/json"
	"github.com/cbdavid14/ms-api-go-banking/logger"
	"net/http"
	"net/url"
)

type AuthRepository interface {
	IsAuthorizedFor(role string, routeName string, vars map[string]string) bool
}

type RemoteAuthRepository struct {
}

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{}
}

func (r RemoteAuthRepository) IsAuthorizedFor(token string, routeName string, vars map[string]string) bool {
	u := buildVerifyURL(token, routeName, vars)
	if response, err := http.Get(u); err != nil {
		logger.Error("Error while sending..." + err.Error())
		return false
	} else {
		m := map[string]bool{}
		if err := json.NewDecoder(response.Body).Decode(&m); err != nil {
			logger.Error("Error while decoding response from auth server: " + err.Error())
			return false
		}
		return m["isAuthorized"]
	}
}

// Sample: /auth/verify?token=aaaa.bbbb.cccc&routeName=MakeTransaction&customer_id=2000&account_id=95470
func buildVerifyURL(token string, routeName string, vars map[string]string) string {
	u := url.URL{Host: "localhost:8001", Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
