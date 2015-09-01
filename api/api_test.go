package api

import (
	"fmt"
	"github.com/Centny/fvm/conf"
	_ "github.com/Centny/fvm/test"
	"github.com/Centny/gwf/routing"
	"github.com/Centny/gwf/routing/httptest"
	"github.com/Centny/gwf/util"
	"net/http"
	"os"
	"testing"
)

func TestUpload(t *testing.T) {
	ReloadFVM()
	ts := httptest.NewServer(func(hs *routing.HTTPSession) routing.HResult {
		hs.ParseQuery()
		return ULoad(hs)
	})
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
	//
}

func TestFVM(t *testing.T) {
	ReloadFVM()
	// os.RemoveAll("/tmp/fvm")
	ts := httptest.NewMuxServer()
	Handle(ts.Mux)
	ts.Mux.Handler("^.*$", http.FileServer(http.Dir(conf.WDir())))
	fmt.Println(ts.PostF2("/api/uload?name=a1&ver=0.0.0", "file", "../fvm/fvm.go", nil))
	fmt.Println(ts.PostF2("/api/uload?name=a1&ver=1.0.0", "file", "../fvm/fvm.go", nil))
	fmt.Println(ts.PostF2("/api/uload?name=a2&ver=0.0.0", "file", "../srv/srv_test.go", nil))
	fmt.Println(ts.PostF2("/api/uload?name=a3&ver=0.0.0", "file", "../api/api.go", nil))
	fmt.Println(ts.PostF2("/api/uload?name=a4&ver=0.0.0", "file", "../api/api_test.go", nil))
	fmt.Println(ts.PostF2("/api/uload?name=a5&ver=0.0.0", "file", "../api/store.go", nil))
	fmt.Println(ts.PostF2("/api/uload?name=a6&ver=0.0.0", "file", "../conf/conf.go", nil))
	fmt.Println(ts.PostF2("/api/uload?name=a61&ver=0.0.0", "file", "../conf/conf.go", nil))
	util.FWrite("t.txt", "abccc->>>")
	util.Zip("t.zip", ".", "t.txt")
	fmt.Println(ts.PostF2("/api/uload?name=a7&ver=0.0.0", "file", "t.zip", nil))
	//
	//test raw download.
	err := util.DLoad("/tmp/aaa.x", "%v/raw/a1", ts.URL)
	if err != nil {
		t.Error(err.Error())
		return
	}
	util.DLoad("/tmp/aaa.x", "%v/raw/asdd", ts.URL)
	FVM.SetVal("kjy", util.Map{
		"VER": "10.0",
	})
	util.DLoad("/tmp/aaa.x", "%v/raw/kjy", ts.URL)
	//
	//
	os.RemoveAll("/tmp/fvm_a")
	os.RemoveAll("/tmp/fvm/a6")
	os.MkdirAll("/tmp/fvm_a", os.ModePerm)
	util.FWrite("/tmp/fvm_a/fvm.json", util.S2Json(
		map[string]interface{}{
			"srv": ts.URL,
			"fvm": map[string]interface{}{
				"a1":      ">=0.0.0",
				"a2":      "0.0.0",
				"a3":      ">=1.0.0",
				"a4":      "1.0.0",
				"a5":      ">=a.0.0",
				"ax":      "0.0.0",
				"ab":      map[string]string{},
				"a6":      "0.0.0",
				"a61@zip": "0.0.0",
				"a7@zip":  "0.0.0",
			},
		}))
	FVM_C("/tmp/fvm_a")
	FVM_C("/tmp/fvm_a")
	//
	os.Remove("/tmp/fvm_a/fvm.json")
	util.FWrite("/tmp/fvm_a/fvm.json", util.S2Json(
		map[string]interface{}{
			"srv": "ts.URL",
			"fvm": map[string]interface{}{
				"a1": ">=0.0.0",
				"a2": "0.0.0",
			},
		}))
	FVM_C("/tmp/fvm_a")
	//
	os.Remove("/tmp/fvm_a/fvm.json")
	util.FWrite("/tmp/fvm_a/fvm.json", util.S2Json(
		map[string]interface{}{
			"srv": ts.URL,
			"fvm": map[string]interface{}{
				"a1": ">=0.0.0",
				"a2": "0.0.0",
			},
		}))
	nm, _ := util.NewMap("/tmp/fvm_a/.fvm")
	nm["skk"] = util.Map{}
	util.FWrite("/tmp/fvm_a/.fvm", util.S2Json(nm))
	FVM_C("/tmp/fvm_a")
	FVM_C("/tmp/fvm_b")
	//
	os.Remove("/tmp/fvm_a/fvm.json")
	util.FWrite("/tmp/fvm_a/fvm.json", util.S2Json(
		map[string]interface{}{
			"fvm": map[string]interface{}{
				"a1": ">=0.0.0",
				"a2": "0.0.0",
			},
		}))
	FVM_C("/tmp/fvm_a")
	//
	os.Remove("/tmp/fvm_a/fvm.json")
	util.FWrite("/tmp/fvm_a/fvm.json", util.S2Json(
		map[string]interface{}{
			"srv": ts.URL,
		}))
	FVM_C("/tmp/fvm_a")
	//
	err = FVM_U(ts.URL, "xxx", "0.0.0", "api.go")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(FVM_U(ts.URL, "xxx", "", "api.go"))
	fmt.Println(FVM_U("", "xxx", "", "api.go"))
}

// func TestFVM2(t *testing.T) {
// 	fmt.Println(FVM_U("http://192.168.2.30:9921", "xxx", "1.0.0", "api.go"))
// }
