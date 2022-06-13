package others

import "fmt"

// type Country struct {
// 	Name string
// }

// type City struct {
// 	Name string
// }

// type Printable interface {
// 	PrintStr()
// }

// func (c Country) PrintStr() {
// 	fmt.Println(c.Name)
// }

// func (c City) PrintStr() {
// 	fmt.Println(c.Name)
// }

// type WithName struct {
// 	Name string
// }

// type Country struct {
// 	WithName
// }

// type City struct {
// 	WithName
// }

// type Printable interface {
// 	PrintStr()
// }

// func (w WithName) PrintStr() {
// 	fmt.Println(w.Name)
// }


type Country struct {
	Name string
}
    
type City struct {
	Name string
}
    
type Stringable interface {
	ToString() string
}
    
func (c Country) ToString() string {
	return "Country = " + c.Name
}
    
func (c City) ToString() string{
	return "City = " + c.Name
    
}
    
    
func PrintStr(p Stringable) {
	fmt.Println(p.ToString())
    
}
    
