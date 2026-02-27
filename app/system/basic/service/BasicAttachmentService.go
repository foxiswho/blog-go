package service

import (
	"bytes"
	"context"

	"github.com/foxiswho/blog-go/app/system/basic/model/modBasicAttachment"
	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBasic"
	"github.com/foxiswho/blog-go/middleware/components/attachmentPg/modAttachment"
	"github.com/foxiswho/blog-go/middleware/components/attachmentPg/types"
	"github.com/foxiswho/blog-go/pkg/configPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"

	"io"
	"net/http"
	"path"
	"reflect"

	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/dbPg/pagePg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
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
