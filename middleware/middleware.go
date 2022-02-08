package middleware

import (
	"net/http"
	"refund-api/constants"
)

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set(constants.ContentType, constants.ApplicationJson)
		next.ServeHTTP(writer, request)
	})
}
