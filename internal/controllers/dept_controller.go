package controllers

import (
	"net/http"
	"strconv"

	"main/internal/models"
	"main/internal/services"

	"github.com/gin-gonic/gin"
)

// DeptController 部门控制器
type DeptController struct {
	deptService services.DeptService
}

// NewDeptController 创建部门控制器实例
func NewDeptController() *DeptController {
	return &DeptController{
		deptService: services.NewDeptService(),
	}
}

// GetAllDepts 获取所有部门
func (c *DeptController) GetAllDepts(ctx *gin.Context) {
	depts, err := c.deptService.GetAllDepts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse("获取部门列表失败"))
		return
	}
	ctx.JSON(http.StatusOK, models.NewSuccessResponse(depts))
}

// GetDeptByID 根据ID获取部门
func (c *DeptController) GetDeptByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse("无效的部门ID"))
		return
	}

	dept, err := c.deptService.GetDeptByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewErrorResponse("部门不存在"))
		return
	}

	ctx.JSON(http.StatusOK, models.NewSuccessResponse(dept))
}

// CreateDept 创建部门
func (c *DeptController) CreateDept(ctx *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse("无效的请求参数"))
		return
	}

	err := c.deptService.CreateDept(req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse("创建部门失败"))
		return
	}

	ctx.JSON(http.StatusOK, models.NewSuccessResponse(nil))
}

// UpdateDept 更新部门
func (c *DeptController) UpdateDept(ctx *gin.Context) {
	var req struct {
		ID   int    `json:"id" binding:"required"`
		Name string `json:"name" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse("无效的请求参数"))
		return
	}

	err := c.deptService.UpdateDept(req.ID, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse("更新部门失败"))
		return
	}

	ctx.JSON(http.StatusOK, models.NewSuccessResponse(nil))
}

// DeleteDept 删除部门
func (c *DeptController) DeleteDept(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse("无效的部门ID"))
		return
	}

	err = c.deptService.DeleteDept(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse("删除部门失败"))
		return
	}

	ctx.JSON(http.StatusOK, models.NewSuccessResponse(nil))
}
