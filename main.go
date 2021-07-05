package main

import (
	"fmt"
	"go/golang/entity"
)

func main() {
	companies := entity.InitCompanies()
	fmt.Println(companies)
	fmt.Print("The number of Employees is: ")
	fmt.Println(CountofEmployees(companies, "Atoss", "Timisoara"))

}
func CountofEmployees(companies []entity.Company, name string, city string) int {
	counter := 0
	for _, company := range companies {
		if company.Name == name {
			for _, employee := range company.List_of_Employees {
				if employee.Location.GetCity() == city {
					counter++
				}
			}
		}
	}
	return counter
}
