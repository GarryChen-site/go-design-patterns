package others

import "testing"

func Test_kubectl(t *testing.T) {
	info := Info{}

	var v AnotherVisitor = &info

	v = LogVisitor{v}
	v = NameVisitor{v}
	v = OtherThingsVisitor{v}

	//info := Info{}
	//var v AnotherVisitor = &info
	//v = NewDecoratedVisitor(v, NameVisitorFun, OtherThinsVisitorFun)

	loadFile := func(info *Info, err error) error {
		info.Name = "Tao Bai Bai"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}

	v.Visit(loadFile)
}
