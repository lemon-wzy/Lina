package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type GiftInfoGetPageReq struct {
	dto.Pagination `search:"-"`
	Title          string `form:"title" search:"type:contains;column:title;table:gift_info" comment:"名称"`
	StartPrice     string `form:"startPrice" search:"type:gte;column:price;table:gift_info" comment:"价格"`
	EndPrice       string `form:"endPrice" search:"type:lte;column:price;table:gift_info" comment:"价格"`
	GiftInfoOrder
}

type GiftInfoOrder struct {
	Id            string `form:"idOrder"  search:"type:order;column:id;table:gift_info"`
	Title         string `form:"titleOrder"  search:"type:order;column:title;table:gift_info"`
	Type          string `form:"typeOrder"  search:"type:order;column:type;table:gift_info"`
	Parametric    string `form:"parametricOrder"  search:"type:order;column:parametric;table:gift_info"`
	Images        string `form:"imagesOrder"  search:"type:order;column:images;table:gift_info"`
	Price         string `form:"priceOrder"  search:"type:order;column:price;table:gift_info"`
	NetPrice      string `form:"netPriceOrder"  search:"type:order;column:net_price;table:gift_info"`
	BoxNum        string `form:"boxNumOrder"  search:"type:order;column:box_num;table:gift_info"`
	IfBox         string `form:"ifBoxOrder"  search:"type:order;column:if_box;table:gift_info"`
	IfTax         string `form:"ifTaxOrder"  search:"type:order;column:if_tax;table:gift_info"`
	ShippingPlace string `form:"shippingPlaceOrder"  search:"type:order;column:shipping_place;table:gift_info"`
	TagsId        string `form:"tagsIdOrder"  search:"type:order;column:tags_id;table:gift_info"`
	TagsName      string `form:"tagsNameOrder"  search:"type:order;column:tags_name;table:gift_info"`
}

func (m *GiftInfoGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type GiftInfoInsertReq struct {
	Id            int    `json:"-" comment:"主键"` // 主键
	Title         string `json:"title" comment:"名称"`
	Type          string `json:"type" comment:"型号"`
	Parametric    string `json:"parametric" comment:"参数"`
	Images        string `json:"images" comment:"图片"`
	Price         string `json:"price" comment:"价格"`
	NetPrice      string `json:"netPrice" comment:"网上参考价格"`
	BoxNum        string `json:"boxNum" comment:"装箱数"`
	IfBox         string `json:"ifBox" comment:"是否彩盒"`
	IfTax         string `json:"ifTax" comment:"是否税"`
	ShippingPlace string `json:"shippingPlace" comment:"发货地"`
	TagsId        string `json:"tagsId" comment:"标签id"`
	TagsName      string `json:"tagsName" comment:"标签名称"`
	common.ControlBy
}

func (s *GiftInfoInsertReq) Generate(model *models.GiftInfo) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Type = s.Type
	model.Parametric = s.Parametric
	model.Images = s.Images
	model.Price = s.Price
	model.NetPrice = s.NetPrice
	model.BoxNum = s.BoxNum
	model.IfBox = s.IfBox
	model.IfTax = s.IfTax
	model.ShippingPlace = s.ShippingPlace
	model.TagsId = s.TagsId
	model.TagsName = s.TagsName
}

func (s *GiftInfoInsertReq) GetId() interface{} {
	return s.Id
}

type GiftInfoUpdateReq struct {
	Id            int    `uri:"id" comment:"主键"` // 主键
	Title         string `json:"title" comment:"名称"`
	Type          string `json:"type" comment:"型号"`
	Parametric    string `json:"parametric" comment:"参数"`
	Images        string `json:"images" comment:"图片"`
	Price         string `json:"price" comment:"价格"`
	NetPrice      string `json:"netPrice" comment:"网上参考价格"`
	BoxNum        string `json:"boxNum" comment:"装箱数"`
	IfBox         string `json:"ifBox" comment:"是否彩盒"`
	IfTax         string `json:"ifTax" comment:"是否税"`
	ShippingPlace string `json:"shippingPlace" comment:"发货地"`
	TagsId        string `json:"tagsId" comment:"标签id"`
	TagsName      string `json:"tagsName" comment:"标签名称"`
	common.ControlBy
}

func (s *GiftInfoUpdateReq) Generate(model *models.GiftInfo) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Type = s.Type
	model.Parametric = s.Parametric
	model.Images = s.Images
	model.Price = s.Price
	model.NetPrice = s.NetPrice
	model.BoxNum = s.BoxNum
	model.IfBox = s.IfBox
	model.IfTax = s.IfTax
	model.ShippingPlace = s.ShippingPlace
	model.TagsId = s.TagsId
	model.TagsName = s.TagsName
}

func (s *GiftInfoUpdateReq) GetId() interface{} {
	return s.Id
}

// GiftInfoGetReq 功能获取请求参数
type GiftInfoGetReq struct {
	Id int `uri:"id"`
}

func (s *GiftInfoGetReq) GetId() interface{} {
	return s.Id
}

// GiftInfoDeleteReq 功能删除请求参数
type GiftInfoDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *GiftInfoDeleteReq) GetId() interface{} {
	return s.Ids
}
