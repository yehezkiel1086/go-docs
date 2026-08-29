package util

import (
	"fmt"
)

func GeneratePath(base, path string, param uint) (string, error) {
	url := fmt.Sprintf("%s%s/%d", base, path, param)
	return url, nil
}
