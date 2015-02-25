package api

import (
	"fmt"
	"github.com/Centny/fvm/conf"
	_ "github.com/Centny/fvm/test"
	"github.com/Centny/gwf/routing/httptest"
	"github.com/Centny/gwf/util"
	"net/http"
	"os"
	"testing"
)

func TestUpload(t *testing.T) {
	ReloadFVM()
	ts := httptest.NewServer(ULoad)
	mv, err := ts.PostF2("?name=abc&ver=0.0.0", "file", "api.go", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if mv.IntVal("code") != 0 {
		t.Error(mv)
		return
	}
	mv, err = ts.PostF2("?name=abc&ver=0.0.0", "file", "api.go", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if mv.IntVal("code") != 0 {
		t.Error(mv)
		return
	}
	mv, err = ts.PostF2("?name=abc&ver=1.0.0", "file", "api.go", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if mv.IntVal("code") != 0 {
		t.Error(mv)
		return
	}
	fmt.Println(ts.G("?name=abc"))
	fmt.Println(ts.G("?name=abc&ver=sdsf.sds"))
	fmt.Println(ts.G("?name=abc&ver=sdsf.sds"))
	fmt.Println(ts.G("?name=abc&ver=123456789012345678901234567890"))
	fmt.Println(ts.G("?name=abc&ver=1.0.0"))
	os.Chmod("/tmp/fvm/fvm.json", 0)
	mv, err = ts.PostF2("?name=abc&ver=2.0.0", "file", "api.go", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if mv.IntVal("code") == 0 {
		t.Error(mv)
		return
	}

	//
	func() {
		defer func() {
			fmt.Println(recover())
		}()
		ReloadFVM()
	}()
	os.Chmod("/tmp/fvm/fvm.json", os.ModePerm)
}

func TestFVM_C(t *testing.T) {
	ReloadFVM()
	// os.RemoveAll("/tmp/fvm")
	ts := httptest.NewMuxServer()
	ts.Mux.HFunc("^/api/uload(\\?.*)?$", ULoad)
	ts.Mux.Handler("^.*$", http.FileServer(http.Dir(conf.WDir())))
	fmt.Println(ts.PostF2("/api/uload?name=a1&ver=0.0.0", "file", "../fvm/fvm.go", nil))
	fmt.Println(ts.PostF2("/api/uload?name=a1&ver=1.0.0", "file", "../fvm/fvm.go", nil))
	fmt.Println(ts.PostF2("/api/uload?name=a2&ver=0.0.0", "file", "../srv/srv_test.go", nil))
	fmt.Println(ts.PostF2("/api/uload?name=a3&ver=0.0.0", "file", "../api/api.go", nil))
	fmt.Println(ts.PostF2("/api/uload?name=a4&ver=0.0.0", "file", "../api/api_test.go", nil))
	fmt.Println(ts.PostF2("/api/uload?name=a5&ver=0.0.0", "file", "../api/store.go", nil))
	fmt.Println(ts.PostF2("/api/uload?name=a6&ver=0.0.0", "file", "../conf/conf.go", nil))
	os.RemoveAll("/tmp/fvm_a")
	os.RemoveAll("/tmp/fvm/a6")
	os.MkdirAll("/tmp/fvm_a", os.ModePerm)
	util.FWrite("/tmp/fvm_a/fvm.json", util.S2Json(map[string]interface{}{
		"a1": ">=0.0.0",
		"a2": "0.0.0",
		"a3": ">=1.0.0",
		"a4": "1.0.0",
		"a5": ">=a.0.0",
		"ax": "0.0.0",
		"ab": map[string]string{},
		"a6": "0.0.0",
	}))
	err := FVM_C(ts.URL, "/tmp/fvm_a")
	if err != nil {
		t.Error(err.Error())
	}
	err = FVM_C(ts.URL, "/tmp/fvm_a")
	if err != nil {
		t.Error(err.Error())
	}
	os.Remove("/tmp/fvm_a/fvm.json")
	util.FWrite("/tmp/fvm_a/fvm.json", util.S2Json(map[string]interface{}{
		"a1": ">=0.0.0",
		"a2": "0.0.0",
	}))
	nm, _ := util.NewMap("/tmp/fvm_a/.fvm")
	nm["skk"] = util.Map{}
	util.FWrite("/tmp/fvm_a/.fvm", util.S2Json(nm))
	err = FVM_C(ts.URL, "/tmp/fvm_a")
	if err != nil {
		t.Error(err.Error())
	}
	FVM_C(ts.URL+"/ss", "/tmp/fvm_a")
	FVM_C(ts.URL, "/tmp/fvm_b")
}
