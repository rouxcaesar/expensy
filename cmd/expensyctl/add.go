package expensyctl

import (
	"fmt"
)

func AddCommand(args []string) {
	fmt.Println(args)
	fmt.Println(len(args))
}
