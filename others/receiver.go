package others

import "fmt"


type Person struct {
	Name	string
	Sexual	string
	Age	int
}

func PrintPerson(p *Person) {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",p.Name, p.Sexual, p.Age)
}
    
func (p *Person) Print() {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",p.Name, p.Sexual, p.Age)
}
    
