package v1

type CheckTariffRequest struct {
	UserID          string  `json:"userid"`
	Command         string  `json:"command"`
	OriginCode      string  `json:"origin_code"`
	DestinationCode string  `json:"destination_code"`
	Weight          float32 `json:"weight"`
}

type CheckTariffResponse struct {
	Data    TariffList `json:"data"`
	Status  string     `json:"status"`
	Message string     `json:"message"`
}

type TariffList struct {
	OriginCode      string         `json:"origin_code"`
	DestinationCode string         `json:"destination_code"`
	Weight          float32        `json:"weight"`
	Tariff          []TariffDetail `json:"tarif"`
}

type TariffDetail struct {
	DeliveryType string `json:"delivery_type"`
	Tariff       int64  `json:"tarif"`
}

type CalculateInsuranceRequest struct {
	UserID      string `json:"userid"`
	Command     string `json:"command"`
	ParcelValue string `json:"parcel_value"`
}

type CalculateInsuranceResponse struct {
	Data    InsuranceList `json:"data"`
	Status  string        `json:"status"`
	Message string        `json:"message"`
}

type InsuranceList struct {
	ParcelValue    string `json:"parcel_value"`
	InsuranceValue int64  `json:"insurance_value"`
}

type PickupRequest struct {
	UserID  string     `json:"userid"`
	Command string     `json:"command"`
	Data    PickupData `json:"data"`
}

type PickupData struct {
	RefID               string `json:"refid"`
	OriginCode          string `json:"origin_code"`
	Amount              string `json:"amount"`
	PickupRequestDate   string `json:"pickup_request_date"`
	PickupMerchantCode  string `json:"pickup_merchant_code"`
	PickupMerchantName  string `json:"pickup_merchant_name"`
	PickupAddress       string `json:"pickup_address"`
	PickupCity          string `json:"pickup_city"`
	PickupMerchantPhone string `json:"pickup_merchant_phone"`
	PickupMerchantEmail string `json:"pickup_merchant_email"`
	DeliveryType        string `json:"delivery_type"`
	ParcelCategory      string `json:"parcel_category"`
	ParcelContent       string `json:"parcel_content"`
	ParcelValue         string `json:"parcel_value"`
	ParcelWeight        string `json:"parcel_weight"`
	RecipientName       string `json:"recipient_name"`
	RecipientPhone      string `json:"recipient_phone"`
	RecipientAddress    string `json:"recipient_address"`
	RecipientCity       string `json:"recipient_city"`
	RecipientProvince   string `json:"recipient_province"`
	RecipientDistrict   string `json:"recipient_district"`
	RecipientZipCode    string `json:"recipient_zipcode"`
	DestinationCode     string `json:"destination_code"`
	Tariff              string `json:"tarif"`
	ShipperName         string `json:"shipper_name"`
	ShipperPhone        string `json:"shipper_phone"`
	ShipperAddress      string `json:"shipper_address"`
	ShipperProvice      string `json:"shipper_province"`
	ShipperCity         string `json:"shipper_city"`
	ShipperDistrict     string `json:"shipper_district"`
	ShipperZipCode      string `json:"shipper_zipcode"`
	Insurance           string `json:"insurance"`
	Notes               string `json:"notes"`
}

type PickupResponse struct {
	Data    PickupDetail `json:"data"`
	Status  string       `json:"status"`
	Message string       `json:"message"`
}

type PickupDetail struct {
	Awb    string `json:"awb"`
	Amount int64  `json:"amount"`
	RefID  string `json:"refid"`
}
