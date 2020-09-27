package request

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/xztaityozx/dbasectl/config"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	as := assert.New(t)
	var dummy EndPoint = "dummyEndPoint"
	expectEmpty := Request{}

	t.Run("tokenがないコンフィグで生成したとき", func(t *testing.T) {
		cfg := config.Config{Name: "name"}

		req, err := New(cfg, http.MethodPost, dummy)
		as.Equal(expectEmpty, req, "空なRequestが帰ってくるべき")
		as.Error(err, "エラーが返されるべき")
	})

	t.Run("nameがないコンフィグで生成した時", func(t *testing.T) {
		cfg := config.Config{Token: "token"}

		req, err := New(cfg, http.MethodPost, dummy)
		as.Equal(expectEmpty, req, "空なRequestが帰ってくるべき")
		as.Error(err, "エラーが返されるべき")
	})

	t.Run("間違ったエンドポイントが指定されたとき", func(t *testing.T) {
		cfg := config.Config{Token: "token", Name:"name"}

		req, err := New(cfg, http.MethodPost, dummy)
		as.Equal(expectEmpty, req, "空なRequestが帰ってくるべき")
		as.Error(err, "エラーが返されるべき")
	})

	t.Run("エンドポイントとコンフィグが正しいとき", func(t *testing.T) {
		cfg := config.Config{Token: "token", Name:"name"}

		for method, eps := range allowedDictionary {
			for _, ep := range eps {
				req, err := New(cfg, method, ep)

				as.Nil(err)
				as.Equal(Request{
					cfg:    cfg,
					req:    nil,
					url:    fmt.Sprintf("https://api.docbase.io/teams/%s/%s", cfg.Name, ep),
					logger: nil,
				}, req, "Requestが返される")
			}
		}
	})
}

func TestRequest_WithLogger(t *testing.T) {
	req := Request{logger: nil}

	l := logrus.New()
	req.WithLogger(l)

	assert.Equal(t, l, req.logger)
}

func TestRequest_SetBody(t *testing.T) {
	body := []byte("This is Body")
	req := Request{body: nil}
	as := assert.New(t)

	req.SetBody(bytes.NewBuffer(body))

	reqBody, err := ioutil.ReadAll(req.body)
	as.Nil(err)

	assert.Equal(t, body, reqBody)
}
