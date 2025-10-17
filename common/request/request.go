package request

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Request interface {
	Get() error
	Post() error
}

func NewHttpRequest() HttpRequest {
	return HttpRequest{}
}

type HttpRequest struct {
	url      string
	baseurl  string
	pathurl  string
	Response string
}

func (r *HttpRequest) Get(url string, auth string, cookies string) (string, error) {
	client := &http.Client{}

	req, err := r.initRequest(url, auth, cookies)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode > 300 {
		return "", fmt.Errorf("status code: %v", resp.StatusCode)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	r.Response = string(body)
	return r.Response, nil
}

func (r *HttpRequest) initRequest(url string, auth string, cookies string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &http.Request{}, err
	}

	req.Header.Set("Accept", "*/*")
	// 改为空 否则请求会乱码
	req.Header.Set("Accept-Encoding", "")
	req.Header.Set("Accept-Language", "zh-CN,zh-Hans;q=0.9")
	req.Header.Set("Authorization", auth)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Cookie", cookies)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.1 Safari/605.1.15")
	return req, nil
}

func (r *HttpRequest) InitUrl(baseurl string, pathurl string) (*url.URL, error) {
	r.baseurl = baseurl
	r.pathurl = pathurl
	r.url = fmt.Sprintf("%v%v", r.baseurl, r.pathurl)

	parsedUrl, err := url.Parse(r.url)
	if err != nil {
		return nil, err
	}

	return parsedUrl, nil
}

func (r *HttpRequest) GetUrl() string {
	return r.url
}

func (r *HttpRequest) Post() error {
	return nil
}
