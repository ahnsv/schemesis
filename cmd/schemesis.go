package main

import (
	"fmt"

	"github.com/ahnsv/schemesis/cli"
)

func main() {
	err := cli.RootCmd.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
