
#GO语言博客
基本功能已 实现

#功能说明
 * 文章 增删该查
 * 图片 上传
 * markdown 编辑器
 * 管理员密码修改
 * 站点属性配置
 * 类别数据 增删改查
 * 博客前台显示 分页
 * 七牛云存储

#未来
 * 其他模块继续完善
    * 省市区
    * 角色和权限
    * 管理员
    * 菜单
    * 缓存
    * 标签
    * 附件
    * 。。。。。。
    
#问题
 * 1.beego 配置文件竟然不支持注释，是不是还没找到正确的注释使用方式呢？

#编译
##GO环境变量
根据你自己目录设置
```shell
export GOROOT=/usr/local/go
export GOPKG=$GOROOT/pkg/tool/darwin_amd64
export GOARCH=amd64
export GOOS=darwin
export GOPATH=/Volumes/work/go/blog-go
export GOBIN=$GOPATH/bin
export PATH=.:$PATH:$GOBIN:$GOPKG:$GOPATH/bin 
```

##新环境

先安装 beego 和其他依赖
```go
go get github.com/astaxie/beego
go get github.com/beego/bee
go get github.com/astaxie/beego/cache
go get github.com/go-xorm/xorm
go get -u github.com/xormplus/xorm
go get github.com/go-xorm/cmd/xorm
go get -u qiniupkg.com/api.v7
go get github.com/russross/blackfriday
```
>如果`golang.org/x/net/context`无法下载使用下面下载地址
http://www.golangtc.com/static/download/packages/golang.org.x.net.tar.gz
或者使用下面网址中按步骤下载
http://www.golangtc.com/download/package

然后进入项目目录
```go
cd src/fox
bee run    #beego 要先安装
```
##环境已安装过了
直接进入项目目录 编译
```go
cd src/fox
bee run    #beego 要先安装
```
#后台用户
用户名：admin

密码：111111

登陆地址 : /admin/login

数据库文件在:src/fox/db/blog-go.sql.zip中

#src/fox/db/说明
www.foxwho.com.start.sh 为项目启动文件
start.sh 为自动部署编译文件
blog_go.sql.zip 数据库文件
www.foxwho.com.conf 为nginx配置文件
#用到组件
go 框架：Beego

orm框架：xorm和xormplus

后台框架：Bootstrap

编辑器 ：Markdown

上传   ：webuploader


#后台说明
## RESTFUL
detail :查看页面

get: 编辑页面  模版文件名 get.html

post: 添加数据

put:更新数据

delete:删除数据

案例

```html
test.com/admin/type/detail/15   [get] 查看 id为15 的数据 页面【查看】
test.com/admin/type/15          [get] 编辑 id为15 的数据 页面【修改】
test.com/admin/type/15          [put] 编辑保存 id为15 的数据 【修改】
test.com/admin/type/15          [delete] 删除 id为15 的数据 【删除】

test.com/admin/type/add         [get] 添加 页面  【添加】
test.com/admin/type             [post] 保存 数据 【添加】
test.com/admin/type             [get] 列表 页面  【查询】
```
#ORM 使用 xorm 和xormplus
安装（注意是2个）

网址：http://www.xorm.io/

网址：https://github.com/xormplus/xorm

库安装
```shell
go get github.com/go-xorm/xorm
go get -u github.com/xormplus/xorm
```
工具安装
```shell
go get github.com/go-xorm/cmd/xorm
```
#生成模型
>templates/goxorm 可以修改此模版
先把 src/github.com/go-xorm/cmd/xorm/templates 目录，复制到你的项目目录里，例如 我的项目目录为 src/fox 那么就复制到该目录下，然后执行此命令
该目录下多余的文件夹可以删除（C++,objc,go）

```shell
xorm reverse mysql root:root@/blog_go?charset=utf8 templates/goxorm
```
>如果要增加更多自定义tag 可以修改源码 src/github.com/go-xorm/cmd/xorm/go.go 第267行