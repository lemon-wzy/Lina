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

type GiftInfo struct {
	api.Api
}

// GetPage 获取商品列表
// @Summary 获取商品列表
// @Description 获取商品列表
// @Tags 商品
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.GiftInfo}} "{"code": 200, "data": [...]}"
// @Router /api/v1/gift-info [get]
// @Security Bearer
func (e GiftInfo) GetPage(c *gin.Context) {
	// req := dto.GiftInfoGetPageReq{}
	req := dto.CustomGiftInfoGetPageReq{}
	s := service.GiftInfo{}
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
	list := make([]models.GiftInfo, 0)
	var count int64

	// err = s.GetPage(&req, p, &list, &count)
	err = s.CustomGetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取商品失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取商品
// @Summary 获取商品
// @Description 获取商品
// @Tags 商品
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.GiftInfo} "{"code": 200, "data": [...]}"
// @Router /api/v1/gift-info/{id} [get]
// @Security Bearer
func (e GiftInfo) Get(c *gin.Context) {
	req := dto.GiftInfoGetReq{}
	s := service.GiftInfo{}
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
	var object models.GiftInfo

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取商品失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建商品
// @Summary 创建商品
// @Description 创建商品
// @Tags 商品
// @Accept application/json
// @Product application/json
// @Param data body dto.GiftInfoInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/gift-info [post]
// @Security Bearer
func (e GiftInfo) Insert(c *gin.Context) {
	req := dto.GiftInfoInsertReq{}
	s := service.GiftInfo{}
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
		e.Error(500, err, fmt.Sprintf("创建商品失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改商品
// @Summary 修改商品
// @Description 修改商品
// @Tags 商品
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.GiftInfoUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/gift-info/{id} [put]
// @Security Bearer
func (e GiftInfo) Update(c *gin.Context) {
	req := dto.GiftInfoUpdateReq{}
	s := service.GiftInfo{}
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
		e.Error(500, err, fmt.Sprintf("修改商品失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除商品
// @Summary 删除商品
// @Description 删除商品
// @Tags 商品
// @Param data body dto.GiftInfoDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/gift-info [delete]
// @Security Bearer
func (e GiftInfo) Delete(c *gin.Context) {
	s := service.GiftInfo{}
	req := dto.GiftInfoDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除商品失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
