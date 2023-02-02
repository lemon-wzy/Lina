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

type GiftInfo struct {
	service.Service
}

// GetPage 获取GiftInfo列表
func (e *GiftInfo) GetPage(c *dto.GiftInfoGetPageReq, p *actions.DataPermission, list *[]models.GiftInfo, count *int64) error {
	var err error
	var data models.GiftInfo

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("GiftInfoService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取GiftInfo对象
func (e *GiftInfo) Get(d *dto.GiftInfoGetReq, p *actions.DataPermission, model *models.GiftInfo) error {
	var data models.GiftInfo

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetGiftInfo error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建GiftInfo对象
func (e *GiftInfo) Insert(c *dto.GiftInfoInsertReq) error {
    var err error
    var data models.GiftInfo
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("GiftInfoService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改GiftInfo对象
func (e *GiftInfo) Update(c *dto.GiftInfoUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.GiftInfo{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("GiftInfoService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除GiftInfo
func (e *GiftInfo) Remove(d *dto.GiftInfoDeleteReq, p *actions.DataPermission) error {
	var data models.GiftInfo

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveGiftInfo error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
