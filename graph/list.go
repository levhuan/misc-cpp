package list

import(
	"fmt"
)

var (
)

const (
)

type ListData_t interface {}

type ListElement_t struct {
	next	*ListElement_t
	data	ListData_t
}

