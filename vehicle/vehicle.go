package vehicle

import "errors"

type Vehicle struct {
	RegistrationNumber string
	Color              Color
}

func NewVehicle(registrationNumber string, color Color) (*Vehicle, error) {
	if registrationNumber == "" {
		return nil, errors.New("Registration Number cannot be empty")
	}
	if color < WHITE || color > BLUE {
		return nil, errors.New("Invalid color")
	}
	return &Vehicle{
		RegistrationNumber: registrationNumber,
		Color:              color,
	}, nil
}
