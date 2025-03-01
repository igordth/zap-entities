package writer

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"time"
)

var (
	HttpDefaultClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    5,
			IdleConnTimeout: time.Minute,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: time.Second * 3,
	}

	ErrHttpStatusCode = errors.New("status code not OK")
)

type wHttp struct {
	Client  *http.Client
	BaseUrl string
	Method  string
}

func NewHttp(client *http.Client, url, method string) io.Writer {
	return &wHttp{client, url, method}
}

func (w wHttp) getUrl(p []byte) string {
	switch w.Method {
	case http.MethodPost, http.MethodPut:
		return w.BaseUrl
	default:
		return fmt.Sprintf("%s?%s", w.BaseUrl, string(p))
	}
}

func (w wHttp) getBody(p []byte) io.Reader {
	switch w.Method {
	case http.MethodPost, http.MethodPut:
		return bytes.NewReader(p)
	default:
		return nil
	}
}

func (w wHttp) Write(p []byte) (n int, err error) {
	// request
	req, err := http.NewRequest(w.Method, w.BaseUrl, w.getBody(p))
	if err != nil {
		return 0, errors.Wrap(err, "http.NewRequest")
	}

	// send request
	res, err := w.Client.Do(req)
	if err != nil {
		return 0, errors.Wrap(err, "http.Client.Do")
	}
	err = res.Body.Close()
	if err != nil {
		return 0, errors.Wrap(err, "response.Body.Close")
	}

	// check request status
	switch {
	case res.StatusCode != 200:
		return 0, errors.Wrap(ErrHttpStatusCode, "response.Status = "+res.Status)
	}

	return len(p), nil
}
