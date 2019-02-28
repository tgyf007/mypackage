package weibo

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/json-iterator/go"
)

const (
	BASE_HOST  = `https://api.weibo.com/`
	APP_KEY    = ``
	APP_SECRET = ``
)

type Oauth struct {
}

type AToken struct {
	AccessToken string `json:"access_token"`
	RemindId    string `json:"remind_in"`
	ExpiresIn   uint64 `json:"expires_in"`
	Uid         string `json:"uid"`
	IsRealName  string `json:"isRealName"`
}

type TokenInfo struct {
	Uid       uint64 `json:"uid"`
	AppKey    string `json:"appkey"`
	Scope     string `json:"scope"`
	CreateAt  uint64 `json:"create_at"`
	Expire_in uint64 `json:"expire_in"`
}

/**
@link https://open.weibo.com/wiki/Oauth2/authorize
*/
func (o Oauth) AuthorizeLoginUrl(redirect_url, response_code, state, display string) string {
	return BASE_HOST + `oauth2/authorize` + `?client_id=` + APP_KEY + `&redirect_uri=` + redirect_url + `&response_code=` + response_code + `&state=` + state + `&display=` + display
}

/**
@link https://open.weibo.com/wiki/Oauth2/access_token
*/
func (o Oauth) AccessToken(grant_type string, k map[string]string) (data *AToken, err error) {
	url := BASE_HOST + `oauth2/access_token`
	r := `client_id=` + APP_KEY + `&client_secret=` + APP_SECRET
	if grant_type == `authorization_code` {
		_, ok := k[`code`]
		_, ok1 := k[`redirect_uri`]
		if !ok || !ok1 {
			err = errors.New(`错误的参数`)
		}
		r += `&grant_type=` + grant_type + `&code=` + k[`code`] + `&redirect_uri=` + k[`redirect_uri`]
	}

	body, err := o.post(url, r)
	if err != nil {
		return
	}

	err = jsoniter.Unmarshal(body, &data)
	if err != nil {
		return
	}
	return
}

/**
@link https://open.weibo.com/wiki/Oauth2/get_token_info
*/
func (o Oauth) GetTokenInfo(token string) (data *TokenInfo, err error) {
	url := BASE_HOST + `oauth2/get_token_info`
	r := `access_token=` + token

	body, err := o.post(url, r)
	if err != nil {
		return
	}

	err = jsoniter.Unmarshal(body, &data)
	if err != nil {
		return
	}
	return
}

/**
@link https://open.weibo.com/wiki/Oauth2/revokeoauth2
*/
func (o Oauth) RevokeOauth(token string) (err error) {
	url := BASE_HOST + `oauth2/revokeoauth2`
	r := `access_token=` + token
	body, err := o.post(url, r)
	if err != nil {
		return
	}
	data := struct {
		Result string `json:"result"`
	}{}
	err = jsoniter.Unmarshal(body, &data)
	if err != nil {
		return
	}
	if data.Result != `true` {
		err = errors.New(`revoke failed`)
		return
	}
	return
}

func (o Oauth) post(url, r string) (body []byte, err error) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(r))
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
