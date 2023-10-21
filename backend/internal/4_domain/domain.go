package domain

import (
	"context"

	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("enterprise_business_rule", "domain")
}

// NewDomain ...
func NewDomain() ToDomain {
	return &domain{}
}



// ParseOrder ...
func (dmn *domain) ParseOrder(
	ctx context.Context,
) error {

	return nil
}

// CookHamburgers ...
func (dmn *domain) CookHamburgers(
	ctx context.Context,
) error {
	return nil
}
