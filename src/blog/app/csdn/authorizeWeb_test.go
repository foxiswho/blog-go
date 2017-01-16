package csdn
import (
	"testing"
	"fmt"
)
func TestWeb(t *testing.T) {
	web:=NewAuthorizeWeb()
	ok,err:=web.SetConfig()
	if err !=nil{
		t.Fatal(err)
	}
	fmt.Println("status:",ok);
	web.SetRedirectUri("/oath_qiniu")
	web.GetAuthorizeUrl()
}