package handlers

import (
	"net/http"
	"sync/atomic"
)

func readyz(isReady *atomic.Value) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if isReady == nil || !isReady.Load().(bool) {
			http.Error(writer, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}
}