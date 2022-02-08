package service

import (
	"encoding/json"
	"net/http"
	"refund-api/constants"
	"refund-api/database"
	"refund-api/models"
	"refund-api/utils"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetAllRefund(writer http.ResponseWriter, _ *http.Request) {

	var refunds []models.Refund

	database.DB.Find(&refunds)

	json.NewEncoder(writer).Encode(refunds)
}

func GetRefundById(writer http.ResponseWriter, request *http.Request) {

	var refund models.Refund

	var id = mux.Vars(request)[constants.ID]

	database.DB.Find(&refund, id)

	if utils.RefundIsValid(writer, refund) {
		return
	}

	json.NewEncoder(writer).Encode(refund)
}

func UpdateRefund(writer http.ResponseWriter, request *http.Request) {

	var refund models.Refund

	var id = mux.Vars(request)[constants.ID]

	database.DB.Find(&refund, id)

	if utils.RefundIsValid(writer, refund) {
		return
	}

	json.NewDecoder(request.Body).Decode(&refund)

	database.DB.Save(&refund)

	json.NewEncoder(writer).Encode(refund)
}

func AddRefund(writer http.ResponseWriter, request *http.Request) {

	var newRefund models.Refund

	json.NewDecoder(request.Body).Decode(&newRefund)

	database.DB.Create(&newRefund)

	json.NewEncoder(writer).Encode(newRefund)
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

	if start != constants.Empty && end != constants.Empty {
		database.FindByAgencyIdAndPerPeriod(agency, start, end, refunds)
	} else {
		database.FindByAgencyId(agency, refunds)
	}
	if refunds == nil {
		utils.NotFound(writer)
		return
	}

	json.NewEncoder(writer).Encode(refunds)
}

func funcName(agency string, refunds []models.Refund) *gorm.DB {
	return database.DB.Where("agency_id = ? ", agency).Find(&refunds)
}

func FindRefundsPerTicketNumber(writer http.ResponseWriter, request *http.Request) {

	var refund models.Refund
	number := request.URL.Query().Get(constants.Number)
	database.DB.Where("ticket_number = ? ", number).Find(&refund)

	if refund.Id == 0 {
		utils.NotFound(writer)
		return
	}

	json.NewEncoder(writer).Encode(refund)
}
