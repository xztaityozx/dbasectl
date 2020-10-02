package request

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/xztaityozx/dbasectl/config"
)

type EndPoint string

const (
	Upload EndPoint = "attachments"
)

type Request struct {
	cfg    config.Config
	req    *http.Request
	url    string
	logger *logrus.Logger
	body   io.Reader
	method string
}

// New は DocBaseの ep に method でアクセスする Request を返す
func New(cfg config.Config, method string, ep EndPoint) (Request, error) {
	if cfg.Token == "" {
		return Request{}, fmt.Errorf("access token is empty")
	}

	if cfg.Name == "" {
		return Request{}, fmt.Errorf("team name is empty")
	}

	if !isAllowedEndPoint(method, ep) {
		return Request{}, fmt.Errorf("%s method is not allowed to %s endpoint", method, ep)
	}

	url := fmt.Sprintf("https://api.docbase.io/teams/%s/%s", cfg.Name, ep)

	return Request{cfg: cfg, req: nil, url: url, logger: nil, method: method}, nil
}

// WithLogger は logrus.Logger をセットした Request を返す
func (r *Request) WithLogger(logger *logrus.Logger) *Request {
	r.logger = logger
	return r
}

// SetBody はリクエストボディをセットする
func (r *Request) SetBody(body io.Reader) *Request {
	r.body = body
	return r
}

// Build は http.Request を生成する
func (r *Request) Build() error {

	if r.method == "" {
		return fmt.Errorf("empty http method")
	}

	if _, ok := allowedDictionary[r.method]; !ok {
		return fmt.Errorf("invalid http method: %s", r.method)
	} else {
		r.info("method: ", r.method)
	}

	// セットされているURLが正しいっぽいかどうか
	// 少なくとも以下の正規表現を満たしていないといけない
	if ok, err := regexp.MatchString("https://api.docbase.io/teams/.+/.+", r.url); err != nil {
		return err
	} else if !ok {
		return fmt.Errorf("invalid URL: %s", r.url)
	} else {
		r.info("url: ", r.url)
	}

	var err error
	r.req, err = http.NewRequest(r.method, r.url, r.body)
	r.req.Header.Set("X-DocBaseToken", r.cfg.Token)

	// POST系はJSONを投げるので、Content-Typeを指定しておく
	if r.method == http.MethodPost {
		r.req.Header.Set("Content-Type", "application/json")
		r.info("Content-Type: application/json")
		if r.body == nil {
			r.warn("request body size is zero")
		}
	}

	return err
}

// Do は DocBaseのAPIにアクセスして、そのレスポンスボディを返す
func (r *Request) Do(ctx context.Context) (responseBody io.ReadCloser, err error) {
	r.req.WithContext(ctx)

	client := &http.Client{}
	if r.cfg.Timeout >= 0 {
		client.Timeout = time.Duration(r.cfg.Timeout * time.Millisecond)
		r.info("timeout: ", r.cfg.Timeout)
	}

	r.info("request started at ", time.Now())
	res, err := client.Do(r.req)
	r.info("request ended at ", time.Now())

	return res.Body, err
}

func (r *Request) info(args ...interface{}) {
	if r.logger != nil {
		r.logger.Info(args...)
	}
}

func (r *Request) warn(args ...interface{}) {
	if r.logger != nil {
		r.logger.Warn(args...)
	}
}

var allowedDictionary = map[string][]EndPoint{
	http.MethodPost: {Upload},
}

func isAllowedEndPoint(method string, ep EndPoint) bool {
	d, ok := allowedDictionary[method]
	if !ok {
		return false
	}

	for _, v := range d {
		if v == ep {
			return true
		}
	}
	return false
}
