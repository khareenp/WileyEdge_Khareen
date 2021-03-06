package notes
import "fmt"


type contactInfo struct {
	Email   string
	ZipCode int
}

type Person1 struct {
	FirstName string
	LastName  string
	// ContactInfo is a pointer to a contactInfo struct
	contactInfo
}

type Person struct {
	FirstName string
	LastName  string
	// Contact is a pointer to a contactInfo struct
	Contact contactInfo
}
func PointertoStruct() {
    /*
	// Create a new Person struct (Method 1)
	alex := Person{"Alex", "Anderson", contactInfo{}}
	alex.print()

	// Create a new Person struct (Method 2)
	alex = Person{FirstName: "Alex", LastName: "Anderson", Contact: contactInfo{}}
	alex.print()

	// Create a new Person struct (Method 3). Assigns ZERO Values to the fields
	var alex2 Person
	alex2.FirstName = "Alex"
	alex2.LastName = "Anderson"
	alex2.Contact.Email = "a.b@c.com"
	alex2.Contact.ZipCode = 12345
	alex2.print()
    
	jim := Person{
		FirstName: "Jim",
		LastName:  "Anderson",
		Contact: contactInfo{
			Email:   "a.b@c.com",
			ZipCode: 12345,
		},
	}
	jim.print()
   */
	jill := Person1{
		FirstName: "Jill",
		LastName:  "Jackson",
		contactInfo: contactInfo{
			Email:   "Jill.Jackson@Sample.com",
			ZipCode: 12345,
		},
	}

	// It will not update the FirstName field. As this is Pass By Value
	jill.updateFirstName("Jilly")
	jill.print()

	// It will update the FirstName field. As this is Pass By Reference
	jillPointer := &jill   //Assign address of jill to jillPointer
	jillPointer.updateFirstNameV1("Jilly")
	jill.print()

	// It will update the FirstName field. As this is Pass By Reference
	//(&jill).updateFirstNameV1("Jilly V1")
	//jill.print()

	// It will update the FirstName field. As this is Pass By Reference
	//jill.updateFirstNameV1("Jilly V2")
	//jill.print()
}

func (p Person) print() {
	fmt.Println(p)
	fmt.Printf("%+v \n", p)
}

func (p Person1) print() {
	fmt.Println(p)
	fmt.Printf("%+v \n", p)
}
// Passing value to function
func (p Person1) updateFirstName(firstName string) {
	p.FirstName = firstName
}

func (p *Person1) updateFirstNameV1(firstName string) {
	(*p).FirstName = firstName
}