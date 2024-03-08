package service_test

import (
	"app/internal"
	"app/internal/repository"
	"app/internal/service"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestServiceVehicleDefault_FindByColorAndYear is a test that checks the method FindByColorAndYear of the service ServiceVehicleDefault
func TestServiceVehicleDefault_FindByColorAndYear(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// arrange
		vehicleMap := map[int]internal.Vehicle{
			1: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Ford",
					Model:           "Fiesta",
					Registration:    "0",
					FabricationYear: 2010,
					Color:           "red",
					Capacity:        4,
					MaxSpeed:        200,
					FuelType:        "gasoline",
					Transmission:    "manual",
					Weight:          1000,
					Dimensions: internal.Dimensions{
						Height: 100,
						Length: 200,
						Width:  300,
					},
				},
			},
		}

		rp := repository.NewVehicleMock()
		rp.On("FindByColorAndYear", "red", 2010).Return(vehicleMap, nil)

		s := service.NewServiceVehicleDefault(rp)
		// act
		v, err := s.FindByColorAndYear("red", 2010)
		// assert
		require.NoError(t, err)
		require.Len(t, v, 1)
		require.Equal(t, 1, v[1].Id)
		require.Equal(t, vehicleMap[1], v[1])
	})
}
