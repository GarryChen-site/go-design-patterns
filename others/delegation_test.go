package others

import (
	"fmt"
	"testing"
)

func Test_delegation(t *testing.T) {

	label := Label{Widget{10, 10}, "State:"}
	button1 := Button{Label{Widget{10, 70}, "OK"}}
	button2 := NewButton(50, 70, "Cancel")
	listBox := ListBox{Widget{10, 40}, []string{"AL", "AK", "AZ", "AR"}, 0}
	// 多态
	for _, painter := range []Painter{label, button1, button2, listBox} {
		painter.Paint()
	}
	fmt.Println("**************************")

	for _, widget := range []interface{}{label, listBox, button1, button2} {
		widget.(Painter).Paint()
		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
		fmt.Println() // print a empty line
	}

}
