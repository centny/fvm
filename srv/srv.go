//Package provide server function.
//Author:Centny
package srv

import (
	"fmt"
	"github.com/Centny/fvm/conf"
	"github.com/Centny/gwf/log"
	"net/http"
	"sync"
)

var lock sync.WaitGroup
var s_running bool

func run(args []string) {
	defer StopSrv()
	cfile := "conf/fvm.properties"
	if len(args) > 1 {
		cfile = args[1]
	}
	fmt.Println("Using config file:", cfile)
	err := conf.Cfg.InitWithFilePath(cfile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//
	log.I("Config:\n%v", conf.Cfg.Show())
	//test connect
	//
	mux := http.NewServeMux()
	HSrvMux(mux, "", conf.WDir())
	log.D("running web server on %s", conf.ListenAddr())
	s := http.Server{Addr: conf.ListenAddr(), Handler: mux}
	err = s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

//run the server.
func RunSrv(args []string) {
	s_running = true
	lock.Add(1)
	go run(args)
	lock.Wait()
	s_running = false
}

//stop the server.
func StopSrv() {
	if s_running {
		lock.Done()
	}
}
