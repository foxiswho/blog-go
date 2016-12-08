package service

import (
	"fox/models"
)

type adminService struct {

}

func (this *adminService) GetById(id int64) (user *models.Admin, err error) {
	user = &models.Admin{Id:id}
	o.Read
	if err := o.Read(); err != nil {

	}
}