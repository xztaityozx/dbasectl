package result

type Stringer interface {
	String() string
}

type Result struct {
	Id int
}
