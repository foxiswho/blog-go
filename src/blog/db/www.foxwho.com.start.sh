#!/bin/bash
nohup /www/wwwroot/www.foxwho.com/fox >/www/wwwroot/www.foxwho.com.log 2>&1 &

#没有任何日志
##nohup /www/wwwroot/www.foxwho.com/fox >/dev/null 2>&1 &