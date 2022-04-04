package tests

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"refund-api/constants"
	"refund-api/controller"
	"refund-api/database"
	"refund-api/models"
	"strconv"
	"testing"
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
	CreateRefundMock()
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

func CreateRefundMock() {
	refund := models.Refund{
		AgencyId: 9876,
		Passenger: "TESTE TESTE",
		TicketNumber: 9570000111000,
		ReservationCode: "12345690",
		UserId: 123819,
		Branch: 281,
		StatusCode: 1,
		RefundNumber: 328327748,
		ReservationId: 8877667,
		InvoiceNumber: 89876889,
		ConsolidatorId: 7,
		IssueConsolidatorId: 7,
		NetValue: 12903,
		Internal: false,
		Released: false,
		Processed: false,
		NotifyBackoffice: false,
	}

	database.DB.Create(&refund)
	ID = refund.Id
}
