package clase4

import "fmt"

type Task struct {
	Name     string
	Priority int
	Users    []Users
}

type Users struct {
	Name string
}

func (u Users) UserAsigned() string {
	return u.Name
}

func (t Task) UserAssined() []Users {
	return t.Users
}

func (u *Users) ChangeName(name string) {
	u.Name = name
}

func Show() {
	task := Task{Name: "HU-45", Priority: 1, Users: []Users{{Name: "Pedro"}, {Name: "Julian"}, {Name: "Carlos"}}}

	fmt.Println(fmt.Sprintf("Assigned user: %v ", task.UserAssined()))

	contador := 1
	for _, value := range task.Users {
		contador++
		value.ChangeName(fmt.Sprintf("Usuario%v", contador))
		fmt.Println(value.UserAsigned())
	}

}
