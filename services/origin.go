package services

import (
	v1 "github.com/naufalihsan/sicepat/models/v1"
	"net/http"
)

func GetProvinceList(request *v1.GetProvinceListRequest) (*v1.GetProvinceListResponse, *error) {
	response := &v1.GetProvinceListResponse{}
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

func GetCityList(request *v1.GetCityListRequest) (*v1.GetCityListResponse, *error) {
	response := &v1.GetCityListResponse{}
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

func GetDistrictList(request *v1.GetDistrictListRequest) (*v1.GetDistrictListResponse, *error) {
	response := &v1.GetDistrictListResponse{}
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

func GetSubDistrictList(request *v1.GetSubDistrictListRequest) (*v1.GetSubDistrictListResponse, *error) {
	response := &v1.GetSubDistrictListResponse{}
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
