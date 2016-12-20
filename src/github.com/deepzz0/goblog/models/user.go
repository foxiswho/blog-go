package models

import (
	"sort"
	"sync"
	"time"

	"github.com/deepzz0/go-com/log"
	db "github.com/deepzz0/go-com/mongo"
	"github.com/deepzz0/goblog/RS"
	"github.com/deepzz0/goblog/helper"
	"gopkg.in/mgo.v2/bson"
)

// 账户
type User struct {
	UserName   string          // 账户ID
	PassWord   string          // 账户密码
	Email      string          // email
	Salt       string          //
	Sex        string          // 1:男， 2：女
	PNumber    int64           // 手机号
	Address    string          // 住址
	Education  string          // 教育
	RealName   string          // 真实姓名
	CreateTime time.Time       // 创建时间
	LoginTime  time.Time       // 登录时间
	LoginIp    string          // 登录ip
	LogoutTime time.Time       // 登出时间
	BlogName   string          // 博客名
	Introduce  string          // 个人简介
	HeadIcon   string          // 头像
	Tags       map[string]*Tag // 标签
	Categories SortCategory    // 分类
	Socials    SortSocial      // social
	Blogrolls  SortBlogroll    // 友情链接
}

type UserMgr struct {
	lock sync.Mutex
	// userid --> *User
	Users map[string]*User
}

func NewUM() *UserMgr { return &UserMgr{Users: make(map[string]*User)} }

func (m *UserMgr) loadUsers() {
	var users []*User
	err := db.FindAll(DB, C_USER, nil, &users)
	if err != nil {
		panic(err)
	}
	log.Debug(len(users))
	for _, u := range users {
		m.Users[u.UserName] = u
	}
}

func (m *UserMgr) Register(user *User) int {
	m.lock.Lock()
	defer m.lock.Unlock()
	err := db.Update(DB, C_USER, bson.M{"username": user.UserName}, *user)
	if err != nil {
		log.Warn(err)
		return RS.RS_register_failed
	}
	m.Users[user.UserName] = user
	return RS.RS_success
}

func (m *UserMgr) FoundPass(name, email string) int {
	if user, found := m.Users[name]; !found {
		return RS.RS_user_inexistence
	} else {
		log.Debug(user.UserName)
	}

	// 发送邮件
	// ....
	return RS.RS_success
}

func (m *UserMgr) Login(name, passwd string) int {
	user := m.Users[name]
	if user == nil {
		return RS.RS_user_inexistence
	}
	if user.PassWord != helper.EncryptPasswd(name, passwd, user.Salt) {
		return RS.RS_password_error
	}
	user.LoginTime = time.Now()
	return RS.RS_success
}

func (m *UserMgr) Logout(name string) int {
	user := m.Users[name]
	if user == nil {
		return RS.RS_user_inexistence
	}
	user.LogoutTime = time.Now()
	db.Update(DB, C_USER, bson.M{"username": name}, *m.Users[name])
	return RS.RS_success
}

func (m *UserMgr) Get(name string) *User {
	return m.Users[name]
}

func (m *UserMgr) Update() int {
	for _, u := range m.Users {
		err := db.Update(DB, C_USER, bson.M{"username": u.UserName}, *u)
		if err != nil {
			return RS.RS_update_failed
		}
	}
	return RS.RS_success
}

//-------------------------------user-------------------------------------
func (u *User) GetCategoryByID(id string) *Category {
	for _, c := range u.Categories {
		if c.ID == id {
			return c
		}
	}
	return nil
}
func (u *User) GetValidCategory() []*Category {
	var temps []*Category
	for _, v := range u.Categories {
		if v.IsCat {
			temps = append(temps, v)
		}
	}
	return temps
}
func (u *User) AddCategory(cat *Category) int {
	if u.GetCategoryByID(cat.ID) == nil {
		u.Categories = append(u.Categories, cat)
		sort.Sort(u.Categories)
		return RS.RS_success
	}
	return RS.RS_failed
}
func (u *User) DelCatgoryByID(id string) int {
	for i, cat := range u.Categories {
		if id == cat.ID {
			u.Categories = append(u.Categories[:i], u.Categories[i+1:]...)
			delete(TMgr.GroupByCategory, id)
			return RS.RS_success
		}
	}
	return RS.RS_failed
}
func (u *User) ReduceCategoryCount(id string) {
	category := u.GetCategoryByID(id)
	if category != nil {
		category.reduceCount()
	}
}
func (u *User) AddCategoryCount(id string) {
	u.GetCategoryByID(id).addCount()
}
func (u *User) DelTagByID(id string) int {
	if u.GetTagByID(id) != nil {
		delete(TMgr.GroupByTag, id)
		delete(u.Tags, id)
		return RS.RS_success
	}
	return RS.RS_failed
}
func (u *User) ReduceTagCount(id string) {
	tag := u.GetTagByID(id)
	if tag != nil {
		tag.reduceCount()
		if tag.Count <= 0 {
			delete(TMgr.GroupByTag, id)
			delete(u.Tags, id)
		}
	}
}
func (u *User) GetTagByID(id string) *Tag {
	return u.Tags[id]
}
func (u *User) AddTagCount(id string) {
	u.GetTagByID(id).addCount()
}
func (u *User) AddTag(tag *Tag) int {
	if t := u.GetTagByID(tag.ID); t == nil {
		tag.addCount()
		u.Tags[tag.ID] = tag
		return RS.RS_success
	}
	return RS.RS_tag_exist
}
func (u *User) GetSocialByID(id string) *Social {
	for _, s := range u.Socials {
		if s.ID == id {
			return s
		}
	}
	return nil
}
func (u *User) AddSocial(social *Social) int {
	if u.GetSocialByID(social.ID) == nil {
		u.Socials = append(u.Socials, social)
		sort.Sort(u.Socials)
		return RS.RS_success
	}
	return RS.RS_failed
}
func (u *User) DelSocialByID(id string) int {
	for i, social := range u.Socials {
		if id == social.ID {
			u.Socials = append(u.Socials[:i], u.Socials[i+1:]...)
			return RS.RS_success
		}
	}
	return RS.RS_failed
}
func (u *User) GetBlogrollByID(id string) *Blogroll {
	for _, br := range u.Blogrolls {
		if br.ID == id {
			return br
		}
	}
	return nil
}
func (u *User) AddBlogroll(br *Blogroll) int {
	if u.GetBlogrollByID(br.ID) == nil {
		u.Blogrolls = append(u.Blogrolls, br)
		sort.Sort(u.Blogrolls)
		return RS.RS_success
	}
	return RS.RS_failed
}
func (u *User) DelBlogrollByID(id string) int {
	for i, br := range u.Blogrolls {
		if id == br.ID {
			u.Blogrolls = append(u.Blogrolls[:i], u.Blogrolls[i+1:]...)
			return RS.RS_success
		}
	}
	return RS.RS_failed
}
func (u *User) ChangePassword(newP string) {
	u.PassWord = helper.EncryptPasswd(u.UserName, newP, u.Salt)
}

// -------------------------------Tags-------------------------------------
const (
	Label_default = iota
	Label_primary
	Label_success
	Label_info
	Label_warning = 10
)

var TagStyle = map[int]string{
	Label_default: "label-default",
	Label_primary: "label-primary",
	Label_success: "label-success",
	Label_info:    "label-info",
	Label_warning: "label-warning",
}

type Tag struct {
	ID    string // 内部ID
	Count int    // 数量
	Extra string // 链接
	Text  string // 显示名称
}

func NewTag() *Tag {
	return &Tag{}
}

func (t *Tag) addCount() {
	t.Count++
}
func (t *Tag) reduceCount() {
	t.Count--
}
func (t *Tag) TagStyle() string {
	rand := helper.GetRand()
	style := TagStyle[rand.Intn(4)]
	return "<span class='label " + style + "'>" + t.ID + "</span>"
}

// -------------------------------category-------------------------------------
type SortCategory []*Category

func (sc SortCategory) Len() int           { return len(sc) }
func (sc SortCategory) Less(i, j int) bool { return sc[i].SortID < sc[j].SortID }
func (sc SortCategory) Swap(i, j int)      { sc[i], sc[j] = sc[j], sc[i] }

type Category struct {
	ID         string    // 内部ID
	Count      int       // 文章数量
	IsCat      bool      // 是否分类
	SortID     int       // 排序ID
	Title      string    // 显示说明
	Extra      string    // 链接地址
	Text       string    // 显示名称
	CreateTime time.Time // 创建时间
}

func NewCategory() *Category {
	return &Category{CreateTime: time.Now()}
}
func (c *Category) addCount() {
	c.Count++
}
func (c *Category) reduceCount() {
	c.Count--
}

// -------------------------------social-------------------------------------
type SortSocial []*Social

func (ss SortSocial) Len() int           { return len(ss) }
func (ss SortSocial) Less(i, j int) bool { return ss[i].SortID < ss[j].SortID }
func (ss SortSocial) Swap(i, j int)      { ss[i], ss[j] = ss[j], ss[i] }

type Social struct {
	ID         string    // 内部ID
	SortID     int       // 排序ID
	Title      string    // 显示说明
	Extra      string    // 链接地址
	Icon       string    // icon
	CreateTime time.Time // 创建时间
}

func NewSocial() *Social {
	return &Social{CreateTime: time.Now()}
}

// -------------------------------blogroll-------------------------------------
type SortBlogroll []*Blogroll

func (sb SortBlogroll) Len() int           { return len(sb) }
func (sb SortBlogroll) Less(i, j int) bool { return sb[i].SortID < sb[j].SortID }
func (sb SortBlogroll) Swap(i, j int)      { sb[i], sb[j] = sb[j], sb[i] }

type Blogroll struct {
	ID         string    // 内部ID
	SortID     int       // 排序ID
	Title      string    // 显示说明
	Extra      string    // 链接
	Text       string    // 显示名称
	CreateTime time.Time // 创建时间
}

func NewBlogroll() *Blogroll {
	return &Blogroll{CreateTime: time.Now()}
}

// -------------------------------icon-------------------------------------
type Icon struct {
	Data []byte
	Time time.Time
}

func cleanIcons() {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()
	for k, v := range Icons {
		if v.Time.Before(time.Now().AddDate(0, 0, -2)) {
			Icons[k] = nil
			delete(Icons, k)
		}
	}
}
