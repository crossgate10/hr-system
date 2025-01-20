package user

import "gorm.io/gorm"

type Repository interface {
	GetApproversByRoleAndDepartment(roleName string, departmentName string) ([]Approver, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetApproversByRoleAndDepartment(roleName string, departmentName string) ([]Approver, error) {
	var approvers []Approver

	// 子查詢，用於獲取 role_id 和 department_id
	var role Role
	var department Department

	if err := r.db.Where("role_name = ?", roleName).First(&role).Error; err != nil {
		return nil, err
	}

	if err := r.db.Where("department_name = ?", departmentName).First(&department).Error; err != nil {
		return nil, err
	}

	// 主查詢，使用子查詢結果來過濾
	if err := r.db.Table("Role_Department AS rd").
		Joins("JOIN Approvers AS a ON rd.approver_id = a.id").
		Where("rd.role_id = ? AND rd.department_id = ?", role.ID, department.ID).
		Order("a.seq").
		Select("a.*").
		Find(&approvers).Error; err != nil {
		return nil, err
	}

	return approvers, nil
}
