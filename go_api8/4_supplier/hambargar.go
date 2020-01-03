package supplier

// type StandardHambargar struct {
// 	ID    int
// 	Name  string
// 	Stock string
// }

// StandardHambargar ...
type StandardHambargar struct {
	bans    bool
	cheese  bool
	tomato  bool
	lettuce bool
}

// StandardHambargars ...
type StandardHambargars []StandardHambargar
