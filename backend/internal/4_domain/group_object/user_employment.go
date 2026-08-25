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
		return nil, err
	}

	userEmployment.companyID, err = typeObject.NewID(args.CompanyID)
	if err != nil {
		return nil, err
	}

	userEmployment.departmentID, err = typeObject.NewID(args.DepartmentID)
	if err != nil {
		return nil, err
	}

	userEmployment.positionID, err = typeObject.NewID(args.PositionID)
	if err != nil {
		return nil, err
	}

	userEmployment.officeLocationID, err = typeObject.NewID(args.OfficeLocationID)
	if err != nil {
		return nil, err
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

	return userEmployment, nil
}

func (receiver UserEmployment) UserID() typeObject.ID {
	return receiver.userID
}

func (receiver UserEmployment) CompanyID() typeObject.ID {
	return receiver.companyID
}

func (receiver UserEmployment) DepartmentID() typeObject.ID {
	return receiver.departmentID
}

func (receiver UserEmployment) PositionID() typeObject.ID {
	return receiver.positionID
}

func (receiver UserEmployment) OfficeLocationID() typeObject.ID {
	return receiver.officeLocationID
}

func (receiver UserEmployment) EmployeeCode() string {
	return receiver.employeeCode
}

func (receiver UserEmployment) EmploymentType() string {
	return receiver.employmentType
}

func (receiver UserEmployment) IsPrimary() bool {
	return receiver.isPrimary
}

func (receiver UserEmployment) EnsureReadyToAssign() error {
	if receiver.userID.GetValue() <= 0 {
		return fmt.Errorf("user employment user identity is required")
	}
	if receiver.companyID.GetValue() <= 0 {
		return fmt.Errorf("user employment company identity is required")
	}
	if receiver.departmentID.GetValue() <= 0 {
		return fmt.Errorf("user employment department identity is required")
	}
	if receiver.positionID.GetValue() <= 0 {
		return fmt.Errorf("user employment position identity is required")
	}

	return nil
}

func (receiver UserEmployment) EnsureBelongsTo(
	user User,
) error {
	if err := receiver.EnsureReadyToAssign(); err != nil {
		return err
	}
	if !user.HasIdentity() {
		return fmt.Errorf("user identity is required")
	}
	if receiver.userID.GetValue() != user.Identity().GetValue() {
		return fmt.Errorf("user employment must belong to the user")
	}

	return nil
}
