package api

import (
	"os"
	"os/user"
	"testing"
)

func TestFVM_A(t *testing.T) {
	usr, _ := user.Current()
	fp := usr.HomeDir + "/.fvm.json"
	FVM_D("/tmp/assx")
	os.Remove(fp)
	FVM_A("/tmp/asss")
	FVM_A("/tmp/assx")
	FVM_ALL()
	FVM_D("/tmp/asss")
	FVM_ALL()
	os.Chmod(fp, 0)
	FVM_A("/tmp/asss")
	FVM_D("/tmp/assx")
	FVM_ALL()
	os.Remove(fp)
}

// func TestXX(t *testing.T) {
// 	usr, _ := user.Current()
// 	fp := usr.HomeDir + "/.fvm.json"
// 	os.Chmod(fp, 0710)
// 	fmt.Println(FVM_D("/tmp/assx"))
// }
