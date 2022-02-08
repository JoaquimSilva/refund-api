package utils

import (
	"log"
	"net/http"
	"refund-api/constants"
	"refund-api/models"
)

func CatchError(err error) {
	if err != nil {
		log.Panic(err.Error())
		return
	}
}
func CatchErrorDB(err error) {
	if err != nil {
		log.Panicf("Erro connect to database, erro: %s", err.Error())
		return
	}
}

func NotFound(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusNotFound)
	return
}

func RefundIsValid(writer http.ResponseWriter, refund models.Refund) bool {
	if refund.Id == 0 || refund.Id < 0 {
		writer.WriteHeader(http.StatusNotFound)
		_, err := writer.Write([]byte(constants.IdNotFound))
		CatchError(err)
		return true
	}
	return false
}
