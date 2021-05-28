package domain

import "context"

type (
	domain struct{}

	// Order ...
	Order struct {
		Combos     []Combo     `json:"combos"`
		Hamburgers []Hamburger `json:"hamburgers"`
		SideMenus  []SideMenu  `json:"side_menus"`
		Drinks     []Drink     `json:"drinks"`
	}

	// Combo ...
	Combo struct {
		Hamburger *Hamburger `json:"hamburger"`
		SideMenu  *SideMenu  `json:"side_menu"`
		Drink     *Drink     `json:"drink"`
	}

	// Hamburger ...
	Hamburger struct {
		// bans
		Top    int `json:"top"`
		Middle int `json:"middle"`
		Bottom int `json:"bottom"`
		// patty
		Beef    int `json:"beef"`
		Chicken int `json:"chicken"`
		Fish    int `json:"fish"`
		//vegetable
		Lettuce int `json:"lettuce"`
		Tomato  int `json:"tomato"`
		Onion   int `json:"onion"`
		//ingredient
		Cheese  int `json:"cheese"`
		Pickles int `json:"pickles"`
	}

	// SideMenu ...
	SideMenu struct {
		Name string `json:"name"`
	}

	// Drink ...
	Drink struct {
		Name string `json:"name"`
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
