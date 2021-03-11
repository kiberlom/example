package main

import "fmt"

type User struct {
	ID   *int
	Name *string
}

func (u *User) setID(i int) {
	u.ID = &i
}

func (u *User) setName(i string) {
	u.Name = &i
}

func main() {

	u := User{}

	u.setID(12)
	u.setName("Миша")

	if u.ID != nil {
		fmt.Println(*u.ID)
	}

	if u.Name != nil {
		fmt.Println(*u.Name)
	}

}
