package main

import (
	"fmt"
	"github.com/Centny/fvm/api"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:fvm <server addr> [<target dir>]")
		return
	}
	if len(os.Args) > 2 {
		api.FVM_C(os.Args[1], os.Args[2])
	} else {
		api.FVM_C(os.Args[1], ".")
	}
}
