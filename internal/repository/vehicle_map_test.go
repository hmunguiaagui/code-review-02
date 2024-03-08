package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestRepositoryReadVehicleMap_FindAll is a test that checks the method FindAll of the repository RepositoryReadVehicleMap
func TestRepositoryReadVehicleMap_FindAll(t *testing.T) {
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

		rp := repository.NewRepositoryReadVehicleMap(vehicleMap)
		// act
		v, err := rp.FindAll()
		// assert
		require.NoError(t, err)
		require.Len(t, v, 1)
		require.Equal(t, 1, v[1].Id)
		require.Equal(t, vehicleMap, v)
	})
	t.Run("error", func(t *testing.T) {
		// arrange
		emptyVehicleMap := map[int]internal.Vehicle{}

		rp := repository.NewRepositoryReadVehicleMap(emptyVehicleMap)
		// act
		v, err := rp.FindAll()
		// assert
		require.NoError(t, err)
		require.Len(t, v, 0)
		require.Equal(t, emptyVehicleMap, v)
	})
}
