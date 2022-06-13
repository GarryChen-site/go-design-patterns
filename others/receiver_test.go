package others

import "testing"

func Test_receiver(t *testing.T) {
	var p = Person{
		Name: "Hao Chen",        
		Sexual: "Male",        
		Age: 44,    
	}    
	PrintPerson(&p)    
	p.Print()
}