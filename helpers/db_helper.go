package helpers

import (
	"errors"
	"fmt"
)

func GenerateId(n int) (res string, err error) {
	if n > 999 {
		err = errors.New("department no is reach maximum")
	}

	res = fmt.Sprintf("d%03d", n+1)
	return
}
