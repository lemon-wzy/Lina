package models

import (
	"go-admin/common/models"
)

type SysTags struct {
	models.Model

	Title string `json:"title" gorm:"type:varchar(255);comment:标签名称"`
	models.ModelTime
	models.ControlBy
}

func (SysTags) TableName() string {
	return "sys_tags"
}

func (e *SysTags) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysTags) GetId() interface{} {
	return e.Id
}
