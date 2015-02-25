package main

import (
	"fmt"
	"github.com/Centny/fvm/api"
	"os"
)

func usage() {
	fmt.Println(`Usage:
			fvm -c <server addr> [<target dir>]
			fvm -u <server addr> <name> <version> <file path>
			`)
}
func main() {
	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "-c":
		if len(os.Args) > 3 {
			api.FVM_C(os.Args[2], os.Args[3])
		} else {
			api.FVM_C(os.Args[2], ".")
		}
	case "-u":
		if len(os.Args) < 6 {
			usage()
			os.Exit(1)
		}
		api.FVM_U(os.Args[2], os.Args[3], os.Args[4], os.Args[5])
	}
}
