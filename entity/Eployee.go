package entity

type Employee struct {
	Name     string
	Age      int
	Location Location
	Gender   string
}
//d
func (employee Employee) GetCity() string {
	return employee.Location.city
}
