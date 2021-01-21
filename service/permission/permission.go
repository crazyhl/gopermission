package permission

import (
	"github.com/crazyhl/gopermission/config"
	"github.com/crazyhl/gopermission/models"
)

func FindById(id int) *models.Permission {
	db := config.GetConfig().GormDb

	p := models.Permission{}
	db.First(&p, id)

	return &p
}

func FindByIdString(id string) *models.Permission {
	db := config.GetConfig().GormDb

	p := models.Permission{}
	db.First(&p, id)

	return &p
}

func FindByIds(ids []int) []models.Permission {
	db := config.GetConfig().GormDb
	var permissions []models.Permission
	db.Where("id IN ?", ids).Find(&permissions)
	return permissions
}

func FindByIdsString(ids []string) []models.Permission {
	db := config.GetConfig().GormDb
	var permissions []models.Permission
	db.Where("id IN ?", ids).Find(&permissions)
	return permissions
}
