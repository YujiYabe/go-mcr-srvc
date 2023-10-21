package domain

import "backend/pkg"

func init() {
	myErr = pkg.NewMyErr("enterprise_business_rule", "domain")
}
