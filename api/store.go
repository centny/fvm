package api

import (
	"fmt"
	"github.com/Centny/fvm/conf"
	"github.com/Centny/gwf/log"
	"github.com/Centny/gwf/util"
	"os"
	"path/filepath"
	"strings"
)

var FVM util.Map

func ReloadFVM() {
	var err error
	FVM, err = util.NewMap(conf.WDir() + "/fvm.json")
	if err != nil {
		panic(err.Error())
	}
}

func StoreFVM() error {
	return util.FWrite(conf.WDir()+"/fvm.json", util.S2Json(FVM))
}

func FVM_C(srv, tp string) error {
	n_fvm_l := util.Map{}
	tfvm, err := util.NewMap(tp + "/fvm.json")
	if err != nil {
		log.E("read local fvm.json error:%v", err.Error())
		return err
	}
	sfvm, err := util.HGet2(srv + "/fvm.json")
	if err != nil {
		log.E("load remote fvm.json error:%v", err.Error())
		return err
	}
	for tfn, tfv_ := range tfvm {
		tfv, ok := tfv_.(string)
		if !ok {
			log.W("invalid value for %v", tfn)
			continue
		}
		tfn = strings.Trim(tfn, " \t\n")
		tfv = strings.Trim(tfv, " \t\n")
		if !sfvm.Exist(tfn) {
			log.W("%v not exist in repository", tfn)
			continue
		}
		sfn_v := sfvm.MapVal(tfn)
		var sfv string
		if strings.HasPrefix(tfv, ">=") {
			sfv = sfn_v.StrVal("VER")
			tfv = strings.TrimPrefix(tfv, ">=")
			iv, err := util.ChkVer(sfv, tfv)
			if err != nil {
				log.W("check version err(%v) for %v ", tfn)
				continue
			}
			if iv < 0 {
				log.W("version not found for %v,expected(%v),having(%v)", tfn, tfv, sfv)
				continue
			}
		} else {
			sfv = tfv
		}
		sfn_tv := sfn_v.MapVal(sfv)
		if sfn_tv == nil {
			log.W("version not found for %v,expected(%v)", tfn, sfv)
			continue
		}
		fpath := sfn_tv.StrVal("PATH")
		_, fn := filepath.Split(fpath)
		n_fvm_l[tfn] = fn
		spath := tp + "/" + fn
		tsha1, err := util.Sha1(spath)
		if err == nil && tsha1 == sfn_tv.StrVal("SHA1") {
			log.D("%v not update", fn)
			continue
		}
		dpath := fmt.Sprintf("%v/%v", srv, sfn_tv.StrVal("PATH"))
		err = util.DLoad(spath, dpath)
		if err == nil {
			log.D("download file(%v) success", dpath)
		} else {
			log.W("download file(%v) error:%v", dpath, err.Error())
		}
	}
	o_fvm_l, err := util.NewMap(tp + "/.fvm")
	defer util.FWrite(tp+"/.fvm", util.S2Json(n_fvm_l))
	if err != nil { //old fvm not found
		return nil
	}
	for tfn, tfv_ := range o_fvm_l {
		tfv, ok := tfv_.(string)
		if !ok {
			continue
		}
		tfn = strings.Trim(tfn, " \t\n")
		tfv = strings.Trim(tfv, " \t\n")
		if n_fvm_l.StrVal(tfn) == tfv {
			continue
		}
		os.Remove(tp + "/" + tfv)
	}
	return nil
}

func FVM_U(srv, name, ver, fp string) error {
	mv, err := util.HPostF2(fmt.Sprintf("%v/api/uload?name=%v&ver=%v", srv, name, ver), nil, "file", fp)
	if err != nil {
		log.E("upload error:%v", err.Error())
		return err
	}
	if mv.IntVal("code") == 0 {
		return nil
	} else {
		log.E("upload error(%v),response:%v", mv.IntVal("code"), mv.StrVal("dmsg"))
		return util.Err("response:%v", mv.StrVal("dmsg"))
	}
}
