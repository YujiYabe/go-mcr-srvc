package domain

import "context"

type (
	domain struct{}

	// Order ...
	Order struct {
		Combos     []Combo
		Hamburgers []Hamburger
		SideMenus  []SideMenu
		Drinks     []Drink
	}

	// Combo ...
	Combo struct {
		Hamburger *Hamburger
		SideMenu  *SideMenu
		Drink     *Drink
	}

	// Hamburger ...
	Hamburger struct {
		BansTop      int
		BansMiddle   int
		BansBottom   int
		BeefPatty    int
		ChickenPatty int
		FishPatty    int
		lettuce      int
		pickles      int
	}

	// SideMenu ...
	SideMenu struct {
		name string
	}

	// Drink ...
	Drink struct {
		name string
	}
)

// NewDomain ...
func NewDomain() *domain {
	return &domain{}
}

// Dummy ...
func (domain *domain) Dummy(ctx context.Context) error {
	return nil
}
