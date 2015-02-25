package test

import (
	"github.com/Centny/fvm/conf"
	"github.com/Centny/gwf/util"
	"os"
)

var Cfg util.Fcfg = util.Fcfg{}

func init() {
	if !Cfg.Exist("LISTEN_ADDR") {
		Cfg.SetVal("LISTEN_ADDR", "9921")
	}
	if !Cfg.Exist("W_DIR") {
		Cfg.SetVal("W_DIR", "/tmp/fvm")
	}
	conf.Cfg = Cfg
	os.RemoveAll("/tmp/fvm")
	os.Mkdir("/tmp/fvm", os.ModePerm)
	// util.FWrite("/tmp/fvm/fvm.json", "{}")
}
