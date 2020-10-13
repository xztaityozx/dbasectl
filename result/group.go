package result

import "time"

type Groups []Group

func (g Groups) String() string {
	panic("implement me")
}

type Group struct {
	Result
	Name string
}

type GroupDetail struct {
	Group
	Description    string
	PostsCount     int
	LastActivityAt *time.Time
	CreatedAt      *time.Time
	Users          []User
}
