package models

type Refund struct {
	Id                  int     `json:"id" gorm:"primaryKey;index"`
	AgencyId            int64   `json:"agencyId"`
	RequestedDate       string  `json:"dateRequested" gorm:"autoCreateTime"`
	ShippingDate        string  `json:"shippingDate"`
	DueDate             string  `json:"dueDate"`
	Branch              int64   `json:"branch"`
	TicketNumber        int64   `json:"ticketNumber"`
	Passenger           string  `json:"passenger"`
	ReservationCode     string  `json:"reservationCode"`
	StatusCode          int     `json:"statusCode"`
	ConsolidatorId      int     `json:"consolidatorId"`
	IssueConsolidatorId int     `json:"issueConsolidatorId"`
	InvoiceNumber       int64   `json:"invoiceNumber"`
	RefundNumber        int64   `json:"refundNumber" gorm:"autoIncrement;index"`
	ReservationId       int64   `json:"reservationId"`
	UserId              int64   `json:"userId"`
	NetValue            float64 `json:"netValue"`
	Processed           bool    `json:"processed"`
	Internal            bool    `json:"internal"`
	NotifyBackoffice    bool    `json:"notifyBackoffice"`
	Released            bool    `json:"released"`
}

type RefundDto struct {
	AgencyId            int64
	Branch              int64
	TicketNumber        int64
	Passenger           string
	ReservationCode     string
	ConsolidatorId      int
	IssueConsolidatorId int
	ReservationId       int64
	UserId              int64
	Internal            bool
	NotifyBackoffice    bool
}
