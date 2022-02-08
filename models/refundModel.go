package models

type Refund struct {
	Id                  int     `json:"id"`
	AgencyId            int64   `json:"agencyId"`
	DateRequested       string  `json:"dateRequested"`
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
	RefundNumber        int64   `json:"refundNumber"`
	ReservationId       int64   `json:"reservationId"`
	UserId              int64   `json:"userId"`
	NetValue            float64 `json:"netValue"`
	Processed           bool    `json:"processed"`
	Internal            bool    `json:"internal"`
	NotifyBackoffice    bool    `json:"notifyBackoffice"`
	Released            bool    `json:"released"`
}
