package models

import (
	"go-admin/common/models"
)

type GiftInfo struct {
	models.Model

	Title         string `json:"title" gorm:"type:varchar(500);comment:名称"`
	Type          string `json:"type" gorm:"type:varchar(255);comment:型号"`
	Parametric    string `json:"parametric" gorm:"type:varchar(1000);comment:参数"`
	Images        string `json:"images" gorm:"type:varchar(3000);comment:图片"`
	Price         string `json:"price" gorm:"type:int;comment:价格"`
	NetPrice      string `json:"netPrice" gorm:"type:varchar(255);comment:网上参考价格"`
	BoxNum        string `json:"boxNum" gorm:"type:varchar(255);comment:装箱数"`
	IfBox         string `json:"ifBox" gorm:"type:int(2);comment:是否彩盒"`
	IfTax         string `json:"ifTax" gorm:"type:int(2);comment:是否税"`
	ShippingPlace string `json:"shippingPlace" gorm:"type:double;comment:发货地"`
	models.ModelTime
	models.ControlBy
}

func (GiftInfo) TableName() string {
	return "gift_info"
}

func (e *GiftInfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *GiftInfo) GetId() interface{} {
	return e.Id
}
