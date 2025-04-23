package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/hhdms/msjx/internal/models"
	"github.com/hhdms/msjx/internal/services"
)

// EmpController 员工控制器
type EmpController struct {
	empService services.EmpService
}

// NewEmpController 创建员工控制器实例
func NewEmpController() *EmpController {
	return &EmpController{
		empService: services.NewEmpService(),
	}
}

// GetAllEmps 获取所有员工列表
func (c *EmpController) GetAllEmps(ctx *gin.Context) {
	emps, err := c.empService.GetAllEmps()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse("获取员工列表失败"))
		return
	}

	ctx.JSON(http.StatusOK, models.NewSuccessResponse(emps))
}

// GetEmpPage 获取员工分页列表
func (c *EmpController) GetEmpPage(ctx *gin.Context) {
	var query models.EmpQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse("无效的查询参数"))
		return
	}

	// 设置默认值
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}

	result, err := c.empService.GetEmpPage(&query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse("获取员工列表失败"))
		return
	}

	ctx.JSON(http.StatusOK, models.NewSuccessResponse(result))
}

// GetEmpByID 根据ID获取员工
func (c *EmpController) GetEmpByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse("无效的员工ID"))
		return
	}

	emp, err := c.empService.GetEmpByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewErrorResponse("员工不存在"))
		return
	}

	ctx.JSON(http.StatusOK, models.NewSuccessResponse(emp))
}

// CreateEmp 创建员工
func (c *EmpController) CreateEmp(ctx *gin.Context) {
	var emp models.Emp
	if err := ctx.ShouldBindJSON(&emp); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse("无效的请求参数: "+err.Error()))
		return
	}

	// 处理image字段中可能存在的反引号
	emp.Image = strings.ReplaceAll(emp.Image, "`", "")
	// 去除image字段中可能存在的多余空格
	emp.Image = strings.TrimSpace(emp.Image)

	err := c.empService.CreateEmp(&emp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse("创建员工失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, models.NewSuccessResponse(nil))
}

// UpdateEmp 更新员工
func (c *EmpController) UpdateEmp(ctx *gin.Context) {
	var emp models.Emp
	if err := ctx.ShouldBindJSON(&emp); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse("无效的请求参数"))
		return
	}

	// 处理image字段中可能存在的反引号
	emp.Image = strings.Trim(emp.Image, " `")

	err := c.empService.UpdateEmp(&emp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse("更新员工失败"))
		return
	}

	ctx.JSON(http.StatusOK, models.NewSuccessResponse(nil))
}

// DeleteEmp 删除员工
func (c *EmpController) DeleteEmp(ctx *gin.Context) {
	idsStr := ctx.Query("ids")
	if idsStr == "" {
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse("未指定要删除的员工ID"))
		return
	}

	// 解析ID列表
	idStrList := strings.Split(idsStr, ",")
	ids := make([]int, 0, len(idStrList))
	for _, idStr := range idStrList {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewErrorResponse("无效的员工ID"))
			return
		}
		ids = append(ids, id)
	}

	err := c.empService.DeleteEmp(ids)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse("删除员工失败"))
		return
	}

	ctx.JSON(http.StatusOK, models.NewSuccessResponse(nil))
}
