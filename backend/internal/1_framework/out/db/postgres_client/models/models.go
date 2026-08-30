package models

import "database/sql"

type User struct {
	ID          int            `gorm:"column:id;primaryKey"`
	Auth0UserID sql.NullString `gorm:"column:auth0_user_id"`
	Email       string         `gorm:"column:email"`
	FullName    sql.NullString `gorm:"column:full_name"`
	CreatedAt   sql.NullTime   `gorm:"column:created_at"`
	UpdatedAt   sql.NullTime   `gorm:"column:updated_at"`
}

func (User) TableName() (
	tableName string,
) {
	return "users"
}

type Role struct {
	ID          int            `gorm:"column:id;primaryKey"`
	RoleName    string         `gorm:"column:role_name"`
	Description sql.NullString `gorm:"column:description"`
}

func (Role) TableName() (
	tableName string,
) {
	return "roles"
}

type UserRole struct {
	ID     int           `gorm:"column:id;primaryKey"`
	UserID sql.NullInt64 `gorm:"column:user_id"`
	RoleID sql.NullInt64 `gorm:"column:role_id"`
}

func (UserRole) TableName() (
	tableName string,
) {
	return "user_roles"
}

type UserEmployment struct {
	ID               int            `gorm:"column:id;primaryKey"`
	UserID           int            `gorm:"column:user_id"`
	CompanyID        int            `gorm:"column:company_id"`
	DepartmentID     int            `gorm:"column:department_id"`
	PositionID       int            `gorm:"column:position_id"`
	OfficeLocationID sql.NullInt64  `gorm:"column:office_location_id"`
	EmployeeCode     sql.NullString `gorm:"column:employee_code"`
	EmploymentType   sql.NullString `gorm:"column:employment_type"`
	JoinedOn         sql.NullTime   `gorm:"column:joined_on"`
	LeftOn           sql.NullTime   `gorm:"column:left_on"`
	IsPrimary        bool           `gorm:"column:is_primary"`
	CreatedAt        sql.NullTime   `gorm:"column:created_at"`
	UpdatedAt        sql.NullTime   `gorm:"column:updated_at"`
}

func (UserEmployment) TableName() (
	tableName string,
) {
	return "user_employments"
}

type ValidationWordRule struct {
	ID          int          `gorm:"column:id;primaryKey"`
	TargetType  string       `gorm:"column:target_type"`
	IsBlacklist bool         `gorm:"column:is_blacklist"`
	Word        string       `gorm:"column:word"`
	MatchType   string       `gorm:"column:match_type"`
	Enabled     bool         `gorm:"column:enabled"`
	CreatedAt   sql.NullTime `gorm:"column:created_at"`
	UpdatedAt   sql.NullTime `gorm:"column:updated_at"`
}

func (ValidationWordRule) TableName() (
	tableName string,
) {
	return "validation_word_rules"
}
