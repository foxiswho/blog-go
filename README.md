
#GO语言博客
基本功能已 实现

#功能改进
 * 已去除beego orm,改用 xorm
 * 改进where查询（我是比较懒的，弄了一个简单的where查询）

#未来
 * 用其他的 orm 重新完善一下
 * 其他模块继续完善
    * 省市区
    * 角色和权限
    * 管理员
    * 菜单
    * 缓存
    * 标签
    * 附件
    * 。。。。。。


#后台用户
用户名：admin

密码：111111

登陆地址 : /admin/login

数据库文件在:src/fox/db/blog-go.sql.zip中

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