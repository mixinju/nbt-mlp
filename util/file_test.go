package util

import (
	"fmt"
	"testing"

	"nbt-mlp/util/errno"
)

func TestReadUserFormExcel(t *testing.T) {
	path := "./demo.xlsx"

	u, err := ReadUserInfoFromExcel(path)
	if err != *errno.OK {
		log.Sugar().Infof("{%s}", err.Message)
	}

	for _, uu := range u {
		fmt.Println(uu.Id, uu.Name, uu.ClassName)
	}
}
