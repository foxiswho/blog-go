#DB Filter
对db 操作封装了一下,暂时添加说明，部分内容还没有验证

```SHELL
var where map[string]interface{}
//字段空值 筛选
where["title"]=""     // SQL语句    title=""
where["title=?"]=""     // SQL语句    title=''

//
where["id"]=3     // SQL语句    id=3
where["id=?"]=3         //id=3
//数组
arr :=[...] int {1,2,3,4,5}
where["id"]= arr    // SQL语句    id in (1,2,3,4,5)

q:="查询字符串"
where["title LIKE ? "]= "%" + q + "%"        // title like "%查询字符串%"

where["num>?"]= 20        // num>20
where["num<?"]= 30        // num<30
where["num!=?"]= 30        // num!=30

```
