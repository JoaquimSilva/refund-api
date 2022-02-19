package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"refund-api/constants"
	"refund-api/database"
	"refund-api/models"
	"refund-api/utils"
	"time"
)

func GetAllRefund(writer http.ResponseWriter, _ *http.Request) {

	var refunds []models.Refund

	database.DB.Find(&refunds)

	err := json.NewEncoder(writer).Encode(refunds)
	utils.CatchError(err)
}

func GetRefundById(writer http.ResponseWriter, request *http.Request) {

	var refund models.Refund
	var id = mux.Vars(request)[constants.ID]

	database.DB.Find(&refund, id)

	if utils.RefundIsValid(writer, refund) {
		return
	}

	err := json.NewEncoder(writer).Encode(refund)
	utils.CatchError(err)
}

func UpdateRefund(writer http.ResponseWriter, request *http.Request) {

	var refund models.Refund
	var id = mux.Vars(request)[constants.ID]

	database.DB.Find(&refund, id)

	if utils.RefundIsValid(writer, refund) {
		return
	}

	err := json.NewDecoder(request.Body).Decode(&refund)
	utils.CatchError(err)
	database.DB.Save(&refund)

	err2 := json.NewEncoder(writer).Encode(refund)
	utils.CatchError(err2)
}

func AddRefund(writer http.ResponseWriter, request *http.Request) {

	var newRefund models.Refund
	var refundBase models.Refund
	var refundPersist models.Refund
	err := json.NewDecoder(request.Body).Decode(&newRefund)
	utils.CatchError(err)

	ticketNumber := newRefund.TicketNumber
	database.FindByTicket64(ticketNumber, refundBase)

	if refundBase.Id > 0 {
		utils.FindRegisterByTicket(writer)
		return
	} else {

		//date := time.Now()
		refundPersist.AgencyId = newRefund.AgencyId
		refundPersist.RequestedDate = time.Now().Format(time.RFC3339)
		refundPersist.ShippingDate = time.RFC3339
		refundPersist.DueDate = time.RFC3339
		refundPersist.Branch = newRefund.Branch
		refundPersist.TicketNumber = newRefund.TicketNumber
		refundPersist.Passenger = newRefund.Passenger
		refundPersist.ReservationCode = newRefund.ReservationCode
		refundPersist.StatusCode = 1
		refundPersist.ConsolidatorId = newRefund.ConsolidatorId
		refundPersist.IssueConsolidatorId = newRefund.IssueConsolidatorId
		refundPersist.InvoiceNumber = 0
		refundPersist.ReservationId = newRefund.ReservationId
		refundPersist.UserId = newRefund.UserId
		refundPersist.NetValue = 0.00
		refundPersist.Processed = false
		refundPersist.Internal = false
		refundPersist.NotifyBackoffice = true
		refundPersist.Released = false

		database.DB.Create(&refundPersist)
	}

	err2 := json.NewEncoder(writer).Encode(refundPersist)
	utils.CatchError(err2)
}

func DeleteRefund(_ http.ResponseWriter, request *http.Request) {

	var refund models.Refund
	var id = mux.Vars(request)[constants.ID]
	database.DB.Delete(&refund, id)
}

func GetRefundByAgency(writer http.ResponseWriter, request *http.Request) {

	var refunds []models.Refund

	agency := request.URL.Query().Get(constants.Agency)
	start := request.URL.Query().Get(constants.Start)
	end := request.URL.Query().Get(constants.End)

	if agency != constants.Empty &&
		start == constants.Empty &&
		end == constants.Empty {
		database.FindByAgencyId(agency, refunds)
	}
	if start != constants.Empty &&
		end != constants.Empty {
		database.FindByAgencyIdAndPerPeriod(agency, start, end, refunds)
	}
	if refunds == nil {
		utils.NotFound(writer)
		return
	}

	err := json.NewEncoder(writer).Encode(refunds)
	utils.CatchError(err)
}

func FindRefundsPerTicketNumber(writer http.ResponseWriter, request *http.Request) {

	var refund models.Refund
	number := request.URL.Query().Get(constants.Number)
	database.FindByTicket(number, refund)

	if refund.Id == 0 {
		utils.NotFound(writer)
		return
	}

	err := json.NewEncoder(writer).Encode(refund)
	utils.CatchError(err)
}
