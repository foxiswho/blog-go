package test

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/foxiswho/blog-go/blog/fox/log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	web.TestBeegoInit(apppath)
}


// TestMain is a sample to run an endpoint test
func Test(t *testing.T) {
	log.Info("xxxx")
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(w, r)

	//web.Trace("testing", "TestMain", "Code[%d]\n%s", w.Code, w.Body.String())
	//
	//Convey("Subject: Test Station Endpoint\n", t, func() {
	//       Convey("Status Code Should Be 200", func() {
	//               So(w.Code, ShouldEqual, 200)
	//       })
	//       Convey("The Result Should Not Be Empty", func() {
	//               So(w.Body.Len(), ShouldBeGreaterThan, 0)
	//       })
	//})
}

