package result_test

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/xztaityozx/dbasectl/result"
	"math/big"
	mrand "math/rand"
	"strings"
	"testing"
	"time"
)

func randomString(size int) string {
	r := make([]byte, size)
	for i := 0; i < size; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(255))
		r[i] = byte(n.Int64())
	}

	return base64.StdEncoding.EncodeToString(r)
}

func TestFile_String(t *testing.T) {
	v := time.Now()
	file := result.File{
		Result:    result.Result{Id: mrand.Int()},
		Name:      randomString(10),
		Size:      mrand.Int(),
		Url:       randomString(40),
		Markdown:  randomString(40),
		CreatedAt: &v,
	}

	assert.Equal(t,
		fmt.Sprint(file.Id, file.Name, file.Size, file.Url, file.Markdown, file.CreatedAt.Format(time.RFC3339)),
		file.String())
}

func TestFiles_String(t *testing.T) {
	var files result.Files
	var expect []string
	for i := 0; i < 10; i++ {
		v := time.Now()
		f :=
			result.File{
				Result:    result.Result{Id: mrand.Int()},
				Name:      randomString(10),
				Size:      mrand.Int(),
				Url:       randomString(40),
				Markdown:  randomString(40),
				CreatedAt: &v,
			}
		files = append(files, f)

		expect = append(expect, f.String())
	}

	if files == nil {
		t.Fail()
		return
	}

	assert.Equal(t, strings.Join(expect, "\n"), files.String())
}
