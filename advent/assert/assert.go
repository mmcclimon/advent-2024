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
