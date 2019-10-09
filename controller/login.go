package controller

import (
	"github.com/pefish/go-core/api-session"
)

type LoginControllerClass struct {
}

var LoginController = LoginControllerClass{}


type LoginParam struct {
	Username string `json:"username" validate:"required" desc:"用户名"`
}
func (this *LoginControllerClass) Login(apiSession *api_session.ApiSessionClass) interface{} {
	params := LoginParam{}
	apiSession.ScanParams(&params)



	return params
}


func (this *LoginControllerClass) LoginGet(apiSession *api_session.ApiSessionClass) interface{} {
	apiSession.Ctx.Header("Content-Type", "text/html; charset=utf-8")
	str := `
	<html><title>title</title>
		<body>
			<h1 style="color: red;">测试</h1>
		</body>
	</html>
`
	apiSession.Ctx.Write([]byte(str))
	return nil
}
