package entity

type Location struct {
	Longitude float32
	Latitude  float32
	city      string
}

func (location Location) GetCity() string {
	return location.city
}
