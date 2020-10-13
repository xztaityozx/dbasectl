package result

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	Result
	Name            string
	ProfileImageUrl string
}

func (u User) String() string {
	return fmt.Sprint(u.Id, u.Name, u.ProfileImageUrl)
}

type Users []User

func (u Users) String() string {
	var sb []string
	for _, v := range u {
		sb = append(sb, v.String())
	}

	return strings.Join(sb, "\n")
}

type UserDetail struct {
	User
	Username              string
	Role                  string
	PostsCount            int
	LastAccessAt          *time.Time
	TowStepAuthentication bool
	Groups                Groups
}

type UserDetails []UserDetail

func (u UserDetails) String() string {
	panic("implement me")
}
