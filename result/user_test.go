package result_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/xztaityozx/dbasectl/result"
	"math/rand"
	"strings"
	"testing"
)

func TestUser_String(t *testing.T) {
	u := result.User{
		Result:          result.Result{Id: rand.Int()},
		Name:            randomString(10),
		ProfileImageUrl: randomString(10),
	}

	assert.Equal(t, fmt.Sprint(u.Id, u.Name, u.ProfileImageUrl), u.String())
}

func TestUsers_String(t *testing.T) {
	var us result.Users
	var expect []string

	for i := 0; i < 20; i++ {
		u := result.User{
			Result:          result.Result{Id: rand.Int()},
			Name:            randomString(10),
			ProfileImageUrl: randomString(10),
		}
		us = append(us, u)
		expect = append(expect, u.String())
	}

	if us == nil {
		t.Fail()
		return
	}

	assert.Equal(t, strings.Join(expect, "\n"), us.String())
}
