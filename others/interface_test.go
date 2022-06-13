package others

import "testing"

func Test_interface(t *testing.T) {

	// c1 := Country {"China"}
	// c2 := City {"Beijing"}
	// c1.PrintStr()
	// c2.PrintStr()

	// c1 := Country {WithName{ "China"}}
	// c2 := City { WithName{"Beijing"}}
	// c1.PrintStr()
	// c2.PrintStr()

	d1 := Country {"USA"}
	d2 := City{"Los Angeles"}
	PrintStr(d1)
	PrintStr(d2)
}