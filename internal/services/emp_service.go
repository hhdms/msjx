package services

import (
	"gorm.io/gorm"

	"github.com/hhdms/msjx/internal/models"
	"github.com/hhdms/msjx/internal/repositories"
)

// EmpService 员工服务接口
type EmpService interface {
	GetEmpPage(query *models.EmpQuery) (*models.EmpPageResult, error)
	GetAllEmps() ([]models.Emp, error)
	GetEmpByID(id int) (*models.Emp, error)
	CreateEmp(emp *models.Emp) error
	UpdateEmp(emp *models.Emp) error
	DeleteEmp(ids []int) error
}

// EmpServiceImpl 员工服务实现
type EmpServiceImpl struct {
	empRepo repositories.EmpRepository
}

// NewEmpService 创建员工服务实例
func NewEmpService() EmpService {
	return &EmpServiceImpl{
		empRepo: repositories.NewEmpRepository(),
	}
}

// GetEmpPage 分页查询员工
func (s *EmpServiceImpl) GetEmpPage(query *models.EmpQuery) (*models.EmpPageResult, error) {
	return s.empRepo.FindPage(query)
}

// GetAllEmps 获取所有员工
func (s *EmpServiceImpl) GetAllEmps() ([]models.Emp, error) {
	// 使用分页查询但设置一个很大的页面大小来获取所有记录
	query := &models.EmpQuery{
		Page:     1,
		PageSize: 1000, // 设置一个足够大的值以获取所有记录
	}

	result, err := s.empRepo.FindPage(query)
	if err != nil {
		return nil, err
	}

	return result.Rows, nil
}

// GetEmpByID 根据ID查询员工
func (s *EmpServiceImpl) GetEmpByID(id int) (*models.Emp, error) {
	return s.empRepo.FindByID(id)
}

// CreateEmp 创建员工
func (s *EmpServiceImpl) CreateEmp(emp *models.Emp) error {
	return s.empRepo.Create(emp)
}

// UpdateEmp 更新员工
func (s *EmpServiceImpl) UpdateEmp(emp *models.Emp) error {
	// 先查询员工是否存在
	existEmp, err := s.empRepo.FindByID(emp.ID)
	if err != nil {
		return err
	}

	// 确保员工存在
	if existEmp == nil {
		return gorm.ErrRecordNotFound
	}

	// 使用事务更新员工信息和工作经历
	return s.empRepo.Update(emp)
}

// DeleteEmp 删除员工
func (s *EmpServiceImpl) DeleteEmp(ids []int) error {
	return s.empRepo.Delete(ids)
}
