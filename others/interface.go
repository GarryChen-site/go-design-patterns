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

// *********************************

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

// ********************************

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

func (c City) ToString() string {
	return "City = " + c.Name
}

func PrintStr(p Stringable) {
	fmt.Println(p.ToString())
}

// 在这段代码中，我们可以看到，我们使用了一个叫Stringable 的接口，
//我们用这个接口把“业务类型” Country 和 City 和“控制逻辑” Print() 给解耦了。
//于是，只要实现了Stringable 接口，都可以传给 PrintStr() 来使用。
