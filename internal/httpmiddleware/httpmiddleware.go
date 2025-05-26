package httpmiddleware

import "net/http"

func AcceptJSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accept := r.Header.Get("Accept")
		if accept != "" && accept != "*/*" && accept != "application/json" {
			http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
			return
		}
		next.ServeHTTP(w, r)
	})
}
