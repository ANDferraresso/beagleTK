package paginator

import (
)

type Paginator struct {
	PagesNum  int64
	Page      int64
	FirstPage int64
	PrevPage  int64
	NextPage  int64
	LastPage  int64
}
