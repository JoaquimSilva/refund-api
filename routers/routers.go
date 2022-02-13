package routers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"refund-api/constants"
	"refund-api/controller"
	"refund-api/database"
	"refund-api/middleware"
)

func Initialization() {
	router := mux.NewRouter()
	router.Use(middleware.ContentTypeMiddleware)

	router.HandleFunc(constants.RefundPath, controller.GetAllRefund).Methods(http.MethodGet)

	router.HandleFunc(constants.RefundPathById, controller.GetRefundById).Methods(http.MethodGet)

	router.HandleFunc(constants.RefundPathById, controller.UpdateRefund).Methods(http.MethodPut)

	router.HandleFunc(constants.RefundPath, controller.AddRefund).Methods(http.MethodPost)

	router.HandleFunc(constants.RefundPathById, controller.DeleteRefund).Methods(http.MethodDelete)

	router.HandleFunc(constants.RefundPathAgency, controller.GetRefundByAgency).Methods(http.MethodGet)

	router.HandleFunc(constants.RefundPathByTicket, controller.FindRefundsPerTicketNumber).Methods(http.MethodGet)

	database.ConnectDB()
	log.Fatal(http.ListenAndServe(constants.Port, router))
}
