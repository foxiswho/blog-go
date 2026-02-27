package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"path"
	"reflect"
	"strings"

	"github.com/duke-git/lancet/v2/strutil"
	"github.com/foxiswho/blog-go/app/manage/domainBasic/model/modBasicAttachment"
	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBasic"
	"github.com/foxiswho/blog-go/middleware/components/attachmentPg/modAttachment"
	"github.com/foxiswho/blog-go/middleware/components/attachmentPg/types"
	"github.com/foxiswho/blog-go/pkg/configPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/noPg"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/dbPg/pagePg"
	"github.com/pangu-2/go-tools/tools/numberPg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
	"gorm.io/gorm"
)

func init() {
	gs.Provide(new(BasicAttachmentService)).Init(func(s *BasicAttachmentService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// BasicAttachmentService 附件上传
// @Description:
type BasicAttachmentService struct {
	sv          *repositoryBasic.BasicAttachmentRepository `autowire:"?"`
	FileService types.FileProvider                         `autowire:"?"`
	log         *log2.Logger                               `autowire:"?"`
	server      configPg.Server                            `value:"${server}"`
}

// Upload 上传
func (c *BasicAttachmentService) Upload(ctx *gin.Context) (rt rg.Rs[modBasicAttachment.OkVo]) {
	file, err := ctx.FormFile("file")
	if err != nil {
		return rt.ErrorMessage("上传文件不能为空")
	}

	f, err := file.Open()
	if err != nil {
		return rt.ErrorMessage("上传文件错误")
	}
	defer func() {
		// 打开的资源。一定要记住主动关闭
		_ = f.Close()
	}()

	c.log.Infof("file.Filename=%+v, file.Size = %+v", file.Filename, file.Size)
	atta, err := c.FileService.PutObject(f, modAttachment.PutFile(file.Filename, file.Size), nil)
	c.log.Infof("atta=%#v", atta)
	if err == nil {
		var vo modBasicAttachment.OkVo
		copier.Copy(&vo, &atta)
		vo.Domain = c.server.Domain
		vo.Url = vo.Domain + vo.Url
		return rg.OkData(vo)
	} else {
		return rt.ErrorMessage(err.Error())
	}
}

// UploadMore 多文件上传
func (c *BasicAttachmentService) UploadMore(ctx *gin.Context) (rt rg.Rs[map[int]modBasicAttachment.OkVo]) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return rt.ErrorMessage("上传文件不能为空")
	}
	headers := form.File["file"]
	if nil == headers || len(headers) <= 0 {
		return rt.ErrorMessage("上传文件不能为空")
	}
	data := make(map[int]modBasicAttachment.OkVo, 0)
	for i, file := range headers {
		f, err := file.Open()
		defer func() {
			// 打开的资源。一定要记住主动关闭
			_ = f.Close()
		}()
		if err != nil {
			data[i+1] = modBasicAttachment.OkVo{Error: file.Filename + " 上传文件错误:" + err.Error()}
			continue
		}
		c.log.Infof("file.Filename=%+v, file.Size = %+v", file.Filename, file.Size)
		atta, err := c.FileService.PutObject(f, modAttachment.PutFile(file.Filename, file.Size), nil)
		if err == nil {
			var vo modBasicAttachment.OkVo
			copier.Copy(&vo, &atta)
			vo.Domain = c.server.Domain
			vo.Url = vo.Domain + vo.Url
			data[i+1] = vo
		} else {
			data[i+1] = modBasicAttachment.OkVo{Error: file.Filename + " 上传文件错误: " + err.Error()}
		}
	}
	return rg.OkData(data)
}

// UploadLink 多url文件上传
func (c *BasicAttachmentService) UploadLink(ctx *gin.Context, ct modBasicAttachment.WebUrlCt) (rt rg.Rs[map[int]modBasicAttachment.OkVo]) {
	c.log.Infof("ct=%+v", ct)
	if nil == ct.Url || len(ct.Url) == 0 {
		return rt.ErrorMessage("上传文件不能为空")
	}
	data := make(map[int]modBasicAttachment.OkVo, 0)
	for i, url := range ct.Url {
		if "" == url {
			data[i+1] = modBasicAttachment.OkVo{Error: "地址错误：" + url}
			continue
		}
		res, err := http.Get(url)
		if err != nil {
			data[i+1] = modBasicAttachment.OkVo{Error: "远程下载错误：" + err.Error()}
			continue
		}
		// defer后的为延时操作，通常用来释放相关变量
		defer res.Body.Close()

		content, err := io.ReadAll(res.Body)
		if err != nil {
			data[i+1] = modBasicAttachment.OkVo{Error: "远程下载错误：" + err.Error()}
			continue
		}
		reader := bytes.NewReader(content)
		filename := path.Base(url)
		size := int64(len(content))
		c.log.Infof("file.Filename=%+v, file.Size = %+v", filename, size)
		atta, err := c.FileService.PutObject(reader, modAttachment.PutFile(filename, size), nil)
		if err == nil {
			var vo modBasicAttachment.OkVo
			copier.Copy(&vo, &atta)
			vo.Domain = c.server.Domain
			vo.Url = vo.Domain + vo.Url
			data[i+1] = vo
		} else {
			data[i+1] = modBasicAttachment.OkVo{Error: filename + " 上传文件错误: " + err.Error()}
		}
	}
	return rg.OkData(data)
}

// ListByOwner 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicAttachmentService) ListByOwner(ctx *gin.Context, ct modBasicAttachment.ListFileOwnerCt) (rt rg.Rs[modBasicAttachment.ListFIleOwnerVo]) {
	data := modBasicAttachment.ListFIleOwnerVo{}
	data.GroupData = make(map[string][]modBasicAttachment.Vo)
	data.Data = make([]modBasicAttachment.Vo, 0)
	//
	fileOwner := make([]string, 0)
	mapOwner := make(map[string]string)
	key := ""
	if nil != ct.GroupData {
		for _, item := range ct.GroupData {
			key = strings.TrimSpace(item.FileOwner)
			if strPg.IsNotBlank(key) && strPg.IsNotBlank(item.Group) {
				fileOwner = append(fileOwner, key)
				mapOwner[key] = strings.TrimSpace(item.Group)
			}
		}
	} else if nil != ct.FileOwner {
		for _, item := range ct.FileOwner {
			key = strings.TrimSpace(item)
			if strPg.IsNotBlank(key) {
				fileOwner = append(fileOwner, key)
			}
		}
	}
	if len(fileOwner) > 0 {
		r := c.sv
		var query entityBasic.BasicAttachmentEntity
		//
		query.State = enumStatePg.ENABLE.Index()
		//
		var con repositoryPg.Condition = func(db *gorm.DB) *gorm.DB {
			db = db.Order("create_at desc")
			db.Where("file_owner in ?", fileOwner)
			return db
		}
		infos := r.FindAll(query, con)
		if nil != infos {
			//字段赋值
			for _, item := range infos {
				var vo modBasicAttachment.Vo
				copier.Copy(&vo, &item)
				vo.Url = item.Domain + item.File
				//分组
				if len(mapOwner) > 0 {
					if group, ok := mapOwner[item.FileOwner]; ok {
						if _, b := data.GroupData[group]; !b {
							data.GroupData[group] = make([]modBasicAttachment.Vo, 0)
						}
						data.GroupData[group] = append(data.GroupData[group], vo)
					}
				} else {
					data.Data = append(data.Data, vo)
				}
			}
		}
	}

	rt.Data = data
	return rt.Ok()
}

// DelByOwner 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicAttachmentService) DelByOwner(ctx *gin.Context, ct modBasicAttachment.DelFileOwnerCt) (rt rg.Rs[string]) {
	if strPg.IsBlank(ct.FileOwner) {
		return rt.ErrorMessage("文件拥有者参数错误")
	}
	if nil == ct.Nos || len(ct.Nos) < 1 {
		return rt.ErrorMessage("请选择要删除的文件")
	}
	nos := make([]string, 0)
	for _, item := range ct.Nos {
		item = strings.TrimSpace(item)
		if strPg.IsNotBlank(item) {
			nos = append(nos, item)
		}
	}
	if len(nos) < 1 {
		return rt.ErrorMessage("请选择要删除的文件")
	}
	if c.sv.Config().Data.Delete {
		c.sv.DeleteByIdAndFileOwner(nos, strings.TrimSpace(ct.FileOwner))
	} else {
		c.sv.UpdateByIdAndFileOwnerSetState13(nos, strings.TrimSpace(ct.FileOwner))
	}
	rt.Data = "删除成功"
	return rt.Ok()
}

// Query 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicAttachmentService) Query(ctx *gin.Context, ct modBasicAttachment.QueryCt) (rt rg.Rs[pagePg.PaginatorPg[modBasicAttachment.Vo]]) {
	var query entityBasic.BasicAttachmentEntity
	copier.Copy(&query, &ct)
	r := c.sv
	slice := make([]modBasicAttachment.Vo, 0)
	rt.Data.Data = slice
	page, err := r.FindAllPageQuery(query, func(p *pagePg.PageCondition[*entityBasic.BasicAttachmentEntity]) {
		p.PageOption = func(c *pagePg.PaginatorPg[*entityBasic.BasicAttachmentEntity]) {
			c.PageNum = ct.PageNum
			c.PageSize = ct.PageSize
			if c.PageSize <= 0 {
				c.PageSize = 20
			}
			if c.PageNum <= 0 {
				c.PageNum = 1
			}
		}
		p.Condition = r.DbModel().Order("create_at DESC")
		//自定义查询
		if "" != ct.Wd {
			p.Condition.Where("name like ?", "%"+ct.Wd+"%").Where("source_name like ?", "%"+ct.Wd+"%")
		}
	})
	if nil != err {
		return rt.Ok()
	}
	if page.Total > 0 && page.Data != nil && len(page.Data) > 0 {
		pg := pagePg.NewPaginatorPg(func(c *pagePg.PaginatorPg[modBasicAttachment.Vo]) {
			c.TotalPage = page.TotalPage
			c.Total = page.Total
			c.PageSize = page.PageSize
			c.PageNum = page.PageNum
		})
		//字段赋值
		for _, item := range page.Data {
			var vo modBasicAttachment.Vo
			copier.Copy(&vo, &item)
			vo.Url = item.Domain + item.File
			slice = append(slice, vo)
		}
		pg.Data = slice
		pg.Pageable = page.Pageable
		rt.Data = pg
		return rt.Ok()
	}
	return rt.Ok()
}

// MakeFileOwner 设置文件拥有者/file token/ 文件归属
func (c *BasicAttachmentService) MakeFileOwner(ctx *gin.Context, ct modBasicAttachment.MakeFileOwnerCt) (rt rg.Rs[modBasicAttachment.MakeFileOwner]) {
	vo := modBasicAttachment.MakeFileOwner{}
	prefix := ""
	if len(ct.Mark) > 0 {
		if ":auto" == ct.Mark {
			prefix = "default-"
		} else {
			str := strutil.Substring(ct.Mark, 0, 1)
			if str != ":" {
				prefix = strings.TrimSpace(ct.Mark) + "-"
			}
		}
	} else {
		prefix = "default-"
	}
	vo.FileOwner = noPg.MakeNo(prefix)
	rt.Data = vo
	return rt.Ok()
}

// MakeFileOwnerAll 设置文件拥有者/file token/ 文件归属
func (c *BasicAttachmentService) MakeFileOwnerAll(ctx *gin.Context, ct modBasicAttachment.MakeFileOwnerAllCt) (rt rg.Rs[[]modBasicAttachment.MakeFileOwner]) {
	data := make([]modBasicAttachment.MakeFileOwner, 0)
	if ct.Rule != nil && len(ct.Rule) > 0 {
		for _, item := range ct.Rule {
			vo := modBasicAttachment.MakeFileOwner{}
			prefix := ""
			if len(item.Mark) > 0 {
				if ":auto" == item.Mark {
					prefix = "default-"
				} else {
					str := strutil.Substring(item.Mark, 0, 1)
					if str != ":" {
						prefix = strings.TrimSpace(item.Mark) + "-"
					}
				}
			} else {
				prefix = "default-"
			}
			vo.FileOwner = noPg.MakeNo(prefix)
			vo.Mark = item.Mark
			//
			data = append(data, vo)
		}
	} else if ct.Num > 0 || ct.Num == 0 {
		num := ct.Num
		if num <= 0 {
			num = 1
		}
		prefix := ""
		for i := int32(0); i < num; i++ {
			vo := modBasicAttachment.MakeFileOwner{}
			prefix = ""
			if len(ct.Mark) > 0 {
				if ":auto" == ct.Mark {
					prefix = "default-"
				} else {
					str := strutil.Substring(ct.Mark, 0, 1)
					if str != ":" {
						prefix = strings.TrimSpace(ct.Mark) + "-"
					}
				}
			} else {
				prefix = "default-"
			}
			vo.FileOwner = noPg.MakeNo(prefix)
			vo.Mark = ct.Mark
			//
			data = append(data, vo)
		}
	}
	rt.Data = data
	return rt.Ok()
}

// UpdateByFileOwner 批量更新文件拥有者
func (c *BasicAttachmentService) UpdateByFileOwner(ctx *gin.Context, ct modBasicAttachment.UpdateByFileOwner) (rt rg.Rs[string]) {
	if ct.Data == nil || len(ct.Data) < 1 {
		return rt.ErrorMessage("请选择要更新的文件")
	}
	ids := make([]string, 0)
	mapOwner := make(map[string][]string)
	for key, item := range ct.Data {
		if nil != item && len(item) > 0 {
			mapOwner[key] = make([]string, 0)
			for _, str := range item {
				str = strings.TrimSpace(str)
				if strPg.IsNotBlank(str) {
					mapOwner[key] = append(mapOwner[key], str)
					ids = append(ids, str)
				}
			}
		}
	}
	//查询这个id，如果文件已经拥有者，那么复制一个新纪录
	if len(mapOwner) > 0 {
		data := make([]*entityBasic.BasicAttachmentEntity, 0)
		mapAtt := make(map[string]*entityBasic.BasicAttachmentEntity)
		if len(ids) > 0 {
			var query entityBasic.BasicAttachmentEntity
			query.State = enumStatePg.ENABLE.Index()
			//
			infos := c.sv.FindAll(query, repositoryPg.ConditionOption(func(db *gorm.DB) *gorm.DB {
				db = db.Order("create_at desc")
				db.Where("id in ?", ids)
				return db
			}))
			if infos != nil && len(infos) > 0 {
				for _, item := range infos {
					//存在文件拥有者
					if strPg.IsNotBlank(item.FileOwner) {
						mapAtt[numberPg.Int64ToString(item.ID)] = item
					}
				}
			}
		}
		mapOwnerNew := make(map[string][]string)
		filter := make([]string, 0)
		for key, item := range mapOwner {
			if nil != item && len(item) > 0 {
				filter = make([]string, 0)
				for _, str := range item {
					if obj, ok := mapAtt[str]; ok {
						//存在文件拥有者
						obj.ID = 0
						obj.FileOwner = key
						obj.TypeData = "copy"
						data = append(data, obj)
					} else {
						//不存在文件拥有者
						filter = append(filter, str)
					}
				}
				if len(filter) > 0 {
					fmt.Printf("不存在文件拥有者:%+v\n", filter)
					mapOwnerNew[key] = filter
				}
			}
		}
		if len(data) > 0 {
			c.sv.DbModel().Create(data)
		}
		if len(mapOwnerNew) > 0 {
			for key, item := range mapOwnerNew {
				fmt.Printf("不存在文件拥有者:%+v\n", item)
				c.sv.UpdateByIdSetFileOwner(item, key)
			}
		}
		clear(filter)
		clear(data)
		clear(mapAtt)
		clear(mapOwnerNew)
	}
	return rt.Ok()
}
