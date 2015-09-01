package conf

import (
	"github.com/Centny/gwf/util"
)

var Cfg *util.Fcfg = util.NewFcfg3()

func ListenAddr() string {
	return Cfg.Val("LISTEN_ADDR")
}

func WDir() string {
	return Cfg.Val("W_DIR")
}
