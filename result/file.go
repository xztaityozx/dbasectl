package result

import (
	"fmt"
	"strings"
	"time"
)

type File struct {
	Result
	Name      string
	Size      int
	Url       string
	Markdown  string
	CreatedAt *time.Time
}

func (f File) String() string {
	return fmt.Sprint(f.Id, f.Name, f.Size, f.Url, f.Markdown, f.CreatedAt.Format(time.RFC3339))
}

type Files []File

func (f Files) String() string {
	var box []string
	for _, v := range f {
		box = append(box, v.String())
	}

	return strings.Join(box, "\n")
}
