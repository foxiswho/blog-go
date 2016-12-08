package admin

import (
	"strconv"
	"strings"

	"fox/models"

	"github.com/astaxie/beego"
	"fox/util/crypt"
	"fox/service"
)

const (
	SUCCESS = "success"
)

type BaseController struct {
	beego.Controller
	admUser        *models.Admin   // 当前登录的用户id
	controllerName string          // 控制器名
	actionName     string          // 动作名
	openPerm       map[string]bool // 公开的权限
}

/*
登陆鉴权等操作，
测试开发的时候可以注释这个方法，方便测试
*/
func (this *BaseController) Prepare() {
	//this.Ctx.Output.Header("Author", "")
	//this.Ctx.Output.Header("Access-Control-Allow-Origin", "")

	//获取请求方法名称
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = controllerName
	this.actionName = actionName

	//判断是否是不需要鉴权的公共操作
	if this.isOpenPerm() {
		return
	}

	//登录校验
	token := this.Ctx.GetCookie("token")
	if admUser := validateToken(token, this.getClientIp()); admUser == nil {
		this.redirect(beego.URLFor("LoginController.Tologin"))
	} else {
		this.admUser = admUser
	}

	// //TODO 暂时判断如果是admin账号登陆就不执行任何权限校验，后续改为在某个组的用户都不做校验
	// if strings.EqualFold(service.RoleService.IsAdministrator(this.admUser.Id)) {
	// 	return
	// }

	if strings.EqualFold(controllerName, "MainController") {
		return
	}

	//操作权限校验
	if ok, err := this.validateRole(); !ok {
		if this.IsAjax() {
			this.jsonResult(err.Error())
		} else {
			this.redirect(beego.URLFor("MainController.Norole"))
		}
	}

}

/**
初始化开放权限(不需要权限校验的操作,后续如果有不需要权限校验的操作都可以写在这里)
*/
func (this *BaseController) initOpenPerm() {
	this.openPerm = map[string]bool{
		"MainController.LeftMenu": true,
		"MainController.Norole":   true,
	}
}

/**
判断是否是不需要鉴权的公共操作
*/
func (this *BaseController) isOpenPerm() bool {
	//如果是登陆相关操作则不进行登陆鉴权和权限鉴权等操作
	if strings.EqualFold(this.controllerName, "logincontroller") {
		return true
	}
	this.initOpenPerm()
	key := this.controllerName + "." + this.actionName
	if this.openPerm[key] {
		return true
	}
	return false
}

/**
token 校验，判断是否登录
*/
func x(token, currentIp string) *models.Admin {
	Dtoken, err := crypt.DeCrypt(token)
	if err != nil {
		beego.Debug("token 解密失败")
		return nil
	}
	array := strings.Split(Dtoken, "|")
	if len(array) != 3 {
		beego.Debug("token 校验失败")
		return nil
	}
	uid := array[0]
	ip := array[2]
	if !strings.EqualFold(ip, currentIp) {
		//IP发生变化 强制重新登录
		beego.Debug("ip chenged")
		return nil
	}
	intid, _ := strconv.ParseInt(uid, 10, 64)
	admuser, err := service.AdmUserService.GetUserById(intid)
	if err != nil || admuser.Id < 0 {
		beego.Debug("ID error")
		return nil
	}
	return admuser
}

/**
校验权限
*/
func (this *BaseController) validateRole() (bool, error) {
	if err := service.RoleService.ValidateRole(this.controllerName, this.actionName, this.admUser.Id); err != nil {
		return false, err
	}
	return true, nil
}

/**
重定向
*/
func (this *BaseController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

/*
指定页面，并且返回公共参数
*/
func (this *BaseController) show(url string) {
	this.Data["staticUrl"] = beego.AppConfig.String("staticUrl")
	this.TplName = url
}

/**
把需要返回的结构序列化成json 输出
*/
func (this *BaseController) jsonResult(result interface{}) {
	this.Data["json"] = result
	this.ServeJSON()
	this.StopRun()
}

type Empty struct {
}

/*
 用于分页展示列表的时候的 输出json
*/
func (this *BaseController) jsonResultPager(count int, roles interface{}) {
	beego.Debug("分页数据：", count, roles)
	resultMap := make(map[string]interface{}, 1)
	if count == 0 || roles == nil {
		beego.Debug("查询分页数据为空，返回默认json")
		//这里默认totle设置为1是因为easyui分页控件如果totle 为0会出现错乱
		resultMap["total"] = 1
		resultMap["rows"] = make([]Empty, 0)
	} else {
		resultMap["total"] = count
		resultMap["rows"] = roles
	}
	this.Data["json"] = resultMap
	this.ServeJSON()
	this.StopRun()
}
/**
token 校验，判断是否登录
*/
func validateToken(token, currentIp string) *model.Admuser {

	Dtoken, err := crypt.DeCrypt(token)
	if err != nil {
		beego.Debug("token 解密失败")
		return nil
	}
	array := strings.Split(Dtoken, "|")
	if len(array) != 3 {
		beego.Debug("token 校验失败")
		return nil
	}
	userid := array[0]
	ip := array[2]
	if !strings.EqualFold(ip, currentIp) {
		//IP发生变化 强制重新登录
		beego.Debug("ip chenged")
		return nil
	}
	intid, _ := strconv.ParseInt(userid, 10, 64)
	admuser, err := service.AdmUserService.GetUserById(intid)
	if err != nil || admuser.Id < 0 {
		beego.Debug("ID error")
		return nil
	}
	return admuser
}