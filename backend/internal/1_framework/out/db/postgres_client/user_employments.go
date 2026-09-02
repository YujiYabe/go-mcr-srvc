package postgres_client

import (
	"context"
	"database/sql"

	"backend/internal/1_framework/out/db/postgres_client/models"
	groupObject "backend/internal/4_domain/group_object"
)

func (receiver *PostgresClient) UpdateUserEmployment(
	ctx context.Context,
	userEmployment groupObject.UserEmployment,
) (
	err error,
) {
	if returnedErr := userEmployment.EnsureReadyToAssign(); returnedErr != nil {
		err = returnedErr
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	record := models.UserEmployment{
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
		Model(&models.UserEmployment{}).
		Where("user_id = ? AND is_primary = ?", record.UserID, record.IsPrimary).
		Select(
			"company_id",
			"department_id",
			"position_id",
			"office_location_id",
			"employee_code",
			"employment_type",
			"is_primary",
			"left_on",
		).
		Updates(&record)
	if result.Error != nil {
		err = result.Error
		return //nolint:nakedret // Use the project-wide named return convention.
	}
	if result.RowsAffected > 0 {
		err = nil
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	err = receiver.conn(ctx).
		Omit("JoinedOn", "LeftOn", "CreatedAt", "UpdatedAt").
		Create(&record).
		Error
	return //nolint:nakedret // Use the project-wide named return convention.
}
