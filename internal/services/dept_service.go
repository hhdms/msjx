package services

import (
	"github.com/hhdms/msjx/internal/models"
	"github.com/hhdms/msjx/internal/repositories"
)

// DeptService 部门服务接口
type DeptService interface {
	GetAllDepts() ([]models.Dept, error)
	GetDeptByID(id int) (*models.Dept, error)
	CreateDept(name string) error
	UpdateDept(id int, name string) error
	DeleteDept(id int) error
}

// DeptServiceImpl 部门服务实现
type DeptServiceImpl struct {
	deptRepo repositories.DeptRepository
}

// NewDeptService 创建部门服务实例
func NewDeptService() DeptService {
	return &DeptServiceImpl{
		deptRepo: repositories.NewDeptRepository(),
	}
}

// GetAllDepts 获取所有部门
func (s *DeptServiceImpl) GetAllDepts() ([]models.Dept, error) {
	return s.deptRepo.FindAll()
}

// GetDeptByID 根据ID获取部门
func (s *DeptServiceImpl) GetDeptByID(id int) (*models.Dept, error) {
	return s.deptRepo.FindByID(id)
}

// CreateDept 创建部门
func (s *DeptServiceImpl) CreateDept(name string) error {
	dept := &models.Dept{
		Name: name,
	}
	return s.deptRepo.Create(dept)
}

// UpdateDept 更新部门
func (s *DeptServiceImpl) UpdateDept(id int, name string) error {
	dept, err := s.deptRepo.FindByID(id)
	if err != nil {
		return err
	}
	dept.Name = name
	return s.deptRepo.Update(dept)
}

// DeleteDept 删除部门
func (s *DeptServiceImpl) DeleteDept(id int) error {
	return s.deptRepo.Delete(id)
}
