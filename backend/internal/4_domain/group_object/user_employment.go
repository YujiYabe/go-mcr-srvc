package group_object

import (
	"fmt"

	typeObject "backend/internal/4_domain/type_object"
)

type UserEmployment struct {
	userID           typeObject.ID
	companyID        typeObject.ID
	departmentID     typeObject.ID
	positionID       typeObject.ID
	officeLocationID typeObject.ID
	employeeCode     string
	employmentType   string
	isPrimary        bool
}

type NewUserEmploymentArgs struct {
	UserID           *int
	CompanyID        *int
	DepartmentID     *int
	PositionID       *int
	OfficeLocationID *int
	EmployeeCode     *string
	EmploymentType   *string
	IsPrimary        *bool
}

func NewUserEmployment(
	args *NewUserEmploymentArgs,
) (
	userEmployment *UserEmployment,
	err error,
) {
	userEmployment = &UserEmployment{}
	if args == nil {
		args = &NewUserEmploymentArgs{}
	}

	userEmployment.userID, err = typeObject.NewID(args.UserID)
	if err != nil {
		userEmployment = nil
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	userEmployment.companyID, err = typeObject.NewID(args.CompanyID)
	if err != nil {
		userEmployment = nil
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	userEmployment.departmentID, err = typeObject.NewID(args.DepartmentID)
	if err != nil {
		userEmployment = nil
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	userEmployment.positionID, err = typeObject.NewID(args.PositionID)
	if err != nil {
		userEmployment = nil
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	userEmployment.officeLocationID, err = typeObject.NewID(args.OfficeLocationID)
	if err != nil {
		userEmployment = nil
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	if args.EmployeeCode != nil {
		userEmployment.employeeCode = *args.EmployeeCode
	}
	if args.EmploymentType != nil {
		userEmployment.employmentType = *args.EmploymentType
	}
	if args.IsPrimary != nil {
		userEmployment.isPrimary = *args.IsPrimary
	}

	err = nil
	return //nolint:nakedret // Use the project-wide named return convention.
}

func (receiver UserEmployment) UserID() (
	iD typeObject.ID,
) {
	iD = receiver.userID
	return
}

func (receiver UserEmployment) CompanyID() (
	iD typeObject.ID,
) {
	iD = receiver.companyID
	return
}

func (receiver UserEmployment) DepartmentID() (
	iD typeObject.ID,
) {
	iD = receiver.departmentID
	return
}

func (receiver UserEmployment) PositionID() (
	iD typeObject.ID,
) {
	iD = receiver.positionID
	return
}

func (receiver UserEmployment) OfficeLocationID() (
	iD typeObject.ID,
) {
	iD = receiver.officeLocationID
	return
}

func (receiver UserEmployment) EmployeeCode() (
	value string,
) {
	value = receiver.employeeCode
	return
}

func (receiver UserEmployment) EmploymentType() (
	value string,
) {
	value = receiver.employmentType
	return
}

func (receiver UserEmployment) IsPrimary() (
	isPrimary bool,
) {
	isPrimary = receiver.isPrimary
	return
}

func (receiver UserEmployment) EnsureReadyToAssign() (
	err error,
) {
	if receiver.userID.GetValue() <= 0 {
		err = fmt.Errorf("user employment user identity is required")
		return
	}
	if receiver.companyID.GetValue() <= 0 {
		err = fmt.Errorf("user employment company identity is required")
		return
	}
	if receiver.departmentID.GetValue() <= 0 {
		err = fmt.Errorf("user employment department identity is required")
		return
	}
	if receiver.positionID.GetValue() <= 0 {
		err = fmt.Errorf("user employment position identity is required")
		return
	}

	err = nil
	return
}
