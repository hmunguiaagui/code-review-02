package repository

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

// NewVehicleMock creates a new instance of VehicleMock
func NewVehicleMock() *VehicleMock {
	return &VehicleMock{}
}

// VehicleMock is a struct that represents a vehicle mock
type VehicleMock struct {
	// mock is the mock of the vehicle
	mock.Mock
}

// FindAll is a method that returns a map of all vehicles
func (m *VehicleMock) FindAll() (v map[int]internal.Vehicle, err error) {
	args := m.Called()
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

// FindByColorAndYear is a method that returns a map of vehicles that match the color and fabrication year
func (m *VehicleMock) FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
	args := m.Called(color, fabricationYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

// FindByBrandAndYearRange is a method that returns a map of vehicles that match the brand and a range of fabrication years
func (m *VehicleMock) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	args := m.Called(brand, startYear, endYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

// FindByBrand is a method that returns a map of vehicles that match the brand
func (m *VehicleMock) FindByBrand(brand string) (v map[int]internal.Vehicle, err error) {
	args := m.Called(brand)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

// FindByWeightRange is a method that returns a map of vehicles that match a range of weight
func (m *VehicleMock) FindByWeightRange(startWeight float64, endWeight float64) (v map[int]internal.Vehicle, err error) {
	args := m.Called(startWeight, endWeight)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}
