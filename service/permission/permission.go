package permission

import (
	"github.com/crazyhl/gopermission/v1/config"
	"github.com/crazyhl/gopermission/v1/models"
)

func FindById(id int) *models.Permission {
	db := config.GetConfig().GormDb

	p := models.Permission{}
	db.First(&p, id)

	return &p
}
