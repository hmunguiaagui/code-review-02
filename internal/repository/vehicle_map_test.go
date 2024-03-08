package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestRepositoryVehicleMap_FindByBrand is a method that returns a map of vehicles that match the brand
func TestRepositoryVehicleMap_FindByBrand(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// /vehicles/average_speed/brand/{brand}
		// arrange
		vehicleMap := map[int]internal.Vehicle{
			1: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Ford",
					Model:           "Fiesta",
					Registration:    "ABC-1234",
					Color:           "Red",
					FabricationYear: 2010,
					Capacity:        5,
					MaxSpeed:        180,
					FuelType:        "Gasoline",
					Transmission:    "Manual",
					Weight:          1000,
					Dimensions: internal.Dimensions{
						Height: 1.5,
						Length: 4,
						Width:  1.8,
					},
				},
			},
		}
		rp := repository.NewRepositoryReadVehicleMap(vehicleMap)
		// act
		v, err := rp.FindByBrand("Ford")
		// assert
		require.NoError(t, err)
		require.Len(t, v, 1)
		require.Equal(t, 1, v[1].Id)
		require.Equal(t, vehicleMap, v)
	})
	t.Run("error", func(t *testing.T) {
		// arrange
		vehicleMap := map[int]internal.Vehicle{}
		rp := repository.NewRepositoryReadVehicleMap(vehicleMap)
		// act
		v, err := rp.FindByBrand("Ford")
		// assert
		require.Error(t, err)
		require.EqualError(t, err, "not found")
		require.Len(t, v, 0)
		require.Equal(t, vehicleMap, v)
	})
}
