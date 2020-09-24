package request

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/xztaityozx/dbasectl/config"
)

type EndPoint string

type Request struct {
	cfg config.Config
}

func New(cfg config.Config) Request {
	return Request{cfg: cfg}
}

// Post は ep に body をPOSTする
func (r Request) Post(ctx context.Context, ep EndPoint, body io.Reader) (string, error) {
	// POSTできるEndPointかどうかチェック
	if func() bool {
		for _, v := range []EndPoint{Upload} {
			if v == ep {
				return true
			}
		}
		return false
	}() {
		return "", errors.New(fmt.Sprintf("%s is not allowed EndPoint for POST method", ep))
	}

	url, err := r.getUrl(ep)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return "", err
	}

	// アクセストークンをセット
	req.Header.Set("X-DocBaseToken", r.cfg.Token)
	// Postでは jsonを投げる
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	return string(result), err
}

func (r Request) getUrl(ep EndPoint) (string, error) {
	panic("not Imple")
}
