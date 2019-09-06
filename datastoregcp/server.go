package datastoregcp


import (
	"context"
	"net/http"
	httptransport "github.com/go-kit/kit/transport/http"
    "github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.Methods("POST").Path("/getObject").Handler(httptransport.NewServer(
		endpoints.GetEndpoint,
		decodeGetRequest,
		encodeResponse,
	))
	r.Methods("POST").Path("/storeObject").Handler(httptransport.NewServer(
		endpoints.StoreEndpoint,
		decodeStoreRequest,
		encodeResponse,
	))
	return r
}


func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}