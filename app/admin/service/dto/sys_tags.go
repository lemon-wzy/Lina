package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysTagsGetPageReq struct {
	dto.Pagination `search:"-"`
	Id             int `form:"id" search:"type:exact;column:id;table:sys_tags" comment:""`
	SysTagsOrder
}

type SysTagsOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:sys_tags"`
	Title     string `form:"titleOrder"  search:"type:order;column:title;table:sys_tags"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:sys_tags"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sys_tags"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:sys_tags"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:sys_tags"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:sys_tags"`
}

func (m *SysTagsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysTagsInsertReq struct {
	Id    int    `json:"-" comment:"主键"` // 主键
	Title string `json:"title" comment:"标签名称"`
	common.ControlBy
}

func (s *SysTagsInsertReq) Generate(model *models.SysTags) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *SysTagsInsertReq) GetId() interface{} {
	return s.Id
}

type SysTagsUpdateReq struct {
	Id    int    `uri:"id" comment:"主键"` // 主键
	Title string `json:"title" comment:"标签名称"`
	common.ControlBy
}

func (s *SysTagsUpdateReq) Generate(model *models.SysTags) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *SysTagsUpdateReq) GetId() interface{} {
	return s.Id
}

// SysTagsGetReq 功能获取请求参数
type SysTagsGetReq struct {
	Id int `uri:"id"`
}

func (s *SysTagsGetReq) GetId() interface{} {
	return s.Id
}

// SysTagsDeleteReq 功能删除请求参数
type SysTagsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysTagsDeleteReq) GetId() interface{} {
	return s.Ids
}
