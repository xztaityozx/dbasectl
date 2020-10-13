package result

import "fmt"

type Meta struct {
	PreviousPage string
	NextPage     string
	Total        int
}

func (m Meta) String() string {
	return fmt.Sprint(m.PreviousPage, m.NextPage, m.Total)
}
