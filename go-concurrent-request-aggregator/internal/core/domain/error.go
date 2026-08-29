package domain

import (
	"fmt"
)

var (
	ErrInternalServer error = fmt.Errorf("internal server error")
	ErrNotFound error = fmt.Errorf("not found")
)
