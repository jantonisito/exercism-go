// Package census simulates a system used to collect census data.
package census

import "fmt"

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	fmt.Println(name, age, address)
	res := new(Resident)
	res.Name = name
	res.Age = age
	res.Address = make(map[string]string)
	if len(address) == 0 {
		res.Address = nil
		return res
	}
	for k, v := range address {
		res.Address[k] = v
	}
	return res
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	switch {
	case r.Name == "":
		return false

	case r.Age < 0:
		return false
	case len(r.Address) == 0:
		return false
	}
	if _, exists := r.Address["street"]; !exists {
		return false
	}
	for _, addr := range r.Address {
		if len(addr) == 0 {
			return false
		}
	}
	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	total := 0
	for _, res := range residents {
		if res.HasRequiredInfo() {
			total++
		}
	}
	return total
}
