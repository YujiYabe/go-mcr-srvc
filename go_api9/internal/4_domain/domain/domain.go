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
		// bans
		Top    int
		Middle int
		Bottom int
		// patty
		Beef    int
		Chicken int
		Fish    int
		//vegetable
		lettuce int
		tomato  int
		pickles int
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
