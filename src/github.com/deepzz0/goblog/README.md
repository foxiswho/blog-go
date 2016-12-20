## Goblog 

Goblog 是基于beego框架开发的博客系统。主张简约、简单、轻量。数据存储mongodb，cdn推荐七牛，多功能后台，可docker部署。

#### 关于版本
``` sh
bee   :1.4.1
beego :1.6.1
Go    :go version go1.6.2 darwin/amd64
```

#### 获取项目 
```
go get -u github.com/deepzz0/goblog
```
#### 部署步骤
###### Mac OS X
######安装mongodb数据库
```
brew install mongodb
```
若没有<code>brew</code>，可自行百度谷歌，安装。
###### 配置mongod
mongodb默认读取环境变量<code>MGO</code>，追加<code>export MGO="127.0.0.1"</code>到<code>~/.bash_profile</code>末尾。如果没有该文件可手动创建。
``` sh
echo 'export MGO="127.0.0.1"' >> .bash_profile
```

##### 修改配置
所有配置文件均在项目目录下的<code>conf</code>下。  
1. <code>app.conf</code>，这里是beego框架的配置文件，采用ini配置方式，如果你了解过beego，应该能看懂。beego框架地址:[http://beego.me](http://beego.me) .
runmode选择你要运行的模式，对应下面的dev、prod、test。线上模式一般为prod，因为该模式不会输出beego的调试log。enablehttps开启https，依次是https监听端口，https相关证书(证书位置替换为自己的)，http的监听端口(当然你也可以关掉http，enablehttp=false)。mydomain是生成feed需要的域名地址
``` ini
appname = goblog
runmode = dev 

[dev]
httpport = 8080
# 用于生成feed
mydomain = 127.0.0.1:8080
[prod]
enablehttp = true
httpport = 80
enablehttps = false
httpsport = 443
httpscertfile = /certs/fullchain.pem
httpskeyfile = /certs/privkey.pem
# 用于生成feed等其它需要用到域名的地方
mydomain = deepzz.com
[test]
httpport = 8888
```
2. <code>17monipdb.dat</code>，该文件是后台解析ip地址的数据库，你可以查看[www.ipip.net](www.ipip.net)
3. <code>qiniu.conf</code>，该文件存放了有关qiniu相关参数，如果没有请创建，暂时没有用到。
4. <code>backleft</code>，该文件是后台相关菜单项，你可以进行删减，并实现对应的功能。
5. <code>backup</code>，该文件夹存放的是账号相关，前台菜单项等相关配置。UserName是后台登陆用户名，PassWord登陆密码(这里填明文)，数据库存储是加密过的，可以在后台进行修改。Salt这是加密的随机盐，由于是手动注册，请随便填写。HeadIcon头像图片，这个地址填图片链接地址.
``` json
{
    "UserName": "deepzz",
    "PassWord": "deepzz",
    "Email": "chenqijing2@163.com",
    "Salt": "__(f",
    "Sex": "男",
    "PNumber": 15100000000,
    "Address": "山西胡同",
    "Education": "本科",
    "RealName": "陈",
    "CreateTime": "2016-01-31T17:38:40.431+08:00",
    "LoginTime": "2016-01-31T17:38:40.431+08:00",
    "LoginIp": "127.0.0.1",
    "LogoutTime": "2016-01-31T17:38:40.431+08:00",
    "BlogName": "Deepzz",
    "Introduce": "程序路上一直坚持的小小码农。",
    "HeadIcon": "https://o7msui8ho.qnssl.com/img/headpic/headpic.jpg",
    "Tags": {},)"
}
```
6. <code>models/model.go</code>修改默认用户，将deepzz替换成你的用户名。
``` go
	UMgr.loadUsers()
	Blogger = UMgr.Get("deepzz")
```

##### 多说评论框架  
``` js
<!-- 多说评论框 start -->
    <div class="ds-thread" data-order="desc" data-limit="20" data-form-position="top" data-thread-key="{{.ID}}" data-title="{{.Title}}" data-url="{{$.Domain}}/{{.URL}}"></div>
    <!-- 多说评论框 end -->
    <!-- 多说公共JS代码 start (一个网页只需插入一次) -->
    <script type="text/javascript">
      var duoshuoQuery = {short_name:"deepzz"};
      (function() {
        var ds = document.createElement('script');
        ds.type = 'text/javascript';ds.async = true;
        ds.src = (document.location.protocol == 'https:' ? 'https:' : 'http:') + '//static.duoshuo.com/embed.js';
        ds.charset = 'UTF-8';
        (document.getElementsByTagName('head')[0] || document.getElementsByTagName('body')[0]).appendChild(ds);
      })();
    </script>
    <!-- 多说公共JS代码 end -->
```
博客集成多说评论框架，你需要到多说官网获取的你网站的shot_name，将上面的deepzz替换掉。该代码段嵌到多个页面，你需要一一替换。由于多说多渠道登陆头像是基于http协议。固如果博客部署为https，将自动进行头像代理为https，实现全站https。

##### 统计相关  
``` js
<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','https://www.google-analytics.com/analytics.js','ga');

  ga('create', 'UA-77251712-1', 'auto');
  ga('send', 'pageview');

</script>
```

这里博主用了Google的数据统计，你可以将上面的代码替换成百度的统计代码等其它统计方式。在<code>views/homelayout.html</code>

##### 插件相关  
暂时做了一个小工具，以后其它工具都会写高该目录下<code>controllers/plugin</code>。

1. UserAgent Parser，基于GO的用户代理解析器。可以到<code>domain:port/plugin/useragent.html</code>访问。

##### 其它说明  
关于持久化，现阶段有sitemap.xml和feed.xml需要持久化，我把它们存放于<code>/data/goblog</code>目录下，自我测试时可能会报错，你可以随意指定该路径<code>models/model.go</code>中。
```
FeedFile     = "/data/goblog/feed.xml"
SiteFile     = "/data/goblog/sitemap.xml"
```

<code>static</code>目录下:

1. <code>feedTemplate.xml</code>是生成feed.xml的模版，你可以通过访问<code>domain:port/feed</code>查看，每小时自动更新。
2. <code>robots.txt</code>，网络爬虫排除协议。
3. <code>sitemap.xml</code>，网站地图，用于搜索引擎快速收录，现只能通过后台手动配置，你也可以通过自己的方式处理。访问两种方式<code>domain:port/sitemap</code>和<code>domain:port/sitemap.xml</code>。
4. 邮件报错，<code>models/model</code>下<code>init</code>的<code>log.SetEmail(Blogger.Email)</code>，如果不开启请注释掉。

所有都配置完成，在根目录下运行<code>bee run</code>

#### 基本页面 
可以到我的博客[http://blog.deepzz.com](http://blog.deepzz.com)查看，https访问[https://blog.deepzz.com](https://blog.deepzz.com)，国外服务器网速稍慢。  
前端页面
![show](https://o7msui8ho.qnssl.com/img/home.png)
后台登陆
![login](https://o7msui8ho.qnssl.com/img/login.png)
首页统计
![analysis](https://o7msui8ho.qnssl.com/img/analysis.png)
博文修改
![modify](https://o7msui8ho.qnssl.com/img/modify.png)
