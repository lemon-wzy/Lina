package version

import (
	"go-admin/cmd/migrate/migration"
	common "go-admin/common/models"
	"runtime"
	"time"

	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"gorm.io/gorm"
)

type Menu struct {
	MenuId     int        `json:"menuId" gorm:"primaryKey;autoIncrement"`
	MenuName   string     `json:"menuName" gorm:"size:128;"`
	Title      string     `json:"title" gorm:"size:128;"`
	Icon       string     `json:"icon" gorm:"size:128;"`
	Path       string     `json:"path" gorm:"size:128;"`
	Paths      string     `json:"paths" gorm:"size:128;"`
	MenuType   string     `json:"menuType" gorm:"size:1;"`
	Action     string     `json:"action" gorm:"size:16;"`
	Permission string     `json:"permission" gorm:"size:255;"`
	ParentId   int        `json:"parentId" gorm:"size:11;"`
	NoCache    bool       `json:"noCache" gorm:"size:8;"`
	Breadcrumb string     `json:"breadcrumb" gorm:"size:255;"`
	Component  string     `json:"component" gorm:"size:255;"`
	Sort       int        `json:"sort" gorm:"size:4;"`
	Visible    string     `json:"visible" gorm:"size:1;"`
	CreateBy   string     `json:"createBy" gorm:"size:128;"`
	UpdateBy   string     `json:"updateBy" gorm:"size:128;"`
	IsFrame    string     `json:"isFrame" gorm:"size:1;DEFAULT:0;"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
}

func (Menu) TableName() string {
	return "sys_menu"
}

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1676283844839Test)
}

func _1676283844839Test(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {

		timeNow := pkg.GetCurrentTime()
		Mmenu := Menu{}
		Mmenu.MenuName = "sys_tagsManage"
		Mmenu.Title = "标签表"
		Mmenu.Icon = "pass"
		Mmenu.Path = "/sys_tags"
		Mmenu.MenuType = "M"
		Mmenu.Action = "无"
		Mmenu.ParentId = 0
		Mmenu.NoCache = false
		Mmenu.Component = "Layout"
		Mmenu.Sort = 0
		Mmenu.Visible = "0"
		Mmenu.IsFrame = "0"
		Mmenu.CreateBy = "1"
		Mmenu.UpdateBy = "1"
		Mmenu.CreatedAt = timeNow
		Mmenu.UpdatedAt = timeNow
		// Mmenu.MenuId, err = Mmenu.Create(db)
		err := tx.Create(&Mmenu).Error
		if err != nil {
			return err
		}
		Cmenu := Menu{}
		Cmenu.MenuName = "sys_tags"
		Cmenu.Title = "标签表"
		Cmenu.Icon = "pass"
		Cmenu.Path = "sys_tags"
		Cmenu.MenuType = "C"
		Cmenu.Action = "无"
		Cmenu.Permission = "admin:sysTags:list"
		Cmenu.ParentId = Mmenu.MenuId
		Cmenu.NoCache = false
		Cmenu.Component = "/sysTags/index"
		Cmenu.Sort = 0
		Cmenu.Visible = "0"
		Cmenu.IsFrame = "0"
		Cmenu.CreateBy = "1"
		Cmenu.UpdateBy = "1"
		Cmenu.CreatedAt = timeNow
		Cmenu.UpdatedAt = timeNow
		// Cmenu.MenuId, err = Cmenu.Create(db)
		err = tx.Create(&Cmenu).Error
		if err != nil {
			return err
		}

		MList := Menu{}
		MList.MenuName = ""
		MList.Title = "分页获取标签表"
		MList.Icon = ""
		MList.Path = "sys_tags"
		MList.MenuType = "F"
		MList.Action = "无"
		MList.Permission = "admin:sysTags:query"
		MList.ParentId = Cmenu.MenuId
		MList.NoCache = false
		MList.Sort = 0
		MList.Visible = "0"
		MList.IsFrame = "0"
		MList.CreateBy = "1"
		MList.UpdateBy = "1"
		MList.CreatedAt = timeNow
		MList.UpdatedAt = timeNow
		// MList.MenuId, err = MList.Create(db)
		err = tx.Create(&MList).Error
		if err != nil {
			return err
		}

		MCreate := Menu{}
		MCreate.MenuName = ""
		MCreate.Title = "创建标签表"
		MCreate.Icon = ""
		MCreate.Path = "sys_tags"
		MCreate.MenuType = "F"
		MCreate.Action = "无"
		MCreate.Permission = "admin:sysTags:add"
		MCreate.ParentId = Cmenu.MenuId
		MCreate.NoCache = false
		MCreate.Sort = 0
		MCreate.Visible = "0"
		MCreate.IsFrame = "0"
		MCreate.CreateBy = "1"
		MCreate.UpdateBy = "1"
		MCreate.CreatedAt = timeNow
		MCreate.UpdatedAt = timeNow
		// MCreate.MenuId, err = MCreate.Create(db)
		err = tx.Create(&MCreate).Error
		if err != nil {
			return err
		}

		MUpdate := Menu{}
		MUpdate.MenuName = ""
		MUpdate.Title = "修改标签表"
		MUpdate.Icon = ""
		MUpdate.Path = "sys_tags"
		MUpdate.MenuType = "F"
		MUpdate.Action = "无"
		MUpdate.Permission = "admin:sysTags:edit"
		MUpdate.ParentId = Cmenu.MenuId
		MUpdate.NoCache = false
		MUpdate.Sort = 0
		MUpdate.Visible = "0"
		MUpdate.IsFrame = "0"
		MUpdate.CreateBy = "1"
		MUpdate.UpdateBy = "1"
		MUpdate.CreatedAt = timeNow
		MUpdate.UpdatedAt = timeNow
		// MUpdate.MenuId, err = MUpdate.Create(db)
		err = tx.Create(&MUpdate).Error
		if err != nil {
			return err
		}

		MDelete := Menu{}
		MDelete.MenuName = ""
		MDelete.Title = "删除标签表"
		MDelete.Icon = ""
		MDelete.Path = "sys_tags"
		MDelete.MenuType = "F"
		MDelete.Action = "无"
		MDelete.Permission = "admin:sysTags:remove"
		MDelete.ParentId = Cmenu.MenuId
		MDelete.NoCache = false
		MDelete.Sort = 0
		MDelete.Visible = "0"
		MDelete.IsFrame = "0"
		MDelete.CreateBy = "1"
		MDelete.UpdateBy = "1"
		MDelete.CreatedAt = timeNow
		MDelete.UpdatedAt = timeNow
		// MDelete.MenuId, err = MDelete.Create(db)
		err = tx.Create(&MDelete).Error
		if err != nil {
			return err
		}

		var InterfaceId = 63
		Amenu := Menu{}
		Amenu.MenuName = "sys_tags"
		Amenu.Title = "标签表"
		Amenu.Icon = "bug"
		Amenu.Path = "sys_tags"
		Amenu.MenuType = "M"
		Amenu.Action = "无"
		Amenu.ParentId = InterfaceId
		Amenu.NoCache = false
		Amenu.Sort = 0
		Amenu.Visible = "1"
		Amenu.IsFrame = "0"
		Amenu.CreateBy = "1"
		Amenu.UpdateBy = "1"
		Amenu.CreatedAt = timeNow
		Amenu.UpdatedAt = timeNow
		// Amenu.MenuId, err = Amenu.Create(db)
		err = tx.Create(&Amenu).Error
		if err != nil {
			return err
		}

		AList := Menu{}
		AList.MenuName = ""
		AList.Title = "分页获取标签表"
		AList.Icon = "bug"
		AList.Path = "/api/v1/sys-tags"
		AList.MenuType = "A"
		AList.Action = "GET"
		AList.ParentId = Amenu.MenuId
		AList.NoCache = false
		AList.Sort = 0
		AList.Visible = "1"
		AList.IsFrame = "0"
		AList.CreateBy = "1"
		AList.UpdateBy = "1"
		AList.CreatedAt = timeNow
		AList.UpdatedAt = timeNow
		// AList.MenuId, err = AList.Create(db)
		err = tx.Create(&AList).Error
		if err != nil {
			return err
		}

		AGet := Menu{}
		AGet.MenuName = ""
		AGet.Title = "根据id获取标签表"
		AGet.Icon = "bug"
		AGet.Path = "/api/v1/sys-tags/:id"
		AGet.MenuType = "A"
		AGet.Action = "GET"
		AGet.ParentId = Amenu.MenuId
		AGet.NoCache = false
		AGet.Sort = 0
		AGet.Visible = "1"
		AGet.IsFrame = "0"
		AGet.CreateBy = "1"
		AGet.UpdateBy = "1"
		AGet.CreatedAt = timeNow
		AGet.UpdatedAt = timeNow
		// AGet.MenuId, err = AGet.Create(db)
		err = tx.Create(&AGet).Error
		if err != nil {
			return err
		}

		ACreate := Menu{}
		ACreate.MenuName = ""
		ACreate.Title = "创建标签表"
		ACreate.Icon = "bug"
		ACreate.Path = "/api/v1/sys-tags"
		ACreate.MenuType = "A"
		ACreate.Action = "POST"
		ACreate.ParentId = Amenu.MenuId
		ACreate.NoCache = false
		ACreate.Sort = 0
		ACreate.Visible = "1"
		ACreate.IsFrame = "0"
		ACreate.CreateBy = "1"
		ACreate.UpdateBy = "1"
		ACreate.CreatedAt = timeNow
		ACreate.UpdatedAt = timeNow
		// ACreate.MenuId, err = ACreate.Create(db)
		err = tx.Create(&ACreate).Error
		if err != nil {
			return err
		}

		AUpdate := Menu{}
		AUpdate.MenuName = ""
		AUpdate.Title = "修改标签表"
		AUpdate.Icon = "bug"
		AUpdate.Path = "/api/v1/sys-tags/:id"
		AUpdate.MenuType = "A"
		AUpdate.Action = "PUT"
		AUpdate.ParentId = Amenu.MenuId
		AUpdate.NoCache = false
		AUpdate.Sort = 0
		AUpdate.Visible = "1"
		AUpdate.IsFrame = "0"
		AUpdate.CreateBy = "1"
		AUpdate.UpdateBy = "1"
		AUpdate.CreatedAt = timeNow
		AUpdate.UpdatedAt = timeNow
		// AUpdate.MenuId, err = AUpdate.Create(db)
		err = tx.Create(&AUpdate).Error
		if err != nil {
			return err
		}

		ADelete := Menu{}
		ADelete.MenuName = ""
		ADelete.Title = "删除标签表"
		ADelete.Icon = "bug"
		ADelete.Path = "/api/v1/sys-tags"
		ADelete.MenuType = "A"
		ADelete.Action = "DELETE"
		ADelete.ParentId = Amenu.MenuId
		ADelete.NoCache = false
		ADelete.Sort = 0
		ADelete.Visible = "1"
		ADelete.IsFrame = "0"
		ADelete.CreateBy = "1"
		ADelete.UpdateBy = "1"
		ADelete.CreatedAt = timeNow
		ADelete.UpdatedAt = timeNow
		//ADelete.MenuId, err = ADelete.Create(db)
		err = tx.Create(&ADelete).Error
		if err != nil {
			return err
		}

		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
