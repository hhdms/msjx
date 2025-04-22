package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/hhdms/msjx/internal/controllers"
	"github.com/hhdms/msjx/internal/middleware"
)

// RegisterRoutes 注册API路由
func RegisterRoutes(r *gin.Engine) {
	// 创建认证控制器
	authController := controllers.NewAuthController()

	// 登录路由 - 无需认证
	r.POST("/login", authController.Login)

	// 创建文件上传控制器
	uploadController := controllers.NewUploadController()

	// 文件上传路由 - 无需认证
	r.POST("/upload", uploadController.Upload)

	// 需要认证的路由
	authorized := r.Group("/")
	authorized.Use(middleware.JWTAuth())
	{
		// 创建部门控制器
		deptController := controllers.NewDeptController()

		// 部门管理路由
		authorized.GET("/depts", deptController.GetAllDepts)
		authorized.GET("/depts/:id", deptController.GetDeptByID)
		authorized.POST("/depts", deptController.CreateDept)
		authorized.PUT("/depts", deptController.UpdateDept)
		authorized.DELETE("/depts/:id", deptController.DeleteDept)

		// 创建员工控制器
		empController := controllers.NewEmpController()

		// 员工管理路由
		authorized.GET("/emps", empController.GetEmpPage)
		authorized.GET("/emps/list", empController.GetAllEmps)
		authorized.GET("/emps/:id", empController.GetEmpByID)
		authorized.POST("/emps", empController.CreateEmp)
		authorized.PUT("/emps", empController.UpdateEmp)
		authorized.DELETE("/emps", empController.DeleteEmp)
	}
}
