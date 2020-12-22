package v1

type GetProvinceListRequest struct {
	UserID  string `json:"userid"`
	Command string `json:"command"`
}

type GetProvinceListResponse struct {
	Data    ProvinceList `json:"data"`
	Status  string       `json:"status"`
	Message string       `json:"message"`
}

type ProvinceList struct {
	Province []string `json:"province"`
}

type GetCityListRequest struct {
	UserID   string `json:"userid"`
	Command  string `json:"command"`
	Province string `json:"province"`
}

type GetCityListResponse struct {
	Data    CityList `json:"data"`
	Status  string   `json:"status"`
	Message string   `json:"message"`
}

type CityList struct {
	Province string   `json:"province"`
	City     []string `json:"city"`
}

type GetDistrictListRequest struct {
	UserID   string `json:"userid"`
	Command  string `json:"command"`
	Province string `json:"province"`
	City     string `json:"city"`
}

type GetDistrictListResponse struct {
	Data    DistrictList `json:"data"`
	Status  string       `json:"status"`
	Message string       `json:"message"`
}

type DistrictList struct {
	Province string           `json:"province"`
	City     string           `json:"city"`
	District []DistrictDetail `json:"district"`
}

type DistrictDetail struct {
	District   string `json:"district"`
	OriginCode string `json:"origin_code"`
}

type GetSubDistrictListRequest struct {
	UserID   string `json:"userid"`
	Command  string `json:"command"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
}

type GetSubDistrictListResponse struct {
	Data    SubDistrictList `json:"data"`
	Status  string       `json:"status"`
	Message string       `json:"message"`
}

type SubDistrictList struct {
	Province    string              `json:"province"`
	City        string              `json:"city"`
	District    string              `json:"district"`
	SubDistrict []SubDistrictDetail `json:"subdistrict"`
}

type SubDistrictDetail struct {
	SubDistrict string `json:"subdistrict"`
	ZipCode     string `json:"zipcode"`
	OriginCode  string `json:"origin_code"`
}
