package api

import (
	"os"
	"os/user"
	"testing"
)

func TestFVM_A(t *testing.T) {
	usr, _ := user.Current()
	fp := usr.HomeDir + "/.fvm.json"
	os.Remove(fp)
	FVM_A("/tmp/asss")
	FVM_ALL()
	os.Chmod(fp, 0)
	FVM_A("/tmp/asss")
	FVM_ALL()
	os.Remove(fp)
}
