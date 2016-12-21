package admin


type IndexV2Controller struct {
	BaseController
}

func (this *IndexV2Controller)Get() {
	this.TplName = "admin/index_v2/get.html"
}