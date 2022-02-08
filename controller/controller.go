package controller

import (
	"log"
	"net/http"
	"refund-api/constants"
	"refund-api/database"
	"refund-api/middleware"
	"refund-api/service"

	"github.com/gorilla/mux"
)

func Initialization() {
	router := mux.NewRouter()
	router.Use(middleware.ContentTypeMiddleware)

	router.HandleFunc(constants.RefundPath, service.GetAllRefund).Methods(http.MethodGet)

	router.HandleFunc(constants.RefundPathById, service.GetRefundById).Methods(http.MethodGet)

	router.HandleFunc(constants.RefundPathById, service.UpdateRefund).Methods(http.MethodPut)

	router.HandleFunc(constants.RefundPath, service.AddRefund).Methods(http.MethodPost)

	router.HandleFunc(constants.RefundPathById, service.DeleteRefund).Methods(http.MethodDelete)

	router.HandleFunc(constants.RefundPathAgency, service.GetRefundByAgency).Methods(http.MethodGet)

	router.HandleFunc(constants.RefundPathByTicket, service.FindRefundsPerTicketNumber).Methods(http.MethodGet)

	database.ConnectDB()
	log.Fatal(http.ListenAndServe(constants.Port, router))
}
