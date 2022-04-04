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

	var newRefundDto models.RefundDto
	var refundBase models.Refund
	var refund2Persist models.Refund
	err := json.NewDecoder(request.Body).Decode(&newRefundDto)
	utils.CatchError(err)

	ticketNumber := newRefundDto.TicketNumber
	database.FindByTicket64(ticketNumber, refundBase)

	if refundBase.Id > 0 {
		utils.FindRegisterByTicket(writer)
		return
	} else {

		//date := time.Now()
		refund2Persist.AgencyId = newRefundDto.AgencyId
		refund2Persist.RequestedDate = time.Now().Format(time.RFC3339)
		refund2Persist.ShippingDate = time.RFC3339
		refund2Persist.DueDate = time.RFC3339
		refund2Persist.Branch = newRefundDto.Branch
		refund2Persist.TicketNumber = newRefundDto.TicketNumber
		refund2Persist.Passenger = newRefundDto.Passenger
		refund2Persist.ReservationCode = newRefundDto.ReservationCode
		refund2Persist.StatusCode = 1
		refund2Persist.ConsolidatorId = newRefundDto.ConsolidatorId
		refund2Persist.IssueConsolidatorId = newRefundDto.IssueConsolidatorId
		refund2Persist.InvoiceNumber = 0
		refund2Persist.ReservationId = newRefundDto.ReservationId
		refund2Persist.UserId = newRefundDto.UserId
		refund2Persist.NetValue = 0.00
		refund2Persist.Processed = false
		refund2Persist.Internal = false
		refund2Persist.NotifyBackoffice = newRefundDto.NotifyBackoffice
		refund2Persist.Released = false

		database.DB.Create(&refund2Persist)
	}

	err2 := json.NewEncoder(writer).Encode(refund2Persist)
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
