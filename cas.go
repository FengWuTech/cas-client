package cas_client

import (
	"crypto/md5"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type Cas struct {
	BaseURL string
	AppID   string
	ApiKey  string
}

func New(baseURL string, appID string, apiKey string) *Cas {
	return &Cas{
		BaseURL: baseURL,
		AppID:   appID,
		ApiKey:  apiKey,
	}
}

func (cas *Cas) HttpGet(requestURI string, param map[string]interface{}, ret interface{}) {
	token := fmt.Sprintf("%x", md5.Sum([]byte(cas.AppID+cas.ApiKey)))
	gorequest.New().
		Get(cas.BaseURL+requestURI).
		Query(param).
		AppendHeader("token", token).
		AppendHeader("appcode", cas.AppID).
		EndStruct(ret)
}

func (cas *Cas) HttpPost(requestURI string, param map[string]interface{}, body interface{}, ret interface{}) {
	token := fmt.Sprintf("%x", md5.Sum([]byte(cas.AppID+cas.ApiKey)))
	gorequest.New().
		Post(cas.BaseURL+requestURI).
		Send(body).
		Query(param).
		AppendHeader("token", token).
		AppendHeader("appcode", cas.AppID).
		EndStruct(ret)
}
