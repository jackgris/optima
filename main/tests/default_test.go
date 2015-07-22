package test

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	_ "github.com/jackgris/optima/routers"

	"github.com/astaxie/beegae"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beegae.TestBeegoInit(apppath)
}

// TestRegister test the GET request of registration form
func TestRegister(t *testing.T) {
	r, _ := http.NewRequest("GET", "/register", nil)
	w := httptest.NewRecorder()
	beegae.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Test registration request\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

// TestHome test the GET request of home
func TestHome(t *testing.T) {
	r, _ := http.NewRequest("GET", "/home", nil)
	w := httptest.NewRecorder()
	beegae.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Test registration request\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

// TestNewUser test the GET request of new user form
func TestNewUser(t *testing.T) {
	r, _ := http.NewRequest("GET", "/user/newuser", nil)
	w := httptest.NewRecorder()
	beegae.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Test registration request\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
