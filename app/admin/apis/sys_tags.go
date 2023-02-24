package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type SysTags struct {
	api.Api
}

// GetPage 获取标签表列表
// @Summary 获取标签表列表
// @Description 获取标签表列表
// @Tags 标签表
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysTags}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-tags [get]
// @Security Bearer
func (e SysTags) GetPage(c *gin.Context) {
	req := dto.SysTagsGetPageReq{}
	s := service.SysTags{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.SysTags, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取标签表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.Logger.Debug(req.GetPageIndex())
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// GetList 获取标签表列表
// @Summary 获取标签表列表
// @Description 获取标签表列表
// @Tags 标签表列表
// @Success 200 {object} response.Response{data=response.Response{list=[]models.SysTags}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-tags/list [get]
// @Security Bearer
func (e SysTags) GetList(c *gin.Context) {
	req := dto.SysTagsGetPageReq{}
	s := service.SysTags{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	list := make([]models.SysTags, 0)

	err = s.GetList(&req, &list)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取标签表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(list, "查询成功")
}

// Get 获取标签表
// @Summary 获取标签表
// @Description 获取标签表
// @Tags 标签表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SysTags} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-tags/{id} [get]
// @Security Bearer
func (e SysTags) Get(c *gin.Context) {
	req := dto.SysTagsGetReq{}
	s := service.SysTags{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SysTags

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取标签表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建标签表
// @Summary 创建标签表
// @Description 创建标签表
// @Tags 标签表
// @Accept application/json
// @Product application/json
// @Param data body dto.SysTagsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sys-tags [post]
// @Security Bearer
func (e SysTags) Insert(c *gin.Context) {
	req := dto.SysTagsInsertReq{}
	s := service.SysTags{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建标签表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改标签表
// @Summary 修改标签表
// @Description 修改标签表
// @Tags 标签表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysTagsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sys-tags/{id} [put]
// @Security Bearer
func (e SysTags) Update(c *gin.Context) {
	req := dto.SysTagsUpdateReq{}
	s := service.SysTags{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改标签表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除标签表
// @Summary 删除标签表
// @Description 删除标签表
// @Tags 标签表
// @Param data body dto.SysTagsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sys-tags [delete]
// @Security Bearer
func (e SysTags) Delete(c *gin.Context) {
	s := service.SysTags{}
	req := dto.SysTagsDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除标签表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
