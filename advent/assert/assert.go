package assert

import (
	"fmt"
	"os"
)

func Nil(err error) {
	if err == nil {
		return
	}

	fmt.Println("fatal:", err)
	os.Exit(1)
}

func True(cond bool, msg string) {
	if cond {
		return
	}

	fmt.Println("assertion failure:", msg)
	os.Exit(1)
}
