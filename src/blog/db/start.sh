#!/bin/bash

#网址
SITE=www.foxwho.com
#部署根路径
DIR=/www/deploy/$SITE
#备份路径
BACKUP=/wwwData/backup/$SITE
#站点根路径
WWWROOT=/www/wwwroot
#站点路径
SITEPATH=$WWWROOT/$SITE

#gopanth
GO_PATH=/www/go

#解压缩后文件夹
DEPLOYPATH=$DIR/source
#当前时间 
DATETIME=$(date +%Y-%m-%d-%H-%M-%S)

#项目端口号
P_PORT=8091
#数据库用户
P_DB_USER=root
#数据库密码
P_DB_PWD=root
#数据库名
P_DB_NAME=blog_go
#数据库地址
P_DB_HOST=127.0.0.1


echo "备份开始===="
echo "只备份源码,Uploads 目录下不备份"
echo "备份文件存放于${BACKUP}/source/$DATETIME.tar.gz"
#创建备份目录
[ ! -d "$BACKUP" ] && mkdir -p "$BACKUP"
[ ! -d "$BACKUP/source" ] && mkdir -p $BACKUP/source
#创建解压缩目录
[ ! -d "$DEPLOYPATH" ] && mkdir -p $DEPLOYPATH

#备份源码
cd $WWWROOT
#tar -zcpf $BACKUP/source/$DATETIME.tar.gz --exclude=$SITE/Uploads --exclude=$SITE/runtime $SITE

#删除解压缩目录内文件
rm -rf $DEPLOYPATH/*

cd $DIR
#解压缩 整理 git 文件
tar -zxf  package.tgz -C $DEPLOYPATH
cd $DEPLOYPATH
mv blog-go-* tttttt
mv $DEPLOYPATH/tttttt/* $DEPLOYPATH/
rm -rf $DEPLOYPATH/tttttt 
#/////////////////////////////////////////////
#编译
#
#删除本项目源码
rm -rf $GO_PATH/src/fox
#
#复制最新源码到项目里
mv $DEPLOYPATH/src/fox $GO_PATH/src/

######################
#数据库相关替换
#配置文件
DB_FILE=$GO_PATH/src/fox/conf/app.conf
#替换数据库
sed -i "s:db_user.*=.*:db_user=\"${P_DB_USER}\":g" $DB_FILE
sed -i "s:db_pass.*=.*:db_pass=\"${P_DB_PWD}\":g" $DB_FILE
sed -i "s:db_name.*=.*:db_name=\"${P_DB_NAME}\":g" $DB_FILE
sed -i "s:db_host.*=.*:db_host=\"${P_DB_HOST}\":g" $DB_FILE
#替换项目端口
sed -i "s:httpport.*=.*:httpport=${P_PORT}:g" $DB_FILE

######################
#进入项目目录
cd $GO_PATH/src/fox
#使用beego 打包
bee pack

######################
#解压缩打包文件
cd $GO_PATH/src/fox
#fox是项目名称
rm -rf $DEPLOYPATH/*

tar -zxf fox.tar.gz -C $DEPLOYPATH
PACK_PATH=$DEPLOYPATH
######################
#删除不相干文件
rm -rf $PACK_PATH/db
######################
#设置文件权限

#权限赋值
chmod -R 750 $PACK_PATH
chown -R www:www $PACK_PATH
######################
#部署目录是否存在，不存在则创建和设置权限
if [ ! -d "$SITEPATH" ]; then
    mkdir -p $SITEPATH
    chmod -R 777 $SITEPATH
    chown -R www:www $SITEPATH
fi
#######################
#更改项目名称
mv $PACK_PATH/fox $PACK_PATH/fox_2
#复制文件
cp -auR $PACK_PATH/* $SITEPATH/

#上传目录检测，如果不存在则创建
if [ ! -d "$SITEPATH/uploads" ]; then
    mkdir -p $SITEPATH/uploads/image
    mkdir -p $SITEPATH/uploads/attachment
    chmod -R 777 $SITEPATH/uploads
    chown -R www:www $SITEPATH/uploads
fi
#############################
#结束进程fox
ps -ef |grep /fox|awk '{print $2}'|xargs kill -9
############
#删除原项目
rm -rf $SITEPATH/fox
#名称恢复
ls -lh ${SITEPATH}
mv "${SITEPATH}/fox_2" "${SITEPATH}/fox"
############
#启动项目
SH="${SITEPATH}.start.sh"
$SH

echo "SUCCESS"
