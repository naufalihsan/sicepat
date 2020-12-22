package v1

type GetDestinationListRequest struct {
	UserID  string `json:"userid"`
	Command string `json:"command"`
	City    string `json:"city"`
}

type GetDestinationListResponse struct {
	Data    []DestinationList `json:"destination"`
	Status  string            `json:"status"`
	Message string            `json:"message"`
}

type DestinationList struct {
	Province        string `json:"province"`
	City            string `json:"city"`
	District        string `json:"district"`
	DestinationCode string `json:"destination_code"`
}
