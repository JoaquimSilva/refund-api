package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"refund-api/constants"
	"refund-api/controller"
	"refund-api/database"
	"refund-api/models"
	"strconv"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var ID int

func Router() *mux.Router {
	router := mux.NewRouter()
	return router
}

func TestGetAllRefunds(t *testing.T) {
	router := Router()
	database.ConnectDB()
	router.HandleFunc(constants.RefundPath, controller.GetAllRefund).Methods(http.MethodGet)

	req := httptest.NewRequest(http.MethodGet, constants.RefundPath, nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Error in test")
}

func TestRefundByID(t *testing.T) {
	database.ConnectDB()
	CreateDtoMock()
	defer DeleteRefundMock()
	router := Router()
	router.HandleFunc(constants.RefundPathById, controller.GetRefundById).Methods(http.MethodGet)

	id := strconv.Itoa(ID)
	req := httptest.NewRequest(http.MethodGet, "/refunds/"+id, nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Error in test")
	fmt.Println()
}

func DeleteRefundMock() {
	var refund models.Refund
	database.DB.Delete(&refund, ID)
}

func CreateDtoMock() {
	dto := models.RefundDto{

		AgencyId:            5190,
		Branch:              100,
		TicketNumber:        9570000111112,
		Passenger:           "Silva/Whashinton",
		ReservationCode:     "XTPTPO",
		ConsolidatorId:      7,
		IssueConsolidatorId: 7,
		ReservationId:       23123123,
		UserId:              990022,
		Internal:            true,
	}

	var toDataBase models.Refund

	toDataBase.AgencyId = dto.AgencyId
	toDataBase.RequestedDate = time.Now().Format("2006-01-02 15:04:05")
	toDataBase.ShippingDate = time.Date(9999, 12, 31, 00, 00, 00, 00, time.UTC).String()
	toDataBase.DueDate = time.Date(9999, 12, 31, 00, 00, 00, 00, time.UTC).String()
	toDataBase.Branch = dto.Branch
	toDataBase.TicketNumber = dto.TicketNumber
	toDataBase.Passenger = dto.Passenger
	toDataBase.ReservationCode = dto.ReservationCode
	toDataBase.StatusCode = 1
	toDataBase.ConsolidatorId = dto.ConsolidatorId
	toDataBase.IssueConsolidatorId = dto.IssueConsolidatorId
	toDataBase.InvoiceNumber = 0
	toDataBase.ReservationId = dto.ReservationId
	toDataBase.UserId = dto.UserId
	toDataBase.NetValue = 0.00
	toDataBase.Processed = false
	toDataBase.Internal = false
	toDataBase.NotifyBackoffice = true
	toDataBase.Released = false

	database.DB.Create(&toDataBase)
	ID = toDataBase.Id
}
