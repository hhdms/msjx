package repositories

import (
	"gorm.io/gorm"

	"github.com/hhdms/msjx/internal/app"
	"github.com/hhdms/msjx/internal/models"
)

// EmpRepository 员工仓库接口
type EmpRepository interface {
	FindPage(query *models.EmpQuery) (*models.EmpPageResult, error)
	FindByID(id int) (*models.Emp, error)
	FindByUsername(username string) (*models.Emp, error)
	FindByPhone(phone string) (*models.Emp, error)
	Create(emp *models.Emp) error
	Update(emp *models.Emp) error
	Delete(ids []int) error
}

// EmpRepositoryImpl 员工仓库实现
type EmpRepositoryImpl struct{}

// NewEmpRepository 创建员工仓库实例
func NewEmpRepository() EmpRepository {
	return &EmpRepositoryImpl{}
}

// FindPage 分页查询员工
func (r *EmpRepositoryImpl) FindPage(query *models.EmpQuery) (*models.EmpPageResult, error) {
	db := app.DB.Model(&models.Emp{})

	// 构建查询条件
	if query.Name != "" {
		db = db.Where("name LIKE ?", "%"+query.Name+"%")
	}

	if query.Gender > 0 {
		db = db.Where("gender = ?", query.Gender)
	}

	if query.Begin != "" && query.End != "" {
		db = db.Where("hire_date BETWEEN ? AND ?", query.Begin, query.End)
	} else if query.Begin != "" {
		db = db.Where("hire_date >= ?", query.Begin)
	} else if query.End != "" {
		db = db.Where("hire_date <= ?", query.End)
	}

	// 查询总记录数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	// 分页查询
	var emps []models.Emp
	offset := (query.Page - 1) * query.PageSize
	result := db.Offset(offset).Limit(query.PageSize).Find(&emps)
	if result.Error != nil {
		return nil, result.Error
	}

	// 查询部门名称
	for i := range emps {
		if emps[i].DeptID > 0 {
			var dept models.Dept
			if err := app.DB.Select("name").First(&dept, emps[i].DeptID).Error; err == nil {
				emps[i].DeptName = dept.Name
			}
		}
	}

	return &models.EmpPageResult{
		Total: total,
		Rows:  emps,
	}, nil
}

// FindByID 根据ID查询员工
func (r *EmpRepositoryImpl) FindByID(id int) (*models.Emp, error) {
	var emp models.Emp
	result := app.DB.First(&emp, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 查询部门名称
	if emp.DeptID > 0 {
		var dept models.Dept
		if err := app.DB.Select("name").First(&dept, emp.DeptID).Error; err == nil {
			emp.DeptName = dept.Name
		}
	}

	// 查询工作经历
	if err := app.DB.Where("emp_id = ?", id).Find(&emp.ExprList).Error; err != nil {
		return nil, err
	}

	return &emp, nil
}

// FindByUsername 根据用户名查询员工
func (r *EmpRepositoryImpl) FindByUsername(username string) (*models.Emp, error) {
	var emp models.Emp
	result := app.DB.Where("username = ?", username).First(&emp)
	if result.Error != nil {
		return nil, result.Error
	}

	// 查询部门名称
	if emp.DeptID > 0 {
		var dept models.Dept
		if err := app.DB.Select("name").First(&dept, emp.DeptID).Error; err == nil {
			emp.DeptName = dept.Name
		}
	}

	return &emp, nil
}

// FindByPhone 根据手机号查询员工
func (r *EmpRepositoryImpl) FindByPhone(phone string) (*models.Emp, error) {
	var emp models.Emp
	result := app.DB.Where("phone = ?", phone).First(&emp)
	if result.Error != nil {
		return nil, result.Error
	}
	return &emp, nil
}

// Create 创建员工
func (r *EmpRepositoryImpl) Create(emp *models.Emp) error {
	return app.DB.Transaction(func(tx *gorm.DB) error {
		// 创建员工
		if err := tx.Create(emp).Error; err != nil {
			return err
		}

		// 创建工作经历
		for i := range emp.ExprList {
			emp.ExprList[i].EmpID = emp.ID
			if err := tx.Create(&emp.ExprList[i]).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// Update 更新员工
func (r *EmpRepositoryImpl) Update(emp *models.Emp) error {
	return app.DB.Transaction(func(tx *gorm.DB) error {
		// 更新员工基本信息
		if err := tx.Model(emp).Omit("CreateTime", "UpdateTime").Updates(emp).Error; err != nil {
			return err
		}

		// 如果有工作经历，先删除原有的，再添加新的
		if len(emp.ExprList) > 0 {
			// 删除原有工作经历
			if err := tx.Where("emp_id = ?", emp.ID).Delete(&models.EmpExpr{}).Error; err != nil {
				return err
			}

			// 添加新的工作经历
			for i := range emp.ExprList {
				// 设置员工ID并清除工作经历ID，让数据库自动生成ID
				emp.ExprList[i].EmpID = emp.ID
				emp.ExprList[i].ID = 0 // 重置ID为0，让数据库自动生成新的ID
				if err := tx.Create(&emp.ExprList[i]).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})
}

// Delete 删除员工
func (r *EmpRepositoryImpl) Delete(ids []int) error {
	return app.DB.Transaction(func(tx *gorm.DB) error {
		// 删除员工的工作经历
		if err := tx.Where("emp_id IN (?)", ids).Delete(&models.EmpExpr{}).Error; err != nil {
			return err
		}

		// 删除员工
		if err := tx.Delete(&models.Emp{}, ids).Error; err != nil {
			return err
		}

		return nil
	})
}
