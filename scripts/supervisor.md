

debian 案例
## 安装 supervisor

```bash
apt-get update
apt-get install supervisor
```
## 配置 supervisor
配置文件为 `/etc/supervisor/supervisord.conf`

应用配置目录 `/etc/supervisor/conf.d`

## 启动 supervisor
```bash
systemctl enable supervisor
```

## 启动 应用
```bash
supervisorctl start all
```
## 把 blog 应用加入
复制 `supervisor.blog.conf` 到 `/etc/supervisor/conf.d`
启动 blog 应用
```bash
supervisorctl start foxwhoBlog
```

## 其他命令
```bash
#关闭所有任务
supervisorctl shutdown 
#关闭指定任务
supervisorctl stop|start program_name 
#查看所有任务状态
supervisorctl status 
#加载新的配置
supervisorctl update 
#重启所有任务
supervisorctl reload
```