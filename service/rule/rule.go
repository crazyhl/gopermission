package rule

import (
	"github.com/crazyhl/gopermission/config"
	"github.com/crazyhl/gopermission/models"
)

func FindById(id int) *models.Rule {
	db := config.GetConfig().GormDb

	r := models.Rule{}
	db.First(&r, id)

	return &r
}

func FindByIdString(id string) *models.Rule {
	db := config.GetConfig().GormDb

	r := models.Rule{}
	db.First(&r, id)

	return &r
}

func FindByIds(ids []int) []models.Rule {
	db := config.GetConfig().GormDb
	var rules []models.Rule
	db.Where("id IN ?", ids).Find(&rules)
	return rules
}

func FindByIdsString(ids []string) []models.Rule {
	db := config.GetConfig().GormDb
	var rules []models.Rule
	db.Where("id IN ?", ids).Find(&rules)
	return rules
}
