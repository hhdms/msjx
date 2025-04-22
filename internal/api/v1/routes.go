package v1

import (
	"github.com/gin-gonic/gin"
	
	"github.com/hhdms/msjx/internal/controllers"
)

// RegisterRoutes 注册API路由
func RegisterRoutes(r *gin.Engine) {
	// 创建部门控制器
	deptController := controllers.NewDeptController()

	// 部门管理路由
	r.GET("/depts", deptController.GetAllDepts)
	r.GET("/depts/:id", deptController.GetDeptByID)
	r.POST("/depts", deptController.CreateDept)
	r.PUT("/depts", deptController.UpdateDept)
	r.DELETE("/depts/:id", deptController.DeleteDept)

	// 创建员工控制器
	empController := controllers.NewEmpController()

	// 员工管理路由
	r.GET("/emps", empController.GetEmpPage)
	r.GET("/emps/list", empController.GetAllEmps)
	r.GET("/emps/:id", empController.GetEmpByID)
	r.POST("/emps", empController.CreateEmp)
	r.PUT("/emps", empController.UpdateEmp)
	r.DELETE("/emps", empController.DeleteEmp)
}
