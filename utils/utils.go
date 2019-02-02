package utils

import (
	"fmt"
	"github.com/jinzhu/copier"
	"strings"
)

func IsDuplicateError(err error) bool {
	return strings.Contains(err.Error(), "Error 1062: Duplicate entry")
}

func MustCopy(to, from interface{}) {
	err := copier.Copy(to, from)
	if err != nil {
		panic(fmt.Sprintf("copy error: %+v\n", err))
	}
}
