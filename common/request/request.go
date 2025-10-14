package request

import (
	"fmt"
	"io"
	"net/http"
)

type Request interface {
	Get() error
	Post() error
}

func NewHttpRequest() HttpRequest {
	return HttpRequest{}
}

type HttpRequest struct {
	url     string
	baseurl string
	pathurl string
}

const AUTHORIZATION = "KrS5n3w6vPyicLgUE-NYBnpr332NPT93ma2GuXot:RlcQxEvIvljfn8oMN46U-iGjJ4Y=:eyJob3N0IjoiIiwiZXhwaXJlIjoxNzYwNTc5MzI3LCJlbWFpbCI6IiJ9"

// func (r *HttpRequest) Init() {
// 	r.InitUrl(baseurl string, pathurl string)
// }

func (r *HttpRequest) Get() (string, error) {
	client := &http.Client{}

	req, err := r.initRequest()
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (r *HttpRequest) initRequest() (*http.Request, error) {
	req, err := http.NewRequest("GET", r.GetUrl(), nil)
	if err != nil {
		return &http.Request{}, err
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh-Hans;q=0.9")
	req.Header.Set("Authorization", AUTHORIZATION)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Cookie", "sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%22zhaoshihao%40qiniu.com%22%2C%22%24device_id%22%3A%22193190fd36025f-055169b43b0c62-4e193001-1296000-193190fd361465%22%2C%22props%22%3A%7B%22%24latest_referrer%22%3A%22%22%2C%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfY29va2llX2lkIjoiMTkzMWE4ZDI2ZDNiZDEtMGQ1NmMxNmMxNmMxNmMtMzg2MjZiNGItMTI5NjAwMC0xOTMxYThkMjZkNGIwOSIsIiRpZGVudGl0eV9sb2dpbl9pZCI6InpoYW9zaGloYW9AcWluaXUuY29tIn0%3D%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%22zhaoshihao%40qiniu.com%22%7D%2C%22first_id%22%3A%22193190fd36025f-055169b43b0c62-4e193001-1296000-193190fd361465%22%7D")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.1 Safari/605.1.15")
	return req, nil
}

func (r *HttpRequest) InitUrl(baseurl string, pathurl string) {
	r.baseurl = baseurl
	r.pathurl = pathurl
	r.url = fmt.Sprintf("%v%v", r.baseurl, r.pathurl)
}

func (r *HttpRequest) GetUrl() string {
	return r.url
}

func (r *HttpRequest) Post() error {
	return nil
}
