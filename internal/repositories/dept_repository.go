package repositories

import (
	"main/internal/app"
	"main/internal/models"
)

// DeptRepository 部门仓库接口
type DeptRepository interface {
	FindAll() ([]models.Dept, error)
	FindByID(id int) (*models.Dept, error)
	Create(dept *models.Dept) error
	Update(dept *models.Dept) error
	Delete(id int) error
}

// DeptRepositoryImpl 部门仓库实现
type DeptRepositoryImpl struct{}

// NewDeptRepository 创建部门仓库实例
func NewDeptRepository() DeptRepository {
	return &DeptRepositoryImpl{}
}

// FindAll 查询所有部门
func (r *DeptRepositoryImpl) FindAll() ([]models.Dept, error) {
	var depts []models.Dept
	result := app.DB.Find(&depts)
	return depts, result.Error
}

// FindByID 根据ID查询部门
func (r *DeptRepositoryImpl) FindByID(id int) (*models.Dept, error) {
	var dept models.Dept
	result := app.DB.First(&dept, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dept, nil
}

// Create 创建部门
func (r *DeptRepositoryImpl) Create(dept *models.Dept) error {
	return app.DB.Create(dept).Error
}

// Update 更新部门
func (r *DeptRepositoryImpl) Update(dept *models.Dept) error {
	return app.DB.Save(dept).Error
}

// Delete 删除部门
func (r *DeptRepositoryImpl) Delete(id int) error {
	return app.DB.Delete(&models.Dept{}, id).Error
}
