package postgres_client

import (
	"context"
	"database/sql"

	groupObject "backend/internal/4_domain/group_object"
)

type userEmploymentRecord struct {
	UserID           int
	CompanyID        int
	DepartmentID     int
	PositionID       int
	OfficeLocationID sql.NullInt64
	EmployeeCode     sql.NullString
	EmploymentType   sql.NullString
	IsPrimary        bool
}

func (receiver *PostgresClient) UpdateUserEmployment(
	ctx context.Context,
	userEmployment groupObject.UserEmployment,
) error {
	if err := userEmployment.EnsureReadyToAssign(); err != nil {
		return err
	}

	record := userEmploymentRecord{
		UserID:         userEmployment.UserID().GetValue(),
		CompanyID:      userEmployment.CompanyID().GetValue(),
		DepartmentID:   userEmployment.DepartmentID().GetValue(),
		PositionID:     userEmployment.PositionID().GetValue(),
		EmployeeCode:   stringToNullString(userEmployment.EmployeeCode()),
		EmploymentType: stringToNullString(userEmployment.EmploymentType()),
		IsPrimary:      userEmployment.IsPrimary(),
	}
	if officeLocationID := userEmployment.OfficeLocationID().GetValue(); officeLocationID > 0 {
		record.OfficeLocationID = sql.NullInt64{
			Int64: int64(officeLocationID),
			Valid: true,
		}
	}

	result := receiver.conn(ctx).
		Table("user_employments").
		Where("user_id = ? AND is_primary = ?", record.UserID, record.IsPrimary).
		Updates(map[string]interface{}{
			"company_id":          record.CompanyID,
			"department_id":       record.DepartmentID,
			"position_id":         record.PositionID,
			"office_location_id":  record.OfficeLocationID,
			"employee_code":       record.EmployeeCode,
			"employment_type":     record.EmploymentType,
			"is_primary":          record.IsPrimary,
			"employment_end_date": nil,
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected > 0 {
		return nil
	}

	return receiver.conn(ctx).
		Table("user_employments").
		Create(record).
		Error
}
