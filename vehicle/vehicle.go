package vehicle

type Vehicle struct {
	RegistrationNumber string
	Color              Color
}

func NewVehicle(registrationNumber string, color Color) *Vehicle {
	return &Vehicle{
		RegistrationNumber: registrationNumber,
		Color:              color,
	}
}
