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