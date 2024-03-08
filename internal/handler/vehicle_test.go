package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/service"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

// TestHandlerVehicle_FindByColorAndYear is a test that checks the method FindByColorAndYear of the handler HandlerVehicle
func TestHandlerVehicle_FindByColorAndYear(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// arrenge
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
		sv := service.NewServiceVehicleDefaultMock()
		sv.On("FindByColorAndYear", "red", 2010).Return(vehicleMap, nil)

		h := handler.NewHandlerVehicle(sv)
		hFindByColorAndYear := h.FindByColorAndYear()

		// expected
		expectedStatus := http.StatusOK
		expectedBody := `
		{
			"data": {
				"1": {
					"Id": 1,
					"Brand": "Ford",
					"Model": "Fiesta",
					"Registration": "0",
					"Color": "red",
					"FabricationYear": 2010,
					"Capacity": 4,
					"MaxSpeed": 200,
					"FuelType": "gasoline",
					"Transmission": "manual",
					"Weight": 1000,
					"Height": 100,
					"Length": 200,
					"Width": 300
				}
			},
			"message": "vehicles found"
		}
		`
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		// act
		request := httptest.NewRequest(http.MethodGet, "/vehicles/color/{color}/year/{year}", nil)
		response := httptest.NewRecorder()

		chiContext := chi.NewRouteContext()
		chiContext.URLParams.Add("color", "red")
		chiContext.URLParams.Add("year", "2010")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, chiContext))
		hFindByColorAndYear.ServeHTTP(response, request)
		// assert
		require.Equal(t, expectedStatus, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
	})
	t.Run("error", func(t *testing.T) {
		// arrenge
		emptyVehicleMap := map[int]internal.Vehicle{}

		sv := service.NewServiceVehicleDefaultMock()
		sv.On("FindByColorAndYear", "red", 2010).Return(emptyVehicleMap, nil)

		h := handler.NewHandlerVehicle(sv)
		hFindByColorAndYear := h.FindByColorAndYear()

		// expected
		expectedStatus := http.StatusOK
		expectedBody := `
		{
			"data": {},
			"message": "vehicles found"
		}
		`
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}

		// act
		request := httptest.NewRequest(http.MethodGet, "/vehicles/color/{color}/year/{year}", nil)
		reponse := httptest.NewRecorder()

		chiContext := chi.NewRouteContext()
		chiContext.URLParams.Add("color", "red")
		chiContext.URLParams.Add("year", "2010")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, chiContext))
		hFindByColorAndYear.ServeHTTP(reponse, request)

		// assert
		require.Equal(t, expectedStatus, reponse.Code)
		require.JSONEq(t, expectedBody, reponse.Body.String())
		require.Equal(t, expectedHeaders, reponse.Header())
	})
}
