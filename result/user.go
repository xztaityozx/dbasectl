package result

import "time"

type User struct {
	Result
	Name            string
	ProfileImageUrl string
}

type Users []User

func (u Users) String() string {
	panic("implement me")
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
