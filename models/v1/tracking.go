package v1

type TrackingRequest struct {
	UserID  string `json:"userid"`
	Command string `json:"command"`
	Awb     string `json:"awb"`
}

type TrackingResponse struct {
	Data    TrackingDetail `json:"data"`
	Status  string         `json:"status"`
	Message string         `json:"message"`
}

type TrackingDetail struct {
	Awb              string               `json:"awb"`
	DeliveryType     string               `json:"delivery_type"`
	Weight           string               `json:"weight"`
	ShipperName      string               `json:"shipper_name"`
	ShipperAddress   string               `json:"shipper_address"`
	RecipientName    string               `json:"recipient_name"`
	RecipientAddress string               `json:"recipient_address"`
	SendDate         string               `json:"send_date"`
	TrackHistory     []TrackHistoryDetail `json:"track_history"`
}

type TrackHistoryDetail struct {
	DateTime string `json:"date_time"`
	Status   string `json:"status"`
	City     string `json:"city"`
}
