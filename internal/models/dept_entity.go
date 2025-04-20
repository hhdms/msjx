package models

import (
	"time"
)

// Dept 部门实体
type Dept struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement;comment:'部门ID（唯一标识，自增主键）'"`
	Name       string    `json:"name" gorm:"type:varchar(10);not null;unique;comment:'部门名称'"`
	CreateTime time.Time `json:"createTime" gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateTime time.Time `json:"updateTime" gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;comment:'最后操作时间'"`
}

// TableName 设置表名
func (Dept) TableName() string {
	return "dept"
}
