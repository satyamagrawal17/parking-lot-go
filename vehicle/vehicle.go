package vehicle

import "errors"

type Vehicle struct {
	registrationNumber string
	color              Color
}

func NewVehicle(registrationNumber string, color Color) (*Vehicle, error) {
	if registrationNumber == "" {
		return nil, errors.New("registration Number cannot be empty")
	}
	if color < WHITE || color > BLUE {
		return nil, errors.New("invalid color")
	}
	return &Vehicle{
		registrationNumber: registrationNumber,
		color:              color,
	}, nil
}
