package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/deepzz0/go-com/log"
	"github.com/deepzz0/go-com/monitor"
	"github.com/deepzz0/goblog/RS"
	"github.com/deepzz0/goblog/helper"
	"github.com/wangtuanjie/ip17mon"
)

const (
	// database name
	DB = "newblog"
	// user table
	C_USER  = "user"
	C_TOPIC = "topic"
	// id
	C_TOPIC_ID = "topic_id"
	// request
	C_REQUEST = "request"
	// config
	C_CONFIG = "config"
)

const (
	TemplateFile = "./static/feedTemplate.xml"
	RobotsFile   = "./static/robots.txt"
	FeedFile     = "/data/goblog/feed.xml"
	SiteFile     = "/data/goblog/sitemap.xml"
)

var UMgr = NewUM()
var TMgr = NewTM()
var path, _ = os.Getwd()
var Blogger *User
var Icons = make(map[string]*Icon, 500)

func init() {
	if err := ip17mon.Init(path + "/conf/17monipdb.dat"); err != nil {
		log.Fatal(err)
	}
	//
	UMgr.loadUsers()
	Blogger = UMgr.Get("deepzz")
	// not found account
	if Blogger == nil {
		initAccount()
	}
	TMgr.loadTopics()
	// open error mail，email addr : Blogger.Email
	log.SetEmail(Blogger.Email)
	ManageData.LoadData()
	monitor.HookOnExit("flushdata", flushdata)
	monitor.Startup()

	go RequestM.Saver()
	go timer()
}

func initAccount() {
	b, err := ioutil.ReadFile(path + "/conf/init/user.json")
	if err != nil {
		panic(err)
	}
	user := User{}
	err = json.Unmarshal(b, &user)
	if err != nil {
		panic(err)
	}
	user.PassWord = helper.EncryptPasswd(user.UserName, user.PassWord, user.Salt)
	UMgr.Register(&user)
	code := UMgr.Update()
	if code != RS.RS_success {
		panic("init failed。")
	}
	Blogger = UMgr.Get("deepzz")
}

func timer() {
	t := time.NewTicker(time.Minute)
	Today := time.Now()

	tUser := time.NewTicker(time.Hour)
	tTopic := time.NewTicker(time.Minute * 10)
	tIcon := time.NewTicker(time.Hour * 12)
	for {
		select {
		case <-t.C:
			if time.Now().Day() != Today.Day() {
				Today = time.Now()
				ManageData.LoadData()
				ManageData.CleanData(Today)
			}
		case <-tUser.C:
			UMgr.Update()
		case <-tTopic.C:
			TMgr.DoDelete(time.Now())
		case <-tIcon.C:
			cleanIcons()
		}
	}
}

func flushdata() {
	UMgr.Update()
	TMgr.Update()
	ManageConf.UpdateConf()
}
