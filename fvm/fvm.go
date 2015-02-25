package main

import (
	"fmt"
	"github.com/Centny/fvm/api"
	"os"
)

func usage() {
	fmt.Println(`Usage:
	fvm [-c [<target dir>]]					sync remote file to local.
	fvm -u <server addr> <name> <version> <file path>	upload file to reponsity.
			`)
}
func main() {
	if len(os.Args) < 2 {
		api.FVM_C(".")
		return
	}
	switch os.Args[1] {
	case "-c":
		if len(os.Args) > 2 {
			api.FVM_C(os.Args[2])
		} else {
			api.FVM_C(".")
		}
	case "-u":
		if len(os.Args) < 6 {
			usage()
			os.Exit(1)
		}
		api.FVM_U(os.Args[2], os.Args[3], os.Args[4], os.Args[5])
	default:
		usage()
	}
}
