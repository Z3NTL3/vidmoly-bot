package globals

import (
	"fmt"
	"os"
)

func ErrHandler(err error) {
	fmt.Println("\033[31m", err, "\033[0m")
	os.Exit(-1)
}