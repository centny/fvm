package api

import (
	"github.com/Centny/fvm/conf"
	"github.com/Centny/gwf/log"
	"github.com/Centny/gwf/routing"
	"github.com/Centny/gwf/util"
	"path/filepath"
)

func ULoad(hs *routing.HTTPSession) routing.HResult {
	var name string
	var ver string
	err := hs.ValidCheckVal(`
		name,R|S,L:0;
		ver,R|S,P:^\d*[\d\.]*\d*$;
		`, &name, &ver)
	if err != nil {
		return hs.MsgResErr2(1, "arg-err", err)
	}
	mv := FVM.MapVal(name)
	if mv == nil {
		mv = util.Map{}
	}
	iv, err := util.ChkVer(ver, mv.StrVal("VER"))
	if err != nil {
		return hs.MsgResErr2(1, "arg-err", err)
	}
	fn, _, sha1, _, err := hs.RecFv2("file", conf.WDir()+"/"+name+"/"+ver+"/")
	if err != nil {
		log.E("add new version by name(%v),ver(%v) error:%v", name, ver, err.Error())
		return hs.MsgResErr2(1, "srv-err", err)
	}
	fpath := name + "/" + ver + "/" + fn
	mv.SetVal(ver, util.Map{
		"SHA1": sha1,
		"PATH": fpath,
	})
	if iv > -1 {
		mv.SetVal("VER", ver)
		mv.SetVal("SHA1", sha1)
	}
	FVM.SetVal(name, mv)
	err = StoreFVM()
	if err == nil {
		log.D("add new version by name(%v),ver(%v),sha1(%v) to path(%v)->OK", name, ver, sha1, fpath)
		return hs.MsgRes("OK")
	} else {
		log.E("add new version by name(%v),ver(%v),sha1(%v) to path(%v)->ERR:%v", name, ver, sha1, fpath, err.Error())
		return hs.MsgResErr2(1, "srv-err", err)
	}
}

func Raw(hs *routing.HTTPSession) routing.HResult {
	_, fn := filepath.Split(hs.R.URL.Path)
	mv := FVM.MapVal(fn)
	if mv == nil {
		hs.W.WriteHeader(404)
		return routing.HRES_RETURN
	}
	tfv := mv.MapVal(mv.StrVal("VER"))
	if tfv == nil {
		hs.W.WriteHeader(404)
		return routing.HRES_RETURN
	}
	hs.Redirect("../" + tfv.StrVal("PATH"))
	return routing.HRES_RETURN
}
