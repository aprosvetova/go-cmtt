package cmtt

import (
	"errors"
	"fmt"
	"gopkg.in/resty.v1"
)

//NewAnonymous creates Cmtt instance without auth, not all methods available
func NewAnonymous(platform string) (*Cmtt, error) {
	cmtt, err := newInstance(platform)
	if err != nil {
		return nil, err
	}
	return cmtt, nil
}

//NewWithQr creates Cmtt instance, authenticates via QR (save token with GetToken and use NewWithToken next times)
func NewWithQr(token, platform string) (*Cmtt, error) {
	cmtt, err := newInstance(platform)
	if err != nil {
		return nil, err
	}
	_, err = cmtt.authQr(token)
	if err != nil {
		return nil, err
	}
	return cmtt, nil
}

//NewWithToken creates Cmtt instance authorized with token. Use only if you already have correct token
func NewWithToken(token, platform string) (*Cmtt, error) {
	cmtt, err := newInstance(platform)
	if err != nil {
		return nil, err
	}
	cmtt.token = token
	cmtt.client.SetHeader("X-Device-Token", token)
	_, err = cmtt.GetMe()
	if err != nil {
		return nil, err
	}
	return cmtt, nil
}

//GetToken returns current token. Useful if you authenticated via NewWithQr and want to keep session
func (api *Cmtt) GetToken() string {
	return api.token
}

//SetProxy sets proxy for current instance. Use http://user@pass:host:port format. Pass empty string to disable proxy
func (api *Cmtt) SetProxy(proxyURL string) {
	if proxyURL == "" {
		api.client.RemoveProxy()
	} else {
		api.client.SetProxy(proxyURL)
	}
}

var authRequiredMethods = map[string]struct{}{
	"user/me":                    {},
	"user/me/comments":           {},
	"user/me/entries":            {},
	"user/me/favorites/comments": {},
	"user/me/favorites/entries":  {},
	"user/push/topic":            {},
	"payments/check":             {},
	"entry/{id}/likes":           {},
	"entry/{id}/comments":        {},
}

var platformBoundMethods = map[string]string{
	"vacancies/widget":    "vc",
	"tweets/{mode}":       "tjournal",
	"rates":               "vc",
	"news/default/recent": "vc",
}

func newInstance(platform string) (*Cmtt, error) {
	if platform != "tjournal" && platform != "dtf" && platform != "vc" {
		return nil, errors.New("not supported platform, use tjournal/dtf/vc")
	}
	client := resty.New()
	client.SetHostURL(getBaseURL(platform))
	client.SetHeader("User-Agent", fmt.Sprintf("%s-app/1.0 (go-cmtt; go-cmtt/1.0; ru_RU; 0x0)", platform))
	cmtt := &Cmtt{
		client:   client,
		platform: platform,
	}
	cmtt.client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		token := resp.Header().Get("X-Device-Token")
		if token != "" {
			cmtt.client.SetHeader("X-Device-Token", token)
			cmtt.token = token
		}
		return nil
	})
	return cmtt, nil
}

func (api *Cmtt) execute(path string, pathParams, queryParams, body map[string]string) (interface{}, error) {
	if err := api.canRun(path); err != nil {
		return nil, err
	}
	req := api.client.R().SetResult(&responseResult{}).SetError(&responseError{})
	if body != nil {
		req = req.SetFormData(body)
	}
	if pathParams != nil {
		req = req.SetPathParams(pathParams)
	}
	if queryParams != nil {
		req = req.SetQueryParams(queryParams)
	}
	method := resty.MethodGet
	if body != nil {
		method = resty.MethodPost
	}
	res, err := req.Execute(method, path)
	if err != nil {
		return nil, err
	}
	if res.StatusCode() == 503 || res.StatusCode() == 504 {
		return nil, errors.New("timeout")
	}
	error := res.Error().(*responseError)
	if error.Error.Code != 0 {
		return nil, errors.New(error.Message)
	}
	result := res.Result().(*responseResult).Result
	return result, nil
}

func (api *Cmtt) canRun(path string) error {
	if api.token == "" {
		if _, ok := authRequiredMethods[path]; ok {
			return fmt.Errorf("authentication is required for the `%s` method", path)
		}
	}
	platform := platformBoundMethods[path]
	if platform != "" && platform != api.platform {
		return fmt.Errorf("the `%s` method is available on the `%s` platform only", path, platform)
	}
	return nil
}

func getBaseURL(platform string) string {
	return fmt.Sprintf("https://api.%s.ru/v1.4/", platform)
}
