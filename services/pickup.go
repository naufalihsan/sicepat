package services

import (
	v1 "github.com/naufalihsan/sicepat/models/v1"
	"net/http"
)

func CheckTariff(request *v1.CheckTariffRequest) (*v1.CheckTariffResponse, *error) {
	response := &v1.CheckTariffResponse{}
	header := &http.Header{}

	err := Call(
		header,
		request,
		response,
	)

	if err != nil {
		panic(err)
	}

	return response, nil

}

func CalculateInsurance(request *v1.CalculateInsuranceRequest) (*v1.CalculateInsuranceResponse, *error) {
	response := &v1.CalculateInsuranceResponse{}
	header := &http.Header{}

	err := Call(
		header,
		request,
		response,
	)

	if err != nil {
		panic(err)
	}

	return response, nil

}

func Pickup(request *v1.PickupRequest) (*v1.PickupResponse, *error) {
	response := &v1.PickupResponse{}
	header := &http.Header{}

	err := Call(
		header,
		request,
		response,
	)

	if err != nil {
		panic(err)
	}

	return response, nil

}
