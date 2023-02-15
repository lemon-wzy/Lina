package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type SysTags struct {
	service.Service
}

// GetPage 获取SysTags列表
func (e *SysTags) GetPage(c *dto.SysTagsGetPageReq, p *actions.DataPermission, list *[]models.SysTags, count *int64) error {
	var err error
	var data models.SysTags

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SysTagsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// GetPage 获取SysTags列表
func (e *SysTags) GetList(c *dto.SysTagsGetPageReq, list *[]models.SysTags) error {
	var err error
	var data models.SysTags

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Get 获取SysTags对象
func (e *SysTags) Get(d *dto.SysTagsGetReq, p *actions.DataPermission, model *models.SysTags) error {
	var data models.SysTags

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysTags error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysTags对象
func (e *SysTags) Insert(c *dto.SysTagsInsertReq) error {
	var err error
	var data models.SysTags
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SysTagsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SysTags对象
func (e *SysTags) Update(c *dto.SysTagsUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.SysTags{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("SysTagsService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SysTags
func (e *SysTags) Remove(d *dto.SysTagsDeleteReq, p *actions.DataPermission) error {
	var data models.SysTags

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSysTags error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
