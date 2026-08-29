package models

import "database/sql"

type User struct {
	ID          int            `gorm:"column:id;primaryKey" json:"id"`
	Auth0UserID sql.NullString `gorm:"column:auth0_user_id" json:"auth0_user_id"`
	Email       string         `gorm:"column:email" json:"email"`
	FullName    sql.NullString `gorm:"column:full_name" json:"full_name"`
	CreatedAt   sql.NullTime   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   sql.NullTime   `gorm:"column:updated_at" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

type Role struct {
	ID          int            `gorm:"column:id;primaryKey" json:"id"`
	RoleName    string         `gorm:"column:role_name" json:"role_name"`
	Description sql.NullString `gorm:"column:description" json:"description"`
}

func (Role) TableName() string {
	return "roles"
}

type UserRole struct {
	ID     int           `gorm:"column:id;primaryKey" json:"id"`
	UserID sql.NullInt64 `gorm:"column:user_id" json:"user_id"`
	RoleID sql.NullInt64 `gorm:"column:role_id" json:"role_id"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
