package entity

type Location struct {
	Longitude float32
	Latitude  float32
	city      string
}
//d
func (location Location) GetCity() string {
	return location.city
}
