package models

import (
	"time"
)

// Emp 员工实体
type Emp struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement;comment:'员工ID'"`
	Username   string    `json:"username" gorm:"type:varchar(20);not null;unique;comment:'用户名（2-20位）'"`
	Password   string    `json:"password,omitempty" gorm:"type:varchar(64);default:'123456';comment:'登录密码'"`
	Name       string    `json:"name" gorm:"type:varchar(10);not null;comment:'姓名（2-10位）'"`
	Gender     int8      `json:"gender" gorm:"type:tinyint;comment:'性别, 1: 男, 2: 女'"`
	Phone      string    `json:"phone" gorm:"type:char(11);not null;unique;comment:'手机号（11位）'"`
	Position   int8      `json:"position" gorm:"type:tinyint;comment:'职位, 1: 班主任, 2: 讲师, 3: 学工主管, 4: 教研主管, 5: 咨询师'"`
	Salary     int       `json:"salary" gorm:"type:int;comment:'薪资（整数存储）'"`
	Image      string    `json:"image" gorm:"type:varchar(255);comment:'头像路径'"`
	HireDate   time.Time `json:"hireDate" gorm:"type:date;comment:'入职日期'"`
	DeptID     int       `json:"deptId" gorm:"type:int;comment:'所属部门ID'"`
	CreateTime time.Time `json:"createTime" gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateTime time.Time `json:"updateTime" gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;comment:'最后操作时间'"`
	DeptName   string    `json:"deptName" gorm:"-"` // 部门名称，不映射到数据库
	ExprList   []EmpExpr `json:"exprList,omitempty" gorm:"foreignKey:EmpID"` // 工作经历列表
}

// TableName 设置表名
func (Emp) TableName() string {
	return "emp"
}

// EmpExpr 员工工作经历实体
type EmpExpr struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;comment:'经历ID'"`
	EmpID     int       `json:"empId" gorm:"type:int;comment:'关联员工的ID'"`
	Company   string    `json:"company" gorm:"type:varchar(50);comment:'公司名称'"`
	Position  string    `json:"position" gorm:"type:varchar(50);comment:'担任职位'"`
	StartDate time.Time `json:"startDate" gorm:"type:date;comment:'开始日期'"`
	EndDate   time.Time `json:"endDate" gorm:"type:date;comment:'结束日期'"`
}

// TableName 设置表名
func (EmpExpr) TableName() string {
	return "emp_expr"
}

// EmpQuery 员工查询参数
type EmpQuery struct {
	Name     string `form:"name"`
	Gender   int8   `form:"gender"`
	Begin    string `form:"begin"`
	End      string `form:"end"`
	Page     int    `form:"page,default=1"`
	PageSize int    `form:"pageSize,default=10"`
}

// EmpPageResult 员工分页结果
type EmpPageResult struct {
	Total int64 `json:"total"`
	Rows  []Emp `json:"rows"`
}