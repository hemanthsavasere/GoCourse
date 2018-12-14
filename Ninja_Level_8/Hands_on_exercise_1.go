package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

type persona struct {
	First   string
	Last    string
	Age     string
	Sayings []string
}

func main() {
	u1 := person{
		Name: "Hemanth Savasere",
		Age:  23,
	}

	u2 := person{
		Name: "Steve Jobs",
		Age:  54,
	}

	u3 := person{
		Name: "Sharath Savasere",
		Age:  21,
	}

	people := []person{}
	people = append(people, u1, u2, u3)

	objectEncoded, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error generated")
	}
	fmt.Println(objectEncoded)
	fmt.Println(people)

	holaPeople := []persona{}

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	k := []byte(s)
	err = json.Unmarshal(k, &holaPeople)

	fmt.Println(holaPeople)

}
