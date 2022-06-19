package others

import "fmt"

type VisitorFunc func(*Info, error) error

type AnotherVisitor interface {
	Visit(VisitorFunc) error
}

type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}

func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}

type NameVisitor struct {
	visitor AnotherVisitor
}

func (v NameVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("NameVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
		}
		fmt.Println("NameVisitor() after call function")
		return err
	})
}

type OtherThingsVisitor struct {
	visitor AnotherVisitor
}

func (v OtherThingsVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("OtherThingsVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
		}
		fmt.Println("OtherThingsVisitor() after call function")
		return err
	})
}

type LogVisitor struct {
	visitor AnotherVisitor
}

func (v LogVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("LogVisitor() before call function")
		err = fn(info, err)
		fmt.Println("LogVisitor() after call function")
		return err
	})
}

type DecorateVisitor struct {
	visitor    AnotherVisitor
	decorators []VisitorFunc
}

func NewDecoratedVisitor(v AnotherVisitor, fn ...VisitorFunc) AnotherVisitor {
	if len(fn) == 0 {
		return v
	}
	return DecorateVisitor{v, fn}
}

// Visit implements Visitor
func (v DecorateVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}
		if err := fn(info, nil); err != nil {
			return err
		}
		for i := range v.decorators {
			if err := v.decorators[i](info, err); err != nil {
				return err
			}
		}
		return nil
	})
}

func NameVisitorFun(info *Info, err error) error {
	fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
	return err
}

func OtherThinsVisitorFun(info *Info, err error) error {
	fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
	return err
}
