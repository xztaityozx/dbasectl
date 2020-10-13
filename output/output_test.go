package output_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/xztaityozx/dbasectl/output"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

type TestResult struct {
	A string
	B int
}

func (t TestResult) String() string {
	return fmt.Sprint(t.A, t.B)
}

func TestFormat_Print(t *testing.T) {
	tr := TestResult{A: "はい", B: 10}
	as := assert.New(t)

	t.Run("formatがYaml", func(t *testing.T) {
		var b []byte
		buf := bytes.NewBuffer(b)
		f := output.Yaml
		as.Nil(f.Print(tr, buf))

		actual, _ := ioutil.ReadAll(buf)

		expect, _ := yaml.Marshal(tr)
		as.Equal(expect, actual)
	})
	t.Run("formatがJson", func(t *testing.T) {
		var b []byte
		buf := bytes.NewBuffer(b)
		f := output.Json
		as.Nil(f.Print(tr, buf))

		actual, _ := ioutil.ReadAll(buf)
		expect, _ := json.MarshalIndent(tr, "", "  ")
		as.Equal(expect, actual)
	})
	t.Run("formatがText", func(t *testing.T) {
		var b []byte
		buf := bytes.NewBuffer(b)
		f := output.Text
		as.Nil(f.Print(tr, buf))

		actual, _ := ioutil.ReadAll(buf)

		as.Equal([]byte(tr.String()), actual)
	})
	t.Run("よくわからんフォーマットの時", func(t *testing.T) {
		var f output.Format = 10000
		as.Error(f.Print(tr, nil))
	})
}

func TestNew(t *testing.T) {
	for _, v := range []struct {
		a string
		b output.Format
	}{
		{a: "Yaml", b: output.Yaml},
		{a: "yaml", b: output.Yaml},
		{a: "Json", b: output.Json},
		{a: "json", b: output.Json},
		{a: "Text", b: output.Text},
		{a: "text", b: output.Text},
		{a: "なんだこれは", b: output.Text},
	} {
		assert.Equal(t, v.b, output.New(v.a))
	}
}
