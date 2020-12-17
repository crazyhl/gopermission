package rule

import (
	"github.com/crazyhl/gopermission/v1/config"
	"github.com/crazyhl/gopermission/v1/models"
)

func FindById(id int) *models.Rule {
	db := config.GetConfig().GormDb

	r := models.Rule{}
	db.First(&r, id)

	return &r
}

func FindByIds(ids []int) []models.Rule {
	db := config.GetConfig().GormDb
	rules := []models.Rule{}
	db.Where("id IN ?", ids).Find(&rules)
	return rules
}
