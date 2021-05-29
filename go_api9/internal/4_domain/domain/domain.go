package domain

import (
	"context"
)

type (
	domain struct{}
)

// infrastructure から来た用
type (
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

// infrastructure へ行く用
type (

	// Assemble ...
	Assemble struct {
		Bans        map[string]int
		Patties     map[string]int
		Vegetables  map[string]int
		Ingredients map[string]int
	}
)

// NewDomain ...
func NewDomain() *domain {
	return &domain{}
}

// Dummy ...
func (dm *domain) Dummy(ctx context.Context) error {
	return nil
}

// ParseOrder ...
func (dm *domain) ParseOrder(ctx context.Context, order Order) *Assemble {
	assemble := &Assemble{
		Bans:        map[string]int{},
		Patties:     map[string]int{},
		Vegetables:  map[string]int{},
		Ingredients: map[string]int{},
	}

	if len(order.Hamburgers) != 0 {
		dm.countAssembleHamburger(ctx, assemble, order.Hamburgers)
	}

	return assemble
}

func (dm *domain) countAssembleHamburger(ctx context.Context, assemble *Assemble, hamburgers []Hamburger) {
	for _, hamburger := range hamburgers {
		// bans
		assemble.Bans["top"] += hamburger.Top
		assemble.Bans["Middle"] += hamburger.Middle
		assemble.Bans["Bottom"] += hamburger.Bottom

		// patty
		assemble.Patties["Beef"] += hamburger.Beef
		assemble.Patties["Chicken"] += hamburger.Chicken
		assemble.Patties["Fish"] += hamburger.Fish

		//vegetable
		assemble.Vegetables["Lettuce"] += hamburger.Lettuce
		assemble.Vegetables["Tomato"] += hamburger.Tomato
		assemble.Vegetables["Onion"] += hamburger.Onion

		//ingredient
		assemble.Ingredients["Cheese"] += hamburger.Cheese
		assemble.Ingredients["Pickles"] += hamburger.Pickles
	}

	return
}
